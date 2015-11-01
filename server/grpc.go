package server

import (
    "log"
    "net"
    "google.golang.org/grpc"
    "golang.org/x/net/context"
    pb "github.com/michalkvasnicak/internal-api-benchmark/grpc"
)

type server struct{}

func (s *server) ReturnLength(ctx context.Context, in *pb.Request) (*pb.Response, error) {
    return &pb.Response{Method: in.Method, PayloadLength:int64(len(in.Payload))}, nil
}

func StartGrpcServer(port string) {
    var lis net.Listener
    var err error

    if lis, err = net.Listen("tcp", "0.0.0.0:" + port); err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterResponderServer(s, &server{})
    s.Serve(lis)
}