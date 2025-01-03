// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/users.proto

package users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Users service

func NewUsersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Users service

type UsersService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Users_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Users_PingPongService, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Users.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Users_StreamService, error) {
	req := c.c.NewRequest(c.name, "Users.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &usersServiceStream{stream}, nil
}

type Users_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type usersServiceStream struct {
	stream client.Stream
}

func (x *usersServiceStream) Close() error {
	return x.stream.Close()
}

func (x *usersServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *usersServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *usersServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *usersServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *usersService) PingPong(ctx context.Context, opts ...client.CallOption) (Users_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Users.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &usersServicePingPong{stream}, nil
}

type Users_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type usersServicePingPong struct {
	stream client.Stream
}

func (x *usersServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *usersServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *usersServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *usersServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *usersServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *usersServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Users service

type UsersHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Users_StreamStream) error
	PingPong(context.Context, Users_PingPongStream) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.UsersHandler.Call(ctx, in, out)
}

func (h *usersHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.UsersHandler.Stream(ctx, m, &usersStreamStream{stream})
}

type Users_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type usersStreamStream struct {
	stream server.Stream
}

func (x *usersStreamStream) Close() error {
	return x.stream.Close()
}

func (x *usersStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *usersStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *usersStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *usersStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *usersHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.UsersHandler.PingPong(ctx, &usersPingPongStream{stream})
}

type Users_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type usersPingPongStream struct {
	stream server.Stream
}

func (x *usersPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *usersPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *usersPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *usersPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *usersPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *usersPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
