// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/proto/configuration/jwt/jwt.proto

package jwt

import (
	eviction "github.com/buildbarn/bb-storage/pkg/proto/configuration/eviction"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthorizationHeaderParserConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Jwks:
	//
	//	*AuthorizationHeaderParserConfiguration_JwksInline
	//	*AuthorizationHeaderParserConfiguration_JwksFile_
	Jwks                                 isAuthorizationHeaderParserConfiguration_Jwks `protobuf_oneof:"jwks"`
	MaximumCacheSize                     int32                                         `protobuf:"varint,3,opt,name=maximum_cache_size,json=maximumCacheSize,proto3" json:"maximum_cache_size,omitempty"`
	CacheReplacementPolicy               eviction.CacheReplacementPolicy               `protobuf:"varint,4,opt,name=cache_replacement_policy,json=cacheReplacementPolicy,proto3,enum=buildbarn.configuration.eviction.CacheReplacementPolicy" json:"cache_replacement_policy,omitempty"`
	ClaimsValidationJmespathExpression   string                                        `protobuf:"bytes,5,opt,name=claims_validation_jmespath_expression,json=claimsValidationJmespathExpression,proto3" json:"claims_validation_jmespath_expression,omitempty"`
	MetadataExtractionJmespathExpression string                                        `protobuf:"bytes,6,opt,name=metadata_extraction_jmespath_expression,json=metadataExtractionJmespathExpression,proto3" json:"metadata_extraction_jmespath_expression,omitempty"`
}

func (x *AuthorizationHeaderParserConfiguration) Reset() {
	*x = AuthorizationHeaderParserConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationHeaderParserConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationHeaderParserConfiguration) ProtoMessage() {}

func (x *AuthorizationHeaderParserConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationHeaderParserConfiguration.ProtoReflect.Descriptor instead.
func (*AuthorizationHeaderParserConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_jwt_jwt_proto_rawDescGZIP(), []int{0}
}

func (m *AuthorizationHeaderParserConfiguration) GetJwks() isAuthorizationHeaderParserConfiguration_Jwks {
	if m != nil {
		return m.Jwks
	}
	return nil
}

func (x *AuthorizationHeaderParserConfiguration) GetJwksInline() *structpb.Struct {
	if x, ok := x.GetJwks().(*AuthorizationHeaderParserConfiguration_JwksInline); ok {
		return x.JwksInline
	}
	return nil
}

func (x *AuthorizationHeaderParserConfiguration) GetJwksFile() *AuthorizationHeaderParserConfiguration_JwksFile {
	if x, ok := x.GetJwks().(*AuthorizationHeaderParserConfiguration_JwksFile_); ok {
		return x.JwksFile
	}
	return nil
}

func (x *AuthorizationHeaderParserConfiguration) GetMaximumCacheSize() int32 {
	if x != nil {
		return x.MaximumCacheSize
	}
	return 0
}

func (x *AuthorizationHeaderParserConfiguration) GetCacheReplacementPolicy() eviction.CacheReplacementPolicy {
	if x != nil {
		return x.CacheReplacementPolicy
	}
	return eviction.CacheReplacementPolicy(0)
}

func (x *AuthorizationHeaderParserConfiguration) GetClaimsValidationJmespathExpression() string {
	if x != nil {
		return x.ClaimsValidationJmespathExpression
	}
	return ""
}

func (x *AuthorizationHeaderParserConfiguration) GetMetadataExtractionJmespathExpression() string {
	if x != nil {
		return x.MetadataExtractionJmespathExpression
	}
	return ""
}

type isAuthorizationHeaderParserConfiguration_Jwks interface {
	isAuthorizationHeaderParserConfiguration_Jwks()
}

type AuthorizationHeaderParserConfiguration_JwksInline struct {
	JwksInline *structpb.Struct `protobuf:"bytes,7,opt,name=jwks_inline,json=jwksInline,proto3,oneof"`
}

type AuthorizationHeaderParserConfiguration_JwksFile_ struct {
	JwksFile *AuthorizationHeaderParserConfiguration_JwksFile `protobuf:"bytes,8,opt,name=jwks_file,json=jwksFile,proto3,oneof"`
}

func (*AuthorizationHeaderParserConfiguration_JwksInline) isAuthorizationHeaderParserConfiguration_Jwks() {
}

func (*AuthorizationHeaderParserConfiguration_JwksFile_) isAuthorizationHeaderParserConfiguration_Jwks() {
}

type AuthorizationHeaderParserConfiguration_JwksFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilePath        string               `protobuf:"bytes,1,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	RefreshInterval *durationpb.Duration `protobuf:"bytes,2,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
}

