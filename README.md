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
RPC CHAT ROOM/
│
├── client.go      # Client-side program
├── server.go      # RPC server
└── README.md      # Project documentation

---

## ⚙️ How It Works

### 🔧 1. Server Logic
- Listens on port **5000**
- Accepts clients and asks them to send their **ID**
- Stores each client connection in a global map
- Every message is sent into a `broadcast` channel
- A separate goroutine (`broadcaster`) forwards each message to all other clients

### 🔧 2. Client Logic
- Connects to server
- Sends its **ID**
- Receives:
  - Current active users
  - Join/leave notifications
  - Incoming messages
- Sends typed messages instantly

---

## ▶️ How to Run

### 1️⃣ Start the Server

Open a terminal inside the project folder:

```bash
go run server.go
```
Server will start on port `:5000` and print incoming messages.

### 2️⃣ Run the Client

open another terminal (many as you want):

```bash
go run client.go
```

* Enter your **name** when prompted
* Type messages — they’ll be sent to the server
* Type `exit` to leave

## 🖼 Example Output
**Server Terminal:**
```bash
Server running on port 5000...
[SERVER] User 'Ahmed' joined the chat
[SERVER] User 'Omar' joined the chat
[SERVER] User 'Ahmed' left the chat
```

**Client Terminal:**
```bash
Enter your ID: Ahmed
** Current users in chat: Omar, Sarah **

[you] > Hello!
[Omar]: Hi Ahmed!
[Sarah]: Welcome Ahmed!
```
## 🧩 Technologies Used

* [GoLang](https://go.dev/)
* Goroutines & Channels
* Mutex-based concurrency control
* ANSI color-coded terminal output

## 📜 Instructions Applied in This Version

* ✔ Convert RPC system to real-time broadcasting
* ✔ Notify all clients when a user joins
* ✔ Notify all clients when a user leaves
* ✔ Broadcast messages live using goroutines + channels
* ✔ Prevent sender from receiving their own message
* ✔ Use Mutex to safely manage shared clients list
* ✔ No polling, no history — live chat onl

## 👤 Author

**Ahmed Elsafty**
📧 [elsaftyahmed09@gmail.com]

⭐ Feel free to fork, improve, and star this repo!
