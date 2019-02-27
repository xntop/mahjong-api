// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/base.proto

package base

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

type Wind int32

const (
	Wind_WIND_UNSPECIFIED Wind = 0
	Wind_EAST             Wind = 1
	Wind_SOUTH            Wind = 2
	Wind_WEST             Wind = 3
	Wind_NORTH            Wind = 4
)

var Wind_name = map[int32]string{
	0: "WIND_UNSPECIFIED",
	1: "EAST",
	2: "SOUTH",
	3: "WEST",
	4: "NORTH",
}

var Wind_value = map[string]int32{
	"WIND_UNSPECIFIED": 0,
	"EAST":             1,
	"SOUTH":            2,
	"WEST":             3,
	"NORTH":            4,
}

func (x Wind) String() string {
	return proto.EnumName(Wind_name, int32(x))
}

func (Wind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{0}
}

type Limit int32

const (
	Limit_LIMIT_UNSPECIFIED Limit = 0
	Limit_MANGAN            Limit = 1
	Limit_HANEMAN           Limit = 2
	Limit_BAIMAN            Limit = 3
	Limit_SANBAIMAN         Limit = 4
	Limit_YAKUMAN           Limit = 5
)

var Limit_name = map[int32]string{
	0: "LIMIT_UNSPECIFIED",
	1: "MANGAN",
	2: "HANEMAN",
	3: "BAIMAN",
	4: "SANBAIMAN",
	5: "YAKUMAN",
}

var Limit_value = map[string]int32{
	"LIMIT_UNSPECIFIED": 0,
	"MANGAN":            1,
	"HANEMAN":           2,
	"BAIMAN":            3,
	"SANBAIMAN":         4,
	"YAKUMAN":           5,
}

func (x Limit) String() string {
	return proto.EnumName(Limit_name, int32(x))
}

func (Limit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{1}
}

type Opponent int32

const (
	Opponent_OPPONENT_UNSPECIFIED Opponent = 0
	Opponent_SELF                 Opponent = 1
	Opponent_RIGHT                Opponent = 2
	Opponent_FRONT                Opponent = 3
	Opponent_LEFT                 Opponent = 4
)

var Opponent_name = map[int32]string{
	0: "OPPONENT_UNSPECIFIED",
	1: "SELF",
	2: "RIGHT",
	3: "FRONT",
	4: "LEFT",
}

var Opponent_value = map[string]int32{
	"OPPONENT_UNSPECIFIED": 0,
	"SELF":                 1,
	"RIGHT":                2,
	"FRONT":                3,
	"LEFT":                 4,
}

func (x Opponent) String() string {
	return proto.EnumName(Opponent_name, int32(x))
}

func (Opponent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{2}
}

// Special holder to be used in repeat.
type Instances struct {
	Instances            []int64  `protobuf:"varint,1,rep,packed,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instances) Reset()         { *m = Instances{} }
func (m *Instances) String() string { return proto.CompactTextString(m) }
func (*Instances) ProtoMessage()    {}
func (*Instances) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{0}
}

func (m *Instances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instances.Unmarshal(m, b)
}
func (m *Instances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instances.Marshal(b, m, deterministic)
}
func (m *Instances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instances.Merge(m, src)
}
func (m *Instances) XXX_Size() int {
	return xxx_messageInfo_Instances.Size(m)
}
func (m *Instances) XXX_DiscardUnknown() {
	xxx_messageInfo_Instances.DiscardUnknown(m)
}

var xxx_messageInfo_Instances proto.InternalMessageInfo

func (m *Instances) GetInstances() []int64 {
	if m != nil {
		return m.Instances
	}
	return nil
}

type Meld struct {
	// Oponent is set, related to one who declared
	Opponent       Opponent `protobuf:"varint,1,opt,name=opponent,proto3,enum=base.Opponent" json:"opponent,omitempty"`
	CalledInstance int64    `protobuf:"varint,2,opt,name=called_instance,json=calledInstance,proto3" json:"called_instance,omitempty"`
	// Only for upgraded pon
	UpgradeInstance int64      `protobuf:"varint,3,opt,name=upgrade_instance,json=upgradeInstance,proto3" json:"upgrade_instance,omitempty"`
	Hand            *Instances `protobuf:"bytes,4,opt,name=hand,proto3" json:"hand,omitempty"`
	// To identify melds on the table.
	// In case of pon upgrade the meld_id will correspond to upgraded pon.
	MeldId int64 `protobuf:"varint,5,opt,name=meld_id,json=meldId,proto3" json:"meld_id,omitempty"`
	// When closed kan declared (or upgrade)
	IncludesTsumo        bool     `protobuf:"varint,6,opt,name=includes_tsumo,json=includesTsumo,proto3" json:"includes_tsumo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meld) Reset()         { *m = Meld{} }
func (m *Meld) String() string { return proto.CompactTextString(m) }
func (*Meld) ProtoMessage()    {}
func (*Meld) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66ec2e140567106, []int{1}
}

