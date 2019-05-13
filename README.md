# go-rpc-calculator

This is a very basic example of how to implement a simple calcualtor service in [gRPC](https://grpc.io) that adds two numbers.

It's composed of:

- `/calcpb`: the protocol buffer package which contains the gRPC message and the service definitions.
- `/server`: is the gRPC server that does the performs the add operation.
- `/client`: is the gRPC client which communicate to the server by sending a request and receive the result of the add operation as response.
- `generate.sh` is a script to generate the protocol buffers go code from the .proto file. 


## Usage

1. Clone the repo inside your `$GOPATH`

2. From the root of the project, install the required dependencies
```
go get
```

3. Run the server:
```bash
go run server/server.go
```

4. Run the client:
 ```bash
go run client/client.go
```

