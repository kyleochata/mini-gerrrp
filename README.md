# gRPC Demo Project

This project demonstrates a simple gRPC client-server application. The project defines a basic service using Protocol Buffers (`.proto` file), where the client sends a request to the server, and the server responds with the appropriate message.

## Prerequisites

To run this project, you will need:

- [Go](https://golang.org/dl/) (if you're using Go for the implementation)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- gRPC and Protobuf plugins for Go (or your chosen language):
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

## Generating gRPC Code from Proto File

Once you have defined your service in the `service.proto` file, generate the necessary Go files by running:

```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

This will generate the following files:

`service.pb.go`: Contains the message types defined in the .proto file.
`service_grpc.pb.go`: Contains the gRPC-specific code for the service.

## Running the Application

Start the gRPC server:

In the server directory, run the server:

```bash

go run main.go
```

Run the gRPC client:

In the client directory, run the client to interact with the server:

```bash

go run main.go
```

## License

Please refer to the LICENSE file in the repository.

## Acknowledgements

- [Akhil Sharma](https://github.com/AkhilSharma90) for his breakdown on grpcs
