// Code generated by protoc-gen-go.
// source: nodeSession.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNodeSessionRequest struct {
	// hex encoded DevAddr
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,3,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppSKey
	AppSKey string `protobuf:"bytes,4,opt,name=appSKey" json:"appSKey,omitempty"`
	// hex encoded NwkSKey
	NwkSKey     string   `protobuf:"bytes,5,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	FCntUp      uint32   `protobuf:"varint,6,opt,name=fCntUp" json:"fCntUp,omitempty"`
	FCntDown    uint32   `protobuf:"varint,7,opt,name=fCntDown" json:"fCntDown,omitempty"`
	RxDelay     uint32   `protobuf:"varint,8,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,9,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,10,rep,packed,name=cFList" json:"cFList,omitempty"`
}

func (m *CreateNodeSessionRequest) Reset()                    { *m = CreateNodeSessionRequest{} }
func (m *CreateNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeSessionRequest) ProtoMessage()               {}
func (*CreateNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type CreateNodeSessionResponse struct {
}

func (m *CreateNodeSessionResponse) Reset()                    { *m = CreateNodeSessionResponse{} }
func (m *CreateNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeSessionResponse) ProtoMessage()               {}
func (*CreateNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

type GetNodeSessionRequest struct {
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
}

func (m *GetNodeSessionRequest) Reset()                    { *m = GetNodeSessionRequest{} }
func (m *GetNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeSessionRequest) ProtoMessage()               {}
func (*GetNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type GetNodeSessionResponse struct {
	// hex encoded DevAddr
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,3,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppSKey
	AppSKey string `protobuf:"bytes,4,opt,name=appSKey" json:"appSKey,omitempty"`
	// hex encoded NwkSKey
	NwkSKey     string   `protobuf:"bytes,5,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	FCntUp      uint32   `protobuf:"varint,6,opt,name=fCntUp" json:"fCntUp,omitempty"`
	FCntDown    uint32   `protobuf:"varint,7,opt,name=fCntDown" json:"fCntDown,omitempty"`
	RxDelay     uint32   `protobuf:"varint,8,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,9,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,10,rep,packed,name=cFList" json:"cFList,omitempty"`
}

