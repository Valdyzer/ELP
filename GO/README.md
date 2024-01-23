# Matrix Multiplication Server and Client

This project implements a TCP server and client in Go for matrix multiplication. The server accepts connections from clients, receives large matrices, performs matrix multiplication, and sends back the result.

## Server Code

### `server.go`

The server code is organized as follows:

- `server.go`: Entry point of the server application.
- `Server` struct: Manages the server state and connections.
- `CreateServer`: Initializes a new server instance.
- `Start`: Initiates the server and starts listening for connections.
- `acceptConnection`: Accepts incoming client connections.
- `readLoop`: Reads and processes data from connected clients.
- Other utility functions for matrix operations.

## Client Code

### `client.go`

The client code is organized as follows:

- `client.go`: Entry point of the client application.
- `readLoop`: Reads and processes data from the connected server.
- `Byte_To_String`: Converts byte content to string.
- `startup`: Handles user input for matrix size and starts the communication with the server.
- Other utility functions for file operations and communication.

## Connecting Clients

Clients can connect to the server and send matrices for multiplication. The server processes the received matrices and sends back the result. To run properly, the client will ask you to choose from a list of 4 sqaure matrix files of sizes 1000x1000, 1500x1500, 2000x2000 and 3000x3000. We have provided these 4 files in the repo, they are randomly generated matrixes. To run the code as is, please download the files and specify in the client.go code the directory where these files are located.

## Dependencies

Both the server and client use standard Go libraries without external dependencies.

## Contributing

Contributions are welcome! If you find any issues or have suggestions, please open an issue or create a pull request.
