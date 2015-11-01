package server

import (
    zmq "github.com/pebbe/zmq4"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
    "runtime"
    "log"
)

func StartZeromqProtoRouterDealerServer(port string) {
    frontend, _ := zmq.NewSocket(zmq.ROUTER)
    defer frontend.Close()

    backend, _ := zmq.NewSocket(zmq.DEALER)
    defer backend.Close()

    frontend.Bind("tcp://0.0.0.0:" + port)
    backend.Bind("inproc://backend")

    for i := 0; i < runtime.NumCPU(); i++ {
        go func() {
            responder, _ := zmq.NewSocket(zmq.REP)
            defer responder.Close()

            responder.Connect("inproc://backend")

            for {
                request, _ := responder.RecvBytes(0)

                var body pb.Request

                proto.Unmarshal(request, &body)

                response, _ := proto.Marshal(&pb.Response{
                    Method: body.Method,
                    PayloadLength: int64(len(body.Payload)),
                })

                responder.Send(string(response), 0)
            }
        }()
    }

    err := zmq.Proxy(frontend, backend, nil)
    log.Fatalln(err)
}