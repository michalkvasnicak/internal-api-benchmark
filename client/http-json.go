package client

import (
    "log"
    "bytes"
    "sync"
    "encoding/json"
    "net/http"
    "github.com/rcrowley/go-metrics"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
    "strings"
)

func StartHttpJsonClient(address string, clients int, requestsPerClient int, messageSize int, timer metrics.Timer) func(wg *sync.WaitGroup) {
    return func(wg *sync.WaitGroup) {
        var err error
        var rawMessage json.RawMessage

        defer wg.Done()

        request := Request{
            Method: "TEST",
            Payload: strings.Repeat("a", messageSize),
        }

        if rawMessage, err = json.Marshal(request); err != nil {
            log.Fatal(err)
        }

        data := bytes.NewReader(rawMessage)

        for i := 0; i < requestsPerClient; i++ {
            timer.Time(func() {
                _, err = http.Post("http://" + address, "application/json", data)

                if err != nil {
                    log.Fatal(err)
                }
            })
        }
    }
}