// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agrpc.proto

package agrpc

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

type RespStatus struct {
	StatusCode           int32    `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	ErrMessage           string   `protobuf:"bytes,2,opt,name=errMessage,proto3" json:"errMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespStatus) Reset()         { *m = RespStatus{} }
func (m *RespStatus) String() string { return proto.CompactTextString(m) }
func (*RespStatus) ProtoMessage()    {}
func (*RespStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{0}
}

func (m *RespStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespStatus.Unmarshal(m, b)
}
func (m *RespStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespStatus.Marshal(b, m, deterministic)
}
func (m *RespStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespStatus.Merge(m, src)
}
func (m *RespStatus) XXX_Size() int {
	return xxx_messageInfo_RespStatus.Size(m)
}
func (m *RespStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RespStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RespStatus proto.InternalMessageInfo

func (m *RespStatus) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *RespStatus) GetErrMessage() string {
	if m != nil {
		return m.ErrMessage
	}
	return ""
}

type KeyInfo struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	PublicKey            string   `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	CreatedTimestamp     int64    `protobuf:"varint,3,opt,name=createdTimestamp,proto3" json:"createdTimestamp,omitempty"`
	ExpiredTimestamp     int64    `protobuf:"varint,4,opt,name=expiredTimestamp,proto3" json:"expiredTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyInfo) Reset()         { *m = KeyInfo{} }
func (m *KeyInfo) String() string { return proto.CompactTextString(m) }
func (*KeyInfo) ProtoMessage()    {}
func (*KeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{1}
}

func (m *KeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyInfo.Unmarshal(m, b)
}
func (m *KeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyInfo.Marshal(b, m, deterministic)
}
func (m *KeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyInfo.Merge(m, src)
}
func (m *KeyInfo) XXX_Size() int {
	return xxx_messageInfo_KeyInfo.Size(m)
}
func (m *KeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeyInfo proto.InternalMessageInfo

func (m *KeyInfo) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

func (m *KeyInfo) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *KeyInfo) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
	}
	return 0
}

func (m *KeyInfo) GetExpiredTimestamp() int64 {
	if m != nil {
		return m.ExpiredTimestamp
	}
	return 0
}

// GenRsaKeyPair api
type GenRsaKeyPairReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	Timeout              int32    `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenRsaKeyPairReq) Reset()         { *m = GenRsaKeyPairReq{} }
func (m *GenRsaKeyPairReq) String() string { return proto.CompactTextString(m) }
func (*GenRsaKeyPairReq) ProtoMessage()    {}
func (*GenRsaKeyPairReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{2}
}

func (m *GenRsaKeyPairReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenRsaKeyPairReq.Unmarshal(m, b)
}
func (m *GenRsaKeyPairReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenRsaKeyPairReq.Marshal(b, m, deterministic)
}
func (m *GenRsaKeyPairReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenRsaKeyPairReq.Merge(m, src)
}
func (m *GenRsaKeyPairReq) XXX_Size() int {
	return xxx_messageInfo_GenRsaKeyPairReq.Size(m)
}
func (m *GenRsaKeyPairReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GenRsaKeyPairReq.DiscardUnknown(m)
}

var xxx_messageInfo_GenRsaKeyPairReq proto.InternalMessageInfo

func (m *GenRsaKeyPairReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

func (m *GenRsaKeyPairReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type GenRsaKeyPairResp struct {
	Status               *RespStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	KeyInfo              *KeyInfo    `protobuf:"bytes,2,opt,name=keyInfo,proto3" json:"keyInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GenRsaKeyPairResp) Reset()         { *m = GenRsaKeyPairResp{} }
func (m *GenRsaKeyPairResp) String() string { return proto.CompactTextString(m) }
func (*GenRsaKeyPairResp) ProtoMessage()    {}
func (*GenRsaKeyPairResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{3}
}

func (m *GenRsaKeyPairResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenRsaKeyPairResp.Unmarshal(m, b)
}
func (m *GenRsaKeyPairResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenRsaKeyPairResp.Marshal(b, m, deterministic)
}
func (m *GenRsaKeyPairResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenRsaKeyPairResp.Merge(m, src)
}
func (m *GenRsaKeyPairResp) XXX_Size() int {
	return xxx_messageInfo_GenRsaKeyPairResp.Size(m)
}
func (m *GenRsaKeyPairResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenRsaKeyPairResp.DiscardUnknown(m)
}

