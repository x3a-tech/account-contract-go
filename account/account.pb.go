// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: account/account.proto

package accountv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Method int32

const (
	Method_NULL              Method = 0
	Method_IDENTIFY          Method = 1
	Method_GET_ACCOUNT_SHORT Method = 2
)

// Enum value maps for Method.
var (
	Method_name = map[int32]string{
		0: "NULL",
		1: "IDENTIFY",
		2: "GET_ACCOUNT_SHORT",
	}
	Method_value = map[string]int32{
		"NULL":              0,
		"IDENTIFY":          1,
		"GET_ACCOUNT_SHORT": 2,
	}
)

func (x Method) Enum() *Method {
	p := new(Method)
	*p = x
	return p
}

func (x Method) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Method) Descriptor() protoreflect.EnumDescriptor {
	return file_account_account_proto_enumTypes[0].Descriptor()
}

func (Method) Type() protoreflect.EnumType {
	return &file_account_account_proto_enumTypes[0]
}

func (x Method) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Method.Descriptor instead.
func (Method) EnumDescriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          []byte                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Photo         []byte                 `protobuf:"bytes,5,opt,name=photo,proto3,oneof" json:"photo,omitempty"`
	Lang          int32                  `protobuf:"varint,6,opt,name=lang,proto3" json:"lang,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_account_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *Account) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Account) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Account) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Account) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

func (x *Account) GetLang() int32 {
	if x != nil {
		return x.Lang
	}
	return 0
}

func (x *Account) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Account) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type AccountShort struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          []byte                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Photo         []byte                 `protobuf:"bytes,5,opt,name=photo,proto3,oneof" json:"photo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountShort) Reset() {
	*x = AccountShort{}
	mi := &file_account_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountShort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountShort) ProtoMessage() {}

func (x *AccountShort) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountShort.ProtoReflect.Descriptor instead.
func (*AccountShort) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountShort) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *AccountShort) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AccountShort) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AccountShort) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AccountShort) GetPhoto() []byte {
	if x != nil {
		return x.Photo
	}
	return nil
}

type TryRegistryFromTelegramParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TgId          int64                  `protobuf:"varint,1,opt,name=tg_id,json=tgId,proto3" json:"tg_id,omitempty"`
	TgFirstName   string                 `protobuf:"bytes,2,opt,name=tg_first_name,json=tgFirstName,proto3" json:"tg_first_name,omitempty"`
	TgLastName    string                 `protobuf:"bytes,3,opt,name=tg_last_name,json=tgLastName,proto3" json:"tg_last_name,omitempty"`
	TgUsername    string                 `protobuf:"bytes,4,opt,name=tg_username,json=tgUsername,proto3" json:"tg_username,omitempty"`
	TgIsPremium   bool                   `protobuf:"varint,5,opt,name=tg_is_premium,json=tgIsPremium,proto3" json:"tg_is_premium,omitempty"`
	LangCode      string                 `protobuf:"bytes,6,opt,name=lang_code,json=langCode,proto3" json:"lang_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TryRegistryFromTelegramParams) Reset() {
	*x = TryRegistryFromTelegramParams{}
	mi := &file_account_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TryRegistryFromTelegramParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TryRegistryFromTelegramParams) ProtoMessage() {}

func (x *TryRegistryFromTelegramParams) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TryRegistryFromTelegramParams.ProtoReflect.Descriptor instead.
func (*TryRegistryFromTelegramParams) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{2}
}

func (x *TryRegistryFromTelegramParams) GetTgId() int64 {
	if x != nil {
		return x.TgId
	}
	return 0
}

func (x *TryRegistryFromTelegramParams) GetTgFirstName() string {
	if x != nil {
		return x.TgFirstName
	}
	return ""
}

func (x *TryRegistryFromTelegramParams) GetTgLastName() string {
	if x != nil {
		return x.TgLastName
	}
	return ""
}

func (x *TryRegistryFromTelegramParams) GetTgUsername() string {
	if x != nil {
		return x.TgUsername
	}
	return ""
}

func (x *TryRegistryFromTelegramParams) GetTgIsPremium() bool {
	if x != nil {
		return x.TgIsPremium
	}
	return false
}

func (x *TryRegistryFromTelegramParams) GetLangCode() string {
	if x != nil {
		return x.LangCode
	}
	return ""
}

type GetAccountParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          []byte                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountParams) Reset() {
	*x = GetAccountParams{}
	mi := &file_account_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountParams) ProtoMessage() {}

func (x *GetAccountParams) ProtoReflect() protoreflect.Message {
	mi := &file_account_account_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountParams.ProtoReflect.Descriptor instead.
func (*GetAccountParams) Descriptor() ([]byte, []int) {
	return file_account_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountParams) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

var File_account_account_proto protoreflect.FileDescriptor

var file_account_account_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x02, 0x0a,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x70, 0x68,
	0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x01, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x1d, 0x54, 0x72, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x67, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x74, 0x67, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x67, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x74, 0x67, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x67, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x67, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65,
	0x6d, 0x69, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x67, 0x49, 0x73,
	0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x2a, 0x37, 0x0a, 0x06,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x46, 0x59, 0x10, 0x01, 0x12, 0x15,
	0x0a, 0x11, 0x47, 0x45, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x48,
	0x4f, 0x52, 0x54, 0x10, 0x02, 0x32, 0x60, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x57,
	0x0a, 0x17, 0x54, 0x72, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x46, 0x72, 0x6f,
	0x6d, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x17, 0x5a, 0x15, 0x78, 0x33, 0x61, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_account_account_proto_rawDescOnce sync.Once
	file_account_account_proto_rawDescData []byte
)

func file_account_account_proto_rawDescGZIP() []byte {
	file_account_account_proto_rawDescOnce.Do(func() {
		file_account_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_account_proto_rawDesc), len(file_account_account_proto_rawDesc)))
	})
	return file_account_account_proto_rawDescData
}

var file_account_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_account_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_account_account_proto_goTypes = []any{
	(Method)(0),                           // 0: users.Method
	(*Account)(nil),                       // 1: users.Account
	(*AccountShort)(nil),                  // 2: users.AccountShort
	(*TryRegistryFromTelegramParams)(nil), // 3: users.TryRegistryFromTelegramParams
	(*GetAccountParams)(nil),              // 4: users.GetAccountParams
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                 // 6: google.protobuf.Empty
}
var file_account_account_proto_depIdxs = []int32{
	5, // 0: users.Account.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: users.Account.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: users.Users.TryRegistryFromTelegram:input_type -> users.TryRegistryFromTelegramParams
	6, // 3: users.Users.TryRegistryFromTelegram:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_account_proto_init() }
func file_account_account_proto_init() {
	if File_account_account_proto != nil {
		return
	}
	file_account_account_proto_msgTypes[0].OneofWrappers = []any{}
	file_account_account_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_account_proto_rawDesc), len(file_account_account_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_account_proto_goTypes,
		DependencyIndexes: file_account_account_proto_depIdxs,
		EnumInfos:         file_account_account_proto_enumTypes,
		MessageInfos:      file_account_account_proto_msgTypes,
	}.Build()
	File_account_account_proto = out.File
	file_account_account_proto_goTypes = nil
	file_account_account_proto_depIdxs = nil
}
