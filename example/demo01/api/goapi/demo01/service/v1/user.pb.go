// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: user.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height    float64                `protobuf:"fixed64,3,opt,name=height,proto3" json:"height,omitempty"`
	Group     int32                  `protobuf:"varint,4,opt,name=group,proto3" json:"group,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *User) GetGroup() int32 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type UserQuerier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             []int32                `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	IdLower        int32                  `protobuf:"varint,2,opt,name=idLower,proto3" json:"idLower,omitempty"`
	IdUpper        int32                  `protobuf:"varint,3,opt,name=idUpper,proto3" json:"idUpper,omitempty"`
	Name           []string               `protobuf:"bytes,4,rep,name=name,proto3" json:"name,omitempty"`
	SearchName     string                 `protobuf:"bytes,5,opt,name=SearchName,proto3" json:"SearchName,omitempty"`
	HeightLower    float64                `protobuf:"fixed64,6,opt,name=heightLower,proto3" json:"heightLower,omitempty"`
	HeightUpper    float64                `protobuf:"fixed64,7,opt,name=heightUpper,proto3" json:"heightUpper,omitempty"`
	Group          []int32                `protobuf:"varint,8,rep,packed,name=group,proto3" json:"group,omitempty"`
	GroupLower     int32                  `protobuf:"varint,9,opt,name=groupLower,proto3" json:"groupLower,omitempty"`
	GroupUpper     int32                  `protobuf:"varint,10,opt,name=groupUpper,proto3" json:"groupUpper,omitempty"`
	CreatedAtLower *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAtLower,proto3" json:"createdAtLower,omitempty"`
	CreatedAtUpper *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=createdAtUpper,proto3" json:"createdAtUpper,omitempty"`
	UpdatedAtLower *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=updatedAtLower,proto3" json:"updatedAtLower,omitempty"`
	UpdatedAtUpper *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=updatedAtUpper,proto3" json:"updatedAtUpper,omitempty"`
}

func (x *UserQuerier) Reset() {
	*x = UserQuerier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserQuerier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserQuerier) ProtoMessage() {}

func (x *UserQuerier) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserQuerier.ProtoReflect.Descriptor instead.
func (*UserQuerier) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserQuerier) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UserQuerier) GetIdLower() int32 {
	if x != nil {
		return x.IdLower
	}
	return 0
}

func (x *UserQuerier) GetIdUpper() int32 {
	if x != nil {
		return x.IdUpper
	}
	return 0
}

func (x *UserQuerier) GetName() []string {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UserQuerier) GetSearchName() string {
	if x != nil {
		return x.SearchName
	}
	return ""
}

func (x *UserQuerier) GetHeightLower() float64 {
	if x != nil {
		return x.HeightLower
	}
	return 0
}

func (x *UserQuerier) GetHeightUpper() float64 {
	if x != nil {
		return x.HeightUpper
	}
	return 0
}

func (x *UserQuerier) GetGroup() []int32 {
	if x != nil {
		return x.Group
	}
	return nil
}

func (x *UserQuerier) GetGroupLower() int32 {
	if x != nil {
		return x.GroupLower
	}
	return 0
}

func (x *UserQuerier) GetGroupUpper() int32 {
	if x != nil {
		return x.GroupUpper
	}
	return 0
}

func (x *UserQuerier) GetCreatedAtLower() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAtLower
	}
	return nil
}

func (x *UserQuerier) GetCreatedAtUpper() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAtUpper
	}
	return nil
}

func (x *UserQuerier) GetUpdatedAtLower() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAtLower
	}
	return nil
}

func (x *UserQuerier) GetUpdatedAtUpper() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAtUpper
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x65,
	0x6d, 0x6f, 0x30, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcc, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x38, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xaf, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x64, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x69, 0x64, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x55,
	0x70, 0x70, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x64, 0x55, 0x70,
	0x70, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x6f, 0x77, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x70, 0x65, 0x72, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x4c, 0x6f,
	0x77, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x42, 0x0a,
	0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x42, 0x13, 0x5a, 0x11, 0x64, 0x65, 0x6d, 0x6f, 0x30, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: demo01.service.v1.User
	(*UserQuerier)(nil),           // 1: demo01.service.v1.UserQuerier
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_user_proto_depIdxs = []int32{
	2, // 0: demo01.service.v1.User.createdAt:type_name -> google.protobuf.Timestamp
	2, // 1: demo01.service.v1.User.updatedAt:type_name -> google.protobuf.Timestamp
	2, // 2: demo01.service.v1.UserQuerier.createdAtLower:type_name -> google.protobuf.Timestamp
	2, // 3: demo01.service.v1.UserQuerier.createdAtUpper:type_name -> google.protobuf.Timestamp
	2, // 4: demo01.service.v1.UserQuerier.updatedAtLower:type_name -> google.protobuf.Timestamp
	2, // 5: demo01.service.v1.UserQuerier.updatedAtUpper:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserQuerier); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
