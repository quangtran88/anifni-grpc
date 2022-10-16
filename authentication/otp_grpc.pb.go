// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: authentication/otp.proto

package otpGRPC

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

// OTPServiceClient is the client API for OTPService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTPServiceClient interface {
	CheckEmailOTP(ctx context.Context, in *CheckEmailOTPInput, opts ...grpc.CallOption) (*CheckEmailOTPResult, error)
	SendEmailOTP(ctx context.Context, in *SendEmailOTPInput, opts ...grpc.CallOption) (*SendEmailOTPResult, error)
}

type oTPServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOTPServiceClient(cc grpc.ClientConnInterface) OTPServiceClient {
	return &oTPServiceClient{cc}
}

func (c *oTPServiceClient) CheckEmailOTP(ctx context.Context, in *CheckEmailOTPInput, opts ...grpc.CallOption) (*CheckEmailOTPResult, error) {
	out := new(CheckEmailOTPResult)
	err := c.cc.Invoke(ctx, "/otpGRPC.OTPService/CheckEmailOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oTPServiceClient) SendEmailOTP(ctx context.Context, in *SendEmailOTPInput, opts ...grpc.CallOption) (*SendEmailOTPResult, error) {
	out := new(SendEmailOTPResult)
	err := c.cc.Invoke(ctx, "/otpGRPC.OTPService/SendEmailOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OTPServiceServer is the server API for OTPService service.
// All implementations must embed UnimplementedOTPServiceServer
// for forward compatibility
type OTPServiceServer interface {
	CheckEmailOTP(context.Context, *CheckEmailOTPInput) (*CheckEmailOTPResult, error)
	SendEmailOTP(context.Context, *SendEmailOTPInput) (*SendEmailOTPResult, error)
	mustEmbedUnimplementedOTPServiceServer()
}

// UnimplementedOTPServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOTPServiceServer struct {
}

func (UnimplementedOTPServiceServer) CheckEmailOTP(context.Context, *CheckEmailOTPInput) (*CheckEmailOTPResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckEmailOTP not implemented")
}
func (UnimplementedOTPServiceServer) SendEmailOTP(context.Context, *SendEmailOTPInput) (*SendEmailOTPResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmailOTP not implemented")
}
func (UnimplementedOTPServiceServer) mustEmbedUnimplementedOTPServiceServer() {}

// UnsafeOTPServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTPServiceServer will
// result in compilation errors.
type UnsafeOTPServiceServer interface {
	mustEmbedUnimplementedOTPServiceServer()
}

func RegisterOTPServiceServer(s grpc.ServiceRegistrar, srv OTPServiceServer) {
	s.RegisterService(&OTPService_ServiceDesc, srv)
}

func _OTPService_CheckEmailOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckEmailOTPInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).CheckEmailOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otpGRPC.OTPService/CheckEmailOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).CheckEmailOTP(ctx, req.(*CheckEmailOTPInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _OTPService_SendEmailOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailOTPInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OTPServiceServer).SendEmailOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/otpGRPC.OTPService/SendEmailOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OTPServiceServer).SendEmailOTP(ctx, req.(*SendEmailOTPInput))
	}
	return interceptor(ctx, in, info, handler)
}

// OTPService_ServiceDesc is the grpc.ServiceDesc for OTPService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTPService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "otpGRPC.OTPService",
	HandlerType: (*OTPServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckEmailOTP",
			Handler:    _OTPService_CheckEmailOTP_Handler,
		},
		{
			MethodName: "SendEmailOTP",
			Handler:    _OTPService_SendEmailOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/otp.proto",
}
