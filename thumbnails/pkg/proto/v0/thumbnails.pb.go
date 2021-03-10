// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: thumbnails.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

// The file types to which the thumbnail cna get encoded to.
type GetRequest_FileType int32

const (
	GetRequest_PNG GetRequest_FileType = 0 // Represents PNG type
	GetRequest_JPG GetRequest_FileType = 1 // Represents JPG type
)

// Enum value maps for GetRequest_FileType.
var (
	GetRequest_FileType_name = map[int32]string{
		0: "PNG",
		1: "JPG",
	}
	GetRequest_FileType_value = map[string]int32{
		"PNG": 0,
		"JPG": 1,
	}
)

func (x GetRequest_FileType) Enum() *GetRequest_FileType {
	p := new(GetRequest_FileType)
	*p = x
	return p
}

func (x GetRequest_FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetRequest_FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_thumbnails_proto_enumTypes[0].Descriptor()
}

func (GetRequest_FileType) Type() protoreflect.EnumType {
	return &file_thumbnails_proto_enumTypes[0]
}

func (x GetRequest_FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetRequest_FileType.Descriptor instead.
func (GetRequest_FileType) EnumDescriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{0, 0}
}

// A request to retrieve a thumbnail
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the source image
	Filepath string `protobuf:"bytes,1,opt,name=filepath,proto3" json:"filepath,omitempty"`
	// The type to which the thumbnail should get encoded to.
	Filetype GetRequest_FileType `protobuf:"varint,2,opt,name=filetype,proto3,enum=com.owncloud.ocis.thumbnails.v0.GetRequest_FileType" json:"filetype,omitempty"`
	// The etag of the source image
	Etag string `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	// The width of the thumbnail
	Width int32 `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty"`
	// The height of the thumbnail
	Height int32 `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
	// The authorization token
	Authorization string `protobuf:"bytes,6,opt,name=authorization,proto3" json:"authorization,omitempty"`
	// The user requesting the resource.
	Username string `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetFilepath() string {
	if x != nil {
		return x.Filepath
	}
	return ""
}

func (x *GetRequest) GetFiletype() GetRequest_FileType {
	if x != nil {
		return x.Filetype
	}
	return GetRequest_PNG
}

func (x *GetRequest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *GetRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *GetRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *GetRequest) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

func (x *GetRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// The service response
type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The thumbnail as a binary
	Thumbnail []byte `protobuf:"bytes,1,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// The mimetype of the thumbnail
	Mimetype string `protobuf:"bytes,2,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thumbnails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thumbnails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_thumbnails_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetThumbnail() []byte {
	if x != nil {
		return x.Thumbnail
	}
	return nil
}

func (x *GetResponse) GetMimetype() string {
	if x != nil {
		return x.Mimetype
	}
	return ""
}

var File_thumbnails_proto protoreflect.FileDescriptor

var file_thumbnails_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x76, 0x30, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x02, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x50,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6f, 0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x76, 0x30, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x65, 0x74, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x47,
	0x10, 0x01, 0x22, 0x47, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x74, 0x79, 0x70, 0x65, 0x32, 0x7d, 0x0a, 0x10, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12,
	0x2b, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x63, 0x69, 0x73, 0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76,
	0x30, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x63, 0x69, 0x73,
	0x2e, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x30, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xc2, 0x02, 0x5a, 0x12, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x30, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x92, 0x41, 0xaa, 0x02, 0x12, 0xb8, 0x01, 0x0a, 0x22, 0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x20, 0x49, 0x6e, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x20, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x20, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x47, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x20, 0x47, 0x6d, 0x62, 0x48, 0x12, 0x20, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x63, 0x69, 0x73, 0x1a,
	0x14, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x42, 0x0a, 0x0a, 0x41, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2d,
	0x32, 0x2e, 0x30, 0x12, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6f, 0x63, 0x69, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30,
	0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x72, 0x45, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x20, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x31, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6f, 0x77, 0x6e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x69, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thumbnails_proto_rawDescOnce sync.Once
	file_thumbnails_proto_rawDescData = file_thumbnails_proto_rawDesc
)

func file_thumbnails_proto_rawDescGZIP() []byte {
	file_thumbnails_proto_rawDescOnce.Do(func() {
		file_thumbnails_proto_rawDescData = protoimpl.X.CompressGZIP(file_thumbnails_proto_rawDescData)
	})
	return file_thumbnails_proto_rawDescData
}

var file_thumbnails_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thumbnails_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_thumbnails_proto_goTypes = []interface{}{
	(GetRequest_FileType)(0), // 0: com.owncloud.ocis.thumbnails.v0.GetRequest.FileType
	(*GetRequest)(nil),       // 1: com.owncloud.ocis.thumbnails.v0.GetRequest
	(*GetResponse)(nil),      // 2: com.owncloud.ocis.thumbnails.v0.GetResponse
}
var file_thumbnails_proto_depIdxs = []int32{
	0, // 0: com.owncloud.ocis.thumbnails.v0.GetRequest.filetype:type_name -> com.owncloud.ocis.thumbnails.v0.GetRequest.FileType
	1, // 1: com.owncloud.ocis.thumbnails.v0.ThumbnailService.GetThumbnail:input_type -> com.owncloud.ocis.thumbnails.v0.GetRequest
	2, // 2: com.owncloud.ocis.thumbnails.v0.ThumbnailService.GetThumbnail:output_type -> com.owncloud.ocis.thumbnails.v0.GetResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_thumbnails_proto_init() }
func file_thumbnails_proto_init() {
	if File_thumbnails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thumbnails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_thumbnails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
			RawDescriptor: file_thumbnails_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thumbnails_proto_goTypes,
		DependencyIndexes: file_thumbnails_proto_depIdxs,
		EnumInfos:         file_thumbnails_proto_enumTypes,
		MessageInfos:      file_thumbnails_proto_msgTypes,
	}.Build()
	File_thumbnails_proto = out.File
	file_thumbnails_proto_rawDesc = nil
	file_thumbnails_proto_goTypes = nil
	file_thumbnails_proto_depIdxs = nil
}