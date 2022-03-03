// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: echo.proto

package echopb

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

type Echo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Reminder string `protobuf:"bytes,3,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *Echo) Reset() {
	*x = Echo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Echo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echo) ProtoMessage() {}

func (x *Echo) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echo.ProtoReflect.Descriptor instead.
func (*Echo) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Echo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Echo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Echo) GetReminder() string {
	if x != nil {
		return x.Reminder
	}
	return ""
}

type PostEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Echo `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *PostEchoRequest) Reset() {
	*x = PostEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEchoRequest) ProtoMessage() {}

func (x *PostEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEchoRequest.ProtoReflect.Descriptor instead.
func (*PostEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{1}
}

func (x *PostEchoRequest) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

type PostEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Echo  `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *PostEchoResponse) Reset() {
	*x = PostEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostEchoResponse) ProtoMessage() {}

func (x *PostEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostEchoResponse.ProtoReflect.Descriptor instead.
func (*PostEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{2}
}

func (x *PostEchoResponse) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

func (x *PostEchoResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type GetEchoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEchoListRequest) Reset() {
	*x = GetEchoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEchoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEchoListRequest) ProtoMessage() {}

func (x *GetEchoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEchoListRequest.ProtoReflect.Descriptor instead.
func (*GetEchoListRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{3}
}

type GetEchoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echos []*Echo `protobuf:"bytes,1,rep,name=echos,proto3" json:"echos,omitempty"`
	Err   string  `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetEchoListResponse) Reset() {
	*x = GetEchoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEchoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEchoListResponse) ProtoMessage() {}

func (x *GetEchoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEchoListResponse.ProtoReflect.Descriptor instead.
func (*GetEchoListResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{4}
}

func (x *GetEchoListResponse) GetEchos() []*Echo {
	if x != nil {
		return x.Echos
	}
	return nil
}

func (x *GetEchoListResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type GetEchoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEchoByIDRequest) Reset() {
	*x = GetEchoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEchoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEchoByIDRequest) ProtoMessage() {}

func (x *GetEchoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEchoByIDRequest.ProtoReflect.Descriptor instead.
func (*GetEchoByIDRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{5}
}

func (x *GetEchoByIDRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEchoByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Echo  `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetEchoByIDResponse) Reset() {
	*x = GetEchoByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEchoByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEchoByIDResponse) ProtoMessage() {}

func (x *GetEchoByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEchoByIDResponse.ProtoReflect.Descriptor instead.
func (*GetEchoByIDResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{6}
}

func (x *GetEchoByIDResponse) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

func (x *GetEchoByIDResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type PutEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Echo `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
}

func (x *PutEchoRequest) Reset() {
	*x = PutEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEchoRequest) ProtoMessage() {}

func (x *PutEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEchoRequest.ProtoReflect.Descriptor instead.
func (*PutEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{7}
}

func (x *PutEchoRequest) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

type PutEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Echo *Echo  `protobuf:"bytes,1,opt,name=echo,proto3" json:"echo,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *PutEchoResponse) Reset() {
	*x = PutEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutEchoResponse) ProtoMessage() {}

func (x *PutEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutEchoResponse.ProtoReflect.Descriptor instead.
func (*PutEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{8}
}

func (x *PutEchoResponse) GetEcho() *Echo {
	if x != nil {
		return x.Echo
	}
	return nil
}

func (x *PutEchoResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type DeleteEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteEchoRequest) Reset() {
	*x = DeleteEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEchoRequest) ProtoMessage() {}

func (x *DeleteEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEchoRequest.ProtoReflect.Descriptor instead.
func (*DeleteEchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteEchoRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *DeleteEchoResponse) Reset() {
	*x = DeleteEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEchoResponse) ProtoMessage() {}

func (x *DeleteEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEchoResponse.ProtoReflect.Descriptor instead.
func (*DeleteEchoResponse) Descriptor() ([]byte, []int) {
	return file_echo_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteEchoResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_echo_proto protoreflect.FileDescriptor

var file_echo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x63,
	0x68, 0x6f, 0x70, 0x62, 0x22, 0x48, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x33,
	0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x22, 0x46, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x14, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x4b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x65, 0x63, 0x68, 0x6f,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x05, 0x65, 0x63, 0x68, 0x6f, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x24,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f,
	0x70, 0x62, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22,
	0x32, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x04, 0x65,
	0x63, 0x68, 0x6f, 0x22, 0x45, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x11, 0x5a, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_echo_proto_rawDescOnce sync.Once
	file_echo_proto_rawDescData = file_echo_proto_rawDesc
)

func file_echo_proto_rawDescGZIP() []byte {
	file_echo_proto_rawDescOnce.Do(func() {
		file_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_proto_rawDescData)
	})
	return file_echo_proto_rawDescData
}

var file_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_echo_proto_goTypes = []interface{}{
	(*Echo)(nil),                // 0: echopb.Echo
	(*PostEchoRequest)(nil),     // 1: echopb.PostEchoRequest
	(*PostEchoResponse)(nil),    // 2: echopb.PostEchoResponse
	(*GetEchoListRequest)(nil),  // 3: echopb.GetEchoListRequest
	(*GetEchoListResponse)(nil), // 4: echopb.GetEchoListResponse
	(*GetEchoByIDRequest)(nil),  // 5: echopb.GetEchoByIDRequest
	(*GetEchoByIDResponse)(nil), // 6: echopb.GetEchoByIDResponse
	(*PutEchoRequest)(nil),      // 7: echopb.PutEchoRequest
	(*PutEchoResponse)(nil),     // 8: echopb.PutEchoResponse
	(*DeleteEchoRequest)(nil),   // 9: echopb.DeleteEchoRequest
	(*DeleteEchoResponse)(nil),  // 10: echopb.DeleteEchoResponse
}
var file_echo_proto_depIdxs = []int32{
	0, // 0: echopb.PostEchoRequest.echo:type_name -> echopb.Echo
	0, // 1: echopb.PostEchoResponse.echo:type_name -> echopb.Echo
	0, // 2: echopb.GetEchoListResponse.echos:type_name -> echopb.Echo
	0, // 3: echopb.GetEchoByIDResponse.echo:type_name -> echopb.Echo
	0, // 4: echopb.PutEchoRequest.echo:type_name -> echopb.Echo
	0, // 5: echopb.PutEchoResponse.echo:type_name -> echopb.Echo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_echo_proto_init() }
func file_echo_proto_init() {
	if File_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Echo); i {
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
		file_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEchoRequest); i {
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
		file_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostEchoResponse); i {
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
		file_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEchoListRequest); i {
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
		file_echo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEchoListResponse); i {
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
		file_echo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEchoByIDRequest); i {
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
		file_echo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEchoByIDResponse); i {
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
		file_echo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutEchoRequest); i {
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
		file_echo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutEchoResponse); i {
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
		file_echo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEchoRequest); i {
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
		file_echo_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEchoResponse); i {
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
			RawDescriptor: file_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_echo_proto_goTypes,
		DependencyIndexes: file_echo_proto_depIdxs,
		MessageInfos:      file_echo_proto_msgTypes,
	}.Build()
	File_echo_proto = out.File
	file_echo_proto_rawDesc = nil
	file_echo_proto_goTypes = nil
	file_echo_proto_depIdxs = nil
}