package client

import (
    "encoding/json"
    "sync"
    zmq "github.com/pebbe/zmq4"
    "github.com/rcrowley/go-metrics"
    "strings"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
)

func StartZeromqJsonTest(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer, requestSize *int) func(awg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        socket, _ := zmq.NewSocket(zmq.REQ)

        defer socket.Close()
        defer wg.Done()

        socket.Connect("tcp://" + address)

        request, _ := json.Marshal(Request{ Method: "TEST", Payload: strings.Repeat("a", messageSize) })
        *requestSize = len(request)

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                socket.SendBytes(request, 0)

                socket.Recv(0)
            })
        }
    }
}