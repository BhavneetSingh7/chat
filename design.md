Flow
1. ClientX connects to server and sends message in a Chat1.
2. Chat1 receives an event. Must emit to all its readers.
3. Chat1 must be persisted remotely and locally for offline messages.
4. ClientY must receive any message sent in Chat1.
5. All messages in Chat1 must be ordered by timestamp of receiving by a server.

Interconnecting protocol could be WebSocket, RPC or HTTP2.

3 core functionalities
1. How client sends messages?
Client connects using a protocol and delivers a message to server. Server decides what kind of message it is.

2. How server persists messages of a chat?
LIFO order in storage. What to do in case of racing?

3. How message received by 1 client is broadcasted to other clients?
Client must be polling/listening all the time?

# Roadmap

- [ ] User management to distinguish client and chats
- [ ] Client sending messages to server
- [ ] Client joining/initiating/ending chat(room)
- [ ] Server parsing various messages
- [ ] Server persisting in chats
- [ ] Server broadcasting received messages
- [ ] User segregation and management
- [ ] Abstracting away user mgmt
- [ ] Session management
- [ ] Reliable broadcasting via MQs?
- [ ] Multi-device same user - Synchronization

