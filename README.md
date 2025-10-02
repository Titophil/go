
# 📄 README.md

````markdown
# Prompt-Powered Kickstart — Getting Started with Go (Golang)

This repository is part of the **Moringa AI Capstone Project**.  
It demonstrates how to use **Go (Golang)** to create a minimal HTTP JSON API, guided by AI prompts.

---

## 🎯 Objective
- Learn the basics of Go (Golang) using generative AI prompts.  
- Produce a runnable HTTP "Hello World" API in Go.  
- Document clear setup steps so beginners can reproduce it.  

**End Goal:** A tiny HTTP server exposing `/hello` that returns JSON:  
```json
{"message":"Hello, world from Go!"}
````

---

## ⚡ Quick Summary of Go

Go (or Golang) is an open-source, compiled language designed by Google.
It’s known for:

* **Speed & Simplicity** → minimal setup, fast execution.
* **Concurrency** → built-in support for goroutines.
* **Reliability** → widely used in production systems.

**Where it’s used:**
Cloud-native tools, microservices, APIs, CLI tools.

**Example in the wild:** Kubernetes (container orchestration system) is written in Go.

---

## 🖥️ System Requirements

* **OS:** Linux, macOS, or Windows
* **Go:** Version 1.18+
* **Editor:** VS Code or any text editor
* **Tools:** `curl` for testing

Optional: Docker, ngrok (for demos)

---

## ⚙️ Installation & Setup

### 1. Install Go

Follow official guide: [https://go.dev/dl/](https://go.dev/dl/)

Check version:

```bash
go version
```

### 2. Clone this repository

```bash
git clone https://github.com/<your-username>/go-hello.git
cd go-hello
```

### 3. Initialize module (if not already)

```bash
go mod init github.com/<your-username>/go-hello
```

### 4. Run the server

```bash
# Development mode
go run main.go

# Or build binary
go build -o hello
./hello
```

Expected terminal output:

```
Listening on :8080
```

### 5. Test the API

```bash
curl http://localhost:8080/hello
```

Expected response:

```json
{"message":"Hello, world from Go!"}
```

---

## 📂 Project Structure

```
go-hello/
 ├── main.go        # HTTP server code
 ├── go.mod         # Go module file
 ├── README.md      # Project documentation
 └── .gitignore     # Ignore build artifacts
```

---

## 🛠️ Troubleshooting

| Issue                      | Cause                           | Fix                                          |
| -------------------------- | ------------------------------- | -------------------------------------------- |
| `go: not found`            | Go not installed / PATH missing | Install Go & add `/usr/local/go/bin` to PATH |
| `module declares its path` | Wrong `go.mod` path             | Edit `go.mod` or re-run `go mod init`        |
| `address already in use`   | Port 8080 in use                | Kill process or change `PORT` env var        |
| JSON not returned          | Missing `Content-Type`          | Ensure `application/json` is set in handler  |

---

## 🤖 AI Prompt Journal

This project was scaffolded with the help of generative AI prompts.
Example prompts used:

* *“Show me a step-by-step scaffold to create a minimal Go HTTP server project that responds on /hello with JSON.”*
* *“Explain common causes of ‘module declares its path’ errors in Go.”*

Reflections and full prompt list are in `prompts.md`.

---

## 📚 References

* [Go Official Docs](https://go.dev/doc/)
* [Go by Example](https://gobyexample.com/)
* [net/http package](https://pkg.go.dev/net/http)

---

## ✅ License

MIT License – free to use and share.

```

---

👉 Next step: I can also generate the `go.mod`, `.gitignore`, and `prompts.md` files for you so you have a **ready-to-push repo**.  

Do you want me to package all those files (`main.go`, `go.mod`, `README.md`, `.gitignore`, `prompts.md`) into one bundle so you can just upload to GitHub or zip them for submission?
```
