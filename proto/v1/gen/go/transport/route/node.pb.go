// Wiregost - Golang Exploitation Framework
// Copyright © 2020 Para
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.8.0
// source: transport/route/node.proto

package routepb

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Node - A route node, member of a proxy chain.
// Used by implants to send back active routes.
//          DB to save current routes
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int32         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	GhostID     uint32        `protobuf:"varint,2,opt,name=GhostID,proto3" json:"GhostID,omitempty"`      // Not saved to DB, only for live stuff
	GhostName   string        `protobuf:"bytes,3,opt,name=GhostName,proto3" json:"GhostName,omitempty"`   // Each node in a route is served by a ghost implant process, we just use the name
	LocalAddr   string        `protobuf:"bytes,4,opt,name=LocalAddr,proto3" json:"LocalAddr,omitempty"`   // host:port on which the router listens
	RemoteAddr  string        `protobuf:"bytes,5,opt,name=RemoteAddr,proto3" json:"RemoteAddr,omitempty"` // Used for port forwarding
	Host        string        `protobuf:"bytes,6,opt,name=Host,proto3" json:"Host,omitempty"`
	Transport   string        `protobuf:"bytes,7,opt,name=Transport,proto3" json:"Transport,omitempty"`
	ActiveConns []*Connection `protobuf:"bytes,10,rep,name=ActiveConns,proto3" json:"ActiveConns,omitempty"` // Connections currently passing through the node
	Pivots      []string      `protobuf:"bytes,11,rep,name=Pivots,proto3" json:"Pivots,omitempty"`           // How many ghost are pivoted through this node: 1 ghost = 1 connection
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_route_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_transport_route_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_transport_route_node_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Node) GetGhostID() uint32 {
	if x != nil {
		return x.GhostID
	}
	return 0
}

func (x *Node) GetGhostName() string {
	if x != nil {
		return x.GhostName
	}
	return ""
}

func (x *Node) GetLocalAddr() string {
	if x != nil {
		return x.LocalAddr
	}
	return ""
}

func (x *Node) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *Node) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Node) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

func (x *Node) GetActiveConns() []*Connection {
	if x != nil {
		return x.ActiveConns
	}
	return nil
}

func (x *Node) GetPivots() []string {
	if x != nil {
		return x.Pivots
	}
	return nil
}

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalAddr  string `protobuf:"bytes,3,opt,name=LocalAddr,proto3" json:"LocalAddr,omitempty"`
	RemoteAddr string `protobuf:"bytes,4,opt,name=RemoteAddr,proto3" json:"RemoteAddr,omitempty"`
	Transport  string `protobuf:"bytes,6,opt,name=Transport,proto3" json:"Transport,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_route_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_transport_route_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_transport_route_node_proto_rawDescGZIP(), []int{1}
}

func (x *Connection) GetLocalAddr() string {
	if x != nil {
		return x.LocalAddr
	}
	return ""
}

func (x *Connection) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *Connection) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

var File_transport_route_node_proto protoreflect.FileDescriptor

var file_transport_route_node_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x95, 0x02,
	0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3d,
	0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x69, 0x76, 0x6f, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x69, 0x76, 0x6f, 0x74, 0x73, 0x22, 0x68, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x42,
	0x09, 0x5a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_transport_route_node_proto_rawDescOnce sync.Once
	file_transport_route_node_proto_rawDescData = file_transport_route_node_proto_rawDesc
)

func file_transport_route_node_proto_rawDescGZIP() []byte {
	file_transport_route_node_proto_rawDescOnce.Do(func() {
		file_transport_route_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_route_node_proto_rawDescData)
	})
	return file_transport_route_node_proto_rawDescData
}

var file_transport_route_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_route_node_proto_goTypes = []interface{}{
	(*Node)(nil),       // 0: transport.route.Node
	(*Connection)(nil), // 1: transport.route.Connection
}
var file_transport_route_node_proto_depIdxs = []int32{
	1, // 0: transport.route.Node.ActiveConns:type_name -> transport.route.Connection
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transport_route_node_proto_init() }
func file_transport_route_node_proto_init() {
	if File_transport_route_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_route_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transport_route_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_route_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_route_node_proto_goTypes,
		DependencyIndexes: file_transport_route_node_proto_depIdxs,
		MessageInfos:      file_transport_route_node_proto_msgTypes,
	}.Build()
	File_transport_route_node_proto = out.File
	file_transport_route_node_proto_rawDesc = nil
	file_transport_route_node_proto_goTypes = nil
	file_transport_route_node_proto_depIdxs = nil
}
