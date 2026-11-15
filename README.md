# Go with the Flow

An event-driven HTTP server in pure Go where business logic is defined as **configurable workflows**. MongoDB change streams act as triggers; each workflow is a sequence of nodes whose handlers are JavaScript functions executed in a V8 engine (rogchap/v8go), with built-in math helpers and user-defined global functions.

## How it works

1. Modules register authenticated/anonymous routes on Fiber.
2. A background goroutine watches configured MongoDB collections via change streams (insert/update/delete).
3. On a change, the matching `Workflow` (by `currentTrigger` + `collections`) is loaded and its current node's `Handler` runs in V8 against the changed document.
4. Nodes declare `NextNode` (defaults to `currentNode + 1`) and `Collection`/`Trigger`, so flows can chain across collections.

## Features

- Event-driven architecture (MongoDB change streams)
- Configurable workflows stored in MongoDB, no redeploy needed
- JS handlers in V8 with default helpers (`add`, `multiply`, `divide`, `subtract`, `average`) + custom globals
- Modular route registration (auth, search, flow, permission modules)
- JWT auth, rate limiting, metrics endpoint (`/metrics`)
- React + Vite client (`client/`)

## Tech

Go · Fiber · MongoDB · JWT · V8 (rogchap/v8go) · React · Vite · Tailwind

## Quick start

```bash
git clone https://github.com/m3rashid/go-with-the-flow.git
cd go-with-the-flow
cp .env.example .env   # set MONGODB_URI, MONGO_DB_NAME, SERVER_PORT, APP_NAME, JWT_SECRET
go run main.go
# client
cd client && pnpm install && pnpm dev
```

## License

MIT
