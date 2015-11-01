package server

import (
    "encoding/json"
    zmq "github.com/pebbe/zmq4"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
    "runtime"
    "log"
)

func StartZeromqJsonRouterDealerServer(port string) {
    frontend, _ := zmq.NewSocket(zmq.ROUTER)
    defer frontend.Close()

    backend, _ := zmq.NewSocket(zmq.DEALER)
    defer backend.Close()

    frontend.Bind("tcp://0.0.0.0:" + port)
    backend.Bind("inproc://backend")

    // start num cpu request processors
    for i := 0; i < runtime.NumCPU(); i++ {
        go func() {
            responder, _ := zmq.NewSocket(zmq.REP)
            defer responder.Close()

            responder.Connect("inproc://backend")

            for {
                request, _ := responder.RecvBytes(0)

                var body Request

                json.Unmarshal(request, body)

                response, _ := json.Marshal(Response{ Method: body.Method, PayloadLength: len(body.Payload)})

                responder.Send(string(response), 0)
            }
        }()
    }

    err := zmq.Proxy(frontend, backend, nil)
    log.Fatalln(err)
}