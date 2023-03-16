// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: contactspb/contact.proto

package contactspb

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

type ContactIn struct {
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

func (x *ContactIn) Reset() {
	*x = ContactIn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactIn) ProtoMessage() {}

func (x *ContactIn) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ContactIn.ProtoReflect.Descriptor instead.
func (*ContactIn) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactIn) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ContactIn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ContactIn) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ContactIn) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ContactIn) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *ContactIn) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ContactIn) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

type GeneralResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GeneralResponse) Reset() {
	*x = GeneralResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contactspb_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResponse) ProtoMessage() {}

func (x *GeneralResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GeneralResponse.ProtoReflect.Descriptor instead.
func (*GeneralResponse) Descriptor() ([]byte, []int) {
	return file_contactspb_contact_proto_rawDescGZIP(), []int{1}
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

var File_contactspb_contact_proto protoreflect.FileDescriptor

var file_contactspb_contact_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x22, 0xbd, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x49, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x47, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x2f, 0x4a, 0x69, 0x74, 0x65, 0x6e, 0x50, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2d,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_contactspb_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_contactspb_contact_proto_goTypes = []interface{}{
	(*ContactIn)(nil),       // 0: contactspb.ContactIn
	(*GeneralResponse)(nil), // 1: contactspb.GeneralResponse
}
var file_contactspb_contact_proto_depIdxs = []int32{
	0, // 0: contactspb.Contact.Create:input_type -> contactspb.ContactIn
	1, // 1: contactspb.Contact.Create:output_type -> contactspb.GeneralResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_contactspb_contact_proto_init() }
func file_contactspb_contact_proto_init() {
	if File_contactspb_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contactspb_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactIn); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contactspb_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
