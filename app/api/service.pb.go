// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: api/service.proto

package api

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LoginRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{1}
}

type BattleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID            string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MyMonsterID       string `protobuf:"bytes,2,opt,name=myMonsterID,proto3" json:"myMonsterID,omitempty"`
	OpponentMonsterID string `protobuf:"bytes,3,opt,name=opponentMonsterID,proto3" json:"opponentMonsterID,omitempty"`
	CreatedAt         int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *BattleRequest) Reset() {
	*x = BattleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleRequest) ProtoMessage() {}

func (x *BattleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleRequest.ProtoReflect.Descriptor instead.
func (*BattleRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *BattleRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *BattleRequest) GetMyMonsterID() string {
	if x != nil {
		return x.MyMonsterID
	}
	return ""
}

func (x *BattleRequest) GetOpponentMonsterID() string {
	if x != nil {
		return x.OpponentMonsterID
	}
	return ""
}

func (x *BattleRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type BattleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BattleResponse) Reset() {
	*x = BattleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleResponse) ProtoMessage() {}

func (x *BattleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleResponse.ProtoReflect.Descriptor instead.
func (*BattleResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{3}
}

type LevelUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MonsterID string `protobuf:"bytes,2,opt,name=monsterID,proto3" json:"monsterID,omitempty"`
	Level     int64  `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *LevelUpRequest) Reset() {
	*x = LevelUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelUpRequest) ProtoMessage() {}

func (x *LevelUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelUpRequest.ProtoReflect.Descriptor instead.
func (*LevelUpRequest) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{4}
}

func (x *LevelUpRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LevelUpRequest) GetMonsterID() string {
	if x != nil {
		return x.MonsterID
	}
	return ""
}

func (x *LevelUpRequest) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LevelUpRequest) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type LevelUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LevelUpResponse) Reset() {
	*x = LevelUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelUpResponse) ProtoMessage() {}

func (x *LevelUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelUpResponse.ProtoReflect.Descriptor instead.
func (*LevelUpResponse) Descriptor() ([]byte, []int) {
	return file_api_service_proto_rawDescGZIP(), []int{5}
}

var File_api_service_proto protoreflect.FileDescriptor

var file_api_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x44, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x0f,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x95, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x79, 0x4d,
	0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x79, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x6f,
	0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x0e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa8, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x61, 0x63, 0x74, 0x69, 0x63, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_service_proto_rawDescOnce sync.Once
	file_api_service_proto_rawDescData = file_api_service_proto_rawDesc
)

func file_api_service_proto_rawDescGZIP() []byte {
	file_api_service_proto_rawDescOnce.Do(func() {
		file_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_service_proto_rawDescData)
	})
	return file_api_service_proto_rawDescData
}

var file_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_service_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),    // 0: api.LoginRequest
	(*LoginResponse)(nil),   // 1: api.LoginResponse
	(*BattleRequest)(nil),   // 2: api.BattleRequest
	(*BattleResponse)(nil),  // 3: api.BattleResponse
	(*LevelUpRequest)(nil),  // 4: api.LevelUpRequest
	(*LevelUpResponse)(nil), // 5: api.LevelUpResponse
}
var file_api_service_proto_depIdxs = []int32{
	0, // 0: api.ActionService.Login:input_type -> api.LoginRequest
	2, // 1: api.ActionService.Battle:input_type -> api.BattleRequest
	4, // 2: api.ActionService.LevelUp:input_type -> api.LevelUpRequest
	1, // 3: api.ActionService.Login:output_type -> api.LoginResponse
	3, // 4: api.ActionService.Battle:output_type -> api.BattleResponse
	5, // 5: api.ActionService.LevelUp:output_type -> api.LevelUpResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_service_proto_init() }
func file_api_service_proto_init() {
	if File_api_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleRequest); i {
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
		file_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleResponse); i {
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
		file_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelUpRequest); i {
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
		file_api_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelUpResponse); i {
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
			RawDescriptor: file_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_proto_goTypes,
		DependencyIndexes: file_api_service_proto_depIdxs,
		MessageInfos:      file_api_service_proto_msgTypes,
	}.Build()
	File_api_service_proto = out.File
	file_api_service_proto_rawDesc = nil
	file_api_service_proto_goTypes = nil
	file_api_service_proto_depIdxs = nil
}
