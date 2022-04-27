// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/hexagon.proto

package hexcli

import (
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

type Direction int32

const (
	Direction_N  Direction = 0
	Direction_NE Direction = 1
	Direction_E  Direction = 2
	Direction_SE Direction = 3
	Direction_S  Direction = 4
	Direction_SW Direction = 5
	Direction_W  Direction = 6
	Direction_NW Direction = 7
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "N",
		1: "NE",
		2: "E",
		3: "SE",
		4: "S",
		5: "SW",
		6: "W",
		7: "NW",
	}
	Direction_value = map[string]int32{
		"N":  0,
		"NE": 1,
		"E":  2,
		"SE": 3,
		"S":  4,
		"SW": 5,
		"W":  6,
		"NW": 7,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_api_hexagon_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_api_hexagon_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{0}
}

type HexLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X         int64     `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	Y         int64     `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z         int64     `protobuf:"varint,3,opt,name=Z,proto3" json:"Z,omitempty"`
	Direction Direction `protobuf:"varint,4,opt,name=Direction,proto3,enum=hexcloud.Direction" json:"Direction,omitempty"`
	HexID     string    `protobuf:"bytes,5,opt,name=HexID,proto3" json:"HexID,omitempty"`
}

func (x *HexLocation) Reset() {
	*x = HexLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocation) ProtoMessage() {}

func (x *HexLocation) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocation.ProtoReflect.Descriptor instead.
func (*HexLocation) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{1}
}

func (x *HexLocation) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *HexLocation) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *HexLocation) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *HexLocation) GetDirection() Direction {
	if x != nil {
		return x.Direction
	}
	return Direction_N
}

func (x *HexLocation) GetHexID() string {
	if x != nil {
		return x.HexID
	}
	return ""
}

type HexInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Exits uint32 `protobuf:"varint,2,opt,name=Exits,proto3" json:"Exits,omitempty"`
}

func (x *HexInfo) Reset() {
	*x = HexInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexInfo) ProtoMessage() {}

func (x *HexInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexInfo.ProtoReflect.Descriptor instead.
func (*HexInfo) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{2}
}

func (x *HexInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *HexInfo) GetExits() uint32 {
	if x != nil {
		return x.Exits
	}
	return 0
}

type HexInfoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexInfo []*HexInfo `protobuf:"bytes,1,rep,name=hexInfo,proto3" json:"hexInfo,omitempty"`
}

func (x *HexInfoList) Reset() {
	*x = HexInfoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexInfoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexInfoList) ProtoMessage() {}

func (x *HexInfoList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexInfoList.ProtoReflect.Descriptor instead.
func (*HexInfoList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{3}
}

func (x *HexInfoList) GetHexInfo() []*HexInfo {
	if x != nil {
		return x.HexInfo
	}
	return nil
}

type HexLocationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexLoc []*HexLocation `protobuf:"bytes,1,rep,name=hexLoc,proto3" json:"hexLoc,omitempty"`
}

func (x *HexLocationList) Reset() {
	*x = HexLocationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexLocationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexLocationList) ProtoMessage() {}

func (x *HexLocationList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexLocationList.ProtoReflect.Descriptor instead.
func (*HexLocationList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{4}
}

func (x *HexLocationList) GetHexLoc() []*HexLocation {
	if x != nil {
		return x.HexLoc
	}
	return nil
}

type HexagonGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexLoc *HexLocation `protobuf:"bytes,1,opt,name=hexLoc,proto3" json:"hexLoc,omitempty"`
	Radius int64        `protobuf:"varint,2,opt,name=radius,proto3" json:"radius,omitempty"`
	Fill   bool         `protobuf:"varint,3,opt,name=fill,proto3" json:"fill,omitempty"`
}

func (x *HexagonGetRequest) Reset() {
	*x = HexagonGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexagonGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexagonGetRequest) ProtoMessage() {}

func (x *HexagonGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexagonGetRequest.ProtoReflect.Descriptor instead.
func (*HexagonGetRequest) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{5}
}

