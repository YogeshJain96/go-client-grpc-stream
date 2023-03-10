// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package books

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BooksClient is the client API for Books service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BooksClient interface {
	// ValidateBooks
	ValidateBooks(ctx context.Context, opts ...grpc.CallOption) (Books_ValidateBooksClient, error)
}

type booksClient struct {
	cc grpc.ClientConnInterface
}

func NewBooksClient(cc grpc.ClientConnInterface) BooksClient {
	return &booksClient{cc}
}

func (c *booksClient) ValidateBooks(ctx context.Context, opts ...grpc.CallOption) (Books_ValidateBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &Books_ServiceDesc.Streams[0], "/books.Books/ValidateBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &booksValidateBooksClient{stream}
	return x, nil
}

type Books_ValidateBooksClient interface {
	Send(*ValidationReq) error
	CloseAndRecv() (*ValidationRes, error)
	grpc.ClientStream
}

type booksValidateBooksClient struct {
	grpc.ClientStream
}

func (x *booksValidateBooksClient) Send(m *ValidationReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *booksValidateBooksClient) CloseAndRecv() (*ValidationRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ValidationRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BooksServer is the server API for Books service.
// All implementations must embed UnimplementedBooksServer
// for forward compatibility
type BooksServer interface {
	// ValidateBooks
	ValidateBooks(Books_ValidateBooksServer) error
	mustEmbedUnimplementedBooksServer()
}

// UnimplementedBooksServer must be embedded to have forward compatible implementations.
type UnimplementedBooksServer struct {
}

func (UnimplementedBooksServer) ValidateBooks(Books_ValidateBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method ValidateBooks not implemented")
}
func (UnimplementedBooksServer) mustEmbedUnimplementedBooksServer() {}

// UnsafeBooksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BooksServer will
// result in compilation errors.
type UnsafeBooksServer interface {
	mustEmbedUnimplementedBooksServer()
}

func RegisterBooksServer(s grpc.ServiceRegistrar, srv BooksServer) {
	s.RegisterService(&Books_ServiceDesc, srv)
}

func _Books_ValidateBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BooksServer).ValidateBooks(&booksValidateBooksServer{stream})
}

type Books_ValidateBooksServer interface {
	SendAndClose(*ValidationRes) error
	Recv() (*ValidationReq, error)
	grpc.ServerStream
}

type booksValidateBooksServer struct {
	grpc.ServerStream
}

func (x *booksValidateBooksServer) SendAndClose(m *ValidationRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *booksValidateBooksServer) Recv() (*ValidationReq, error) {
	m := new(ValidationReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Books_ServiceDesc is the grpc.ServiceDesc for Books service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Books_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "books.Books",
	HandlerType: (*BooksServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ValidateBooks",
			Handler:       _Books_ValidateBooks_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "books.proto",
}
