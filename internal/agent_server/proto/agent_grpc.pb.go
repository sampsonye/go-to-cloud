// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: agent.proto

package gotocloud

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

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	// rpc GitClone (stream CloneRequest) returns (stream CloneResponse) {}  // 克隆代码(s->c)
	Running(ctx context.Context, opts ...grpc.CallOption) (Agent_RunningClient, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) Running(ctx context.Context, opts ...grpc.CallOption) (Agent_RunningClient, error) {
	stream, err := c.cc.NewStream(ctx, &Agent_ServiceDesc.Streams[0], "/gotocloud.Agent/Running", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRunningClient{stream}
	return x, nil
}

type Agent_RunningClient interface {
	Send(*RunResult) error
	Recv() (*RunRequest, error)
	grpc.ClientStream
}

type agentRunningClient struct {
	grpc.ClientStream
}

func (x *agentRunningClient) Send(m *RunResult) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRunningClient) Recv() (*RunRequest, error) {
	m := new(RunRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	// rpc GitClone (stream CloneRequest) returns (stream CloneResponse) {}  // 克隆代码(s->c)
	Running(Agent_RunningServer) error
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) Running(Agent_RunningServer) error {
	return status.Errorf(codes.Unimplemented, "method Running not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_Running_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentServer).Running(&agentRunningServer{stream})
}

type Agent_RunningServer interface {
	Send(*RunRequest) error
	Recv() (*RunResult, error)
	grpc.ServerStream
}

type agentRunningServer struct {
	grpc.ServerStream
}

func (x *agentRunningServer) Send(m *RunRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRunningServer) Recv() (*RunResult, error) {
	m := new(RunResult)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gotocloud.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Running",
			Handler:       _Agent_Running_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "agent.proto",
}
