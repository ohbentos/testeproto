package testeproto

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ZaryaPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedPermissions []string `protobuf:"bytes,1,rep,name=allowed_permissions,json=allowedPermissions,proto3" json:"allowed_permissions,omitempty"`
	AllowedServices    []string `protobuf:"bytes,2,rep,name=allowed_services,json=allowedServices,proto3" json:"allowed_services,omitempty"`
}

func (x *ZaryaPermissions) Reset() {
	*x = ZaryaPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zarya_rbac_opa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZaryaPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZaryaPermissions) ProtoMessage() {}

func (x *ZaryaPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_zarya_rbac_opa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZaryaPermissions.ProtoReflect.Descriptor instead.
func (*ZaryaPermissions) Descriptor() ([]byte, []int) {
	return file_zarya_rbac_opa_proto_rawDescGZIP(), []int{0}
}

func (x *ZaryaPermissions) GetAllowedPermissions() []string {
	if x != nil {
		return x.AllowedPermissions
	}
	return nil
}

func (x *ZaryaPermissions) GetAllowedServices() []string {
	if x != nil {
		return x.AllowedServices
	}
	return nil
}

var file_zarya_rbac_opa_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*ZaryaPermissions)(nil),
		Field:         1000,
		Name:          "zarya.rbac.authz",
		Tag:           "bytes,1000,opt,name=authz",
		Filename:      "zarya_rbac/opa.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional zarya.rbac.ZaryaPermissions authz = 1000;
	E_Authz = &file_zarya_rbac_opa_proto_extTypes[0]
)

var File_zarya_rbac_opa_proto protoreflect.FileDescriptor

var file_zarya_rbac_opa_proto_rawDesc = []byte{
	0x0a, 0x14, 0x7a, 0x61, 0x72, 0x79, 0x61, 0x5f, 0x72, 0x62, 0x61, 0x63, 0x2f, 0x6f, 0x70, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x7a, 0x61, 0x72, 0x79, 0x61, 0x2e, 0x72, 0x62,
	0x61, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x10, 0x5a, 0x61, 0x72, 0x79, 0x61, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x3a, 0x53, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x12, 0x1e, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe8, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x7a, 0x61, 0x72, 0x79, 0x61, 0x2e, 0x72, 0x62, 0x61,
	0x63, 0x2e, 0x5a, 0x61, 0x72, 0x79, 0x61, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x68, 0x62, 0x65, 0x6e, 0x74, 0x6f, 0x73,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_zarya_rbac_opa_proto_rawDescOnce sync.Once
	file_zarya_rbac_opa_proto_rawDescData = file_zarya_rbac_opa_proto_rawDesc
)

func file_zarya_rbac_opa_proto_rawDescGZIP() []byte {
	file_zarya_rbac_opa_proto_rawDescOnce.Do(func() {
		file_zarya_rbac_opa_proto_rawDescData = protoimpl.X.CompressGZIP(file_zarya_rbac_opa_proto_rawDescData)
	})
	return file_zarya_rbac_opa_proto_rawDescData
}

var file_zarya_rbac_opa_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zarya_rbac_opa_proto_goTypes = []interface{}{
	(*ZaryaPermissions)(nil),           // 0: zarya.rbac.ZaryaPermissions
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_zarya_rbac_opa_proto_depIdxs = []int32{
	1, // 0: zarya.rbac.authz:extendee -> google.protobuf.MethodOptions
	0, // 1: zarya.rbac.authz:type_name -> zarya.rbac.ZaryaPermissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zarya_rbac_opa_proto_init() }
func file_zarya_rbac_opa_proto_init() {
	if File_zarya_rbac_opa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zarya_rbac_opa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZaryaPermissions); i {
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
			RawDescriptor: file_zarya_rbac_opa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_zarya_rbac_opa_proto_goTypes,
		DependencyIndexes: file_zarya_rbac_opa_proto_depIdxs,
		MessageInfos:      file_zarya_rbac_opa_proto_msgTypes,
		ExtensionInfos:    file_zarya_rbac_opa_proto_extTypes,
	}.Build()
	File_zarya_rbac_opa_proto = out.File
	file_zarya_rbac_opa_proto_rawDesc = nil
	file_zarya_rbac_opa_proto_goTypes = nil
	file_zarya_rbac_opa_proto_depIdxs = nil
}
