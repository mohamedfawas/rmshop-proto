// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: v1/admin/admin.proto

package adminv1

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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_v1_admin_admin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_admin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_v1_admin_admin_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetUserDetailsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserDetailsRequest) Reset() {
	*x = GetUserDetailsRequest{}
	mi := &file_v1_admin_admin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsRequest) ProtoMessage() {}

func (x *GetUserDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_admin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetUserDetailsRequest) Descriptor() ([]byte, []int) {
	return file_v1_admin_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserDetailsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserDetailsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserDetailsResponse) Reset() {
	*x = GetUserDetailsResponse{}
	mi := &file_v1_admin_admin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsResponse) ProtoMessage() {}

func (x *GetUserDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_admin_admin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetUserDetailsResponse) Descriptor() ([]byte, []int) {
	return file_v1_admin_admin_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserDetailsResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_v1_admin_admin_proto protoreflect.FileDescriptor

var file_v1_admin_admin_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x5f, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x30, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x32, 0x63, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x68, 0x61, 0x6d, 0x65, 0x64, 0x66, 0x61, 0x77, 0x61,
	0x73, 0x2f, 0x72, 0x6d, 0x73, 0x68, 0x6f, 0x70, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x3b, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_admin_admin_proto_rawDescOnce sync.Once
	file_v1_admin_admin_proto_rawDescData = file_v1_admin_admin_proto_rawDesc
)

func file_v1_admin_admin_proto_rawDescGZIP() []byte {
	file_v1_admin_admin_proto_rawDescOnce.Do(func() {
		file_v1_admin_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_admin_admin_proto_rawDescData)
	})
	return file_v1_admin_admin_proto_rawDescData
}

var file_v1_admin_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_admin_admin_proto_goTypes = []any{
	(*User)(nil),                   // 0: v1.admin.User
	(*GetUserDetailsRequest)(nil),  // 1: v1.admin.GetUserDetailsRequest
	(*GetUserDetailsResponse)(nil), // 2: v1.admin.GetUserDetailsResponse
}
var file_v1_admin_admin_proto_depIdxs = []int32{
	0, // 0: v1.admin.GetUserDetailsResponse.user:type_name -> v1.admin.User
	1, // 1: v1.admin.AdminService.GetUserDetails:input_type -> v1.admin.GetUserDetailsRequest
	2, // 2: v1.admin.AdminService.GetUserDetails:output_type -> v1.admin.GetUserDetailsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v1_admin_admin_proto_init() }
func file_v1_admin_admin_proto_init() {
	if File_v1_admin_admin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_admin_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_admin_admin_proto_goTypes,
		DependencyIndexes: file_v1_admin_admin_proto_depIdxs,
		MessageInfos:      file_v1_admin_admin_proto_msgTypes,
	}.Build()
	File_v1_admin_admin_proto = out.File
	file_v1_admin_admin_proto_rawDesc = nil
	file_v1_admin_admin_proto_goTypes = nil
	file_v1_admin_admin_proto_depIdxs = nil
}
