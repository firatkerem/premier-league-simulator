# Premier League Simulator

This project is a small **Go** application that lets you build and play through a football league from start to finish. One command launches the server, creates a fair fixture list, simulates each match, and stores the results in SQLite. The browser UI then updates in real time—showing the table, week-by-week scores, and each team’s title chances.

![Premier League Simulator screenshot](/screenshot.png)

> *Screenshot: After 3 match-weeks, Liverpool have a 87.5% chance of winning the title,
>  with Arsenal trailing behind them.*

---

## Key Features
| Category | What you get |
| -------- | ------------ |
| **Interactive UI** | Responsive interface with three actions: **Simulate Next Week**, **Simulate All Remaining**, **Reset League**. |
| **Live League Table** | Automatic calculation of **P · W · D · L · GF · GA · Pts** for every club after each simulation. |
| **Persisted State**  | Results, table and probabilities are stored in **SQLite** so you can stop and resume. |

---

## Tech Stack

| Layer         | Technology                        | Why?                               |
| ------------- | --------------------------------- | ---------------------------------- |
| **Backend**   | Go (Gin Framework) [https://go.dev] | Fast, simple, cloud-native. |
| **Database**  | SQLite  [https://www.sqlite.org] | Portable, serverless storage. |
| **Frontend**  | HTML, CSS, JavaScript (Bootstrap)    | Tiny bundle, instant load time. |
| **Tooling**   | `go mod`, `make`, Git   | Reproducible builds & CI. |

---

## Quick Start

```bash
# 1. Clone
git clone https://github.com/<your-user>/premier-league-simulator.git
cd premier-league-simulator

# 2. Install dependencies
go mod download

# 3. Run the server
go run cmd/server/main.go  # defaults to :8080

# 4. Open Web Interface
Open http://localhost:8080 in your browser and start simulating.
