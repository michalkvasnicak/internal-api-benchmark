package server

import (
    "encoding/json"
    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/transport/tcp"
    "github.com/gdamore/mangos/protocol/rep"
    "log"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
)

func StartNanomsgJsonServer(port string) {
    var server mangos.Socket
    var err error

    if server, err = rep.NewSocket(); err != nil {
        log.Fatal(err)
    }

    server.AddTransport(tcp.NewTransport())
    server.Listen("tcp://0.0.0.0:" + port)

    for {
        msg, _ := server.Recv()

        var body Request
        var response []byte

        json.Unmarshal(msg, &body)

        if response, err = json.Marshal(Response{ Method: body.Method, PayloadLength: len(body.Payload) }); err != nil {
            log.Fatal(err)
        }

        server.Send(response)
    }
}