# grpc-gateway example
Fully Implemented gRPC-Gateway example in golang.

# Setup
Run the following Commands to install the following packages into your `$GOPATH`
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
Then, To generate the client and proxy files, run:
```
./gen_grpc_stubs.sh
```
This should place generated file in your `$GOPATH`. Troubleshoot as required.

# Explanation
The way the entire world communicates is REST, with JSON. Google open-sourced gRPC in 2015, and it took the tech industry by storm. It goes perfectly hand in hand in microservice communication and with golang. Grpc-Gateway is an API to assist with the runtime translation of JSON into Protocol Buffers, the exchange format used by gRPC, and back. It generates a proxy attached to your microservice, which recieves the data and parses it into protobuf/json for you. It is a great tool to help companies adapt gRPC without rewriting their entire existing architecture, as well. It is a very powerful translation layer.

## Code Explanation
There are two golang files - `main.go` and `service.go`. `main.go` is the proxy itself that connects to the service running in `service.go`, where the server work is being done, and the request being handled. 

# Running
To Build both programs, run:
```
go build main.go
go build service.go
```
This will generate 2 executables, `main` and `service`. Run `service` first with `./service` because main connects to service. (ie the proxy connects to the service). This is crucial. After running `./service`, run `./main`, and the service as a whole should work.

# Test with Curl
If you have curl installed, run the command:
```
curl -d '{"val1":10, "val2":10, "operation_type" :"Multiplication"}' -H "Content-Type: application/json" -X POST http://localhost:8081/api/v1/perform_operation
```
Here, `"val1"` and `"val2"` are the arguments for the math evaluation, and the `"operation_type"` string specifies the calculation type. The endpoint itself was defined in `example_service.proto`. You should see a response(json) like:
```
{"response":"100","operation_valid":true}
```
The service is being implemented with gRPC, but we communicate with it using REST.

# Author
- Aditya Prerepa 2019.
