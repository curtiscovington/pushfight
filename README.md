# Pushfight

A Go implementation of the board game [PushFight](http://pushfightgame.com/rules.htm). This project exposes a simple WebSocket based server and a small browser client for playing the game.

## Requirements

- Go 1.20 or newer

## Getting started

Clone the repository and run the server from the project root:

```bash
go run ./...
```

The server listens on **localhost:3000**. Visiting that address will serve `index.html`, which contains a minimal interface for the game. Open the page in two separate browser windows to start a two player match.

## Repository layout

- `engine/`  – core game logic for board state, pieces and move validation
- `server/`  – WebSocket server coordinating two players
- `index.html` – browser UI communicating with the server
- `main.go` – entry point that starts the HTTP and WebSocket server

## Gameplay

After two clients connect, each is assigned a colour. Players alternately move or push pieces by clicking on them and the target square. Pushing may knock pieces off the board, resulting in a win for the opposing player.

Use the game log beside the board to follow moves and pushes during play.

## License

This repository is provided as-is without any specific licence.
