// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noderesolver.proto

/*
Package noderesolver is a generated protocol buffer package.

It is generated from these files:
	noderesolver.proto

It has these top-level messages:
	ResolveRequest
	ResolveResponse
*/
package noderesolver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import spire_common_plugin "github.com/spiffe/spire/pkg/common/plugin"
import spire_common "github.com/spiffe/spire/pkg/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type ConfigureRequest spire_common_plugin.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*spire_common_plugin.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*spire_common_plugin.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*spire_common_plugin.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type ConfigureResponse spire_common_plugin.ConfigureResponse

func (m *ConfigureResponse) Reset() { (*spire_common_plugin.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string {
	return (*spire_common_plugin.ConfigureResponse)(m).String()
}
func (*ConfigureResponse) ProtoMessage() {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*spire_common_plugin.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type GetPluginInfoRequest spire_common_plugin.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset() { (*spire_common_plugin.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string {
	return (*spire_common_plugin.GetPluginInfoRequest)(m).String()
}
func (*GetPluginInfoRequest) ProtoMessage() {}

// GetPluginInfoResponse from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type GetPluginInfoResponse spire_common_plugin.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset() { (*spire_common_plugin.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).String()
}
func (*GetPluginInfoResponse) ProtoMessage() {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*spire_common_plugin.GetPluginInfoResponse)(m).GetCompany()
}

// PluginInfoRequest from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type PluginInfoRequest spire_common_plugin.PluginInfoRequest

func (m *PluginInfoRequest) Reset() { (*spire_common_plugin.PluginInfoRequest)(m).Reset() }
func (m *PluginInfoRequest) String() string {
	return (*spire_common_plugin.PluginInfoRequest)(m).String()
}
func (*PluginInfoRequest) ProtoMessage() {}

// PluginInfoReply from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type PluginInfoReply spire_common_plugin.PluginInfoReply

func (m *PluginInfoReply) Reset()         { (*spire_common_plugin.PluginInfoReply)(m).Reset() }
func (m *PluginInfoReply) String() string { return (*spire_common_plugin.PluginInfoReply)(m).String() }
func (*PluginInfoReply) ProtoMessage()    {}
func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	o := (*spire_common_plugin.PluginInfoReply)(m).GetPluginInfo()
	if o == nil {
		return nil
	}
	s := make([]*GetPluginInfoResponse, len(o))
	for i, x := range o {
		s[i] = (*GetPluginInfoResponse)(x)
	}
	return s
}

// StopRequest from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type StopRequest spire_common_plugin.StopRequest

func (m *StopRequest) Reset()         { (*spire_common_plugin.StopRequest)(m).Reset() }
func (m *StopRequest) String() string { return (*spire_common_plugin.StopRequest)(m).String() }
func (*StopRequest) ProtoMessage()    {}

// StopReply from public import github.com/spiffe/spire/pkg/common/plugin/plugin.proto
type StopReply spire_common_plugin.StopReply

func (m *StopReply) Reset()         { (*spire_common_plugin.StopReply)(m).Reset() }
func (m *StopReply) String() string { return (*spire_common_plugin.StopReply)(m).String() }
func (*StopReply) ProtoMessage()    {}

// Empty from public import github.com/spiffe/spire/pkg/common/common.proto
type Empty spire_common.Empty

func (m *Empty) Reset()         { (*spire_common.Empty)(m).Reset() }
func (m *Empty) String() string { return (*spire_common.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// AttestedData from public import github.com/spiffe/spire/pkg/common/common.proto
type AttestedData spire_common.AttestedData

func (m *AttestedData) Reset()          { (*spire_common.AttestedData)(m).Reset() }
func (m *AttestedData) String() string  { return (*spire_common.AttestedData)(m).String() }
func (*AttestedData) ProtoMessage()     {}
func (m *AttestedData) GetType() string { return (*spire_common.AttestedData)(m).GetType() }
func (m *AttestedData) GetData() []byte { return (*spire_common.AttestedData)(m).GetData() }

// Selector from public import github.com/spiffe/spire/pkg/common/common.proto
type Selector spire_common.Selector

func (m *Selector) Reset()           { (*spire_common.Selector)(m).Reset() }
func (m *Selector) String() string   { return (*spire_common.Selector)(m).String() }
func (*Selector) ProtoMessage()      {}
func (m *Selector) GetType() string  { return (*spire_common.Selector)(m).GetType() }
func (m *Selector) GetValue() string { return (*spire_common.Selector)(m).GetValue() }

// Selectors from public import github.com/spiffe/spire/pkg/common/common.proto
type Selectors spire_common.Selectors

func (m *Selectors) Reset()         { (*spire_common.Selectors)(m).Reset() }
func (m *Selectors) String() string { return (*spire_common.Selectors)(m).String() }
func (*Selectors) ProtoMessage()    {}
func (m *Selectors) GetEntries() []*Selector {
	o := (*spire_common.Selectors)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}

// RegistrationEntry from public import github.com/spiffe/spire/pkg/common/common.proto
type RegistrationEntry spire_common.RegistrationEntry

func (m *RegistrationEntry) Reset()         { (*spire_common.RegistrationEntry)(m).Reset() }
func (m *RegistrationEntry) String() string { return (*spire_common.RegistrationEntry)(m).String() }
func (*RegistrationEntry) ProtoMessage()    {}
func (m *RegistrationEntry) GetSelectors() []*Selector {
	o := (*spire_common.RegistrationEntry)(m).GetSelectors()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}
func (m *RegistrationEntry) GetParentId() string {
	return (*spire_common.RegistrationEntry)(m).GetParentId()
}
func (m *RegistrationEntry) GetSpiffeId() string {
	return (*spire_common.RegistrationEntry)(m).GetSpiffeId()
}
func (m *RegistrationEntry) GetTtl() int32 { return (*spire_common.RegistrationEntry)(m).GetTtl() }
func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	return (*spire_common.RegistrationEntry)(m).GetFbSpiffeIds()
}

// RegistrationEntries from public import github.com/spiffe/spire/pkg/common/common.proto
type RegistrationEntries spire_common.RegistrationEntries

func (m *RegistrationEntries) Reset()         { (*spire_common.RegistrationEntries)(m).Reset() }
func (m *RegistrationEntries) String() string { return (*spire_common.RegistrationEntries)(m).String() }
func (*RegistrationEntries) ProtoMessage()    {}
func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	o := (*spire_common.RegistrationEntries)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*RegistrationEntry, len(o))
	for i, x := range o {
		s[i] = (*RegistrationEntry)(x)
	}
	return s
}

// * Represents a request with a list of BaseSPIFFEIDs.
type ResolveRequest struct {
	// * A list of BaseSPIFFE Ids.
	BaseSpiffeIdList []string `protobuf:"bytes,1,rep,name=baseSpiffeIdList" json:"baseSpiffeIdList,omitempty"`
}

func (m *ResolveRequest) Reset()                    { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string            { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()               {}
func (*ResolveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResolveRequest) GetBaseSpiffeIdList() []string {
	if m != nil {
		return m.BaseSpiffeIdList
	}
	return nil
}

// * Represents a response with a map of SPIFFE ID to a list of Selectors.
type ResolveResponse struct {
	// * Map[SPIFFE_ID] => Selectors.
	Map map[string]*spire_common.Selectors `protobuf:"bytes,1,rep,name=map" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ResolveResponse) Reset()                    { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string            { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()               {}
func (*ResolveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResolveResponse) GetMap() map[string]*spire_common.Selectors {
	if m != nil {
		return m.Map
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveRequest)(nil), "spire.server.noderesolver.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "spire.server.noderesolver.ResolveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeResolver service

type NodeResolverClient interface {
	// * Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
	// * Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *spire_common_plugin.ConfigureRequest, opts ...grpc.CallOption) (*spire_common_plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *spire_common_plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*spire_common_plugin.GetPluginInfoResponse, error)
}

type nodeResolverClient struct {
	cc *grpc.ClientConn
}

func NewNodeResolverClient(cc *grpc.ClientConn) NodeResolverClient {
	return &nodeResolverClient{cc}
}

func (c *nodeResolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := grpc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/Resolve", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) Configure(ctx context.Context, in *spire_common_plugin.ConfigureRequest, opts ...grpc.CallOption) (*spire_common_plugin.ConfigureResponse, error) {
	out := new(spire_common_plugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) GetPluginInfo(ctx context.Context, in *spire_common_plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*spire_common_plugin.GetPluginInfoResponse, error) {
	out := new(spire_common_plugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeResolver service

type NodeResolverServer interface {
	// * Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
	// * Responsible for configuration of the plugin.
	Configure(context.Context, *spire_common_plugin.ConfigureRequest) (*spire_common_plugin.ConfigureResponse, error)
	// * Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *spire_common_plugin.GetPluginInfoRequest) (*spire_common_plugin.GetPluginInfoResponse, error)
}

func RegisterNodeResolverServer(s *grpc.Server, srv NodeResolverServer) {
	s.RegisterService(&_NodeResolver_serviceDesc, srv)
}

func _NodeResolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common_plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Configure(ctx, req.(*spire_common_plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(spire_common_plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, req.(*spire_common_plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeResolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.noderesolver.NodeResolver",
	HandlerType: (*NodeResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _NodeResolver_Resolve_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeResolver_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeResolver_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noderesolver.proto",
}

func init() { proto.RegisterFile("noderesolver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x5d, 0x88, 0x7f, 0x58, 0x10, 0xc9, 0x5e, 0xc4, 0x9e, 0x08, 0x89, 0x06, 0x49, 0x68,
	0x13, 0xf0, 0x60, 0x8c, 0x27, 0x0d, 0x31, 0x24, 0xfe, 0x21, 0xe5, 0xc6, 0xc9, 0x02, 0xd3, 0xba,
	0xa1, 0xed, 0xac, 0xbb, 0x5b, 0x12, 0x3e, 0x92, 0x77, 0x3f, 0xa0, 0xa1, 0x5d, 0x08, 0x60, 0x14,
	0x4e, 0xbb, 0x99, 0x79, 0xbf, 0x79, 0xf3, 0xda, 0xa5, 0x2c, 0xc6, 0x09, 0x48, 0x50, 0x18, 0xce,
	0x40, 0xda, 0x42, 0xa2, 0x46, 0x76, 0xa1, 0x04, 0x97, 0x60, 0x2b, 0x90, 0x8b, 0xda, 0xba, 0xc0,
	0xba, 0x09, 0xb8, 0xfe, 0x48, 0x46, 0xf6, 0x18, 0x23, 0x47, 0x09, 0xee, 0xfb, 0xe0, 0x28, 0xc9,
	0x1d, 0x31, 0x0d, 0x9c, 0x31, 0x46, 0x11, 0xc6, 0x8e, 0x08, 0x93, 0x80, 0x2f, 0x8f, 0x6c, 0xa0,
	0xd5, 0xda, 0x49, 0x65, 0x47, 0x26, 0xaf, 0xdf, 0xd3, 0xb2, 0x9b, 0x19, 0xba, 0xf0, 0x99, 0x80,
	0xd2, 0xac, 0x49, 0x2b, 0x23, 0x4f, 0xc1, 0x20, 0x65, 0x7b, 0x93, 0x67, 0xae, 0x74, 0x95, 0xd4,
	0xf2, 0x8d, 0x82, 0xfb, 0xab, 0x5e, 0xff, 0x22, 0xf4, 0x6c, 0x85, 0x2b, 0x81, 0xb1, 0x02, 0xd6,
	0xa5, 0xf9, 0xc8, 0x13, 0x29, 0x52, 0x6c, 0x77, 0xec, 0x3f, 0xf3, 0xd9, 0x5b, 0xa0, 0xfd, 0xe2,
	0x89, 0x6e, 0xac, 0xe5, 0xdc, 0x5d, 0xf0, 0xd6, 0x1b, 0x3d, 0x59, 0x16, 0x58, 0x85, 0xe6, 0xa7,
	0x30, 0xaf, 0x92, 0x1a, 0x69, 0x14, 0xdc, 0xc5, 0x95, 0xb5, 0xe8, 0xe1, 0xcc, 0x0b, 0x13, 0xa8,
	0xe6, 0x6a, 0xa4, 0x51, 0x6c, 0x9f, 0x1b, 0x1b, 0x13, 0x6d, 0x00, 0x21, 0x8c, 0x35, 0x4a, 0xe5,
	0x66, 0xaa, 0xbb, 0xdc, 0x2d, 0x69, 0x7f, 0xe7, 0x68, 0xe9, 0x15, 0x27, 0x60, 0x6c, 0x25, 0x7b,
	0xa7, 0xc7, 0xe6, 0xce, 0xae, 0xf7, 0x59, 0x33, 0xfd, 0x3c, 0x56, 0x73, 0xff, 0x44, 0x6c, 0x48,
	0x0b, 0x8f, 0x18, 0xfb, 0x3c, 0x48, 0x24, 0xb0, 0xcb, 0xcd, 0x1d, 0xcd, 0x4f, 0x5b, 0xf5, 0x97,
	0xf3, 0xaf, 0x76, 0xc9, 0xcc, 0x6c, 0x9f, 0x9e, 0x3e, 0x81, 0xee, 0xa7, 0xed, 0x5e, 0xec, 0xe3,
	0x2a, 0xc3, 0x26, 0xb8, 0xa1, 0xd9, 0xce, 0xf0, 0xaf, 0x34, 0xf3, 0x79, 0x28, 0x0f, 0x4b, 0xeb,
	0x19, 0xfb, 0x07, 0x7d, 0x32, 0x3a, 0x4a, 0xdf, 0x4e, 0xe7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xe4, 0x0b, 0xe9, 0xd1, 0x02, 0x00, 0x00,
}