func (x *HexagonGetRequest) GetHexLoc() *HexLocation {
	if x != nil {
		return x.HexLoc
	}
	return nil
}

func (x *HexagonGetRequest) GetRadius() int64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

func (x *HexagonGetRequest) GetFill() bool {
	if x != nil {
		return x.Fill
	}
	return false
}

type HexIDList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HexID []string `protobuf:"bytes,1,rep,name=HexID,proto3" json:"HexID,omitempty"`
}

func (x *HexIDList) Reset() {
	*x = HexIDList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HexIDList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HexIDList) ProtoMessage() {}

func (x *HexIDList) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HexIDList.ProtoReflect.Descriptor instead.
func (*HexIDList) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{6}
}

func (x *HexIDList) GetHexID() []string {
	if x != nil {
		return x.HexID
	}
	return nil
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{7}
}

func (x *Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_hexagon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_hexagon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_hexagon_proto_rawDescGZIP(), []int{8}
}

func (x *Result) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_hexagon_proto protoreflect.FileDescriptor

var file_api_hexagon_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x78, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x01, 0x58, 0x12, 0x0c, 0x0a, 0x01, 0x59, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x59, 0x12, 0x0c, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x5a,
	0x12, 0x31, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x22, 0x2f, 0x0a, 0x07, 0x48, 0x65, 0x78,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x78, 0x69, 0x74, 0x73, 0x22, 0x3a, 0x0a, 0x0b, 0x48, 0x65,
	0x78, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x68, 0x65, 0x78,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x65, 0x78,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x68,
	0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x40, 0x0a, 0x0f, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x78,
	0x4c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x65, 0x78, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x22, 0x6e, 0x0a, 0x11, 0x48, 0x65, 0x78, 0x61,
	0x67, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a,
	0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x68, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61,
	0x64, 0x69, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x6c, 0x22, 0x21, 0x0a, 0x09, 0x48, 0x65, 0x78, 0x49,
	0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x48, 0x65, 0x78, 0x49, 0x44, 0x22, 0x1a, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x47, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x05, 0x0a, 0x01, 0x4e, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x45, 0x10, 0x02, 0x12, 0x06,
	0x0a, 0x02, 0x53, 0x45, 0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x53, 0x10, 0x04, 0x12, 0x06, 0x0a,
	0x02, 0x53, 0x57, 0x10, 0x05, 0x12, 0x05, 0x0a, 0x01, 0x57, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02,
	0x4e, 0x57, 0x10, 0x07, 0x32, 0xe2, 0x04, 0x0a, 0x0e, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x41,
	0x64, 0x64, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e,
	0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x44, 0x65,
	0x6c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x68,
	0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x40, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x47, 0x65, 0x74, 0x48, 0x65,
	0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x78, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e, 0x66,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6f, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x48, 0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f,
	0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x49, 0x6e,
	0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64,
	0x12, 0x15, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x4d, 0x61, 0x70,
	0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48,
	0x65, 0x78, 0x61, 0x67, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x48, 0x65, 0x78, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12,
	0x0f, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0f, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x68, 0x65, 0x78, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x68, 0x65,
	0x78, 0x63, 0x6c, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_hexagon_proto_rawDescOnce sync.Once
	file_api_hexagon_proto_rawDescData = file_api_hexagon_proto_rawDesc
)

func file_api_hexagon_proto_rawDescGZIP() []byte {
	file_api_hexagon_proto_rawDescOnce.Do(func() {
		file_api_hexagon_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_hexagon_proto_rawDescData)
	})
	return file_api_hexagon_proto_rawDescData
}

