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
// source: db/port.proto

package dbpb

import (
	scanner "../gen/go/scanner"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// PortStatus - Port state constants
type PortStatus int32

const (
	PortStatus_Open       PortStatus = 0
	PortStatus_Closed     PortStatus = 1
	PortStatus_Filtered   PortStatus = 2
	PortStatus_Unfiltered PortStatus = 3
)

// Enum value maps for PortStatus.
var (
	PortStatus_name = map[int32]string{
		0: "Open",
		1: "Closed",
		2: "Filtered",
		3: "Unfiltered",
	}
	PortStatus_value = map[string]int32{
		"Open":       0,
		"Closed":     1,
		"Filtered":   2,
		"Unfiltered": 3,
	}
)

func (x PortStatus) Enum() *PortStatus {
	p := new(PortStatus)
	*p = x
	return p
}

func (x PortStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_db_port_proto_enumTypes[0].Descriptor()
}

func (PortStatus) Type() protoreflect.EnumType {
	return &file_db_port_proto_enumTypes[0]
}

func (x PortStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortStatus.Descriptor instead.
func (PortStatus) EnumDescriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{0}
}

// Port - A port on a host
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// General
	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Number uint32 `protobuf:"varint,2,opt,name=Number,proto3" json:"Number,omitempty"`
	HostID uint32 `protobuf:"varint,3,opt,name=HostID,proto3" json:"HostID,omitempty"`
	// Nmap
	// @inject_tag: xml:"protocol"
	Protocol string `protobuf:"bytes,4,opt,name=Protocol,proto3" json:"Protocol,omitempty" xml:"protocol"`
	// @inject_tag: xml:"owner"
	Owner *Owner `protobuf:"bytes,5,opt,name=Owner,proto3" json:"Owner,omitempty" xml:"owner"`
	// @inject_tag: xml:"service"
	Service *Service `protobuf:"bytes,6,opt,name=Service,proto3" json:"Service,omitempty" xml:"service"`
	// @inject_tag: xml:"state"
	State *State `protobuf:"bytes,7,opt,name=State,proto3" json:"State,omitempty" xml:"state"`
	// @inject_tag: xml:"script"
	Scripts []*scanner.NmapScript `protobuf:"bytes,8,rep,name=Scripts,proto3" json:"Scripts,omitempty" xml:"script"`
	// Timestamp
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,38,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,39,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_port_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_db_port_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Port) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Port) GetHostID() uint32 {
	if x != nil {
		return x.HostID
	}
	return 0
}

func (x *Port) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Port) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Port) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Port) GetState() *State {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Port) GetScripts() []*scanner.NmapScript {
	if x != nil {
		return x.Scripts
	}
	return nil
}

func (x *Port) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Port) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// ExtraPort - Reasons for why a port is closed/filtered
type ExtraPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	HostID uint32 `protobuf:"varint,2,opt,name=HostID,proto3" json:"HostID,omitempty"`
	// @inject_tag: xml:"state"
	State string `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty" xml:"state"`
	// @inject_tag: xml:"count"
	Count int32 `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty" xml:"count"`
	// @inject_tag: xml:"extrareasons"
	Reasons []*Reason `protobuf:"bytes,5,rep,name=Reasons,proto3" json:"Reasons,omitempty" xml:"extrareasons"`
}

func (x *ExtraPort) Reset() {
	*x = ExtraPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_port_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtraPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtraPort) ProtoMessage() {}

func (x *ExtraPort) ProtoReflect() protoreflect.Message {
	mi := &file_db_port_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtraPort.ProtoReflect.Descriptor instead.
func (*ExtraPort) Descriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{1}
}

func (x *ExtraPort) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ExtraPort) GetHostID() uint32 {
	if x != nil {
		return x.HostID
	}
	return 0
}

func (x *ExtraPort) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ExtraPort) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ExtraPort) GetReasons() []*Reason {
	if x != nil {
		return x.Reasons
	}
	return nil
}

// Reason - Extra information on a closed/filtered port
type Reason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ExtraPortID uint32 `protobuf:"varint,2,opt,name=ExtraPortID,proto3" json:"ExtraPortID,omitempty"`
	// @inject_tag: xml:"reason"
	Reason string `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty" xml:"reason"`
	// @inject_tag: xml:"count"
	Count int32 `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty" xml:"count"`
}

func (x *Reason) Reset() {
	*x = Reason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_port_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reason) ProtoMessage() {}

func (x *Reason) ProtoReflect() protoreflect.Message {
	mi := &file_db_port_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reason.ProtoReflect.Descriptor instead.
func (*Reason) Descriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{2}
}

func (x *Reason) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Reason) GetExtraPortID() uint32 {
	if x != nil {
		return x.ExtraPortID
	}
	return 0
}

