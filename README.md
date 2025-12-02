💬 RPC Chat Room (GoLang Project)
A simple chatroom application built using Go's net/rpc package. Clients can connect to a central RPC server, send messages, and fetch full chat history.

🚀 Features
📡 Client–Server communication using Go RPC
💾 Server stores all messages in memory
👥 Multiple clients can send messages (chat history shared)
🔁 Each message returns the full chat history
🧹 Simple and clean console interface
🧱 Project Structure
RPC CHAT ROOM/
│
├── client.go      # Client-side program
├── client+.go     # Improved version with live updates
├── server.go      # RPC server
├── go.mod         # Go module
└── README.md      # Project documentation
⚙️ How to Run
1️⃣ Run the Server
Open a terminal inside the project folder:

go run server.go
Server will start on port :1234 and print incoming messages.

2️⃣ Run the Client
open another terminal (many as you want):

go run client.go
Enter your name when prompted
Type messages — they’ll be sent to the server
The entire chat history will display after each message
Type exit to leave
🧩 client+.go (Improved Version)
A more advanced client version that adds real-time chat updates and a cleaner interface while staying fully compatible with the same server.

✨ Added Features
🔄 Live Updates – messages from other clients appear automatically in real-time without waiting to type.
💬 Cleaner UI – removed redundant prefixes like “You:” and formatted the output neatly.
🧠 Smart Message Filtering – your own messages no longer echo back; you only see new ones from others.
⚙️ Efficient Refresh – only new messages print to the screen instead of reloading the full chat history each second.
⏱️ Non-blocking Updates – the client runs a background process to fetch messages continuously while you type.
▶️ Run the improved client:
go run "client+.go"
🖼 Example Output
Server terminal:

Chat server running on port 1234...
Ahmed: Hello!
Omar: Hi Ahmed, how are you?
Client terminal (Ahmed):

> Hello!
Client terminal (Omar):

Ahmed: Hello!
> Hi Ahmed, how are you?
🧩 Technologies Used
GoLang
net/rpc package
bufio, fmt, log, strings, sync, time
📜 License
This project is licensed under the MIT License – see the LICENSE file for details.

👤 Author
Ahmed Elsafty 📧 [[elsaftyahmed09@gmail.com]

⭐ Feel free to fork, improve, and star this repo!
