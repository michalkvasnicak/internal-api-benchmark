package server

import (
    "encoding/json"
    zmq "github.com/pebbe/zmq4"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
    "log"
)

func StartZeromqJsonRepServer(port string) {
    var socket *zmq.Socket
    var err error

    if socket, err = zmq.NewSocket(zmq.REP); err != nil {
        log.Fatal(err)
    }

    if err = socket.Bind("tcp://0.0.0.0:" + port); err != nil {
        log.Fatal(err)
    }

    for {
        request, _ := socket.RecvBytes(0)

        var body Request

        json.Unmarshal(request, body)

        response, _ := json.Marshal(Response{
            Method: body.Method,
            PayloadLength: len(body.Payload),
        })

        socket.SendBytes(response, 0)
    }
}
