// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: kmgm.proto

package pb

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/kmgm.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionServiceClient interface {
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/kmgm.VersionService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
// All implementations must embed UnimplementedVersionServiceServer
// for forward compatibility
type VersionServiceServer interface {
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedVersionServiceServer()
}

// UnimplementedVersionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVersionServiceServer struct {
}

func (UnimplementedVersionServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedVersionServiceServer) mustEmbedUnimplementedVersionServiceServer() {}

// UnsafeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionServiceServer will
// result in compilation errors.
type UnsafeVersionServiceServer interface {
	mustEmbedUnimplementedVersionServiceServer()
}

func RegisterVersionServiceServer(s grpc.ServiceRegistrar, srv VersionServiceServer) {
	s.RegisterService(&VersionService_ServiceDesc, srv)
}

func _VersionService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.VersionService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionService_ServiceDesc is the grpc.ServiceDesc for VersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _VersionService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateServiceClient interface {
	IssuePreflight(ctx context.Context, in *IssuePreflightRequest, opts ...grpc.CallOption) (*IssuePreflightResponse, error)
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error)
	GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) IssuePreflight(ctx context.Context, in *IssuePreflightRequest, opts ...grpc.CallOption) (*IssuePreflightResponse, error) {
	out := new(IssuePreflightResponse)
	err := c.cc.Invoke(ctx, "/kmgm.CertificateService/IssuePreflight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error) {
	out := new(IssueCertificateResponse)
	err := c.cc.Invoke(ctx, "/kmgm.CertificateService/IssueCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error) {
	out := new(GetCertificateResponse)
	err := c.cc.Invoke(ctx, "/kmgm.CertificateService/GetCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations must embed UnimplementedCertificateServiceServer
// for forward compatibility
type CertificateServiceServer interface {
	IssuePreflight(context.Context, *IssuePreflightRequest) (*IssuePreflightResponse, error)
	IssueCertificate(context.Context, *IssueCertificateRequest) (*IssueCertificateResponse, error)
	GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error)
	mustEmbedUnimplementedCertificateServiceServer()
}

// UnimplementedCertificateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateServiceServer struct {
}

func (UnimplementedCertificateServiceServer) IssuePreflight(context.Context, *IssuePreflightRequest) (*IssuePreflightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssuePreflight not implemented")
}
func (UnimplementedCertificateServiceServer) IssueCertificate(context.Context, *IssueCertificateRequest) (*IssueCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificate not implemented")
}
func (UnimplementedCertificateServiceServer) mustEmbedUnimplementedCertificateServiceServer() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_IssuePreflight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuePreflightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).IssuePreflight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.CertificateService/IssuePreflight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).IssuePreflight(ctx, req.(*IssuePreflightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.CertificateService/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_GetCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).GetCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.CertificateService/GetCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).GetCertificate(ctx, req.(*GetCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssuePreflight",
			Handler:    _CertificateService_IssuePreflight_Handler,
		},
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateService_IssueCertificate_Handler,
		},
		{
			MethodName: "GetCertificate",
			Handler:    _CertificateService_GetCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}
