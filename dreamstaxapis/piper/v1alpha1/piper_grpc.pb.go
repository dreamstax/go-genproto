// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: piper.proto

package piper

import (
	context "context"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Piper_ListPipelineJobs_FullMethodName           = "/kai.piper.v1alpha1.Piper/ListPipelineJobs"
	Piper_GetPipleineJob_FullMethodName             = "/kai.piper.v1alpha1.Piper/GetPipleineJob"
	Piper_CreatePipelineJob_FullMethodName          = "/kai.piper.v1alpha1.Piper/CreatePipelineJob"
	Piper_UpdatePipelineJob_FullMethodName          = "/kai.piper.v1alpha1.Piper/UpdatePipelineJob"
	Piper_DeletePipelineJob_FullMethodName          = "/kai.piper.v1alpha1.Piper/DeletePipelineJob"
	Piper_RunPipelineJob_FullMethodName             = "/kai.piper.v1alpha1.Piper/RunPipelineJob"
	Piper_ListPipelineJobExecutions_FullMethodName  = "/kai.piper.v1alpha1.Piper/ListPipelineJobExecutions"
	Piper_GetPipelineJobExecution_FullMethodName    = "/kai.piper.v1alpha1.Piper/GetPipelineJobExecution"
	Piper_DeletePipelineJobExecution_FullMethodName = "/kai.piper.v1alpha1.Piper/DeletePipelineJobExecution"
)

