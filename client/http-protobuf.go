package client

import (
    "log"
    "bytes"
    "sync"
    "net/http"
    "github.com/rcrowley/go-metrics"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
    "strings"
)

func StartHttpProtobufTest(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer) func(wg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        var err error
        var data []byte

        defer wg.Done()

        request := &pb.Request{
            Method: "TEST",
            Payload: strings.Repeat("a", messageSize),
        }

        if data, err = proto.Marshal(request); err != nil {
            log.Fatal(err)
        }

        dataReader := bytes.NewReader(data)

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                _, err = http.Post("http://" + address, "application/json", dataReader)

                if err != nil {
                    log.Fatal(err)
                }
            })
        }
    }
}