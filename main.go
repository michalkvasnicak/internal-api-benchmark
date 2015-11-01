package main

import (
    "os"
    "flag"
    "log"
    "fmt"
    "sync"
    "github.com/rcrowley/go-metrics"
    "github.com/michalkvasnicak/internal-api-benchmark/client"
    "github.com/michalkvasnicak/internal-api-benchmark/server"
)

var (
    processType string
    clientType string
    serverType string
    clients int
    requestsPerClient int
    messageSize int
)

func init() {
    flag.StringVar(&processType, "type", "server", "server or client")
    flag.StringVar(&serverType, "serverType", "http-json", "Server type (grpc, http-json, http-proto, nanomsg-json, nanomsg-proto, zeromq-json-rep, zeromq-json-router-dealer, zeromq-proto-rep, zeromq-proto-router-dealer)")
    flag.StringVar(&clientType, "clientType", "http-json", "Client type (grpc, http-json, http-proto, nanomsg-json, nanomsg-proto, zeromq-json, zeromq-proto)")
    flag.IntVar(&clients, "clients", 10, "Number of clients (type has to be client)")
    flag.IntVar(&requestsPerClient, "rpc", 1000, "Number of sequential requests per client (type has ot be client)")
    flag.IntVar(&messageSize, "size", 1024, "Message size in bytes (type has to be client)")
}

func startClient() {
    var testFunc func(string, int, int, int, metrics.Timer) func(*sync.WaitGroup)

    switch clientType {
        case "grpc": testFunc = client.StartGrpcClient
        case "http-json": testFunc = client.StartHttpJsonClient
        case "http-proto": testFunc = client.StartHttpProtobufTest
        case "nanomsg-json": testFunc = client.StartNanomsgJsonTest
        case "nanomsg-proto": testFunc = client.StartNanomsgProtobufTest
        case "zeromq-json": testFunc = client.StartZeromqJsonTest
        case "zeromq-proto": testFunc = client.StartZeromqProtobufTest
        default: log.Fatal("Unknown client type")
    }

    var timer = metrics.NewTimer()
    var wg sync.WaitGroup

    metrics.Register("Requests", timer)

    wg.Add(clients)

    var clientRoutine func(*sync.WaitGroup) = testFunc("127.0.0.1:8000", clients, requestsPerClient, messageSize, timer)

    for client := 0; client < clients; client++ {
        go clientRoutine(&wg)
    }

    wg.Wait()
    metrics.WriteOnce(metrics.DefaultRegistry, os.Stdout)
}

func startServer() {
    var startServer func(string)

    switch serverType {
        case "grpc": startServer = server.StartGrpcServer
        case "http-json": startServer = server.StartHttpJsonServer
        case "http-proto": startServer = server.StartHttpProtoServer
        case "nanomsg-json": startServer = server.StartNanomsgJsonServer
        case "nanomsg-proto": startServer = server.StartNanomsgProtoServer
        case "zeromq-json-rep": startServer = server.StartZeromqJsonRepServer
        case "zeromq-json-router-dealer": startServer = server.StartZeromqJsonRouterDealerServer
        case "zeromq-proto-rep": startServer = server.StartZeromqProtoRepServer
        case "zeromq-proto-router-dealer": startServer = server.StartZeromqProtoRouterDealerServer
        default: log.Fatal("Unknown server type")
    }

    fmt.Printf("Starting %s server\n", serverType)
    startServer("8000")
}

func main() {
    flag.Parse()

    if processType == "client" {
        startClient()
    } else {
        startServer()
    }
}
