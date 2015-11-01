package server

import (
    "net/http"
    "log"
    "bytes"
    "github.com/golang/protobuf/proto"
    pb "github.com/michalkvasnicak/internal-api-benchmark/protobuf"
)

func handleProtobufRequest(w http.ResponseWriter, r *http.Request) {
    var err error
    var data []byte

    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    buffer := new(bytes.Buffer)
    buffer.ReadFrom(r.Body)

    var request pb.Request

    if err = proto.Unmarshal(buffer.Bytes(), &request); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    response := &pb.Response{
        Method: request.Method,
        PayloadLength: int64(len(request.Payload)),
    }

    if data, err = proto.Marshal(response); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Add("Content-Type", "application/octet-stream")

    w.Write(data)
}

func StartHttpProtoServer(port string) {
    var err error

    http.HandleFunc("/", handleProtobufRequest)

    if err = http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal("Server could not be started: ", err)
    }
}