func (m *Meld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meld.Unmarshal(m, b)
}
func (m *Meld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meld.Marshal(b, m, deterministic)
}
func (m *Meld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meld.Merge(m, src)
}
func (m *Meld) XXX_Size() int {
	return xxx_messageInfo_Meld.Size(m)
}
func (m *Meld) XXX_DiscardUnknown() {
	xxx_messageInfo_Meld.DiscardUnknown(m)
}

var xxx_messageInfo_Meld proto.InternalMessageInfo

func (m *Meld) GetOpponent() Opponent {
	if m != nil {
		return m.Opponent
	}
	return Opponent_OPPONENT_UNSPECIFIED
}

func (m *Meld) GetCalledInstance() int64 {
	if m != nil {
		return m.CalledInstance
	}
	return 0
}

func (m *Meld) GetUpgradeInstance() int64 {
	if m != nil {
		return m.UpgradeInstance
	}
	return 0
}

func (m *Meld) GetHand() *Instances {
	if m != nil {
		return m.Hand
	}
	return nil
}

func (m *Meld) GetMeldId() int64 {
	if m != nil {
		return m.MeldId
	}
	return 0
}

func (m *Meld) GetIncludesTsumo() bool {
	if m != nil {
		return m.IncludesTsumo
	}
	return false
}

func init() {
	proto.RegisterEnum("base.Wind", Wind_name, Wind_value)
	proto.RegisterEnum("base.Limit", Limit_name, Limit_value)
	proto.RegisterEnum("base.Opponent", Opponent_name, Opponent_value)
	proto.RegisterType((*Instances)(nil), "base.Instances")
	proto.RegisterType((*Meld)(nil), "base.Meld")
}

func init() { proto.RegisterFile("base/base.proto", fileDescriptor_d66ec2e140567106) }

