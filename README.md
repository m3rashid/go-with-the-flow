<div align="center">

# ⚡ Go with the Flow

### Event-driven HTTP server in pure Go — business logic as configurable workflows

[![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev)
[![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge)](https://gofiber.io)
[![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)](https://mongodb.com)
[![React](https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react&logoColor=black)](https://react.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge)](LICENSE)

> *Define your business logic as JSON-configurable workflows. No redeploy. No code changes. Just data.*

</div>

---

## 📖 Overview

**Go with the Flow** is an event-driven HTTP server where business logic is expressed as **workflows** — sequences of nodes stored in MongoDB. A change stream watches configured collections; on insert/update/delete, the matching workflow is loaded and its current node's handler (a JavaScript function) runs inside a V8 engine against the changed document.

This lets you change application behavior by editing workflow documents — no redeploy, no recompile.

## ✨ Features

- ⚡ **Event-driven architecture** — MongoDB change streams trigger workflows
- 🔧 **Configurable workflows** — stored in DB, editable at runtime
- 🟨 **JS handlers in V8** — `rogchap/v8go` executes node handlers with built-in helpers (`add`, `multiply`, `divide`, `subtract`, `average`) plus user-defined globals
- 🧩 **Modular route registration** — auth, search, flow, permission modules plug in
- 🔐 **JWT authentication** with role/permission levels
- 🚦 **Rate limiting** (sliding window, 60 req/min)
- 📊 **Metrics endpoint** (`/metrics`) via Fiber monitor
- 🎨 **React + Vite + Tailwind client** with Recoil state
- 📝 **Audit fields** — `createdBy`/`updatedBy` on every document

## 🏗 Architecture

```
┌────────────┐   change stream   ┌──────────────┐
│  MongoDB   │ ────────────────► │  Flow engine │
│ collections│                   │ (V8 handlers)│
└────────────┘                   └──────┬───────┘
       ▲                                │
       │ writes                         │ triggers
       │                                ▼
┌──────┴──────┐                  ┌──────────────┐
│   Modules   │◄─────────────────│  Workflows   │
│ (auth,flow, │   register       │ (nodes, trig)│
│  search,...)│   routes         └──────────────┘
└──────┬──────┘
       │
       ▼
┌──────────────┐
│ Fiber server │  HTTP API + static + metrics
└──────────────┘
```

### Workflow model

```go
type Workflow struct {
    Nodes           []Node   // sequence of handlers
    CurrentNode     int      // active index
    Collections     []string // watched collections
    CurrentTrigger  string   // insert/update/delete/replace
    GlobalFunctions string   // user JS, injected into V8
}

type Node struct {
    Name       string  // identifier
    Collection string  // target collection
    Trigger    string  // insert/update/delete/replace
    Handler    string  // JS function body
    NextNode   string  // defaults to currentNode + 1
}
```

## 📁 Project structure

```
go-with-the-flow/
├── main.go              # entrypoint, wires modules + change streams
├── modules/
│   ├── main.go          # Module/Controller types, RegisterRoutes
│   ├── auth/            # auth module (users, profiles, JWT)
│   ├── flow/            # workflow engine (schema, run, controller)
│   ├── permission/      # permission module
│   └── search/          # search/resource module
├── db/                  # base schemas, audit helpers
├── controller/          # generic list controller
├── utils/               # auth middleware, errors
├── client/              # React + Vite + Tailwind frontend
└── public/              # static assets
```

## 🛠 Tech stack

| Layer | Technology |
|---|---|
| Backend | Go 1.21, Fiber v2, MongoDB driver |
| Scripting | V8 via `rogchap/v8go` |
| Auth | JWT (golang-jwt), bcrypt |
| Validation | go-playground/validator |
| Frontend | React 18, Vite, Tailwind, Recoil, HeadlessUI |
| Infra | Docker-ready, env-based config |

## 📦 Quick start

```bash
git clone https://github.com/m3rashid/go-with-the-flow.git
cd go-with-the-flow

# configure
cp .env.example .env   # MONGODB_URI, MONGO_DB_NAME, SERVER_PORT, APP_NAME, JWT_SECRET

# run backend
go run main.go

# run frontend
cd client && pnpm install && pnpm dev
```

## 🗺 Roadmap

- [ ] Persist workflow execution state across restarts
- [ ] Visual workflow editor in the client
- [ ] Support for replace/delete change-stream events end-to-end
- [ ] Webhook nodes for external integrations

## 🤝 Contributing

PRs welcome. Please open an issue first for major changes.

## 📄 License

MIT © [m3rashid](https://github.com/m3rashid)
