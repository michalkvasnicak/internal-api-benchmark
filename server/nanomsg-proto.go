package server

import (
    "github.com/gdamore/mangos"
    "github.com/gdamore/mangos/transport/tcp"
    "github.com/gdamore/mangos/protocol/rep"
    "log"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
)

func StartNanomsgProtoServer(port string) {
    var server mangos.Socket
    var err error

    if server, err = rep.NewSocket(); err != nil {
        log.Fatal(err)
    }

    server.AddTransport(tcp.NewTransport())
    server.Listen("tcp://0.0.0.0:" + port)

    for {
        msg, _ := server.Recv()

        var body pb.Request

        proto.Unmarshal(msg, &body)

        data, _ := proto.Marshal(&pb.Response{
            Method: body.Method,
            PayloadLength: int64(len(body.Payload)),
        })

        server.Send(data)
    }
}