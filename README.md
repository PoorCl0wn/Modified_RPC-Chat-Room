# 💬 Real-Time Concurrent Chat System (GoLang)

A fully concurrent **real-time chatroom system** built using Go's standard `net` package.  
The application consists of **TCP server + multiple clients**, with instant message broadcasting using goroutines, channels, and thread-safe shared state.

This project is an upgraded version of a previous RPC-based system — now redesigned for **true live messaging** without history polling.

---

## 🚀 Features

### 🧵 Full Concurrency
- Each client is handled in its own **goroutine**.
- Real-time broadcasting via a central **channel**.

### 📡 Instant Message Broadcast
- Messages sent by a client are instantly delivered to **all other users**.
- No echo to the sender ("no self-feedback").

### 👥 Join/Leave Notifications
- On joining: all clients receive  
  **`User [ID] joined the chat`**
- On leaving:  
  **`User [ID] left the chat`**

### 🔒 Safe Shared State
- Connected clients are stored in a shared map.
- Protected with `sync.Mutex` to avoid race conditions.

### 🎨 Clean Color-Coded UI (Client)
- System messages (join/leave) → blue  
- Other users' messages → yellow  
- Your own messages → green  

---

## 🧱 Project Structure

