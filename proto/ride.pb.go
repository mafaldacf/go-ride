// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ride.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{1}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterResponse struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{2}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type LogoutRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{5}
}

func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (m *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(m, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type AddFundsRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	Amount               uint32   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFundsRequest) Reset()         { *m = AddFundsRequest{} }
func (m *AddFundsRequest) String() string { return proto.CompactTextString(m) }
func (*AddFundsRequest) ProtoMessage()    {}
func (*AddFundsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{6}
}

func (m *AddFundsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFundsRequest.Unmarshal(m, b)
}
func (m *AddFundsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFundsRequest.Marshal(b, m, deterministic)
}
func (m *AddFundsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFundsRequest.Merge(m, src)
}
func (m *AddFundsRequest) XXX_Size() int {
	return xxx_messageInfo_AddFundsRequest.Size(m)
}
func (m *AddFundsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFundsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddFundsRequest proto.InternalMessageInfo

func (m *AddFundsRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *AddFundsRequest) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AddFundsResponse struct {
	Balance              uint32   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFundsResponse) Reset()         { *m = AddFundsResponse{} }
func (m *AddFundsResponse) String() string { return proto.CompactTextString(m) }
func (*AddFundsResponse) ProtoMessage()    {}
func (*AddFundsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{7}
}

func (m *AddFundsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFundsResponse.Unmarshal(m, b)
}
func (m *AddFundsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFundsResponse.Marshal(b, m, deterministic)
}
func (m *AddFundsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFundsResponse.Merge(m, src)
}
func (m *AddFundsResponse) XXX_Size() int {
	return xxx_messageInfo_AddFundsResponse.Size(m)
}
func (m *AddFundsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFundsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddFundsResponse proto.InternalMessageInfo

func (m *AddFundsResponse) GetBalance() uint32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type CheckBalanceRequest struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckBalanceRequest) Reset()         { *m = CheckBalanceRequest{} }
func (m *CheckBalanceRequest) String() string { return proto.CompactTextString(m) }
func (*CheckBalanceRequest) ProtoMessage()    {}
func (*CheckBalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{8}
}

func (m *CheckBalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBalanceRequest.Unmarshal(m, b)
}
func (m *CheckBalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBalanceRequest.Marshal(b, m, deterministic)
}
func (m *CheckBalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBalanceRequest.Merge(m, src)
}
func (m *CheckBalanceRequest) XXX_Size() int {
	return xxx_messageInfo_CheckBalanceRequest.Size(m)
}
func (m *CheckBalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBalanceRequest proto.InternalMessageInfo

func (m *CheckBalanceRequest) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type CheckBalanceResponse struct {
	Balance              uint32   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckBalanceResponse) Reset()         { *m = CheckBalanceResponse{} }
func (m *CheckBalanceResponse) String() string { return proto.CompactTextString(m) }
func (*CheckBalanceResponse) ProtoMessage()    {}
func (*CheckBalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ffa51ead3a087cb, []int{9}
}

func (m *CheckBalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckBalanceResponse.Unmarshal(m, b)
}
func (m *CheckBalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckBalanceResponse.Marshal(b, m, deterministic)
}
func (m *CheckBalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckBalanceResponse.Merge(m, src)
}
func (m *CheckBalanceResponse) XXX_Size() int {
	return xxx_messageInfo_CheckBalanceResponse.Size(m)
}
func (m *CheckBalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckBalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckBalanceResponse proto.InternalMessageInfo

func (m *CheckBalanceResponse) GetBalance() uint32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "LogoutRequest")
	proto.RegisterType((*AddFundsRequest)(nil), "AddFundsRequest")
	proto.RegisterType((*AddFundsResponse)(nil), "AddFundsResponse")
	proto.RegisterType((*CheckBalanceRequest)(nil), "CheckBalanceRequest")
	proto.RegisterType((*CheckBalanceResponse)(nil), "CheckBalanceResponse")
}

func init() {
	proto.RegisterFile("proto/ride.proto", fileDescriptor_4ffa51ead3a087cb)
}

var fileDescriptor_4ffa51ead3a087cb = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0x4f, 0xc2, 0x30,
	0x14, 0xc7, 0x83, 0x09, 0x03, 0x9e, 0x4c, 0x46, 0x45, 0x43, 0x16, 0x0f, 0x64, 0x07, 0xe3, 0x41,
	0x0b, 0x91, 0xa3, 0x27, 0x31, 0x62, 0x4c, 0x3c, 0x4d, 0x4f, 0xde, 0xca, 0xf6, 0x02, 0x0b, 0xd0,
	0xce, 0xb6, 0xd3, 0xf8, 0x8f, 0x7b, 0x36, 0x74, 0xc5, 0xf1, 0x2b, 0x4a, 0xe2, 0x69, 0xfb, 0xee,
	0xb5, 0xef, 0xfb, 0xfa, 0xf9, 0x76, 0xe0, 0xa5, 0x52, 0x68, 0xd1, 0x95, 0x49, 0x8c, 0xd4, 0xbc,
	0x06, 0x15, 0x28, 0xdf, 0xcf, 0x53, 0xfd, 0x19, 0x3c, 0x42, 0x23, 0xc4, 0x71, 0xa2, 0x34, 0xca,
	0x10, 0xdf, 0x32, 0x54, 0x9a, 0xf8, 0x50, 0xcd, 0x14, 0x4a, 0xce, 0xe6, 0xd8, 0x2e, 0x75, 0x4a,
	0x17, 0xb5, 0xf0, 0x47, 0x2f, 0x6a, 0x29, 0x53, 0xea, 0x43, 0xc8, 0xb8, 0x7d, 0x90, 0xd7, 0x96,
	0x3a, 0xe8, 0x81, 0x57, 0xb4, 0x52, 0xa9, 0xe0, 0x0a, 0xc9, 0x19, 0xd4, 0x58, 0xa6, 0x27, 0x2f,
	0x62, 0x8a, 0xdc, 0x36, 0x2b, 0x3e, 0x04, 0x43, 0xa8, 0x3f, 0x89, 0x71, 0xc2, 0xff, 0xeb, 0x7c,
	0x05, 0xae, 0xed, 0xb3, 0x97, 0x6d, 0xbe, 0x5c, 0x64, 0x7a, 0xe9, 0xfb, 0xfb, 0xf2, 0x07, 0x68,
	0xdc, 0xc6, 0xf1, 0x30, 0xe3, 0xb1, 0xda, 0x6b, 0x03, 0x39, 0x05, 0x87, 0xcd, 0x45, 0xc6, 0xb5,
	0x19, 0xd4, 0x0d, 0xad, 0x0a, 0x2e, 0xc1, 0x2b, 0x1a, 0xd9, 0x49, 0xdb, 0x50, 0x19, 0xb1, 0x19,
	0xe3, 0x51, 0x7e, 0x62, 0x37, 0x5c, 0xca, 0xa0, 0x0f, 0xc7, 0x77, 0x13, 0x8c, 0xa6, 0x83, 0x5c,
	0xef, 0x37, 0x6b, 0x0f, 0x5a, 0xeb, 0x9b, 0xfe, 0xb2, 0xb9, 0xfe, 0x2a, 0xc1, 0x61, 0x98, 0xc4,
	0xf8, 0x8c, 0xf2, 0x3d, 0x89, 0x90, 0x74, 0xa1, 0x2a, 0x6d, 0x8a, 0xc4, 0xa3, 0x1b, 0x77, 0xc3,
	0x6f, 0xd2, 0xad, 0x88, 0xcf, 0xa1, 0x3c, 0x5b, 0xc0, 0x27, 0x2e, 0x5d, 0x0d, 0xd3, 0x3f, 0xa2,
	0xeb, 0x99, 0x74, 0xc0, 0x99, 0x19, 0xea, 0xc4, 0x54, 0x0a, 0xfc, 0xbe, 0x43, 0xcd, 0x5d, 0x5c,
	0x58, 0x33, 0xcb, 0x87, 0x78, 0x74, 0x83, 0xb9, 0xdf, 0xa4, 0x5b, 0xf0, 0x6e, 0xa0, 0x1e, 0xad,
	0x9c, 0x96, 0xb4, 0xe8, 0x0e, 0x62, 0xfe, 0x09, 0xdd, 0x85, 0x64, 0x50, 0x7d, 0x75, 0xba, 0xe6,
	0x67, 0x18, 0x39, 0xe6, 0xd1, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x37, 0xf0, 0xb7, 0x27,
	0x03, 0x00, 0x00,
}