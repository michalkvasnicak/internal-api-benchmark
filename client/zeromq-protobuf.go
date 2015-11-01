package client

import (
    "sync"
    zmq "github.com/pebbe/zmq4"
    "github.com/rcrowley/go-metrics"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
    "log"
    "strings"
)

func StartZeromqProtobufTest(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer) func(wg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        var socket *zmq.Socket
        var err error
        var request []byte

        if socket, err = zmq.NewSocket(zmq.REQ); err != nil {
            log.Fatal(err)
        }

        defer socket.Close()
        defer wg.Done()

        if err = socket.Connect("tcp://" + address); err != nil {
            log.Fatal(err)
        }

        if request, err = proto.Marshal(&pb.Request{
            Method: "TEST",
            Payload: strings.Repeat("a", messageSize),
        }); err != nil {
            log.Fatal(err)
        }

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                socket.SendBytes(request, 0)

                socket.Recv(0)
            })
        }
    }
}