// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contactspb/contact.proto

package contactspb

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type ContactCreateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber  string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Status       string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	LastModified int64  `protobuf:"varint,7,opt,name=LastModified,proto3" json:"LastModified,omitempty"`
}

func (x *ContactCreateMessage) Reset() {
	*x = ContactCreateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactCreateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactCreateMessage) ProtoMessage() {}

func (x *ContactCreateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactCreateMessage.ProtoReflect.Descriptor instead.
func (*ContactCreateMessage) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactCreateMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactCreateMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactCreateMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactCreateMessage) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ContactCreateMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContactCreateMessage) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type ContactMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	PhoneNumber  string `protobuf:"bytes,5,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Status       string `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	LastModified int64  `protobuf:"varint,7,opt,name=LastModified,proto3" json:"LastModified,omitempty"`
}

func (x *ContactMessage) Reset() {
	*x = ContactMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactMessage) ProtoMessage() {}

func (x *ContactMessage) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactMessage.ProtoReflect.Descriptor instead.
func (*ContactMessage) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactMessage) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ContactMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactMessage) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactMessage) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactMessage) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ContactMessage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContactMessage) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type ContactMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contacts []*ContactMessage `protobuf:"bytes,1,rep,name=Contacts,proto3" json:"Contacts,omitempty"` // repeated means array
}

func (x *ContactMessages) Reset() {
	*x = ContactMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactMessages) ProtoMessage() {}

func (x *ContactMessages) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactMessages.ProtoReflect.Descriptor instead.
func (*ContactMessages) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{2}
}

func (x *ContactMessages) GetContacts() []*ContactMessage {
	if x != nil {
		return x.Contacts
	}
	return nil
}

type ByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ByIDRequest) Reset() {
	*x = ByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIDRequest) ProtoMessage() {}

func (x *ByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIDRequest.ProtoReflect.Descriptor instead.
func (*ByIDRequest) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{3}
}

func (x *ByIDRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Data *_struct.Struct `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"` // if using map in golang code then there should be google.protobuf.struct
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateRequest) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{5}
}

func (x *GeneralResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GeneralResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NoIn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoIn) Reset() {
	*x = NoIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoIn) ProtoMessage() {}

func (x *NoIn) ProtoReflect() protoreflect.Message {
	mi := &file_contactspb_contact_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoIn.ProtoReflect.Descriptor instead.
func (*NoIn) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{6}
}

var File_contactspb_contact_proto protoreflect.FileDescriptor

var file_contactspb_contact_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22,
	0xc2, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x22,
	0x1d, 0x0a, 0x0b, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4c,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x2b, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x0f,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x06, 0x0a,
	0x04, 0x4e, 0x6f, 0x49, 0x6e, 0x32, 0xcb, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70,
	0x62, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73,
	0x70, 0x62, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x49, 0x6e, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x2f, 0x4a, 0x69, 0x74, 0x65,
	0x6e, 0x50, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2d, 0x71, 0x75, 0x65, 0x73, 0x74, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_contactspb_contact_proto_rawDescOnce sync.Once
	file_contactspb_contact_proto_rawDescData = file_contactspb_contact_proto_rawDesc
)

func file_contactspb_contact_proto_rawDescGZIP() []byte {
	file_contactspb_contact_proto_rawDescOnce.Do(func() {
		file_contactspb_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contactspb_contact_proto_rawDescData)
	})
	return file_contactspb_contact_proto_rawDescData
}

var file_contactspb_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_contactspb_contact_proto_goTypes = []interface{}{
	(*ContactCreateMessage)(nil), // 0: contactspb.ContactCreateMessage
	(*ContactMessage)(nil),       // 1: contactspb.ContactMessage
	(*ContactMessages)(nil),      // 2: contactspb.ContactMessages
	(*ByIDRequest)(nil),          // 3: contactspb.ByIDRequest
	(*UpdateRequest)(nil),        // 4: contactspb.UpdateRequest
	(*GeneralResponse)(nil),      // 5: contactspb.GeneralResponse
	(*NoIn)(nil),                 // 6: contactspb.NoIn
	(*_struct.Struct)(nil),       // 7: google.protobuf.Struct
}
var file_contactspb_contact_proto_depIdxs = []int32{
	1, // 0: contactspb.ContactMessages.Contacts:type_name -> contactspb.ContactMessage
	7, // 1: contactspb.UpdateRequest.Data:type_name -> google.protobuf.Struct
	0, // 2: contactspb.Contact.Create:input_type -> contactspb.ContactCreateMessage
	3, // 3: contactspb.Contact.Delete:input_type -> contactspb.ByIDRequest
	4, // 4: contactspb.Contact.Update:input_type -> contactspb.UpdateRequest
	3, // 5: contactspb.Contact.GetBy:input_type -> contactspb.ByIDRequest
	6, // 6: contactspb.Contact.GetAll:input_type -> contactspb.NoIn
	5, // 7: contactspb.Contact.Create:output_type -> contactspb.GeneralResponse
	5, // 8: contactspb.Contact.Delete:output_type -> contactspb.GeneralResponse
	5, // 9: contactspb.Contact.Update:output_type -> contactspb.GeneralResponse
	1, // 10: contactspb.Contact.GetBy:output_type -> contactspb.ContactMessage
	2, // 11: contactspb.Contact.GetAll:output_type -> contactspb.ContactMessages
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contactspb_contact_proto_init() }
func file_contactspb_contact_proto_init() {
	if File_contactspb_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contactspb_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactCreateMessage); i {
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
		file_contactspb_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactMessage); i {
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
		file_contactspb_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactMessages); i {
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
		file_contactspb_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIDRequest); i {
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
		file_contactspb_contact_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_contactspb_contact_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResponse); i {
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
		file_contactspb_contact_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoIn); i {
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
			RawDescriptor: file_contactspb_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contactspb_contact_proto_goTypes,
		DependencyIndexes: file_contactspb_contact_proto_depIdxs,
		MessageInfos:      file_contactspb_contact_proto_msgTypes,
	}.Build()
	File_contactspb_contact_proto = out.File
	file_contactspb_contact_proto_rawDesc = nil
	file_contactspb_contact_proto_goTypes = nil
	file_contactspb_contact_proto_depIdxs = nil
}
