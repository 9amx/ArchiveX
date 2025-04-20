# 🗃️ ArchiveX

![Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go&logoColor=white)  
![Status](https://img.shields.io/badge/status-active-brightgreen)  
![License](https://img.shields.io/github/license/9amx/ArchiveX)  
![Telegram](https://img.shields.io/badge/Telegram-Bot-blue?logo=telegram)  

> **ArchiveX** is an advanced Wayback Machine scraping tool written in Go, capable of extracting historical URLs of domains/subdomains with filtering, live progress updates, and **Telegram bot integration**.

---

## 🚀 Features

- 🔎 Supports both **single domain** and **bulk domain scanning**
- 🧠 Smart filtering of sensitive filetypes (`.pdf`, `.sql`, `.json`, `.txt`, and more)
- 📦 Results organized by domain/subdomain and timestamp
- 🛡️ Protected by Telegram verification (code required before scan)
- 📤 Automatically sends results to Telegram bot `[@ehh_wayback_bot](https://t.me/ehh_wayback_bot)`
- 🧾 Generates detailed summary report per scan
- ✨ Modern CLI with spinners, progress counter, and clean formatting

---

## 🧑‍💻 Author

Made with 💀 by [@the9am](https://t.me/the9am)  
Tool: `ArchiveX`  
Bot: [@ehh_wayback_bot](https://t.me/ehh_wayback_bot)

---

## 📦 Installation

### Clone & Build

```bash
git clone https://github.com/9amx/ArchiveX.git
cd ArchiveX
go build -o wayback
📲 Telegram Verification
Before each scan, ArchiveX requests a 6-digit verification code from the Telegram bot.

🔐 How it works:
Start the bot → @ehh_wayback_bot

You will receive a 6-digit access code

Run ArchiveX — it will ask you for this code before scanning starts

⚙️ Usage
Scan a single domain
bash
Copy
Edit
./wayback -domain example.com
Scan multiple domains from file
bash
Copy
Edit
./wayback -file domains.txt
Override Telegram Chat ID (optional)
bash
Copy
Edit
./wayback -domain example.com -id 1234567890
📁 Output Structure
css
Copy
Edit
wayback_output/
└── 20250420_123456/
    ├── example_com/
    │   ├── all.txt           # All captured URLs
    │   ├── filtered.txt      # Filtered URLs (sensitive extensions)
    └── summary.txt           # Summary report
🔧 Flags

Flag	Description
-domain	Scan a single domain
-file	File containing list of domains
-output	Output directory (default: wayback_output)
-ext	Custom file extensions to filter (default used for now)
-id	Override Telegram chat ID manually
💡 Example domains.txt
pgsql
Copy
Edit
example.com
admin.targetsite.org
shop.company.xyz
🛠️ To-Do / Roadmap
 Interactive UI / GUI version

 Export as JSON / CSV

 Real-time result streaming to Telegram

 Historical snapshot downloader

📝 License
This project is open-source and available under the MIT License.

⭐️ Support & Community
Found it useful?
Give a ⭐️ on GitHub and share with others.
Join @the9am for updates or suggestions!

yaml
Copy
Edit

---

### ✅ Want this saved directly as `README.md` in your repo?

Let me know, and I’ll generate & write it directly for you.





