# 🔔 BootyCall

> **Your app’s wake-up call on Linux!**  
> One command to turn any executable into a systemd service that auto-starts on boot.

---

## 💡 What is BootyCall?

`BootyCall` is a lightweight Go-powered CLI tool that helps you "daemonize" your executables.  
Forget writing boring `.service` files manually. Just call `BootyCall`, and you're ready to boot!

---

## ⚙️ Features

- 📦 Converts any executable into a `systemd` service
- 🔥 Automatically enables it to run at boot
- 🚀 Instantly starts the service
- 🧼 No messy configs — it's all handled for you

---

## 🚀 Getting Started

### 1. Build the binary

```bash
go build -o bootycall main.go