func (m *GetNodeSessionResponse) Reset()                    { *m = GetNodeSessionResponse{} }
func (m *GetNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeSessionResponse) ProtoMessage()               {}
func (*GetNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type GetNodeSessionByDevEUIRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *GetNodeSessionByDevEUIRequest) Reset()                    { *m = GetNodeSessionByDevEUIRequest{} }
func (m *GetNodeSessionByDevEUIRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeSessionByDevEUIRequest) ProtoMessage()               {}
func (*GetNodeSessionByDevEUIRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type UpdateNodeSessionRequest struct {
	// hex encoded DevAddr
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
	// hex encoded AppEUI
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,3,opt,name=devEUI" json:"devEUI,omitempty"`
	// hex encoded AppSKey
	AppSKey string `protobuf:"bytes,4,opt,name=appSKey" json:"appSKey,omitempty"`
	// hex encoded NwkSKey
	NwkSKey     string   `protobuf:"bytes,5,opt,name=nwkSKey" json:"nwkSKey,omitempty"`
	FCntUp      uint32   `protobuf:"varint,6,opt,name=fCntUp" json:"fCntUp,omitempty"`
	FCntDown    uint32   `protobuf:"varint,7,opt,name=fCntDown" json:"fCntDown,omitempty"`
	RxDelay     uint32   `protobuf:"varint,8,opt,name=rxDelay" json:"rxDelay,omitempty"`
	Rx1DROffset uint32   `protobuf:"varint,9,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	CFList      []uint32 `protobuf:"varint,10,rep,packed,name=cFList" json:"cFList,omitempty"`
}

func (m *UpdateNodeSessionRequest) Reset()                    { *m = UpdateNodeSessionRequest{} }
func (m *UpdateNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeSessionRequest) ProtoMessage()               {}
func (*UpdateNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type UpdateNodeSessionResponse struct {
}

func (m *UpdateNodeSessionResponse) Reset()                    { *m = UpdateNodeSessionResponse{} }
func (m *UpdateNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeSessionResponse) ProtoMessage()               {}
func (*UpdateNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

type DeleteNodeSessionRequest struct {
	// hex encoded DevAddr
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
}

func (m *DeleteNodeSessionRequest) Reset()                    { *m = DeleteNodeSessionRequest{} }
func (m *DeleteNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSessionRequest) ProtoMessage()               {}
func (*DeleteNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

type DeleteNodeSessionResponse struct {
}

func (m *DeleteNodeSessionResponse) Reset()                    { *m = DeleteNodeSessionResponse{} }
func (m *DeleteNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSessionResponse) ProtoMessage()               {}
func (*DeleteNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

type GetRandomDevAddrRequest struct {
}

func (m *GetRandomDevAddrRequest) Reset()                    { *m = GetRandomDevAddrRequest{} }
func (m *GetRandomDevAddrRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRandomDevAddrRequest) ProtoMessage()               {}
func (*GetRandomDevAddrRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

type GetRandomDevAddrResponse struct {
	// hex encoded DevAddr
	DevAddr string `protobuf:"bytes,1,opt,name=devAddr" json:"devAddr,omitempty"`
}

func (m *GetRandomDevAddrResponse) Reset()                    { *m = GetRandomDevAddrResponse{} }
func (m *GetRandomDevAddrResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRandomDevAddrResponse) ProtoMessage()               {}
func (*GetRandomDevAddrResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func init() {
	proto.RegisterType((*CreateNodeSessionRequest)(nil), "api.CreateNodeSessionRequest")
	proto.RegisterType((*CreateNodeSessionResponse)(nil), "api.CreateNodeSessionResponse")
	proto.RegisterType((*GetNodeSessionRequest)(nil), "api.GetNodeSessionRequest")
	proto.RegisterType((*GetNodeSessionResponse)(nil), "api.GetNodeSessionResponse")
	proto.RegisterType((*GetNodeSessionByDevEUIRequest)(nil), "api.GetNodeSessionByDevEUIRequest")
	proto.RegisterType((*UpdateNodeSessionRequest)(nil), "api.UpdateNodeSessionRequest")
	proto.RegisterType((*UpdateNodeSessionResponse)(nil), "api.UpdateNodeSessionResponse")
	proto.RegisterType((*DeleteNodeSessionRequest)(nil), "api.DeleteNodeSessionRequest")
	proto.RegisterType((*DeleteNodeSessionResponse)(nil), "api.DeleteNodeSessionResponse")
	proto.RegisterType((*GetRandomDevAddrRequest)(nil), "api.GetRandomDevAddrRequest")
	proto.RegisterType((*GetRandomDevAddrResponse)(nil), "api.GetRandomDevAddrResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for NodeSession service

type NodeSessionClient interface {
	// Create creates the given node-session. The DevAddr must contain the same NwkID as the configured NetID. Node-sessions will expire automatically after the configured TTL.
	Create(ctx context.Context, in *CreateNodeSessionRequest, opts ...grpc.CallOption) (*CreateNodeSessionResponse, error)
	// Get returns the node-session matching the given DevAddr.
	Get(ctx context.Context, in *GetNodeSessionRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error)
	// GetByDevEUI returns the node-session matching the given DevEUI.
	GetByDevEUI(ctx context.Context, in *GetNodeSessionByDevEUIRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error)
	// Update updates the given node-session.
	Update(ctx context.Context, in *UpdateNodeSessionRequest, opts ...grpc.CallOption) (*UpdateNodeSessionResponse, error)
	// Delete deletes the node-session matching the given DevAddr.
	Delete(ctx context.Context, in *DeleteNodeSessionRequest, opts ...grpc.CallOption) (*DeleteNodeSessionResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into account.
	GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error)
}

type nodeSessionClient struct {
	cc *grpc.ClientConn
}

func NewNodeSessionClient(cc *grpc.ClientConn) NodeSessionClient {
	return &nodeSessionClient{cc}
}

func (c *nodeSessionClient) Create(ctx context.Context, in *CreateNodeSessionRequest, opts ...grpc.CallOption) (*CreateNodeSessionResponse, error) {
	out := new(CreateNodeSessionResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSessionClient) Get(ctx context.Context, in *GetNodeSessionRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error) {
	out := new(GetNodeSessionResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSessionClient) GetByDevEUI(ctx context.Context, in *GetNodeSessionByDevEUIRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error) {
	out := new(GetNodeSessionResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/GetByDevEUI", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSessionClient) Update(ctx context.Context, in *UpdateNodeSessionRequest, opts ...grpc.CallOption) (*UpdateNodeSessionResponse, error) {
	out := new(UpdateNodeSessionResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSessionClient) Delete(ctx context.Context, in *DeleteNodeSessionRequest, opts ...grpc.CallOption) (*DeleteNodeSessionResponse, error) {
	out := new(DeleteNodeSessionResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeSessionClient) GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error) {
	out := new(GetRandomDevAddrResponse)
	err := grpc.Invoke(ctx, "/api.NodeSession/GetRandomDevAddr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeSession service

type NodeSessionServer interface {
	// Create creates the given node-session. The DevAddr must contain the same NwkID as the configured NetID. Node-sessions will expire automatically after the configured TTL.
	Create(context.Context, *CreateNodeSessionRequest) (*CreateNodeSessionResponse, error)
	// Get returns the node-session matching the given DevAddr.
	Get(context.Context, *GetNodeSessionRequest) (*GetNodeSessionResponse, error)
	// GetByDevEUI returns the node-session matching the given DevEUI.
	GetByDevEUI(context.Context, *GetNodeSessionByDevEUIRequest) (*GetNodeSessionResponse, error)
	// Update updates the given node-session.
	Update(context.Context, *UpdateNodeSessionRequest) (*UpdateNodeSessionResponse, error)
	// Delete deletes the node-session matching the given DevAddr.
	Delete(context.Context, *DeleteNodeSessionRequest) (*DeleteNodeSessionResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into account.
	GetRandomDevAddr(context.Context, *GetRandomDevAddrRequest) (*GetRandomDevAddrResponse, error)
}

func RegisterNodeSessionServer(s *grpc.Server, srv NodeSessionServer) {
	s.RegisterService(&_NodeSession_serviceDesc, srv)
}

func _NodeSession_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).Create(ctx, req.(*CreateNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSession_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).Get(ctx, req.(*GetNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSession_GetByDevEUI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeSessionByDevEUIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).GetByDevEUI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/GetByDevEUI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).GetByDevEUI(ctx, req.(*GetNodeSessionByDevEUIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSession_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).Update(ctx, req.(*UpdateNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSession_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).Delete(ctx, req.(*DeleteNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeSession_GetRandomDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomDevAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeSessionServer).GetRandomDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NodeSession/GetRandomDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeSessionServer).GetRandomDevAddr(ctx, req.(*GetRandomDevAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeSession_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NodeSession",
	HandlerType: (*NodeSessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NodeSession_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _NodeSession_Get_Handler,
		},
		{
			MethodName: "GetByDevEUI",
			Handler:    _NodeSession_GetByDevEUI_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NodeSession_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NodeSession_Delete_Handler,
		},
		{
			MethodName: "GetRandomDevAddr",
			Handler:    _NodeSession_GetRandomDevAddr_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor4,
}

func init() { proto.RegisterFile("nodeSession.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x55, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0x5b, 0x96, 0x6d, 0xa7, 0x42, 0x02, 0x23, 0x86, 0x97, 0xad, 0xa3, 0x32, 0x0c, 0x4d,
	0x03, 0xb5, 0x2a, 0x20, 0x21, 0x71, 0x07, 0x0b, 0x20, 0x04, 0x02, 0x29, 0x53, 0x1f, 0x20, 0x10,
	0xb7, 0xaa, 0x18, 0x76, 0x96, 0x98, 0x6d, 0xd5, 0xd4, 0x1b, 0x5e, 0x81, 0x7b, 0xc4, 0x83, 0xf0,
	0x16, 0xbc, 0x02, 0x0f, 0x82, 0x7f, 0x4e, 0x58, 0x09, 0x4e, 0xa7, 0xdd, 0xef, 0xaa, 0x39, 0xfe,
	0xec, 0xef, 0xcb, 0xf9, 0xfa, 0x1d, 0x07, 0xae, 0x0b, 0x99, 0xf2, 0x7d, 0x5e, 0x14, 0x13, 0x29,
	0x7a, 0x59, 0x2e, 0x95, 0x24, 0xad, 0x24, 0x9b, 0x84, 0x9b, 0x63, 0x29, 0xc7, 0x07, 0xbc, 0xaf,
	0x9f, 0xfb, 0x89, 0x10, 0x52, 0x25, 0x4a, 0xef, 0x28, 0xdc, 0x16, 0xf6, 0xa3, 0x09, 0x74, 0x2f,
	0xe7, 0x89, 0xe2, 0xef, 0xce, 0x8e, 0xc7, 0xfc, 0xf0, 0x0b, 0x2f, 0x14, 0xa1, 0xb0, 0x9c, 0xf2,
	0xa3, 0x67, 0x69, 0x9a, 0xd3, 0x46, 0xb7, 0xb1, 0xb3, 0x1a, 0x97, 0x25, 0x59, 0x83, 0x20, 0xc9,
	0xb2, 0x17, 0xc3, 0xd7, 0xb4, 0x69, 0x01, 0xac, 0xcc, 0xba, 0xde, 0x62, 0xd6, 0x5b, 0x6e, 0xdd,
	0x55, 0x86, 0x49, 0xef, 0xd8, 0x7f, 0xc3, 0xa7, 0xf4, 0x8a, 0x63, 0xc2, 0xd2, 0x20, 0xe2, 0xf8,
	0x93, 0x45, 0x96, 0x1c, 0x82, 0xa5, 0xe1, 0x1a, 0xed, 0x09, 0x35, 0xcc, 0x68, 0xa0, 0x81, 0xab,
	0x31, 0x56, 0x24, 0x84, 0x15, 0xf3, 0x14, 0xc9, 0x63, 0x41, 0x97, 0x2d, 0xf2, 0xb7, 0x36, 0x6c,
	0xf9, 0x49, 0xc4, 0x0f, 0x92, 0x29, 0x5d, 0xb1, 0x50, 0x59, 0x92, 0x2e, 0xb4, 0xf3, 0x93, 0x41,
	0x14, 0xbf, 0x1f, 0x8d, 0x0a, 0xae, 0xe8, 0xaa, 0x45, 0xe7, 0x97, 0x8c, 0xde, 0xc7, 0x97, 0x6f,
	0x27, 0x85, 0xa2, 0xd0, 0x6d, 0x19, 0x3d, 0x57, 0xb1, 0x0d, 0x58, 0xf7, 0x38, 0x54, 0x64, 0xda,
	0x44, 0xce, 0x06, 0x70, 0xf3, 0x15, 0x57, 0x17, 0xf1, 0x8e, 0x7d, 0x6f, 0xc2, 0x5a, 0xf5, 0x8c,
	0x63, 0xbb, 0x34, 0xdc, 0x1a, 0xfe, 0x04, 0x3a, 0xff, 0xfa, 0xf3, 0x7c, 0x1a, 0xd9, 0xae, 0x4a,
	0x6f, 0xcf, 0x9a, 0x6e, 0xcc, 0x37, 0x6d, 0xc3, 0x3c, 0xcc, 0xd2, 0xcb, 0x30, 0x2f, 0x0c, 0xb3,
	0xc7, 0x21, 0x0c, 0xf3, 0x63, 0xa0, 0x9a, 0x9f, 0x5f, 0xcc, 0x3e, 0x43, 0xe9, 0x39, 0x85, 0x94,
	0xeb, 0x70, 0x4b, 0xff, 0x97, 0x71, 0x22, 0x52, 0xf9, 0x39, 0x72, 0x07, 0x90, 0xd1, 0xa8, 0xfd,
	0x0f, 0x9d, 0x37, 0x08, 0x0f, 0x7f, 0x2e, 0x41, 0x7b, 0x4e, 0x88, 0x8c, 0x21, 0x70, 0xd3, 0x49,
	0x3a, 0x3d, 0x7d, 0xc5, 0xf5, 0xea, 0x2e, 0xb3, 0x70, 0xab, 0x0e, 0xc6, 0x37, 0xdd, 0xfa, 0xfa,
	0xeb, 0xf7, 0xb7, 0x26, 0x65, 0x37, 0xec, 0x4d, 0x79, 0x34, 0xe8, 0xcf, 0xdd, 0xa7, 0x4f, 0x1b,
	0xbb, 0x84, 0x43, 0x4b, 0xbf, 0x2e, 0x09, 0x2d, 0x8d, 0x77, 0xe6, 0xc3, 0x0d, 0x2f, 0x86, 0xfc,
	0xdb, 0x96, 0xff, 0x36, 0xe9, 0x78, 0xf8, 0xfb, 0xa7, 0xd8, 0xde, 0x8c, 0x9c, 0x42, 0x5b, 0x13,
	0x94, 0x89, 0x27, 0xcc, 0x43, 0x59, 0x19, 0x87, 0xc5, 0xb2, 0xf7, 0xad, 0xec, 0x36, 0xb9, 0xe3,
	0x93, 0x75, 0x81, 0xb6, 0xea, 0xfa, 0x77, 0x46, 0x0e, 0x21, 0x70, 0xe9, 0x40, 0x33, 0xeb, 0x86,
	0x09, 0xcd, 0xac, 0x4f, 0xd2, 0x8e, 0x55, 0x65, 0xe1, 0xe2, 0x66, 0x8d, 0xad, 0x02, 0x02, 0x97,
	0x1e, 0x94, 0xac, 0x0b, 0x20, 0x4a, 0xd6, 0x27, 0x0d, 0xfd, 0xdd, 0x3d, 0xc7, 0xdf, 0x19, 0x5c,
	0xab, 0xa6, 0x8e, 0x6c, 0x96, 0x06, 0xfa, 0x72, 0x1a, 0x76, 0x6a, 0x50, 0xd4, 0x7d, 0x60, 0x75,
	0xef, 0xb1, 0xbb, 0x3e, 0xdd, 0x71, 0xe5, 0xd4, 0x87, 0xc0, 0x7e, 0x76, 0x1f, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x61, 0x59, 0xb2, 0xae, 0x07, 0x00, 0x00,
}
