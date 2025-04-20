<h1 align="center">📦 ArchiveX</h1>
<p align="center">
  <strong>The ultimate Wayback Machine scraper tool</strong> built in Go — fast, elegant, and Telegram-integrated.
</p>

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/9amx/ArchiveX?style=for-the-badge" />
  <img src="https://img.shields.io/github/license/9amx/ArchiveX?style=for-the-badge" />
  <img src="https://img.shields.io/github/stars/9amx/ArchiveX?style=for-the-badge" />
</p>

---

## 🚀 Features

- 🔍 **Scan URLs from the Wayback Machine** for single or multiple domains.
- 🧠 **Smart file filtering** (`.pdf`, `.sql`, `.json`, etc.)
- ⚙️ Auto-detects subdomains and saves results into structured folders.
- 📁 Outputs both **raw and filtered** URLs.
- 🕐 Shows real-time progress and estimated time remaining.
- 🤖 **Secure Start:** Requires a one-time access code from Telegram bot before scanning.
- 📤 Sends results directly to your Telegram via `@ehh_wayback_bot`.
- ⚡ Written in Go — minimal dependencies, blazing fast!

---

## 📥 Installation

> ✅ Requires Go 1.16+

```bash
go install github.com/9amx/ArchiveX@latest
This will download and install ArchiveX into your $GOPATH/bin or $HOME/go/bin.

💡 Usage
Start the bot @ehh_wayback_bot on Telegram first, it will send you a 6-digit code.

Then run the tool:

bash
Copy
Edit
# For a single domain
ArchiveX -domain example.com

# For a list of domains from a file
ArchiveX -file domains.txt

# Optional: specify Telegram Chat ID manually
ArchiveX -domain example.com -id 123456789

# Optional: custom output directory
ArchiveX -domain example.com -output results_folder
📦 Output Structure
css
Copy
Edit
wayback_output/
└── 20250420_123456/
    ├── example_com/
    │   ├── all.txt
    │   └── filtered.txt
    └── summary.txt
📲 Telegram Integration
✅ You will receive all .txt files (All + Filtered URLs) directly via @ehh_wayback_bot.

📩 If you wish to use another chat ID, use -id flag.

🔐 Access Control
To start a scan, the tool will request a one-time 6-digit code from the Telegram bot. This prevents unauthorized usage.

👤 Author
Made with by @the9am
GitHub: github.com/9amx

📄 License
MIT License

yaml
Copy
Edit

---

✅ You can copy this directly into `README.md` and commit:

```bash
echo "<paste content here>" > README.md
git add README.md
git commit -m "Added modern README with go install instructions"
git push
