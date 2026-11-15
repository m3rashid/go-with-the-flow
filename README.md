<div align="center">

# ⚡ Go with the Flow

**Event-driven HTTP server in pure Go — business logic as configurable workflows.**

[![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7)](https://gofiber.io)
[![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?logo=mongodb&logoColor=white)](https://mongodb.com)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

</div>

---

## ✨ What it does

MongoDB change streams act as triggers. Each **workflow** is a sequence of nodes whose handlers are JavaScript functions executed in a V8 engine, with built-in math helpers and user-defined globals. No redeploy needed to change logic — workflows live in MongoDB.

## 🧠 How it works

1. Modules register authenticated/anonymous routes on Fiber.
2. A goroutine watches configured MongoDB collections (`insert`/`update`/`delete`).
3. On a change, the matching `Workflow` (by `currentTrigger` + `collections`) is loaded; its current node's `Handler` runs in V8 against the changed document.
4. Nodes declare `NextNode`, `Collection`, `Trigger` — flows chain across collections.

## 🚀 Features

- ⚡ Event-driven architecture (MongoDB change streams)
- 🔧 Configurable workflows stored in DB
- 🟨 JS handlers in V8 with helpers (`add`, `multiply`, `divide`, `subtract`, `average`) + custom globals
- 🧩 Modular route registration (auth, search, flow, permission)
- 🔐 JWT auth, rate limiting, `/metrics` endpoint
- 🎨 React + Vite client

## 🛠 Tech

`Go` · `Fiber` · `MongoDB` · `JWT` · `V8 (rogchap/v8go)` · `React` · `Vite` · `Tailwind`

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/go-with-the-flow.git
cd go-with-the-flow
cp .env.example .env   # MONGODB_URI, MONGO_DB_NAME, SERVER_PORT, APP_NAME, JWT_SECRET
go run main.go

# client
cd client && pnpm install && pnpm dev
```

## 📄 License

MIT
