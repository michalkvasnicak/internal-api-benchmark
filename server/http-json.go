package server

import (
    "net/http"
    "log"
    "encoding/json"
    . "github.com/michalkvasnicak/internal-api-benchmark/json"
)

func handleJsonRequest(w http.ResponseWriter, r *http.Request) {
    var err error
    var body Request

    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)

    if err = decoder.Decode(&body); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    response := Response{
        Method: body.Method,
        PayloadLength: len(body.Payload),
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type", "application/json")

    encoder := json.NewEncoder(w)
    encoder.Encode(response)
}

func StartHttpJsonServer(port string) {
    var err error

    http.HandleFunc("/", handleJsonRequest)

    if err = http.ListenAndServe("0.0.0.0:" + port, nil); err != nil {
        log.Fatal("Server could not be started: ", err)
    }
}