func (x *AuthorizationHeaderParserConfiguration_JwksFile) Reset() {
	*x = AuthorizationHeaderParserConfiguration_JwksFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationHeaderParserConfiguration_JwksFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationHeaderParserConfiguration_JwksFile) ProtoMessage() {}

func (x *AuthorizationHeaderParserConfiguration_JwksFile) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationHeaderParserConfiguration_JwksFile.ProtoReflect.Descriptor instead.
func (*AuthorizationHeaderParserConfiguration_JwksFile) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_jwt_jwt_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AuthorizationHeaderParserConfiguration_JwksFile) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

func (x *AuthorizationHeaderParserConfiguration_JwksFile) GetRefreshInterval() *durationpb.Duration {
	if x != nil {
		return x.RefreshInterval
	}
	return nil
}

var File_pkg_proto_configuration_jwt_jwt_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_jwt_jwt_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x77, 0x74, 0x2f, 0x6a, 0x77,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x6a, 0x77, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x05, 0x0a, 0x26, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a,
	0x0a, 0x0b, 0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x6a, 0x77, 0x6b, 0x73, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x6b, 0x0a, 0x09, 0x6a, 0x77,
	0x6b, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4c, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4a, 0x77, 0x6b, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x08, 0x6a,
	0x77, 0x6b, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d,
	0x75, 0x6d, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x72, 0x0a, 0x18, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x65, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x52, 0x16, 0x63, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x51, 0x0a, 0x25, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x74, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x27,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x65, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x24, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x6d, 0x0a, 0x08, 0x4a, 0x77, 0x6b, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x44, 0x0a, 0x10,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x6a, 0x77, 0x6b, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62,
	0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x77, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_jwt_jwt_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_jwt_jwt_proto_rawDescData = file_pkg_proto_configuration_jwt_jwt_proto_rawDesc
)

func file_pkg_proto_configuration_jwt_jwt_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_jwt_jwt_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_jwt_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_jwt_jwt_proto_rawDescData)
	})
	return file_pkg_proto_configuration_jwt_jwt_proto_rawDescData
}

var file_pkg_proto_configuration_jwt_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_configuration_jwt_jwt_proto_goTypes = []interface{}{
	(*AuthorizationHeaderParserConfiguration)(nil),          // 0: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration
	(*AuthorizationHeaderParserConfiguration_JwksFile)(nil), // 1: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.JwksFile
	(*structpb.Struct)(nil),                                 // 2: google.protobuf.Struct
	(eviction.CacheReplacementPolicy)(0),                    // 3: buildbarn.configuration.eviction.CacheReplacementPolicy
	(*durationpb.Duration)(nil),                             // 4: google.protobuf.Duration
}
var file_pkg_proto_configuration_jwt_jwt_proto_depIdxs = []int32{
	2, // 0: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.jwks_inline:type_name -> google.protobuf.Struct
	1, // 1: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.jwks_file:type_name -> buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.JwksFile
	3, // 2: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.cache_replacement_policy:type_name -> buildbarn.configuration.eviction.CacheReplacementPolicy
	4, // 3: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration.JwksFile.refresh_interval:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_jwt_jwt_proto_init() }
func file_pkg_proto_configuration_jwt_jwt_proto_init() {
	if File_pkg_proto_configuration_jwt_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationHeaderParserConfiguration); i {
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
		file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationHeaderParserConfiguration_JwksFile); i {
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
	file_pkg_proto_configuration_jwt_jwt_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AuthorizationHeaderParserConfiguration_JwksInline)(nil),
		(*AuthorizationHeaderParserConfiguration_JwksFile_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_jwt_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_jwt_jwt_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_jwt_jwt_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_jwt_jwt_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_jwt_jwt_proto = out.File
	file_pkg_proto_configuration_jwt_jwt_proto_rawDesc = nil
	file_pkg_proto_configuration_jwt_jwt_proto_goTypes = nil
	file_pkg_proto_configuration_jwt_jwt_proto_depIdxs = nil
}
