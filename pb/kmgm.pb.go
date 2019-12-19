// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kmgm.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticationType int32

const (
	AuthenticationType_UNSPECIFIED     AuthenticationType = 0
	AuthenticationType_ANONYMOUS       AuthenticationType = 1
	AuthenticationType_BOOTSTRAP_TOKEN AuthenticationType = 2
	AuthenticationType_CLIENT_CERT     AuthenticationType = 3
)

var AuthenticationType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "ANONYMOUS",
	2: "BOOTSTRAP_TOKEN",
	3: "CLIENT_CERT",
}

var AuthenticationType_value = map[string]int32{
	"UNSPECIFIED":     0,
	"ANONYMOUS":       1,
	"BOOTSTRAP_TOKEN": 2,
	"CLIENT_CERT":     3,
}

func (x AuthenticationType) String() string {
	return proto.EnumName(AuthenticationType_name, int32(x))
}

func (AuthenticationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{0}
}

type HelloRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

type HelloResponse struct {
	ApiVersion           int32              `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	AuthenticationType   AuthenticationType `protobuf:"varint,2,opt,name=authentication_type,json=authenticationType,proto3,enum=kmgm.AuthenticationType" json:"authentication_type,omitempty"`
	AuthenticatedUser    string             `protobuf:"bytes,3,opt,name=authenticated_user,json=authenticatedUser,proto3" json:"authenticated_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetApiVersion() int32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *HelloResponse) GetAuthenticationType() AuthenticationType {
	if m != nil {
		return m.AuthenticationType
	}
	return AuthenticationType_UNSPECIFIED
}

func (m *HelloResponse) GetAuthenticatedUser() string {
	if m != nil {
		return m.AuthenticatedUser
	}
	return ""
}

type GetVersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionRequest) Reset()         { *m = GetVersionRequest{} }
func (m *GetVersionRequest) String() string { return proto.CompactTextString(m) }
func (*GetVersionRequest) ProtoMessage()    {}
func (*GetVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{2}
}

func (m *GetVersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionRequest.Unmarshal(m, b)
}
func (m *GetVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionRequest.Marshal(b, m, deterministic)
}
func (m *GetVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionRequest.Merge(m, src)
}
func (m *GetVersionRequest) XXX_Size() int {
	return xxx_messageInfo_GetVersionRequest.Size(m)
}
func (m *GetVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionRequest proto.InternalMessageInfo

type GetVersionResponse struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Commit               string   `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVersionResponse) Reset()         { *m = GetVersionResponse{} }
func (m *GetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*GetVersionResponse) ProtoMessage()    {}
func (*GetVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{3}
}

func (m *GetVersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVersionResponse.Unmarshal(m, b)
}
func (m *GetVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVersionResponse.Marshal(b, m, deterministic)
}
func (m *GetVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVersionResponse.Merge(m, src)
}
func (m *GetVersionResponse) XXX_Size() int {
	return xxx_messageInfo_GetVersionResponse.Size(m)
}
func (m *GetVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVersionResponse proto.InternalMessageInfo

func (m *GetVersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GetVersionResponse) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

type DistinguishedName struct {
	CommonName           string   `protobuf:"bytes,1,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	Organization         string   `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	OrganizationalUnit   string   `protobuf:"bytes,3,opt,name=organizational_unit,json=organizationalUnit,proto3" json:"organizational_unit,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Locality             string   `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	Province             string   `protobuf:"bytes,6,opt,name=province,proto3" json:"province,omitempty"`
	StreetAddress        string   `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode           string   `protobuf:"bytes,8,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistinguishedName) Reset()         { *m = DistinguishedName{} }
func (m *DistinguishedName) String() string { return proto.CompactTextString(m) }
func (*DistinguishedName) ProtoMessage()    {}
func (*DistinguishedName) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{4}
}

func (m *DistinguishedName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistinguishedName.Unmarshal(m, b)
}
func (m *DistinguishedName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistinguishedName.Marshal(b, m, deterministic)
}
func (m *DistinguishedName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistinguishedName.Merge(m, src)
}
func (m *DistinguishedName) XXX_Size() int {
	return xxx_messageInfo_DistinguishedName.Size(m)
}
func (m *DistinguishedName) XXX_DiscardUnknown() {
	xxx_messageInfo_DistinguishedName.DiscardUnknown(m)
}

var xxx_messageInfo_DistinguishedName proto.InternalMessageInfo

func (m *DistinguishedName) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *DistinguishedName) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *DistinguishedName) GetOrganizationalUnit() string {
	if m != nil {
		return m.OrganizationalUnit
	}
	return ""
}

func (m *DistinguishedName) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *DistinguishedName) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *DistinguishedName) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *DistinguishedName) GetStreetAddress() string {
	if m != nil {
		return m.StreetAddress
	}
	return ""
}

func (m *DistinguishedName) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

type Names struct {
	// FIXME[P1]: should be plural. see style guide.
	Dnsname              []string `protobuf:"bytes,1,rep,name=dnsname,proto3" json:"dnsname,omitempty"`
	Ipaddr               []string `protobuf:"bytes,2,rep,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Names) Reset()         { *m = Names{} }
func (m *Names) String() string { return proto.CompactTextString(m) }
func (*Names) ProtoMessage()    {}
func (*Names) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{5}
}

