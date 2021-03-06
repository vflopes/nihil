// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package analysis

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnalysisServiceClient is the client API for AnalysisService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnalysisServiceClient interface {
	Do(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Axis, error)
}

type analysisServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnalysisServiceClient(cc grpc.ClientConnInterface) AnalysisServiceClient {
	return &analysisServiceClient{cc}
}

func (c *analysisServiceClient) Do(ctx context.Context, in *Pipeline, opts ...grpc.CallOption) (*Axis, error) {
	out := new(Axis)
	err := c.cc.Invoke(ctx, "/nihil.analysis.AnalysisService/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnalysisServiceServer is the server API for AnalysisService service.
// All implementations must embed UnimplementedAnalysisServiceServer
// for forward compatibility
type AnalysisServiceServer interface {
	Do(context.Context, *Pipeline) (*Axis, error)
	mustEmbedUnimplementedAnalysisServiceServer()
}

// UnimplementedAnalysisServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnalysisServiceServer struct {
}

func (*UnimplementedAnalysisServiceServer) Do(context.Context, *Pipeline) (*Axis, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (*UnimplementedAnalysisServiceServer) mustEmbedUnimplementedAnalysisServiceServer() {}

func RegisterAnalysisServiceServer(s *grpc.Server, srv AnalysisServiceServer) {
	s.RegisterService(&_AnalysisService_serviceDesc, srv)
}

func _AnalysisService_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pipeline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnalysisServiceServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nihil.analysis.AnalysisService/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnalysisServiceServer).Do(ctx, req.(*Pipeline))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnalysisService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nihil.analysis.AnalysisService",
	HandlerType: (*AnalysisServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _AnalysisService_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "analysis/analysis.proto",
}
