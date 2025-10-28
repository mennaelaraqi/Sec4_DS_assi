Link for the Recording

https://drive.google.com/file/d/1KFB9gikG6_fLmllms59ea4CFRkPvP5n7/view?usp=sharing


# Assignment 04: Simple RPC Chatroom

This project is a simple implementation of a multi-client chatroom.

---

## Components

The project consists of three main files:

* **`server.go`**: The main server application. It is responsible for:
    * Listening for incoming $TCP$ connections on `port 1234`.
    * Registering the `ChatService` for remote procedures.
    * Maintaining a shared chat history (a `slice` of strings) for all clients.
    * Using a `sync.Mutex` to protect the chat history from concurrent access (race conditions) when multiple clients send messages simultaneously.

* **`client.go`**: The client application. It is responsible for:
    * Dialing the server's $TCP$ address.
    * Prompting the user for their name upon startup.
    * Reading full-line messages from the user (handling spaces).
    * Calling the `ChatService.SendMessage` remote procedure, passing the user's name and message.
    * Receiving the complete chat history as a reply and printing it to the console.

* **`shared.go`**: The common "contract" file. It defines the `MessageArgs` struct, which is used by both the client (for marshaling) and the server (for unmarshaling) to pass data. This ensures both parties agree on the data structure.

---
