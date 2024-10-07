# NetCat

This project recreates the functionality of the NetCat (`nc`) system command in a **Server-Client Architecture** using Go. 

The program runs in server mode to listen for incoming connections or in client mode to connect to a server and transmit messages. 

The end goal is to implement a **group chat** system that mimics NetCat’s behavior while adding extra functionality such as named clients, message timestamps, and client join/leave notifications.

## About NetCat

NetCat (`nc`) is a command-line utility used for network communication over TCP, UDP, or UNIX-domain sockets. 

It is often referred to as the `Swiss army knife` of networking tools because it can be used for various tasks such as opening TCP connections, sending UDP packets, and listening on arbitrary ports.

For more details, refer to the NetCat manual: `man nc`.

## Project Features

1. **TCP Connections**: The server supports multiple clients through TCP connections.
2. **Client Naming**: Clients must provide a non-empty name to join the chat.
3. **Message Broadcast**: Clients can send messages to the group, identified by their name and the timestamp.
   
   Example message format: `[YYYY-MM-DD HH:MM:SS][client.name]:[client.message]`
   
   ```bash
   hamza has joined the chat...
   [2024-10-07 14:51:59][hamza]: hello guys
   ```
4. **Chat History**: New clients receive the full message history when they join.
5. **Client Join/Leave Notifications**: All clients are notified when a new client joins or a client leaves the chat.
   
   ```bash
   [ENTER YOUR NAME]: hamza
   hello guys
   aaron has left the chat...
   ```
6. **Connection Limit**: The server restricts the maximum number of concurrent connections to 10.
7. **Empty Messages**: Empty messages from clients are not broadcasted.
8. **Default Port**: If no port is specified, the server listens on port 8989.
9. **Graceful Exit**: If a client leaves, the remaining clients continue to operate without interruption.

## Usage

The server and client interaction mimics NetCat but adds features for a better group chat experience. Here's how to use the program:

## Server

**Clone the project**

```bash
git clone https://learn.zone01kisumu.ke/git/abrakingoo/net-cat.git
cd net-cat
```

Start the server using the command:

```bash
go run .
```

By default the server with run on port `8989` when no port is provided.

If wrong port is provided or other arguments, the program will display a usage message:

```bash
go run . localhost
[USAGE]: ./TCPChat $port
```

### Client

Connect to the server using NetCat or any client:

```bash
nc <IP> <port>
```

For example

```bash
nc localhost 8989
```

The server will greet you with a welcome message and a Linux logo. You'll be prompted to enter your name:

```bash
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: 
```

Once connected, you can start chatting. Messages will be broadcast to all connected clients.

When a client joins or leaves, the system will notify all connected clients.

## Error Handling

- The server gracefully handles client-side and server-side errors.
- If an error occurs during a client connection or message transmission, an appropriate error message is displayed.

## Project Structure

- The server uses **Go-routines** and **channels** to handle multiple clients concurrently.
- **Mutexes** are used for synchronization to avoid race conditions during message broadcasting and client management.
- The maximum number of concurrent connections is limited to 10.

## Unit Testing

- Unit tests are provided to test both the server and client functionalities, ensuring robustness in various scenarios.
  
  ## License
  
  This project is open-source and available under the MIT License.
