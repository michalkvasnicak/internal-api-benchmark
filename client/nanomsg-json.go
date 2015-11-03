package client

import (
    "encoding/json"
    "sync"
    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/protocol/req"
    "github.com/gdamore/mangos/transport/tcp"
    "github.com/rcrowley/go-metrics"
    "log"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
    "strings"
)

func StartNanomsgJsonTest(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer, requestSize *int) func(wg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        var socket mangos.Socket
        var err error

        if socket, err = req.NewSocket(); err != nil {
            log.Fatal(err)
        }

        defer socket.Close()
        defer wg.Done()

        socket.AddTransport(tcp.NewTransport())

        if err = socket.Dial("tcp://" + address); err != nil {
            log.Fatal(err)
        }

        request, _ := json.Marshal(Request{ Method: "TEST", Payload: strings.Repeat("a", messageSize) })
        *requestSize = len(request)

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                if err = socket.Send(request); err != nil {
                    log.Fatal(err)
                }

                socket.Recv()
            })
        }
    }
}