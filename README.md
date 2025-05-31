# 🥋 Rolling Trends

**AI-powered insights on trending Brazilian Jiu-Jitsu techniques and strategies, sourced from Reddit discussions.**  
Rolling Trends scrapes threads from subreddits like `r/bjj` and `r/grappling`, uses AI to identify the most talked-about techniques, and organizes them into guard styles, takedowns, submissions, sweeps, and more — all displayed in a clean, searchable Nuxt frontend.

---

## 🧠 What It Does

- Scrapes Reddit threads about BJJ (guard, takedowns, passing, etc.)
- Uses OpenAI to extract and categorize community-recommended techniques
- Displays current “meta” trends in BJJ across categories
- Highlights why each technique is popular, all based on real user feedback

---

## 🧩 Project Structure

This app includes three key components:

1. **Go Scraper**  
   Pulls Reddit thread titles and top comments from `r/bjj` and `r/grappling`.

2. **AI Analyzer (OpenAI)**  
   Sends scraped data to OpenAI to extract technique names, categories, and reasoning.

3. **Nuxt Frontend**  
   Displays categorized, ranked trends with summaries and filters.

---

## 🛠️ Tech Stack

| Layer        | Tech                 |
|--------------|----------------------|
| Frontend     | Nuxt 3, TailwindCSS  |
| Backend API  | Go, net/http, goquery |
| AI Analysis  | OpenAI API (`gpt-4`) |
| (Optional) DB| MongoDB or PostgreSQL |

---

## 🚀 Getting Started

### Clone the Repo

```bash
git clone https://github.com/yourusername/rolling-trends.git
cd rolling-trends
```

### Backend (Go)
```Bash
cd backend
go mod tidy
go run main.go
```

### Frontend (Nuxt)
```Bash
cd frontend
npm install
npm run dev
```

---

## 📁 Directory Structure
```text
rolling-trends/
├── backend/           # Go API & scraper logic
│   ├── scraper/       # Reddit scraping logic
│   ├── analyzer/      # OpenAI integration
│   ├── handlers/      # API route handlers
│   └── main.go        # API server entry point
├── frontend/          # Nuxt 3 UI
│   ├── pages/         # e.g. /trends/passing
│   ├── components/    # TrendCard.vue, Filters, etc.
│   └── assets/        # Tailwind styling
├── data/              # (Optional) raw + processed JSON for dev/testing
│   ├── raw/
│   └── processed/
├── .env               # API keys (not committed)
└── README.md
```
---

## 🔮 Roadmap

✅ Phase 1: Proof of Concept
	•	Choose core categories (Guard, Takedowns, Passing, Sweeps, Submissions)
	•	Scrape a few sample Reddit threads manually or with Go
	•	Generate mock JSON data for frontend display
	•	Build Nuxt frontend pages for trends
	•	Display AI-generated summaries per technique

🚧 Phase 2: Full Integration
	•	Automate Reddit scraping by keyword/category
	•	Use OpenAI API to process and categorize techniques
	•	Enable filtering by Gi/No-Gi or beginner/advanced
	•	Style and deploy frontend

🔜 Phase 3: Backend & Persistence
	•	Store data with timestamps in a database
	•	Add cron jobs or scraping triggers
	•	(Optional) Build an admin dashboard

