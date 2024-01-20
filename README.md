# Matrix Multiplication Server

This project implements a TCP server in Go for matrix multiplication. The server accepts connections from clients, receives large matrices, performs matrix multiplication, and sends back the result.

## Features

- Accepts connections from clients.
- Efficiently handles large matrices for multiplication.
- Calculates the product of the received matrix with itself.
- Provides timing information for the server-side calculation.

## Server Code

The server code is organized as follows:

- `server.go`: Entry point of the server application.
- `Server` struct: Manages the server state and connections.
- `CreateServer`: Initializes a new server instance.
- `Start`: Initiates the server and starts listening for connections.
- `acceptConnection`: Accepts incoming client connections.
- `readLoop`: Reads and processes data from connected clients.
- Other utility functions for matrix operations.

## Running the Server

To run the server, execute the following commands:

```bash
go build main.go
./main
