# pulsemon

A service health checker built from scratch in Go. Give it a list of URLs, and it pings all of them every 30 seconds concurrently. If something goes down, you get a Telegram alert immediately.

Think baby Datadog. No frameworks, no external dependencies, just Go stdlib.

---

## What it does

- Reads a list of services from a JSON config file
- Pings every service simultaneously using goroutines
- Logs response time and status for every check
- Sends a Telegram alert the moment a service goes down
- Repeats every 30 seconds forever

---

## How it works

```
config.json (list of services)
      ↓
main.go fires one goroutine per service
      ↓
checker pings the URL, measures response time
      ↓
results collected through a buffered channel
      ↓
logger prints [UP] / [DOWN] with response time
      ↓
notifier sends Telegram alert for anything down
      ↓
waits 30 seconds, repeats
```

---

## Project structure

```
pulsemon/
├── config/
│   ├── config.go   → reads services from config.json
│   └── env.go      → loads .env file (zero dependencies)
├── checker/
│   └── checker.go  → pings one URL, returns result with response time
├── notifier/
│   └── notify.go   → sends Telegram alert when a service is down
├── logger/
│   └── logger.go   → formats and prints check results
└── main.go         → wires everything together
```

---

## Setup

**1. Clone the repo**
```bash
git clone https://github.com/SHAIK14/pulsemon
cd pulsemon
```

**2. Create your config file**

Edit `config.json` with the services you want to monitor:
```json
{
  "services": [
    {"name": "google", "url": "https://google.com"},
    {"name": "github", "url": "https://github.com"},
    {"name": "your-api", "url": "https://your-api.com/health"}
  ]
}
```

**3. Set up Telegram alerts**

- Create a bot via [@BotFather](https://t.me/botfather) on Telegram
- Get your chat ID by messaging your bot and hitting `https://api.telegram.org/bot{TOKEN}/getUpdates`

Create a `.env` file in the project root:
```
TELEGRAM_TOKEN=your_bot_token
TELEGRAM_CHAT_ID=your_chat_id
```

**4. Run it**
```bash
go run .
```

---

## Sample output

```
[UP]   google        200  423ms
[UP]   github        200  141ms
[DOWN] your-api           no such host
```

And on your phone:
```
🚨 Down: your-api
URL: https://your-api.com/health
Error: no such host
```

---

## What I learned building this

- Goroutines and buffered channels for concurrent HTTP checks
- Why you pass loop variables as goroutine arguments (loop capture bug)
- `http.Client` with timeout — non-negotiable in production
- `time.Ticker` for scheduled recurring work
- Separating concerns across packages (config, checker, notifier, logger)
- Loading `.env` files without any external library
- HTTP POST to an external API (Telegram Bot API)