var file_api_hexagon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_hexagon_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_hexagon_proto_goTypes = []interface{}{
	(Direction)(0),            // 0: hexcloud.Direction
	(*Empty)(nil),             // 1: hexcloud.Empty
	(*HexLocation)(nil),       // 2: hexcloud.HexLocation
	(*HexInfo)(nil),           // 3: hexcloud.HexInfo
	(*HexInfoList)(nil),       // 4: hexcloud.HexInfoList
	(*HexLocationList)(nil),   // 5: hexcloud.HexLocationList
	(*HexagonGetRequest)(nil), // 6: hexcloud.HexagonGetRequest
	(*HexIDList)(nil),         // 7: hexcloud.HexIDList
	(*Status)(nil),            // 8: hexcloud.Status
	(*Result)(nil),            // 9: hexcloud.Result
}
var file_api_hexagon_proto_depIdxs = []int32{
	0,  // 0: hexcloud.HexLocation.Direction:type_name -> hexcloud.Direction
	3,  // 1: hexcloud.HexInfoList.hexInfo:type_name -> hexcloud.HexInfo
	2,  // 2: hexcloud.HexLocationList.hexLoc:type_name -> hexcloud.HexLocation
	2,  // 3: hexcloud.HexagonGetRequest.hexLoc:type_name -> hexcloud.HexLocation
	4,  // 4: hexcloud.HexagonService.RepoAddHexagonInfo:input_type -> hexcloud.HexInfoList
	7,  // 5: hexcloud.HexagonService.RepoDelHexagonInfo:input_type -> hexcloud.HexIDList
	7,  // 6: hexcloud.HexagonService.RepoGetHexagonInfo:input_type -> hexcloud.HexIDList
	1,  // 7: hexcloud.HexagonService.RepoGetAllHexagonInfo:input_type -> hexcloud.Empty
	2,  // 8: hexcloud.HexagonService.MapAdd:input_type -> hexcloud.HexLocation
	6,  // 9: hexcloud.HexagonService.MapGet:input_type -> hexcloud.HexagonGetRequest
	5,  // 10: hexcloud.HexagonService.MapRemove:input_type -> hexcloud.HexLocationList
	1,  // 11: hexcloud.HexagonService.GetStatusServer:input_type -> hexcloud.Empty
	1,  // 12: hexcloud.HexagonService.GetStatusStorage:input_type -> hexcloud.Empty
	1,  // 13: hexcloud.HexagonService.GetStatusClients:input_type -> hexcloud.Empty
	9,  // 14: hexcloud.HexagonService.RepoAddHexagonInfo:output_type -> hexcloud.Result
	9,  // 15: hexcloud.HexagonService.RepoDelHexagonInfo:output_type -> hexcloud.Result
	4,  // 16: hexcloud.HexagonService.RepoGetHexagonInfo:output_type -> hexcloud.HexInfoList
	4,  // 17: hexcloud.HexagonService.RepoGetAllHexagonInfo:output_type -> hexcloud.HexInfoList
	9,  // 18: hexcloud.HexagonService.MapAdd:output_type -> hexcloud.Result
	5,  // 19: hexcloud.HexagonService.MapGet:output_type -> hexcloud.HexLocationList
	9,  // 20: hexcloud.HexagonService.MapRemove:output_type -> hexcloud.Result
	8,  // 21: hexcloud.HexagonService.GetStatusServer:output_type -> hexcloud.Status
	8,  // 22: hexcloud.HexagonService.GetStatusStorage:output_type -> hexcloud.Status
	8,  // 23: hexcloud.HexagonService.GetStatusClients:output_type -> hexcloud.Status
	14, // [14:24] is the sub-list for method output_type
	4,  // [4:14] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_hexagon_proto_init() }
func file_api_hexagon_proto_init() {
	if File_api_hexagon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_hexagon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_hexagon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocation); i {
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
		file_api_hexagon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexInfo); i {
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
		file_api_hexagon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexInfoList); i {
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
		file_api_hexagon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexLocationList); i {
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
		file_api_hexagon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexagonGetRequest); i {
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
		file_api_hexagon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HexIDList); i {
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
		file_api_hexagon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_api_hexagon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_api_hexagon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_hexagon_proto_goTypes,
		DependencyIndexes: file_api_hexagon_proto_depIdxs,
		EnumInfos:         file_api_hexagon_proto_enumTypes,
		MessageInfos:      file_api_hexagon_proto_msgTypes,
	}.Build()
	File_api_hexagon_proto = out.File
	file_api_hexagon_proto_rawDesc = nil
	file_api_hexagon_proto_goTypes = nil
	file_api_hexagon_proto_depIdxs = nil
}
