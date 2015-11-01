package server

import (
    "log"
    zmq "github.com/pebbe/zmq4"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
)

func StartZeromqProtoRepServer(port string) {
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

        var body pb.Request

        proto.Unmarshal(request, &body)

        response, _ := proto.Marshal(&pb.Response{
            Method: body.Method,
            PayloadLength: int64(len(body.Payload)),
        })

        socket.SendBytes(response, 0)
    }
}