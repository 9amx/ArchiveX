# 🚀 ArchiveX

![Go Version](https://img.shields.io/badge/go-1.23%2B-blue)  
🔎 **ArchiveX** is a blazing-fast, stylish CLI tool written in Go to scrape historical URLs from the Wayback Machine — built for OSINT, recon, and bug bounty workflows.

## ✨ Features

- 🔐 **Telegram Bot Verification** (6-digit code to start scanning)
- 🧠 **Smart Subdomain & Multi-domain Support**
- 🌀 **Real-time Spinner + Progress Counter**
- 🗃️ **Extension Filtering** (e.g., `.pdf`, `.sql`, `.json`, etc.)
- 📁 **Organized Output by Domain/Subdomain**
- 📬 **Automatic Telegram Delivery of Results**
- 🕰️ **Timestamped Folder & Summary File**
- 💥 Built with performance and clean UX in mind

---

## 🛠️ Installation

> Requires **Go 1.23+**

```bash
go install github.com/9amx/ArchiveX/cmd/archivex@latest
Make sure your $GOPATH/bin is in your PATH.

🧪 Usage
🔍 Scan a Single Domain
bash
Copy
Edit
archivex -domain example.com
📄 Scan Multiple Domains
bash
Copy
Edit
archivex -file domains.txt
🧾 Output Options
bash
Copy
Edit
- output     Output directory (default: "wayback_output")
- ext        File extensions to filter (currently unused)
- id         Override Telegram Chat ID
📦 Example
bash
Copy
Edit
archivex -file targets.txt -output outbox -id 1234567890
You’ll receive a secure code via Telegram bot @ehh_wayback_bot.

Enter the code in your terminal to proceed.

Results will be saved and auto-sent to Telegram.

📬 Telegram Bot Setup
Telegram Bot Token: Predefined in code

Chat ID: Can be overridden via -id flag

Make sure to start the bot on Telegram before scanning.

🧑‍💻 Author
👤 @the9am

🔗 github.com/9amx

📝 License
MIT

yaml
Copy
Edit