// PiperClient is the client API for Piper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PiperClient interface {
	ListPipelineJobs(ctx context.Context, in *ListPipelineJobsRequest, opts ...grpc.CallOption) (*ListPipelineJobsResponse, error)
	GetPipleineJob(ctx context.Context, in *GetPipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error)
	CreatePipelineJob(ctx context.Context, in *CreatePipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error)
	UpdatePipelineJob(ctx context.Context, in *UpdatePipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error)
	DeletePipelineJob(ctx context.Context, in *DeletePipelineJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RunPipelineJob(ctx context.Context, in *RunPipelineJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	ListPipelineJobExecutions(ctx context.Context, in *ListPipelineJobExecutionsRequest, opts ...grpc.CallOption) (*ListPipelineJobExecutionsResponse, error)
	GetPipelineJobExecution(ctx context.Context, in *GetPipelineJobExecutionRequest, opts ...grpc.CallOption) (*PipelineJobExecution, error)
	DeletePipelineJobExecution(ctx context.Context, in *DeletePipelineJobExecutionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type piperClient struct {
	cc grpc.ClientConnInterface
}

func NewPiperClient(cc grpc.ClientConnInterface) PiperClient {
	return &piperClient{cc}
}

func (c *piperClient) ListPipelineJobs(ctx context.Context, in *ListPipelineJobsRequest, opts ...grpc.CallOption) (*ListPipelineJobsResponse, error) {
	out := new(ListPipelineJobsResponse)
	err := c.cc.Invoke(ctx, Piper_ListPipelineJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) GetPipleineJob(ctx context.Context, in *GetPipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error) {
	out := new(PipelineJob)
	err := c.cc.Invoke(ctx, Piper_GetPipleineJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) CreatePipelineJob(ctx context.Context, in *CreatePipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error) {
	out := new(PipelineJob)
	err := c.cc.Invoke(ctx, Piper_CreatePipelineJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) UpdatePipelineJob(ctx context.Context, in *UpdatePipelineJobRequest, opts ...grpc.CallOption) (*PipelineJob, error) {
	out := new(PipelineJob)
	err := c.cc.Invoke(ctx, Piper_UpdatePipelineJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) DeletePipelineJob(ctx context.Context, in *DeletePipelineJobRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Piper_DeletePipelineJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) RunPipelineJob(ctx context.Context, in *RunPipelineJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, Piper_RunPipelineJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) ListPipelineJobExecutions(ctx context.Context, in *ListPipelineJobExecutionsRequest, opts ...grpc.CallOption) (*ListPipelineJobExecutionsResponse, error) {
	out := new(ListPipelineJobExecutionsResponse)
	err := c.cc.Invoke(ctx, Piper_ListPipelineJobExecutions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) GetPipelineJobExecution(ctx context.Context, in *GetPipelineJobExecutionRequest, opts ...grpc.CallOption) (*PipelineJobExecution, error) {
	out := new(PipelineJobExecution)
	err := c.cc.Invoke(ctx, Piper_GetPipelineJobExecution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piperClient) DeletePipelineJobExecution(ctx context.Context, in *DeletePipelineJobExecutionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Piper_DeletePipelineJobExecution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PiperServer is the server API for Piper service.
// All implementations must embed UnimplementedPiperServer
// for forward compatibility
type PiperServer interface {
	ListPipelineJobs(context.Context, *ListPipelineJobsRequest) (*ListPipelineJobsResponse, error)
	GetPipleineJob(context.Context, *GetPipelineJobRequest) (*PipelineJob, error)
	CreatePipelineJob(context.Context, *CreatePipelineJobRequest) (*PipelineJob, error)
	UpdatePipelineJob(context.Context, *UpdatePipelineJobRequest) (*PipelineJob, error)
	DeletePipelineJob(context.Context, *DeletePipelineJobRequest) (*emptypb.Empty, error)
	RunPipelineJob(context.Context, *RunPipelineJobRequest) (*longrunning.Operation, error)
	ListPipelineJobExecutions(context.Context, *ListPipelineJobExecutionsRequest) (*ListPipelineJobExecutionsResponse, error)
	GetPipelineJobExecution(context.Context, *GetPipelineJobExecutionRequest) (*PipelineJobExecution, error)
	DeletePipelineJobExecution(context.Context, *DeletePipelineJobExecutionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPiperServer()
}

// UnimplementedPiperServer must be embedded to have forward compatible implementations.
type UnimplementedPiperServer struct {
}

func (UnimplementedPiperServer) ListPipelineJobs(context.Context, *ListPipelineJobsRequest) (*ListPipelineJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineJobs not implemented")
}
func (UnimplementedPiperServer) GetPipleineJob(context.Context, *GetPipelineJobRequest) (*PipelineJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipleineJob not implemented")
}
func (UnimplementedPiperServer) CreatePipelineJob(context.Context, *CreatePipelineJobRequest) (*PipelineJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipelineJob not implemented")
}
func (UnimplementedPiperServer) UpdatePipelineJob(context.Context, *UpdatePipelineJobRequest) (*PipelineJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePipelineJob not implemented")
}
func (UnimplementedPiperServer) DeletePipelineJob(context.Context, *DeletePipelineJobRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipelineJob not implemented")
}
func (UnimplementedPiperServer) RunPipelineJob(context.Context, *RunPipelineJobRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPipelineJob not implemented")
}
func (UnimplementedPiperServer) ListPipelineJobExecutions(context.Context, *ListPipelineJobExecutionsRequest) (*ListPipelineJobExecutionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPipelineJobExecutions not implemented")
}
func (UnimplementedPiperServer) GetPipelineJobExecution(context.Context, *GetPipelineJobExecutionRequest) (*PipelineJobExecution, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineJobExecution not implemented")
}
func (UnimplementedPiperServer) DeletePipelineJobExecution(context.Context, *DeletePipelineJobExecutionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipelineJobExecution not implemented")
}
func (UnimplementedPiperServer) mustEmbedUnimplementedPiperServer() {}

// UnsafePiperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PiperServer will
// result in compilation errors.
type UnsafePiperServer interface {
	mustEmbedUnimplementedPiperServer()
}

func RegisterPiperServer(s grpc.ServiceRegistrar, srv PiperServer) {
	s.RegisterService(&Piper_ServiceDesc, srv)
}

func _Piper_ListPipelineJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).ListPipelineJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_ListPipelineJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).ListPipelineJobs(ctx, req.(*ListPipelineJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_GetPipleineJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).GetPipleineJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_GetPipleineJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).GetPipleineJob(ctx, req.(*GetPipelineJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_CreatePipelineJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).CreatePipelineJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_CreatePipelineJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).CreatePipelineJob(ctx, req.(*CreatePipelineJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_UpdatePipelineJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePipelineJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).UpdatePipelineJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_UpdatePipelineJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).UpdatePipelineJob(ctx, req.(*UpdatePipelineJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_DeletePipelineJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).DeletePipelineJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_DeletePipelineJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).DeletePipelineJob(ctx, req.(*DeletePipelineJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_RunPipelineJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunPipelineJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).RunPipelineJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_RunPipelineJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).RunPipelineJob(ctx, req.(*RunPipelineJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_ListPipelineJobExecutions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPipelineJobExecutionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).ListPipelineJobExecutions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_ListPipelineJobExecutions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).ListPipelineJobExecutions(ctx, req.(*ListPipelineJobExecutionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_GetPipelineJobExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineJobExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).GetPipelineJobExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_GetPipelineJobExecution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).GetPipelineJobExecution(ctx, req.(*GetPipelineJobExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Piper_DeletePipelineJobExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePipelineJobExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiperServer).DeletePipelineJobExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Piper_DeletePipelineJobExecution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiperServer).DeletePipelineJobExecution(ctx, req.(*DeletePipelineJobExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Piper_ServiceDesc is the grpc.ServiceDesc for Piper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Piper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kai.piper.v1alpha1.Piper",
	HandlerType: (*PiperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPipelineJobs",
			Handler:    _Piper_ListPipelineJobs_Handler,
		},
		{
			MethodName: "GetPipleineJob",
			Handler:    _Piper_GetPipleineJob_Handler,
		},
		{
			MethodName: "CreatePipelineJob",
			Handler:    _Piper_CreatePipelineJob_Handler,
		},
		{
			MethodName: "UpdatePipelineJob",
			Handler:    _Piper_UpdatePipelineJob_Handler,
		},
		{
			MethodName: "DeletePipelineJob",
			Handler:    _Piper_DeletePipelineJob_Handler,
		},
		{
			MethodName: "RunPipelineJob",
			Handler:    _Piper_RunPipelineJob_Handler,
		},
		{
			MethodName: "ListPipelineJobExecutions",
			Handler:    _Piper_ListPipelineJobExecutions_Handler,
		},
		{
			MethodName: "GetPipelineJobExecution",
			Handler:    _Piper_GetPipelineJobExecution_Handler,
		},
		{
			MethodName: "DeletePipelineJobExecution",
			Handler:    _Piper_DeletePipelineJobExecution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "piper.proto",
}