func (m *Names) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Names.Unmarshal(m, b)
}
func (m *Names) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Names.Marshal(b, m, deterministic)
}
func (m *Names) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Names.Merge(m, src)
}
func (m *Names) XXX_Size() int {
	return xxx_messageInfo_Names.Size(m)
}
func (m *Names) XXX_DiscardUnknown() {
	xxx_messageInfo_Names.DiscardUnknown(m)
}

var xxx_messageInfo_Names proto.InternalMessageInfo

func (m *Names) GetDnsname() []string {
	if m != nil {
		return m.Dnsname
	}
	return nil
}

func (m *Names) GetIpaddr() []string {
	if m != nil {
		return m.Ipaddr
	}
	return nil
}

type KeyUsage struct {
	KeyUsage             uint32   `protobuf:"varint,1,opt,name=key_usage,json=keyUsage,proto3" json:"key_usage,omitempty"`
	ExtKeyUsages         []uint32 `protobuf:"varint,2,rep,packed,name=ext_key_usages,json=extKeyUsages,proto3" json:"ext_key_usages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyUsage) Reset()         { *m = KeyUsage{} }
func (m *KeyUsage) String() string { return proto.CompactTextString(m) }
func (*KeyUsage) ProtoMessage()    {}
func (*KeyUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{6}
}

func (m *KeyUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyUsage.Unmarshal(m, b)
}
func (m *KeyUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyUsage.Marshal(b, m, deterministic)
}
func (m *KeyUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyUsage.Merge(m, src)
}
func (m *KeyUsage) XXX_Size() int {
	return xxx_messageInfo_KeyUsage.Size(m)
}
func (m *KeyUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyUsage.DiscardUnknown(m)
}

var xxx_messageInfo_KeyUsage proto.InternalMessageInfo

func (m *KeyUsage) GetKeyUsage() uint32 {
	if m != nil {
		return m.KeyUsage
	}
	return 0
}

func (m *KeyUsage) GetExtKeyUsages() []uint32 {
	if m != nil {
		return m.ExtKeyUsages
	}
	return nil
}

type IssueCertificateRequest struct {
	PublicKey            []byte             `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Subject              *DistinguishedName `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Names                *Names             `protobuf:"bytes,3,opt,name=names,proto3" json:"names,omitempty"`
	NotAfterUnixtime     int64              `protobuf:"varint,4,opt,name=not_after_unixtime,json=notAfterUnixtime,proto3" json:"not_after_unixtime,omitempty"`
	KeyUsage             *KeyUsage          `protobuf:"bytes,5,opt,name=key_usage,json=keyUsage,proto3" json:"key_usage,omitempty"`
	Profile              string             `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IssueCertificateRequest) Reset()         { *m = IssueCertificateRequest{} }
func (m *IssueCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*IssueCertificateRequest) ProtoMessage()    {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{7}
}

