💬 Real-Time Concurrent Chat System (GoLang)

A fully concurrent real-time chatroom system built using Go's standard net package.
The application consists of TCP server + multiple clients, with instant message broadcasting using goroutines, channels, and thread-safe shared state.

This project is an upgraded version of a previous RPC-based system — now redesigned for true live messaging without history polling.

🚀 Features
🧵 Full Concurrency

Each client is handled in its own goroutine.

Real-time broadcasting via a central channel.

📡 Instant Message Broadcast

Messages sent by a client are instantly delivered to all other users.

No echo to the sender ("no self-feedback").

👥 Join/Leave Notifications

On joining: all clients receive
User [ID] joined the chat

On leaving:
User [ID] left the chat

🔒 Safe Shared State

Connected clients are stored in a shared map.

Protected with sync.Mutex to avoid race conditions.

🎨 Clean Color-Coded UI (Client)

System messages (join/leave) → blue

Other users' messages → yellow

Your own messages → green

🧱 Project Structure
REALTIME-CHAT/
│
├── server.go       # TCP chat server
├── client.go       # Interactive chat client
├── go.mod          # Go module file
└── README.md       # Project documentation

⚙️ How It Works
🔧 1. Server Logic

Listens on port 5000

Accepts clients and asks them to send their ID

Stores each client connection in a global map

Every message is sent into a broadcast channel

A separate goroutine (broadcaster) forwards each message to all other clients

🔧 2. Client Logic

Connects to server

Sends its ID

Receives:

Current active users

Join/leave notifications

Incoming messages from all clients

Sends typed messages instantly

▶️ How to Run
1️⃣ Start the Server
go run server.go


You should see:

Server running on port 5000...

2️⃣ Start a Client

Open another terminal:

go run client.go


Enter your ID:

Enter your ID: Ahmed


Start chatting:

[you] > Hello everyone!


Open more terminals for more clients — all will receive messages instantly.

🖼 Example Output
Server Terminal
Server running on port 5000...
[SERVER] User 'Ahmed' joined the chat
[SERVER] User 'Omar' joined the chat
[SERVER] User 'Ahmed' left the chat

Client Terminal
Enter your ID: Ahmed
** Current users in chat: Omar, Sarah **

[you] > Hello!
[Omar]: Hi Ahmed!
[Sarah]: Welcome Ahmed!

🧩 Technologies Used

Go (net, bufio, sync)

Goroutines & Channels

Mutex-based concurrency control

ANSI color-coded terminal output

📜 Instructions Applied in This Version

✔ Convert RPC system to real-time broadcasting
✔ Notify all clients when a user joins
✔ Notify all clients when a user leaves
✔ Broadcast messages live using goroutines + channels
✔ Prevent sender from receiving their own message
✔ Use Mutex to safely manage shared clients list
✔ No polling, no history — live chat only

👤 Author

Ahmed Elsafty
📧 elsaftyahmed09@gmail.com

⭐ Feel free to fork, improve, and star the repo!
