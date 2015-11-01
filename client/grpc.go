package client

import (
    "log"
    "sync"
    "github.com/rcrowley/go-metrics"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    pb "github.com/michalkvasnicak/internal-api-benchmark/grpc"
    "strings"
)

func StartGrpcClient(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer) func(wg *sync.WaitGroup) {
    var connection *grpc.ClientConn
    var err error

    if connection, err = grpc.Dial(address, grpc.WithInsecure()); err != nil {
        log.Fatal(err)
    }

    client := pb.NewResponderClient(connection)
    request := &pb.Request{Method: "TEST", Payload: strings.Repeat("a", messageSize)}

    return func(wg *sync.WaitGroup) {
        defer wg.Done()

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                _, err := client.ReturnLength(context.Background(), request)

                if err != nil {
                    log.Fatal(err)
                }
            })
        }
    }
}