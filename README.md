# Matchmaking System

This project implements a matchmaking system for players in a game. The system includes functionalities for players to join a matchmaking queue, creating a competition, and handle players who have been in the queue for too long.

## Features

- **Join Matchmaking**: Players can join the matchmaking queue.
- **Get Competitions**: The system fetches competitions based on player levels and countries.
- **Handle Expired Players**: Players who have been in the queue for too long are handled separately.

## Usage

1. Run the matchmaking server:
    ```sh
    go run main.go
    ```

2. Use Postman or any other HTTP client to interact with the server.

### API Endpoints

- **Join Matchmaking**
    - **URL**: `/join`
    - **Method**: `POST`
    - **Description**: Adds a player to the matchmaking queue.

### Code Structure

- **main.go**: Entry point of the application.
- **matchmaker/matchmaker.go**: Contains the `Matchmaker` struct and its methods.
- **matchmaker/player.go**: Contains the `Player` struct and its methods.
- **director/director.go**: Contains the `Director` coordinator between matchmaking and other game services.
