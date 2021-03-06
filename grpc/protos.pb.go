// Code generated by protoc-gen-go.
// source: grpc/protos.proto
// DO NOT EDIT!

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	grpc/protos.proto

It has these top-level messages:
	Request
	Response
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Method  string `protobuf:"bytes,1,opt,name=Method" json:"Method,omitempty"`
	Payload string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}

type Response struct {
	Method        string `protobuf:"bytes,1,opt,name=Method" json:"Method,omitempty"`
	PayloadLength int64  `protobuf:"varint,2,opt,name=PayloadLength" json:"PayloadLength,omitempty"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// Client API for Responder service

type ResponderClient interface {
	ReturnLength(ctx context.Context, in *Request, opts ...grpc1.CallOption) (*Response, error)
}

type responderClient struct {
	cc *grpc1.ClientConn
}

func NewResponderClient(cc *grpc1.ClientConn) ResponderClient {
	return &responderClient{cc}
}

func (c *responderClient) ReturnLength(ctx context.Context, in *Request, opts ...grpc1.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc1.Invoke(ctx, "/grpc.Responder/ReturnLength", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Responder service

type ResponderServer interface {
	ReturnLength(context.Context, *Request) (*Response, error)
}

func RegisterResponderServer(s *grpc1.Server, srv ResponderServer) {
	s.RegisterService(&_Responder_serviceDesc, srv)
}

func _Responder_ReturnLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ResponderServer).ReturnLength(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Responder_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.Responder",
	HandlerType: (*ResponderServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "ReturnLength",
			Handler:    _Responder_ReturnLength_Handler,
		},
	},
	Streams: []grpc1.StreamDesc{},
}