var fileDescriptor_d66ec2e140567106 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x6e, 0xda, 0x30,
	0x18, 0x85, 0x17, 0x62, 0x28, 0xfc, 0x55, 0xc1, 0xb3, 0x3a, 0x2d, 0x17, 0xbb, 0x88, 0x3a, 0x4d,
	0x4b, 0x91, 0x06, 0x53, 0xf7, 0x04, 0xe9, 0x9a, 0x94, 0x68, 0xe0, 0xa0, 0x24, 0x08, 0x6d, 0x37,
	0x51, 0xc0, 0x2e, 0x78, 0x4b, 0xec, 0xa8, 0x09, 0x7b, 0xe9, 0xbd, 0xc4, 0xe4, 0x40, 0xa8, 0xb4,
	0xdd, 0x44, 0xe7, 0x7c, 0xff, 0xd1, 0xf1, 0x1f, 0xcb, 0x30, 0xda, 0x64, 0x15, 0x9f, 0xea, 0xcf,
	0xa4, 0x7c, 0x56, 0xb5, 0x22, 0x48, 0xeb, 0x9b, 0x5b, 0x18, 0x04, 0xb2, 0xaa, 0x33, 0xb9, 0xe5,
	0x15, 0x79, 0x07, 0x03, 0xd1, 0x1a, 0xcb, 0xb0, 0x4d, 0xc7, 0x8c, 0x5e, 0xc0, 0xcd, 0x1f, 0x03,
	0xd0, 0x82, 0xe7, 0x8c, 0x8c, 0xa1, 0xaf, 0xca, 0x52, 0x49, 0x2e, 0x6b, 0xcb, 0xb0, 0x0d, 0x67,
	0x78, 0x37, 0x9c, 0x34, 0xc5, 0xe1, 0x89, 0x46, 0xe7, 0x39, 0xf9, 0x08, 0xa3, 0x6d, 0x96, 0xe7,
	0x9c, 0xa5, 0x6d, 0x91, 0xd5, 0xb1, 0x0d, 0xc7, 0x8c, 0x86, 0x47, 0xdc, 0x1e, 0x4e, 0x6e, 0x01,
	0x1f, 0xca, 0xdd, 0x73, 0xc6, 0xf8, 0x4b, 0xd2, 0x6c, 0x92, 0xa3, 0x13, 0x3f, 0x47, 0xdf, 0x03,
	0xda, 0x67, 0x92, 0x59, 0xc8, 0x36, 0x9c, 0xcb, 0xbb, 0xd1, 0xf1, 0xec, 0xf3, 0x5f, 0x44, 0xcd,
	0x90, 0xbc, 0x85, 0x8b, 0x82, 0xe7, 0x2c, 0x15, 0xcc, 0xea, 0x36, 0x35, 0x3d, 0x6d, 0x03, 0x46,
	0x3e, 0xc0, 0x50, 0xc8, 0x6d, 0x7e, 0x60, 0xbc, 0x4a, 0xeb, 0xea, 0x50, 0x28, 0xab, 0x67, 0x1b,
	0x4e, 0x3f, 0xba, 0x6a, 0x69, 0xa2, 0xe1, 0xd8, 0x07, 0xb4, 0x16, 0x92, 0x91, 0x6b, 0xc0, 0xeb,
	0x80, 0x3e, 0xa4, 0x2b, 0x1a, 0x2f, 0xbd, 0xaf, 0x81, 0x1f, 0x78, 0x0f, 0xf8, 0x15, 0xe9, 0x03,
	0xf2, 0xdc, 0x38, 0xc1, 0x06, 0x19, 0x40, 0x37, 0x0e, 0x57, 0xc9, 0x0c, 0x77, 0x34, 0x5c, 0x7b,
	0x71, 0x82, 0x4d, 0x0d, 0x69, 0x18, 0x25, 0x33, 0x8c, 0xc6, 0x29, 0x74, 0xe7, 0xa2, 0x10, 0x35,
	0x79, 0x03, 0xaf, 0xe7, 0xc1, 0x22, 0x48, 0xfe, 0x69, 0x02, 0xe8, 0x2d, 0x5c, 0xfa, 0xe8, 0x52,
	0x6c, 0x90, 0x4b, 0xb8, 0x98, 0xb9, 0xd4, 0x5b, 0xb8, 0x14, 0x77, 0xf4, 0xe0, 0xde, 0x0d, 0xb4,
	0x36, 0xc9, 0x15, 0x0c, 0x62, 0x97, 0x9e, 0x2c, 0xd2, 0xb9, 0xef, 0xee, 0xb7, 0x95, 0x36, 0xdd,
	0x31, 0x85, 0x7e, 0x7b, 0xef, 0xc4, 0x82, 0xeb, 0x70, 0xb9, 0x0c, 0xa9, 0x47, 0x93, 0xff, 0x17,
	0x8e, 0xbd, 0xb9, 0x7f, 0x5c, 0x38, 0x0a, 0x1e, 0x67, 0x09, 0xee, 0x68, 0xe9, 0x47, 0x21, 0xd5,
	0x1b, 0xf7, 0x01, 0xcd, 0x3d, 0x3f, 0xc1, 0xe8, 0xfe, 0xf3, 0x8f, 0xc9, 0x4e, 0xd4, 0xfb, 0xc3,
	0x66, 0xb2, 0x55, 0xc5, 0x94, 0x49, 0xf5, 0x5b, 0xfc, 0x52, 0x4f, 0x4f, 0xd3, 0x22, 0xdb, 0xff,
	0x54, 0x72, 0xf7, 0x29, 0x2b, 0xc5, 0x74, 0xc7, 0x65, 0xf3, 0x88, 0x9a, 0xf7, 0xb4, 0xe9, 0x35,
	0xfa, 0xcb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x61, 0x81, 0x27, 0x63, 0x02, 0x00, 0x00,
}