func (x *Reason) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Reason) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// State - Contains information about a given's port status
type State struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PortID uint32 `protobuf:"varint,2,opt,name=PortID,proto3" json:"PortID,omitempty"`
	// Nmap
	// @inject_tag: xml:"state,attr"
	State string `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty" xml:"state,attr"`
	// @inject_tag: xml:"reason,attr"
	Reason string `protobuf:"bytes,4,opt,name=Reason,proto3" json:"Reason,omitempty" xml:"reason,attr"`
	// @inject_tag: xml:"reason_ip,attr"
	ReasonIP string `protobuf:"bytes,5,opt,name=ReasonIP,proto3" json:"ReasonIP,omitempty" xml:"reason_ip,attr"`
	// @inject_tag: xml:"reason_ttl,attr"
	ReasonTTL float32 `protobuf:"fixed32,6,opt,name=ReasonTTL,proto3" json:"ReasonTTL,omitempty" xml:"reason_ttl,attr"`
	// Timestamp
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,38,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,39,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *State) Reset() {
	*x = State{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_port_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*State) ProtoMessage() {}

func (x *State) ProtoReflect() protoreflect.Message {
	mi := &file_db_port_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use State.ProtoReflect.Descriptor instead.
func (*State) Descriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{3}
}

func (x *State) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *State) GetPortID() uint32 {
	if x != nil {
		return x.PortID
	}
	return 0
}

func (x *State) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *State) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *State) GetReasonIP() string {
	if x != nil {
		return x.ReasonIP
	}
	return ""
}

func (x *State) GetReasonTTL() float32 {
	if x != nil {
		return x.ReasonTTL
	}
	return 0
}

func (x *State) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *State) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_port_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_db_port_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_db_port_proto_rawDescGZIP(), []int{4}
}

var File_db_port_proto protoreflect.FileDescriptor

var file_db_port_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x62, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x64, 0x62, 0x1a, 0x10, 0x64, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2f, 0x6e,
	0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x02, 0x0a, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x48,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x1f, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x64, 0x62, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x64, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6e, 0x65, 0x72, 0x2e, 0x4e, 0x6d, 0x61, 0x70, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x07,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x27,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x09,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x07, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x73, 0x22, 0x68, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8b, 0x02,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x54, 0x54, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x54, 0x54, 0x4c, 0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x27,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x07, 0x0a, 0x05, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x2a, 0x40, 0x0a, 0x0a, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x6e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x65, 0x64, 0x10, 0x03, 0x42, 0x06, 0x5a, 0x04, 0x64, 0x62, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_port_proto_rawDescOnce sync.Once
	file_db_port_proto_rawDescData = file_db_port_proto_rawDesc
)

func file_db_port_proto_rawDescGZIP() []byte {
	file_db_port_proto_rawDescOnce.Do(func() {
		file_db_port_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_port_proto_rawDescData)
	})
	return file_db_port_proto_rawDescData
}

var file_db_port_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_db_port_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_db_port_proto_goTypes = []interface{}{
	(PortStatus)(0),             // 0: db.PortStatus
	(*Port)(nil),                // 1: db.Port
	(*ExtraPort)(nil),           // 2: db.ExtraPort
	(*Reason)(nil),              // 3: db.Reason
	(*State)(nil),               // 4: db.State
	(*Owner)(nil),               // 5: db.Owner
	(*Service)(nil),             // 6: db.Service
	(*scanner.NmapScript)(nil),  // 7: scanner.NmapScript
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_db_port_proto_depIdxs = []int32{
	5, // 0: db.Port.Owner:type_name -> db.Owner
	6, // 1: db.Port.Service:type_name -> db.Service
	4, // 2: db.Port.State:type_name -> db.State
	7, // 3: db.Port.Scripts:type_name -> scanner.NmapScript
	8, // 4: db.Port.CreatedAt:type_name -> google.protobuf.Timestamp
	8, // 5: db.Port.UpdatedAt:type_name -> google.protobuf.Timestamp
	3, // 6: db.ExtraPort.Reasons:type_name -> db.Reason
	8, // 7: db.State.CreatedAt:type_name -> google.protobuf.Timestamp
	8, // 8: db.State.UpdatedAt:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_db_port_proto_init() }
func file_db_port_proto_init() {
	if File_db_port_proto != nil {
		return
	}
	file_db_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_db_port_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_db_port_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtraPort); i {
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
		file_db_port_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reason); i {
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
		file_db_port_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*State); i {
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
		file_db_port_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
			RawDescriptor: file_db_port_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_port_proto_goTypes,
		DependencyIndexes: file_db_port_proto_depIdxs,
		EnumInfos:         file_db_port_proto_enumTypes,
		MessageInfos:      file_db_port_proto_msgTypes,
	}.Build()
	File_db_port_proto = out.File
	file_db_port_proto_rawDesc = nil
	file_db_port_proto_goTypes = nil
	file_db_port_proto_depIdxs = nil
}
