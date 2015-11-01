package client

import (
    "sync"
    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/protocol/req"
    "github.com/gdamore/mangos/transport/tcp"
    "github.com/rcrowley/go-metrics"
    "log"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
    "strings"
)

func StartNanomsgProtobufTest(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer) func(wg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        var socket mangos.Socket
        var err error
        var data []byte

        if socket, err = req.NewSocket(); err != nil {
            log.Fatal(err)
        }

        defer socket.Close()
        defer wg.Done()

        socket.AddTransport(tcp.NewTransport())

        if err = socket.Dial("tcp://" + address); err != nil {
            log.Fatal(err)
        }

        request := &pb.Request{
            Method: "TEST",
            Payload: strings.Repeat("a", messageSize),
        }

        if data, err = proto.Marshal(request); err != nil {
            log.Fatal(err)
        }

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                if err = socket.Send(data); err != nil {
                    log.Fatal(err)
                }

                socket.Recv()
            })
        }
    }
}