func (m *IssueCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueCertificateRequest.Unmarshal(m, b)
}
func (m *IssueCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueCertificateRequest.Marshal(b, m, deterministic)
}
func (m *IssueCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueCertificateRequest.Merge(m, src)
}
func (m *IssueCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_IssueCertificateRequest.Size(m)
}
func (m *IssueCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueCertificateRequest proto.InternalMessageInfo

func (m *IssueCertificateRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *IssueCertificateRequest) GetSubject() *DistinguishedName {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *IssueCertificateRequest) GetNames() *Names {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *IssueCertificateRequest) GetNotAfterUnixtime() int64 {
	if m != nil {
		return m.NotAfterUnixtime
	}
	return 0
}

func (m *IssueCertificateRequest) GetKeyUsage() *KeyUsage {
	if m != nil {
		return m.KeyUsage
	}
	return nil
}

func (m *IssueCertificateRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type IssueCertificateResponse struct {
	Certificate          []byte   `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueCertificateResponse) Reset()         { *m = IssueCertificateResponse{} }
func (m *IssueCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*IssueCertificateResponse) ProtoMessage()    {}
func (*IssueCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5b8bd6158ea4246, []int{8}
}

func (m *IssueCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueCertificateResponse.Unmarshal(m, b)
}
func (m *IssueCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueCertificateResponse.Marshal(b, m, deterministic)
}
func (m *IssueCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueCertificateResponse.Merge(m, src)
}
func (m *IssueCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_IssueCertificateResponse.Size(m)
}
func (m *IssueCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IssueCertificateResponse proto.InternalMessageInfo

func (m *IssueCertificateResponse) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func init() {
	proto.RegisterEnum("kmgm.AuthenticationType", AuthenticationType_name, AuthenticationType_value)
	proto.RegisterType((*HelloRequest)(nil), "kmgm.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "kmgm.HelloResponse")
	proto.RegisterType((*GetVersionRequest)(nil), "kmgm.GetVersionRequest")
	proto.RegisterType((*GetVersionResponse)(nil), "kmgm.GetVersionResponse")
	proto.RegisterType((*DistinguishedName)(nil), "kmgm.DistinguishedName")
	proto.RegisterType((*Names)(nil), "kmgm.Names")
	proto.RegisterType((*KeyUsage)(nil), "kmgm.KeyUsage")
	proto.RegisterType((*IssueCertificateRequest)(nil), "kmgm.IssueCertificateRequest")
	proto.RegisterType((*IssueCertificateResponse)(nil), "kmgm.IssueCertificateResponse")
}

func init() { proto.RegisterFile("kmgm.proto", fileDescriptor_d5b8bd6158ea4246) }

var fileDescriptor_d5b8bd6158ea4246 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x6f, 0x1a, 0x47,
	0x14, 0x2d, 0x60, 0x0c, 0x5c, 0x3e, 0x8c, 0x07, 0xa9, 0x5e, 0x51, 0xb9, 0xa5, 0xab, 0x56, 0xb2,
	0xfa, 0xe1, 0xb6, 0xf4, 0xa9, 0x52, 0xa5, 0x04, 0x63, 0x9c, 0x20, 0xc7, 0x60, 0x2f, 0x10, 0x29,
	0x7e, 0x59, 0x2d, 0xcb, 0x35, 0x9e, 0x78, 0xd9, 0xd9, 0xec, 0xcc, 0x5a, 0x26, 0xaf, 0xf9, 0x29,
	0x79, 0xcb, 0xaf, 0x8c, 0xe6, 0x63, 0xed, 0x25, 0xc8, 0x6f, 0x7b, 0xcf, 0x99, 0x3d, 0x73, 0xcf,
	0xb9, 0x33, 0x03, 0x70, 0xb7, 0x5a, 0xae, 0x8e, 0xa3, 0x98, 0x09, 0x46, 0x76, 0xe4, 0xb7, 0xdd,
	0x80, 0xda, 0x6b, 0x0c, 0x02, 0xe6, 0xe0, 0x87, 0x04, 0xb9, 0xb0, 0xbf, 0xe4, 0xa0, 0x6e, 0x00,
	0x1e, 0xb1, 0x90, 0x23, 0xf9, 0x09, 0xaa, 0x5e, 0x44, 0xdd, 0x7b, 0x8c, 0x39, 0x65, 0xa1, 0x95,
	0xeb, 0xe4, 0x8e, 0x8a, 0x0e, 0x78, 0x11, 0x7d, 0xab, 0x11, 0x32, 0x84, 0x96, 0x97, 0x88, 0x5b,
	0x0c, 0x05, 0xf5, 0x3d, 0x41, 0x59, 0xe8, 0x8a, 0x75, 0x84, 0x56, 0xbe, 0x93, 0x3b, 0x6a, 0x74,
	0xad, 0x63, 0xb5, 0x65, 0x6f, 0x63, 0xc1, 0x74, 0x1d, 0xa1, 0x43, 0xbc, 0x2d, 0x8c, 0xfc, 0x09,
	0x59, 0x14, 0x17, 0x6e, 0xc2, 0x31, 0xb6, 0x0a, 0x9d, 0xdc, 0x51, 0xc5, 0xd9, 0xdf, 0x60, 0x66,
	0x1c, 0x63, 0xbb, 0x05, 0xfb, 0xaf, 0x50, 0x98, 0x3e, 0x52, 0x07, 0x67, 0x40, 0xb2, 0xa0, 0x71,
	0x61, 0x41, 0x29, 0xeb, 0xa0, 0xe2, 0xa4, 0x25, 0xf9, 0x1e, 0x76, 0x7d, 0xb6, 0x5a, 0x51, 0xa1,
	0x3a, 0xae, 0x38, 0xa6, 0xb2, 0x3f, 0xe7, 0x61, 0xff, 0x94, 0x72, 0x41, 0xc3, 0x65, 0x42, 0xf9,
	0x2d, 0x2e, 0x46, 0xde, 0x4a, 0xa5, 0x21, 0x79, 0x16, 0xba, 0xa1, 0xb7, 0x42, 0xa3, 0x05, 0x1a,
	0x52, 0x0b, 0x6c, 0xa8, 0xb1, 0x78, 0xe9, 0x85, 0xf4, 0xa3, 0xb2, 0x65, 0x44, 0x37, 0x30, 0xf2,
	0x17, 0xb4, 0xb2, 0xb5, 0x17, 0xb8, 0x49, 0x48, 0x85, 0xf1, 0x49, 0x36, 0xa9, 0x59, 0x48, 0x85,
	0xec, 0xde, 0x67, 0x49, 0x28, 0xe2, 0xb5, 0xb5, 0xa3, 0xbb, 0x37, 0x25, 0x69, 0x43, 0x39, 0x60,
	0xbe, 0x17, 0x50, 0xb1, 0xb6, 0x8a, 0x8a, 0x7a, 0xac, 0x25, 0x17, 0xc5, 0xec, 0x9e, 0x86, 0x3e,
	0x5a, 0xbb, 0x9a, 0x4b, 0x6b, 0xf2, 0x2b, 0x34, 0xb8, 0x88, 0x11, 0x85, 0xeb, 0x2d, 0x16, 0x31,
	0x72, 0x6e, 0x95, 0xd4, 0x8a, 0xba, 0x46, 0x7b, 0x1a, 0x94, 0x76, 0x23, 0xc6, 0x85, 0x17, 0xb8,
	0x3e, 0x5b, 0xa0, 0x55, 0xd6, 0x76, 0x35, 0xd4, 0x67, 0x0b, 0xb4, 0xff, 0x83, 0xa2, 0xb4, 0xcd,
	0x65, 0x8b, 0x8b, 0x90, 0x9b, 0x50, 0x0a, 0xb2, 0x45, 0x53, 0xca, 0x80, 0x69, 0x24, 0x77, 0xb1,
	0xf2, 0x8a, 0x30, 0x95, 0x7d, 0x01, 0xe5, 0x73, 0x5c, 0xcf, 0xb8, 0xb7, 0x44, 0xf2, 0x03, 0x54,
	0xee, 0x70, 0xed, 0x26, 0xb2, 0x50, 0xa1, 0xd6, 0x9d, 0xf2, 0x5d, 0x4a, 0xfe, 0x02, 0x0d, 0x7c,
	0x10, 0xee, 0xe3, 0x02, 0xae, 0x84, 0xea, 0x4e, 0x0d, 0x1f, 0x44, 0xaa, 0xc0, 0xed, 0x4f, 0x79,
	0x38, 0x18, 0x72, 0x9e, 0x60, 0x1f, 0x63, 0x41, 0x6f, 0xd4, 0x31, 0x31, 0x67, 0x82, 0x1c, 0x02,
	0x44, 0xc9, 0x3c, 0xa0, 0xbe, 0x14, 0x51, 0xfa, 0x35, 0xa7, 0xa2, 0x91, 0x73, 0x5c, 0x93, 0x7f,
	0xa0, 0xc4, 0x93, 0xf9, 0x7b, 0xf4, 0xf5, 0x19, 0xa8, 0x76, 0x0f, 0xf4, 0xa9, 0xdd, 0x1a, 0xbf,
	0x93, 0xae, 0x23, 0x3f, 0x43, 0x51, 0x9a, 0xe3, 0x6a, 0x68, 0xd5, 0x6e, 0x55, 0xff, 0xa0, 0xa2,
	0x70, 0x34, 0x43, 0xfe, 0x00, 0x12, 0x32, 0xe1, 0x7a, 0x37, 0x02, 0x63, 0x39, 0xe0, 0x07, 0x41,
	0x57, 0xa8, 0xe6, 0x57, 0x70, 0x9a, 0x21, 0x13, 0x3d, 0x49, 0xcc, 0x0c, 0x4e, 0x7e, 0xcf, 0x26,
	0x50, 0x54, 0xa2, 0x0d, 0x2d, 0x9a, 0x5a, 0xcc, 0x24, 0x62, 0x41, 0x29, 0x8a, 0xd9, 0x0d, 0x0d,
	0xd2, 0xc1, 0xa6, 0xa5, 0xfd, 0x3f, 0x58, 0xdb, 0x21, 0x98, 0x3b, 0xd0, 0x81, 0xaa, 0xff, 0x04,
	0x9b, 0x18, 0xb2, 0xd0, 0x6f, 0xd7, 0x40, 0xb6, 0x6f, 0x2a, 0xd9, 0x83, 0xea, 0x6c, 0x34, 0xb9,
	0x1c, 0xf4, 0x87, 0x67, 0xc3, 0xc1, 0x69, 0xf3, 0x3b, 0x52, 0x87, 0x4a, 0x6f, 0x34, 0x1e, 0xbd,
	0xbb, 0x18, 0xcf, 0x26, 0xcd, 0x1c, 0x69, 0xc1, 0xde, 0xc9, 0x78, 0x3c, 0x9d, 0x4c, 0x9d, 0xde,
	0xa5, 0x3b, 0x1d, 0x9f, 0x0f, 0x46, 0xcd, 0xbc, 0xfc, 0xa9, 0xff, 0x66, 0x38, 0x18, 0x4d, 0xdd,
	0xfe, 0xc0, 0x99, 0x36, 0x0b, 0xdd, 0x97, 0xe6, 0xa5, 0x99, 0x60, 0x7c, 0x4f, 0x7d, 0x24, 0x7f,
	0x43, 0x51, 0xd5, 0x84, 0x68, 0x9b, 0xd9, 0x67, 0xa8, 0xdd, 0xda, 0xc0, 0x74, 0xff, 0xdd, 0x2b,
	0x68, 0x98, 0x6b, 0x9d, 0x6a, 0xbc, 0x00, 0x78, 0xba, 0xeb, 0xc4, 0x4c, 0x6d, 0xeb, 0x49, 0x68,
	0x5b, 0xdb, 0x84, 0x91, 0x5c, 0x02, 0xc9, 0x24, 0x95, 0xca, 0x5e, 0x41, 0xf3, 0xdb, 0x10, 0xc9,
	0xa1, 0xd6, 0x78, 0xe6, 0x84, 0xb5, 0x7f, 0x7c, 0x8e, 0xd6, 0x1b, 0x9d, 0xec, 0x5c, 0xe7, 0xa3,
	0xf9, 0x7c, 0x57, 0x3d, 0xbd, 0xff, 0x7e, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x36, 0xc2, 0x50,
	0x88, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/kmgm.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionServiceClient interface {
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type versionServiceClient struct {
	cc *grpc.ClientConn
}

func NewVersionServiceClient(cc *grpc.ClientConn) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/kmgm.VersionService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
type VersionServiceServer interface {
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
}

func RegisterVersionServiceServer(s *grpc.Server, srv VersionServiceServer) {
	s.RegisterService(&_VersionService_serviceDesc, srv)
}

func _VersionService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.VersionService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _VersionService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateServiceClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error)
}

type certificateServiceClient struct {
	cc *grpc.ClientConn
}

func NewCertificateServiceClient(cc *grpc.ClientConn) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssueCertificateResponse, error) {
	out := new(IssueCertificateResponse)
	err := c.cc.Invoke(ctx, "/kmgm.CertificateService/IssueCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
type CertificateServiceServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*IssueCertificateResponse, error)
}

func RegisterCertificateServiceServer(s *grpc.Server, srv CertificateServiceServer) {
	s.RegisterService(&_CertificateService_serviceDesc, srv)
}

func _CertificateService_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kmgm.CertificateService/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kmgm.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateService_IssueCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kmgm.proto",
}
