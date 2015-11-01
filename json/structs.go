package json

type Request struct {
    Method, Payload string
}

type Response struct {
    Method string
    PayloadLength int
}