var xxx_messageInfo_GenRsaKeyPairResp proto.InternalMessageInfo

func (m *GenRsaKeyPairResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GenRsaKeyPairResp) GetKeyInfo() *KeyInfo {
	if m != nil {
		return m.KeyInfo
	}
	return nil
}

// RemoveRsaKeyPair api
type RemoveRsaKeyPairReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRsaKeyPairReq) Reset()         { *m = RemoveRsaKeyPairReq{} }
func (m *RemoveRsaKeyPairReq) String() string { return proto.CompactTextString(m) }
func (*RemoveRsaKeyPairReq) ProtoMessage()    {}
func (*RemoveRsaKeyPairReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{4}
}

func (m *RemoveRsaKeyPairReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRsaKeyPairReq.Unmarshal(m, b)
}
func (m *RemoveRsaKeyPairReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRsaKeyPairReq.Marshal(b, m, deterministic)
}
func (m *RemoveRsaKeyPairReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRsaKeyPairReq.Merge(m, src)
}
func (m *RemoveRsaKeyPairReq) XXX_Size() int {
	return xxx_messageInfo_RemoveRsaKeyPairReq.Size(m)
}
func (m *RemoveRsaKeyPairReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRsaKeyPairReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRsaKeyPairReq proto.InternalMessageInfo

func (m *RemoveRsaKeyPairReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

type RemoveRsaKeyPairResp struct {
	Status               *RespStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RemoveRsaKeyPairResp) Reset()         { *m = RemoveRsaKeyPairResp{} }
func (m *RemoveRsaKeyPairResp) String() string { return proto.CompactTextString(m) }
func (*RemoveRsaKeyPairResp) ProtoMessage()    {}
func (*RemoveRsaKeyPairResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{5}
}

func (m *RemoveRsaKeyPairResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRsaKeyPairResp.Unmarshal(m, b)
}
func (m *RemoveRsaKeyPairResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRsaKeyPairResp.Marshal(b, m, deterministic)
}
func (m *RemoveRsaKeyPairResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRsaKeyPairResp.Merge(m, src)
}
func (m *RemoveRsaKeyPairResp) XXX_Size() int {
	return xxx_messageInfo_RemoveRsaKeyPairResp.Size(m)
}
func (m *RemoveRsaKeyPairResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRsaKeyPairResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRsaKeyPairResp proto.InternalMessageInfo

func (m *RemoveRsaKeyPairResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Encrypt api
type EncryptReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	PlainText            string   `protobuf:"bytes,2,opt,name=plainText,proto3" json:"plainText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptReq) Reset()         { *m = EncryptReq{} }
func (m *EncryptReq) String() string { return proto.CompactTextString(m) }
func (*EncryptReq) ProtoMessage()    {}
func (*EncryptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{6}
}

func (m *EncryptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptReq.Unmarshal(m, b)
}
func (m *EncryptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptReq.Marshal(b, m, deterministic)
}
func (m *EncryptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptReq.Merge(m, src)
}
func (m *EncryptReq) XXX_Size() int {
	return xxx_messageInfo_EncryptReq.Size(m)
}
func (m *EncryptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptReq.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptReq proto.InternalMessageInfo

func (m *EncryptReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

func (m *EncryptReq) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

type EncryptResp struct {
	Status               *RespStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	CipherText           string      `protobuf:"bytes,2,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EncryptResp) Reset()         { *m = EncryptResp{} }
func (m *EncryptResp) String() string { return proto.CompactTextString(m) }
func (*EncryptResp) ProtoMessage()    {}
func (*EncryptResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{7}
}

func (m *EncryptResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResp.Unmarshal(m, b)
}
func (m *EncryptResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResp.Marshal(b, m, deterministic)
}
func (m *EncryptResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResp.Merge(m, src)
}
func (m *EncryptResp) XXX_Size() int {
	return xxx_messageInfo_EncryptResp.Size(m)
}
func (m *EncryptResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResp.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResp proto.InternalMessageInfo

func (m *EncryptResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EncryptResp) GetCipherText() string {
	if m != nil {
		return m.CipherText
	}
	return ""
}

// Decrypt api
type DecryptReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	CipherText           string   `protobuf:"bytes,2,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptReq) Reset()         { *m = DecryptReq{} }
func (m *DecryptReq) String() string { return proto.CompactTextString(m) }
func (*DecryptReq) ProtoMessage()    {}
func (*DecryptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{8}
}

func (m *DecryptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptReq.Unmarshal(m, b)
}
func (m *DecryptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptReq.Marshal(b, m, deterministic)
}
func (m *DecryptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptReq.Merge(m, src)
}
func (m *DecryptReq) XXX_Size() int {
	return xxx_messageInfo_DecryptReq.Size(m)
}
func (m *DecryptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptReq.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptReq proto.InternalMessageInfo

func (m *DecryptReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

func (m *DecryptReq) GetCipherText() string {
	if m != nil {
		return m.CipherText
	}
	return ""
}

type DecryptResp struct {
	Status               *RespStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PlainText            string      `protobuf:"bytes,2,opt,name=plainText,proto3" json:"plainText,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DecryptResp) Reset()         { *m = DecryptResp{} }
func (m *DecryptResp) String() string { return proto.CompactTextString(m) }
func (*DecryptResp) ProtoMessage()    {}
func (*DecryptResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{9}
}

func (m *DecryptResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResp.Unmarshal(m, b)
}
func (m *DecryptResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResp.Marshal(b, m, deterministic)
}
func (m *DecryptResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResp.Merge(m, src)
}
func (m *DecryptResp) XXX_Size() int {
	return xxx_messageInfo_DecryptResp.Size(m)
}
func (m *DecryptResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResp.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResp proto.InternalMessageInfo

func (m *DecryptResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DecryptResp) GetPlainText() string {
	if m != nil {
		return m.PlainText
	}
	return ""
}

// CheckKey api
type CheckKeyReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckKeyReq) Reset()         { *m = CheckKeyReq{} }
func (m *CheckKeyReq) String() string { return proto.CompactTextString(m) }
func (*CheckKeyReq) ProtoMessage()    {}
func (*CheckKeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{10}
}

func (m *CheckKeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckKeyReq.Unmarshal(m, b)
}
func (m *CheckKeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckKeyReq.Marshal(b, m, deterministic)
}
func (m *CheckKeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckKeyReq.Merge(m, src)
}
func (m *CheckKeyReq) XXX_Size() int {
	return xxx_messageInfo_CheckKeyReq.Size(m)
}
func (m *CheckKeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckKeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_CheckKeyReq proto.InternalMessageInfo

func (m *CheckKeyReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

type CheckKeyResp struct {
	Status               *RespStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	KeyInfos             []*KeyInfo  `protobuf:"bytes,2,rep,name=keyInfos,proto3" json:"keyInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CheckKeyResp) Reset()         { *m = CheckKeyResp{} }
func (m *CheckKeyResp) String() string { return proto.CompactTextString(m) }
func (*CheckKeyResp) ProtoMessage()    {}
func (*CheckKeyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{11}
}

func (m *CheckKeyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckKeyResp.Unmarshal(m, b)
}
func (m *CheckKeyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckKeyResp.Marshal(b, m, deterministic)
}
func (m *CheckKeyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckKeyResp.Merge(m, src)
}
func (m *CheckKeyResp) XXX_Size() int {
	return xxx_messageInfo_CheckKeyResp.Size(m)
}
func (m *CheckKeyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckKeyResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckKeyResp proto.InternalMessageInfo

func (m *CheckKeyResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CheckKeyResp) GetKeyInfos() []*KeyInfo {
	if m != nil {
		return m.KeyInfos
	}
	return nil
}

type SecretParamInfo struct {
	SecretName           string   `protobuf:"bytes,1,opt,name=secretName,proto3" json:"secretName,omitempty"`
	CreatedTimestamp     int64    `protobuf:"varint,2,opt,name=createdTimestamp,proto3" json:"createdTimestamp,omitempty"`
	ExpiredTimestamp     int64    `protobuf:"varint,3,opt,name=expiredTimestamp,proto3" json:"expiredTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecretParamInfo) Reset()         { *m = SecretParamInfo{} }
func (m *SecretParamInfo) String() string { return proto.CompactTextString(m) }
func (*SecretParamInfo) ProtoMessage()    {}
func (*SecretParamInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{12}
}

func (m *SecretParamInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretParamInfo.Unmarshal(m, b)
}
func (m *SecretParamInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretParamInfo.Marshal(b, m, deterministic)
}
func (m *SecretParamInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretParamInfo.Merge(m, src)
}
func (m *SecretParamInfo) XXX_Size() int {
	return xxx_messageInfo_SecretParamInfo.Size(m)
}
func (m *SecretParamInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretParamInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SecretParamInfo proto.InternalMessageInfo

func (m *SecretParamInfo) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *SecretParamInfo) GetCreatedTimestamp() int64 {
	if m != nil {
		return m.CreatedTimestamp
	}
	return 0
}

func (m *SecretParamInfo) GetExpiredTimestamp() int64 {
	if m != nil {
		return m.ExpiredTimestamp
	}
	return 0
}

// CreateSecret api
type CreateSecretParamReq struct {
	KeyPairId            string   `protobuf:"bytes,1,opt,name=keyPairId,proto3" json:"keyPairId,omitempty"`
	CipherText           string   `protobuf:"bytes,2,opt,name=cipherText,proto3" json:"cipherText,omitempty"`
	SecretName           string   `protobuf:"bytes,3,opt,name=secretName,proto3" json:"secretName,omitempty"`
	Timeout              int32    `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSecretParamReq) Reset()         { *m = CreateSecretParamReq{} }
func (m *CreateSecretParamReq) String() string { return proto.CompactTextString(m) }
func (*CreateSecretParamReq) ProtoMessage()    {}
func (*CreateSecretParamReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{13}
}

func (m *CreateSecretParamReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretParamReq.Unmarshal(m, b)
}
func (m *CreateSecretParamReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretParamReq.Marshal(b, m, deterministic)
}
func (m *CreateSecretParamReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretParamReq.Merge(m, src)
}
func (m *CreateSecretParamReq) XXX_Size() int {
	return xxx_messageInfo_CreateSecretParamReq.Size(m)
}
func (m *CreateSecretParamReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretParamReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretParamReq proto.InternalMessageInfo

func (m *CreateSecretParamReq) GetKeyPairId() string {
	if m != nil {
		return m.KeyPairId
	}
	return ""
}

func (m *CreateSecretParamReq) GetCipherText() string {
	if m != nil {
		return m.CipherText
	}
	return ""
}

func (m *CreateSecretParamReq) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *CreateSecretParamReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type CreateSecretParamResp struct {
	Status               *RespStatus      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	SecretParam          *SecretParamInfo `protobuf:"bytes,2,opt,name=secretParam,proto3" json:"secretParam,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateSecretParamResp) Reset()         { *m = CreateSecretParamResp{} }
func (m *CreateSecretParamResp) String() string { return proto.CompactTextString(m) }
func (*CreateSecretParamResp) ProtoMessage()    {}
func (*CreateSecretParamResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9840c882bb9a7ad, []int{14}
}

func (m *CreateSecretParamResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSecretParamResp.Unmarshal(m, b)
}
func (m *CreateSecretParamResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSecretParamResp.Marshal(b, m, deterministic)
}
func (m *CreateSecretParamResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSecretParamResp.Merge(m, src)
}
func (m *CreateSecretParamResp) XXX_Size() int {
	return xxx_messageInfo_CreateSecretParamResp.Size(m)
}
func (m *CreateSecretParamResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSecretParamResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSecretParamResp proto.InternalMessageInfo

func (m *CreateSecretParamResp) GetStatus() *RespStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateSecretParamResp) GetSecretParam() *SecretParamInfo {
	if m != nil {
		return m.SecretParam
	}
	return nil
}

func init() {
	proto.RegisterType((*RespStatus)(nil), "protos.RespStatus")
	proto.RegisterType((*KeyInfo)(nil), "protos.KeyInfo")
	proto.RegisterType((*GenRsaKeyPairReq)(nil), "protos.GenRsaKeyPairReq")
	proto.RegisterType((*GenRsaKeyPairResp)(nil), "protos.GenRsaKeyPairResp")
	proto.RegisterType((*RemoveRsaKeyPairReq)(nil), "protos.RemoveRsaKeyPairReq")
	proto.RegisterType((*RemoveRsaKeyPairResp)(nil), "protos.RemoveRsaKeyPairResp")
	proto.RegisterType((*EncryptReq)(nil), "protos.EncryptReq")
	proto.RegisterType((*EncryptResp)(nil), "protos.EncryptResp")
	proto.RegisterType((*DecryptReq)(nil), "protos.DecryptReq")
	proto.RegisterType((*DecryptResp)(nil), "protos.DecryptResp")
	proto.RegisterType((*CheckKeyReq)(nil), "protos.CheckKeyReq")
	proto.RegisterType((*CheckKeyResp)(nil), "protos.CheckKeyResp")
	proto.RegisterType((*SecretParamInfo)(nil), "protos.SecretParamInfo")
	proto.RegisterType((*CreateSecretParamReq)(nil), "protos.CreateSecretParamReq")
	proto.RegisterType((*CreateSecretParamResp)(nil), "protos.CreateSecretParamResp")
}

func init() {
	proto.RegisterFile("agrpc.proto", fileDescriptor_d9840c882bb9a7ad)
}

var fileDescriptor_d9840c882bb9a7ad = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xa5, 0x94, 0xcf, 0x5b, 0x0c, 0x30, 0x60, 0xac, 0x2b, 0x92, 0xcd, 0x3c, 0xa1, 0x24, 0x6c,
	0x02, 0x89, 0xe2, 0x23, 0xe0, 0x17, 0xac, 0x1a, 0x32, 0x90, 0x18, 0x7d, 0x21, 0x43, 0xb9, 0x2e,
	0x23, 0xdb, 0x0f, 0x3b, 0xb3, 0x86, 0xbe, 0xf8, 0xee, 0x8b, 0xff, 0xc0, 0x27, 0xff, 0xa8, 0xe9,
	0xf4, 0x6b, 0xd8, 0x2e, 0xac, 0x8d, 0x4f, 0xdd, 0x3d, 0xf7, 0xde, 0x93, 0x33, 0xe7, 0xf6, 0x4c,
	0xc1, 0xe1, 0xbd, 0x38, 0xf2, 0xb6, 0xa2, 0x38, 0x54, 0x21, 0x99, 0xd1, 0x0f, 0x49, 0xdf, 0x01,
	0x30, 0x94, 0xd1, 0x89, 0xe2, 0x6a, 0x20, 0xc9, 0x3a, 0x80, 0xd4, 0xbf, 0x0e, 0xc2, 0x0b, 0x74,
	0xad, 0xb6, 0xb5, 0x31, 0xcd, 0x0c, 0x24, 0xad, 0x63, 0x1c, 0xbf, 0x47, 0x29, 0x79, 0x0f, 0xdd,
	0xc9, 0xb6, 0xb5, 0x31, 0xcf, 0x0c, 0x84, 0xfe, 0xb6, 0x60, 0xb6, 0x8b, 0xc9, 0x61, 0xf0, 0x25,
	0x24, 0x6b, 0x30, 0x7f, 0x85, 0xc9, 0x31, 0x17, 0xf1, 0xe1, 0x85, 0xa6, 0x9a, 0x67, 0x15, 0x90,
	0x56, 0xa3, 0xc1, 0x79, 0x5f, 0x78, 0x5d, 0x4c, 0x72, 0xa2, 0x0a, 0x20, 0x4f, 0x61, 0xc9, 0x8b,
	0x91, 0x2b, 0xbc, 0x38, 0x15, 0x3e, 0x4a, 0xc5, 0xfd, 0xc8, 0xb5, 0xdb, 0xd6, 0x86, 0xcd, 0x6a,
	0x78, 0xda, 0x8b, 0xd7, 0x91, 0x88, 0xcd, 0xde, 0xa9, 0xac, 0x77, 0x18, 0xa7, 0x47, 0xb0, 0xf4,
	0x06, 0x03, 0x26, 0x79, 0x37, 0x13, 0xc2, 0xf0, 0xdb, 0x18, 0x9d, 0x2e, 0xcc, 0x2a, 0xe1, 0x63,
	0x38, 0x50, 0x5a, 0xe5, 0x34, 0x2b, 0xfe, 0xd2, 0xaf, 0xb0, 0x3c, 0xc4, 0x25, 0x53, 0x31, 0x33,
	0x99, 0x5d, 0x9a, 0xc9, 0xd9, 0x26, 0x99, 0xdd, 0x72, 0xab, 0x32, 0x99, 0xe5, 0x1d, 0xe4, 0x09,
	0xcc, 0x5e, 0x65, 0x5e, 0x69, 0x6a, 0x67, 0x7b, 0xb1, 0x68, 0xce, 0x2d, 0x64, 0x45, 0x9d, 0xee,
	0xc0, 0x0a, 0x43, 0x3f, 0xfc, 0x8e, 0x0d, 0xa4, 0xd3, 0x7d, 0x58, 0xad, 0x0f, 0x35, 0xd3, 0x48,
	0xdf, 0x02, 0xbc, 0x0a, 0xbc, 0x38, 0x89, 0xd4, 0x78, 0xab, 0xd2, 0x95, 0xf6, 0xb9, 0x08, 0x4e,
	0xf1, 0x5a, 0x95, 0x2b, 0x2d, 0x00, 0xfa, 0x09, 0x9c, 0x92, 0xa9, 0xa1, 0x51, 0xeb, 0x00, 0x9e,
	0x88, 0x2e, 0x31, 0x36, 0x98, 0x0d, 0x84, 0x1e, 0x01, 0xbc, 0xc4, 0x7f, 0x14, 0x39, 0x8e, 0xeb,
	0x23, 0x38, 0x25, 0x57, 0x43, 0x99, 0x77, 0x9f, 0x7f, 0x13, 0x9c, 0x83, 0x4b, 0xf4, 0xae, 0xba,
	0x98, 0x8c, 0x5f, 0x5d, 0x0f, 0x16, 0xaa, 0xe6, 0x86, 0x32, 0x36, 0x61, 0x2e, 0x7f, 0x6d, 0xa4,
	0x3b, 0xd9, 0xb6, 0x47, 0xbd, 0x57, 0x65, 0x03, 0xfd, 0x69, 0xc1, 0xe2, 0x09, 0x7a, 0x31, 0xaa,
	0x63, 0x1e, 0x73, 0x5f, 0x07, 0x37, 0xbd, 0x04, 0x34, 0xf4, 0x81, 0xfb, 0x98, 0x6b, 0x33, 0x90,
	0x91, 0xe1, 0x9c, 0x6c, 0x10, 0x4e, 0xfb, 0x96, 0x70, 0xfe, 0xb2, 0x60, 0xf5, 0x40, 0x13, 0x18,
	0x8a, 0xfe, 0x7b, 0xa3, 0x43, 0xc7, 0xb1, 0x6b, 0xc7, 0x31, 0x12, 0x3e, 0x75, 0x33, 0xe1, 0x3f,
	0xe0, 0xfe, 0x08, 0x3d, 0x0d, 0xd7, 0xf1, 0x02, 0x1c, 0x59, 0x8d, 0xe7, 0x49, 0x7f, 0x50, 0x0c,
	0x0c, 0x79, 0xcf, 0xcc, 0xde, 0xed, 0x3f, 0x36, 0x38, 0x7b, 0x52, 0x0a, 0xa9, 0xf6, 0x7a, 0x18,
	0x28, 0xf2, 0x1a, 0xee, 0xdd, 0xb8, 0x71, 0x88, 0x5b, 0xd0, 0x0c, 0x5f, 0x6a, 0xad, 0x87, 0xb7,
	0x54, 0x64, 0x44, 0x27, 0x48, 0x17, 0x16, 0x98, 0x6f, 0xd0, 0x3c, 0xaa, 0xe4, 0xd7, 0xee, 0x98,
	0xd6, 0xda, 0xed, 0x45, 0x4d, 0xb6, 0x5b, 0xe6, 0x5a, 0xbb, 0x5d, 0x5a, 0x51, 0x5d, 0x1b, 0xad,
	0x95, 0x1a, 0x56, 0x4c, 0xe6, 0x51, 0xbb, 0x39, 0x59, 0x65, 0xb9, 0x9a, 0x34, 0x32, 0x49, 0x27,
	0xc8, 0x73, 0x98, 0x2b, 0xe2, 0x41, 0xca, 0x16, 0x23, 0x5d, 0xad, 0xd5, 0x3a, 0xa8, 0x07, 0x19,
	0x2c, 0xd7, 0x36, 0x4a, 0xca, 0x13, 0x8e, 0x7a, 0xf9, 0x5a, 0x8f, 0xef, 0xa8, 0xa6, 0x9c, 0xfb,
	0xbb, 0x9f, 0x9f, 0xf5, 0x84, 0xba, 0x1c, 0x9c, 0x6f, 0x79, 0xa1, 0xdf, 0xe1, 0x7d, 0x91, 0x0c,
	0x82, 0xfc, 0x71, 0xc6, 0xf5, 0xf6, 0xce, 0xbc, 0xbe, 0xc0, 0x40, 0x75, 0x78, 0xba, 0xc5, 0x8e,
	0x88, 0xbc, 0x8e, 0xfe, 0x12, 0x9f, 0x67, 0xdf, 0xe0, 0x9d, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0xde, 0x84, 0x5e, 0x99, 0x07, 0x00, 0x00,
}
