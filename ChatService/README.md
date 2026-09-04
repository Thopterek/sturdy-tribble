# ChatServie

ChatService is a Go WebSocket service for realt-time communication between players during a game.

## Features

- JWT authentication
- Game lobby managment
- Real-time lobby message broadcasting
- Duplicate connection protection
- Message validation
- WebSocket heartbeat using ping/pong
- Thread-safe client and lobby managment

## Requirements

- Go
- A shared JWT secret

## Run locally

```bash
SECRET=test-secret go run .
```

The service listens on port `8081`.

## Endpoints

### HealthCheck

```http
GET /ChatHealth
```

Response:

```text
ChatService ist erreichbar
```

### WebSocket

```text
ws://localhost:8081/ws
```
### WebSocket protocol

The client must authenticate imediatly after opening the WebSocket connection.

### Authentication

Client event:

```json
{
    "type": "auth",
    "token": "JWT_TOKEN"
}
```

Successful server response:

```json
{
    "type": "auth_success",
    "user_id": "USER_UUID"
}
```

The autentication event must be sent within 10 seconds. Invalid or expired tokens cause the connection to be closed.

The JWT must be signed with the shared secret using HS256 and must contain the following claims:

```json
{
    "iss": "chirpy",
    "sub": "USER_UUID",
    "lobby_id": "LOBBY_UUID",
    "exp": 1788278400
}
```

- `sub` identifies the authenticated user.
- `lobby_id` identifies the lobby the user is allowed to join.
- `exp` defiones when the token expires.
- The requested lobby ID must match the lobby ID contained in the token.

### Join a lobby

Client event:

```json
{
    "type": "join_lobby",
    "lobby_id": "LOBBY_UUID"
}
```

Successful server response:

```json
{
    "type": "lobby_joined",
    "lobby_id": "LOBBY_UUID"
}
```

A client can only join one lobby per WebSocket connection.

### Send a message

Client event:

```json
{
    "type": "message",
    "lobby_id": "LOBBY_UUID",
    "content": "Hello!"
}
```

the sender receives:

```json
{
    "id": "MESSAGE_UUID",
    "type": "message_sent",
    "sender_id": "USER_UUID",
    "lobby_id": "LOBBY_UUID",
    "content": "Hello!",
    "created_at" : "2026-09-01T10:00:00Z"
}
```

Other connected lobby members recieve:

```json
{
    "id": "MESSAGE_UUID",
    "type": "message",
    "sender_id": "USER_UUID",
    "lobby_id": "LOBBY_UUID",
    "content": "Hello!",
    "created_at": "2026-09-01T10:00:00Z"
}
```

Messages must not be empty and may contain a maximum of 1000 bytes.

### Error response

```json
{
    "type": "error",
    "message": "Error description"
}
```