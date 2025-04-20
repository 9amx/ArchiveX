package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var (
	extensionList string
	summaryMu     sync.Mutex
	summary       []string

	botToken      = "7666125700:AAGP365K1HopYDxYWPL2DWzCnawglv3suY4"
	chatID  int64 = 7162301957
)

func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func askVerificationCode(chatID int64) bool {
	code := generateCode()
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Telegram bot error:", err)
		return false
	}

	message := fmt.Sprintf(
		"🔐 *Verification Code for ArchiveX*\n\nHello! Your secure 6-digit code is: *%s*\n\n👨‍💻 Tool: ArchiveX\n👤 Created by: @the9am\n\nPlease enter this code in your terminal to begin.",
		code,
	)

	msg := tgbotapi.NewMessage(chatID, message)
	msg.ParseMode = "Markdown"
	bot.Send(msg)

	fmt.Println("\n🔐 ArchiveX Security Check")
	fmt.Println("\n👉 Please start the bot @ehh_wayback_bot on Telegram.")
	fmt.Println("📩 You’ll receive a 6-digit code there. Enter it below to begin.")
	fmt.Print("\n🔑 Enter the code: ")

	var input string
	fmt.Scanln(&input)

	return input == code
}

func fetchWithRetry(url string, retries int) (*http.Response, error) {
	var resp *http.Response
	var err error
	for i := 0; i < retries; i++ {
		resp, err = http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			return resp, nil
		}
		time.Sleep(2 * time.Second)
	}
	return nil, errors.New("failed after retries")
}

func fetchAndSave(domain, url, outputFile, label string, counterChan chan int) error {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " " + label
	s.Start()

	resp, err := fetchWithRetry(url, 3)
	if err != nil {
		s.Stop()
		return err
	}
	defer resp.Body.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		s.Stop()
		return err
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(resp.Body)
	count := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			_, _ = outFile.WriteString(line + "\n")
			count++
			if counterChan != nil {
				counterChan <- 1
			}
		}
	}

	s.Stop()
	summaryMu.Lock()
	summary = append(summary, fmt.Sprintf("%s - %s: %d URLs", domain, label, count))
	summaryMu.Unlock()

	return nil
}

func sanitizeFileName(domain string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return reg.ReplaceAllString(domain, "_")
}

func processDomain(domain, timestamp, baseOutput, extRegex string, wg *sync.WaitGroup) {
	defer wg.Done()

	safeDomain := sanitizeFileName(domain)
	outputDir := filepath.Join(baseOutput, safeDomain)
	_ = os.MkdirAll(outputDir, os.ModePerm)

	allFile := filepath.Join(outputDir, "all.txt")
	filteredFile := filepath.Join(outputDir, "filtered.txt")

	allURL := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=*.%s/*&collapse=urlkey&output=text&fl=original", domain)
	filteredURL := fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=*.%s/*&collapse=urlkey&output=text&fl=original&filter=original:.*\\.(xls|xml|xlsx|json|pdf|sql|doc|docx|pptx|txt|git|zip|tar\\.gz|tgz|bak|7z|rar|log|cache|secret|db|backup|yml|gz|config|csv|yaml|md|md5|exe|dll|bin|ini|bat|sh|tar|deb|rpm|iso|img|env|apk|msi|dmg|tmp|crt|pem|key|pub|asc)$", domain)

	counterChan := make(chan int)
	go func() {
		total := 0
		for c := range counterChan {
			total += c
			fmt.Printf("[%s] Processed URLs: %d\r", domain, total)
		}
	}()

	_ = fetchAndSave(domain, allURL, allFile, "All URLs", counterChan)
	_ = fetchAndSave(domain, filteredURL, filteredFile, "Filtered URLs", counterChan)

	close(counterChan)
}

func sendToTelegram(outputDir string) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Telegram bot error:", err)
		return
	}

	err = filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		doc := tgbotapi.FileReader{
			Name:   filepath.Base(path),
			Reader: file,
			Size:   info.Size(),
		}

		msg := tgbotapi.NewDocumentUpload(chatID, doc)
		_, err = bot.Send(msg)
		if err != nil {
			fmt.Println("Failed to send", path, "->", err)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking files:", err)
	}
}

func main() {
	var inputFile, domain, outputDir string
	var overrideID int64

	flag.StringVar(&inputFile, "file", "", "Path to domains.txt file")
	flag.StringVar(&domain, "domain", "", "Single domain to scan (e.g. example.com)")
	flag.StringVar(&outputDir, "output", "wayback_output", "Directory to save output")
	flag.StringVar(&extensionList, "ext", ".pdf,.sql,.json,.txt", "Comma-separated list of file extensions (currently unused)")
	flag.Int64Var(&overrideID, "id", 0, "Override Telegram chat ID")
	flag.Parse()

	if overrideID != 0 {
		chatID = overrideID
	}

	if !askVerificationCode(chatID) {
		fmt.Println("❌ Invalid code. Exiting.")
		return
	}

	timestamp := time.Now().Format("20060102_150405")
	outputPath := filepath.Join(outputDir, "wayback_results_"+timestamp)
	_ = os.MkdirAll(outputPath, os.ModePerm)

	var wg sync.WaitGroup
	if domain != "" {
		wg.Add(1)
		go processDomain(domain, timestamp, outputPath, extensionList, &wg)
	} else if inputFile != "" {
		file, err := os.Open(inputFile)
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			wg.Add(1)
			go processDomain(line, timestamp, outputPath, extensionList, &wg)
		}
	} else {
		fmt.Println("❌ Please provide either -domain or -file")
		return
	}

	wg.Wait()

	summaryPath := filepath.Join(outputPath, "summary.txt")
	summaryFile, _ := os.Create(summaryPath)
	defer summaryFile.Close()
	for _, line := range summary {
		_, _ = summaryFile.WriteString(line + "\n")
	}

	fmt.Printf("\n✅ Done! Sending results to Telegram...\n")
	sendToTelegram(outputPath)
	fmt.Println("✅ All results sent to Telegram.")
}
