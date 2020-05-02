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
// source: ghost/core/shell.proto

package corepb

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

type ShellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GhostID   uint32 `protobuf:"varint,1,opt,name=GhostID,proto3" json:"GhostID,omitempty"`
	Path      string `protobuf:"bytes,2,opt,name=Path,proto3" json:"Path,omitempty"`
	EnablePTY bool   `protobuf:"varint,3,opt,name=EnablePTY,proto3" json:"EnablePTY,omitempty"`
	TunnelID  uint64 `protobuf:"varint,9,opt,name=TunnelID,proto3" json:"TunnelID,omitempty"`
}

func (x *ShellRequest) Reset() {
	*x = ShellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_core_shell_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellRequest) ProtoMessage() {}

func (x *ShellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_core_shell_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellRequest.ProtoReflect.Descriptor instead.
func (*ShellRequest) Descriptor() ([]byte, []int) {
	return file_ghost_core_shell_proto_rawDescGZIP(), []int{0}
}

func (x *ShellRequest) GetGhostID() uint32 {
	if x != nil {
		return x.GhostID
	}
	return 0
}

func (x *ShellRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ShellRequest) GetEnablePTY() bool {
	if x != nil {
		return x.EnablePTY
	}
	return false
}

func (x *ShellRequest) GetTunnelID() uint64 {
	if x != nil {
		return x.TunnelID
	}
	return 0
}

type Shell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Err      string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	TunnelID uint64 `protobuf:"varint,9,opt,name=TunnelID,proto3" json:"TunnelID,omitempty"`
}

func (x *Shell) Reset() {
	*x = Shell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ghost_core_shell_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shell) ProtoMessage() {}

func (x *Shell) ProtoReflect() protoreflect.Message {
	mi := &file_ghost_core_shell_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shell.ProtoReflect.Descriptor instead.
func (*Shell) Descriptor() ([]byte, []int) {
	return file_ghost_core_shell_proto_rawDescGZIP(), []int{1}
}

func (x *Shell) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Shell) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *Shell) GetTunnelID() uint64 {
	if x != nil {
		return x.TunnelID
	}
	return 0
}

var File_ghost_core_shell_proto protoreflect.FileDescriptor

var file_ghost_core_shell_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x68, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0x76, 0x0a, 0x0c, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x54, 0x59, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x54, 0x59,
	0x12, 0x1a, 0x0a, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x05,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x42, 0x08, 0x5a,
	0x06, 0x63, 0x6f, 0x72, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ghost_core_shell_proto_rawDescOnce sync.Once
	file_ghost_core_shell_proto_rawDescData = file_ghost_core_shell_proto_rawDesc
)

func file_ghost_core_shell_proto_rawDescGZIP() []byte {
	file_ghost_core_shell_proto_rawDescOnce.Do(func() {
		file_ghost_core_shell_proto_rawDescData = protoimpl.X.CompressGZIP(file_ghost_core_shell_proto_rawDescData)
	})
	return file_ghost_core_shell_proto_rawDescData
}

var file_ghost_core_shell_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ghost_core_shell_proto_goTypes = []interface{}{
	(*ShellRequest)(nil), // 0: ghost.core.ShellRequest
	(*Shell)(nil),        // 1: ghost.core.Shell
}
var file_ghost_core_shell_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ghost_core_shell_proto_init() }
func file_ghost_core_shell_proto_init() {
	if File_ghost_core_shell_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ghost_core_shell_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellRequest); i {
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
		file_ghost_core_shell_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shell); i {
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
			RawDescriptor: file_ghost_core_shell_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ghost_core_shell_proto_goTypes,
		DependencyIndexes: file_ghost_core_shell_proto_depIdxs,
		MessageInfos:      file_ghost_core_shell_proto_msgTypes,
	}.Build()
	File_ghost_core_shell_proto = out.File
	file_ghost_core_shell_proto_rawDesc = nil
	file_ghost_core_shell_proto_goTypes = nil
	file_ghost_core_shell_proto_depIdxs = nil
}
