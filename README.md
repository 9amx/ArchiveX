# ArchiveX
🚀 Advanced Wayback Machine URL  Scraper for Sensitive Artifacts
# ArchiveX

🚀 ArchiveX is a blazing-fast and stylish Wayback Machine URL scraper that digs deep into archived domains to uncover sensitive files like `.sql`, `.env`, `.bak`, and many more!

## Features
- ⚡ Multi-domain scraping via `.txt` file
- 🎯 Targeted extension filters for sensitive file types
- 🧠 Cool terminal UI with banner, progress bars, and spinners
- 💾 Results saved with timestamp for each domain

## Installation

```bash
git clone https://github.com/yourusername/ArchiveX.git
cd ArchiveX
go build -o archivex main.go

---

### ✅ 5. **Optional: Create a Release**

In GitHub:
- Go to "Releases" tab
- Click **Draft a new release**
- Tag: `v1.0.0`
- Title: `Initial Release`
- Add notes about what it does
- Upload compiled binaries (optional)

---

### ✅ 6. **Add a .gitignore (if not added)**

Create `.gitignore` in root:

```bash
# Go
bin/
*.exe
*.out
*.log
*.txt
*.bak

# OS files
.DS_Store
Thumbs.db
