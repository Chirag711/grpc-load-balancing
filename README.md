# gRPC Load Balancing Example (Go)

This project demonstrates **basic load balancing in gRPC using Go**.
Multiple gRPC servers run the same service, and a client distributes requests across them using **round-robin load balancing**.

The goal of this project is to show how **microservices distribute traffic across multiple service instances**.

---

## Overview

In distributed systems, a single service instance may become overloaded if all requests go to one server. Load balancing distributes requests across multiple servers to improve reliability and scalability.

Example request flow:

Client Request 1 → Server 1
Client Request 2 → Server 2
Client Request 3 → Server 3
Client Request 4 → Server 1

This example uses **client-side round-robin load balancing**.

---

## Project Structure

```id="f6b80e"
grpc-load-balancing
│
├── go.mod
│
├── proto
│   └── greeting.proto
│
├── pb
│   ├── greeting.pb.go
│   └── greeting_grpc.pb.go
│
├── servers
│   ├── server1
│   │   └── main.go
│   │
│   ├── server2
│   │   └── main.go
│   │
│   └── server3
│       └── main.go
│
└── client
    └── main.go
```

---

## Prerequisites

Install the following tools:

* Go 1.20 or later
* Protocol Buffers compiler (protoc)
* gRPC Go plugins

Install plugins:

```id="efb8f3"
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## Install Dependencies

Initialize the module:

```id="21b86b"
go mod init grpc-load-balancing
```

Install gRPC:

```id="2b2d33"
go get google.golang.org/grpc
```

---

## Proto Definition

File: `proto/greeting.proto`

```id="d1e2b2"
syntax = "proto3";

package greeting;

option go_package = "grpc-load-balancing/pb";

service GreetingService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  string name = 1;
}

message HelloResponse {
  string message = 1;
}
```

Generate code:

```id="b227e2"
protoc --go_out=. --go-grpc_out=. proto/greeting.proto
```

---

## Servers

This project runs **three identical servers**. Each server listens on a different port.

Server 1 → Port 50051
Server 2 → Port 50052
Server 3 → Port 50053

Each server returns a message indicating which server handled the request.

Example server response:

Server 1 response:

```
Hello Chirag from Server 1
```

Server 2 response:

```
Hello Chirag from Server 2
```

Server 3 response:

```
Hello Chirag from Server 3
```

---

## Client

The client connects to multiple servers and distributes requests using **round-robin load balancing**.

Connection example:

```
dns:///localhost:50051,localhost:50052,localhost:50053
```

Load balancing policy:

```
round_robin
```

The client sends multiple requests and receives responses from different servers.

---

## Running the Application

### Start Server 1

```
go run ./servers/server1
```

Output:

```
Server 1 running on port 50051
```

---

### Start Server 2

```
go run ./servers/server2
```

Output:

```
Server 2 running on port 50052
```

---

### Start Server 3

```
go run ./servers/server3
```

Output:

```
Server 3 running on port 50053
```

---

### Run the Client

```
go run ./client
```

Example output:

```
Hello Chirag from Server 1
Hello Chirag from Server 2
Hello Chirag from Server 3
Hello Chirag from Server 1
Hello Chirag from Server 2
```

Requests are distributed evenly across the servers.

---

## Concepts Demonstrated

* gRPC client-server communication
* Multiple service instances
* Round-robin load balancing
* Scalable microservice architecture

---

## Types of Load Balancing

Pick First
Client connects to the first available server.

Round Robin
Requests are distributed sequentially across servers.

Weighted Load Balancing
Servers receive traffic based on capacity.

Least Connections
Requests go to the server with the fewest active connections.

---

## Real Production Usage

In production environments, load balancing is usually handled by infrastructure such as:

* Kubernetes Services
* Envoy Proxy
* API Gateways
* Cloud Load Balancers

These systems automatically route traffic across service instances.

---

## Possible Improvements

This project can be extended with:

* Service discovery
* Health checks
* Automatic scaling
* Kubernetes deployment
* Metrics and monitoring
* Circuit breaker integration
