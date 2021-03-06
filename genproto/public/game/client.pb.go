// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public/game/client.proto

package game

import (
	fmt "fmt"
	base "github.com/dnovikoff/mahjong-api/genproto/public/base"
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

// Server will treat any incorrect answer as a cancel event.
// For example, when you try to Drop tile with riichi flag, on a wrong tile
// or when not in tempai - it will lead to tsumogiri.
// When you get Server response with, previously sent suggest_id and cancel=true,
// this means that the event is outdated (ex. priority action by other player or timeout).
// It is still safe to send answers on outdated event, in this case server
// just omits the answer and continues waiting for answer with valid suggest_id (if any).
type Client struct {
	// See Server.suggest.suggest_id
	SuggestId int64 `protobuf:"varint,1,opt,name=suggest_id,json=suggestId,proto3" json:"suggest_id,omitempty"`
	// Types that are valid to be assigned to OneofClient:
	//	*Client_Cancel
	//	*Client_Drop
	//	*Client_Call
	//	*Client_Win
	//	*Client_Draw
	//	*Client_Noten
	//	*Client_ContinueGame
	//	*Client_Settings
	OneofClient          isClient_OneofClient `protobuf_oneof:"oneof_client"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_98669709ec32e53a, []int{0}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetSuggestId() int64 {
	if m != nil {
		return m.SuggestId
	}
	return 0
}

type isClient_OneofClient interface {
	isClient_OneofClient()
}

type Client_Cancel struct {
	Cancel bool `protobuf:"varint,2,opt,name=cancel,proto3,oneof"`
}

type Client_Drop struct {
	Drop *ClientDrop `protobuf:"bytes,3,opt,name=drop,proto3,oneof"`
}

type Client_Call struct {
	Call *base.Instances `protobuf:"bytes,4,opt,name=call,proto3,oneof"`
}

type Client_Win struct {
	Win bool `protobuf:"varint,5,opt,name=win,proto3,oneof"`
}

type Client_Draw struct {
	Draw bool `protobuf:"varint,6,opt,name=draw,proto3,oneof"`
}

type Client_Noten struct {
	Noten bool `protobuf:"varint,7,opt,name=noten,proto3,oneof"`
}

type Client_ContinueGame struct {
	ContinueGame bool `protobuf:"varint,8,opt,name=continue_game,json=continueGame,proto3,oneof"`
}

type Client_Settings struct {
	Settings *Settings `protobuf:"bytes,9,opt,name=settings,proto3,oneof"`
}

func (*Client_Cancel) isClient_OneofClient() {}

func (*Client_Drop) isClient_OneofClient() {}

func (*Client_Call) isClient_OneofClient() {}

func (*Client_Win) isClient_OneofClient() {}

func (*Client_Draw) isClient_OneofClient() {}

func (*Client_Noten) isClient_OneofClient() {}

func (*Client_ContinueGame) isClient_OneofClient() {}

func (*Client_Settings) isClient_OneofClient() {}

func (m *Client) GetOneofClient() isClient_OneofClient {
	if m != nil {
		return m.OneofClient
	}
	return nil
}

func (m *Client) GetCancel() bool {
	if x, ok := m.GetOneofClient().(*Client_Cancel); ok {
		return x.Cancel
	}
	return false
}

func (m *Client) GetDrop() *ClientDrop {
	if x, ok := m.GetOneofClient().(*Client_Drop); ok {
		return x.Drop
	}
	return nil
}

func (m *Client) GetCall() *base.Instances {
	if x, ok := m.GetOneofClient().(*Client_Call); ok {
		return x.Call
	}
	return nil
}

func (m *Client) GetWin() bool {
	if x, ok := m.GetOneofClient().(*Client_Win); ok {
		return x.Win
	}
	return false
}

func (m *Client) GetDraw() bool {
	if x, ok := m.GetOneofClient().(*Client_Draw); ok {
		return x.Draw
	}
	return false
}

func (m *Client) GetNoten() bool {
	if x, ok := m.GetOneofClient().(*Client_Noten); ok {
		return x.Noten
	}
	return false
}

func (m *Client) GetContinueGame() bool {
	if x, ok := m.GetOneofClient().(*Client_ContinueGame); ok {
		return x.ContinueGame
	}
	return false
}

func (m *Client) GetSettings() *Settings {
	if x, ok := m.GetOneofClient().(*Client_Settings); ok {
		return x.Settings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Client) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Client_Cancel)(nil),
		(*Client_Drop)(nil),
		(*Client_Call)(nil),
		(*Client_Win)(nil),
		(*Client_Draw)(nil),
		(*Client_Noten)(nil),
		(*Client_ContinueGame)(nil),
		(*Client_Settings)(nil),
	}
}

type ClientDrop struct {
	Instance             int64    `protobuf:"varint,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Riichi               bool     `protobuf:"varint,2,opt,name=riichi,proto3" json:"riichi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientDrop) Reset()         { *m = ClientDrop{} }
func (m *ClientDrop) String() string { return proto.CompactTextString(m) }
func (*ClientDrop) ProtoMessage()    {}
func (*ClientDrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_98669709ec32e53a, []int{1}
}

func (m *ClientDrop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientDrop.Unmarshal(m, b)
}
func (m *ClientDrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientDrop.Marshal(b, m, deterministic)
}
func (m *ClientDrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientDrop.Merge(m, src)
}
func (m *ClientDrop) XXX_Size() int {
	return xxx_messageInfo_ClientDrop.Size(m)
}
func (m *ClientDrop) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientDrop.DiscardUnknown(m)
}

var xxx_messageInfo_ClientDrop proto.InternalMessageInfo

func (m *ClientDrop) GetInstance() int64 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *ClientDrop) GetRiichi() bool {
	if m != nil {
		return m.Riichi
	}
	return false
}

// Settings will be applied on server, even when disconnected.
// You will get confirmation on settings changes from server (see Server.settings).
type Settings struct {
	DropTsumo  bool `protobuf:"varint,1,opt,name=drop_tsumo,json=dropTsumo,proto3" json:"drop_tsumo,omitempty"`
	AutoWin    bool `protobuf:"varint,2,opt,name=auto_win,json=autoWin,proto3" json:"auto_win,omitempty"`
	NoDeclare  bool `protobuf:"varint,3,opt,name=no_declare,json=noDeclare,proto3" json:"no_declare,omitempty"`
	AutoTempai bool `protobuf:"varint,4,opt,name=auto_tempai,json=autoTempai,proto3" json:"auto_tempai,omitempty"`
	// For optional agari-yame
	AutoEndGame          bool     `protobuf:"varint,5,opt,name=auto_end_game,json=autoEndGame,proto3" json:"auto_end_game,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_98669709ec32e53a, []int{2}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetDropTsumo() bool {
	if m != nil {
		return m.DropTsumo
	}
	return false
}

func (m *Settings) GetAutoWin() bool {
	if m != nil {
		return m.AutoWin
	}
	return false
}

func (m *Settings) GetNoDeclare() bool {
	if m != nil {
		return m.NoDeclare
	}
	return false
}

func (m *Settings) GetAutoTempai() bool {
	if m != nil {
		return m.AutoTempai
	}
	return false
}

func (m *Settings) GetAutoEndGame() bool {
	if m != nil {
		return m.AutoEndGame
	}
	return false
}

func init() {
	proto.RegisterType((*Client)(nil), "mahjong.game.Client")
	proto.RegisterType((*ClientDrop)(nil), "mahjong.game.ClientDrop")
	proto.RegisterType((*Settings)(nil), "mahjong.game.Settings")
}

func init() { proto.RegisterFile("public/game/client.proto", fileDescriptor_98669709ec32e53a) }

var fileDescriptor_98669709ec32e53a = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x52, 0x5d, 0x6b, 0xdb, 0x30,
	0x14, 0x8d, 0x9b, 0xd4, 0x55, 0x6e, 0xd3, 0x3d, 0x88, 0x91, 0x69, 0x85, 0xb1, 0x10, 0x18, 0xe4,
	0xa5, 0x36, 0xec, 0x83, 0xbd, 0x8e, 0xae, 0x63, 0xe9, 0xab, 0x56, 0x18, 0xec, 0xc5, 0x28, 0xb2,
	0xe2, 0x68, 0xb3, 0xaf, 0x8c, 0x2d, 0xaf, 0x7f, 0x69, 0xff, 0x66, 0x7f, 0x69, 0xe8, 0x5a, 0xcd,
	0xda, 0x17, 0xc1, 0x39, 0xf7, 0xdc, 0xaf, 0x73, 0x05, 0xa2, 0x1d, 0x76, 0xb5, 0xd5, 0x79, 0xa5,
	0x1a, 0x93, 0xeb, 0xda, 0x1a, 0xf4, 0x59, 0xdb, 0x39, 0xef, 0xf8, 0xa2, 0x51, 0x87, 0x9f, 0x0e,
	0xab, 0x2c, 0x84, 0x2e, 0x97, 0x51, 0xb7, 0x53, 0xbd, 0xa1, 0x67, 0x54, 0xad, 0xff, 0x9e, 0x40,
	0xfa, 0x99, 0xd2, 0xf8, 0x2b, 0x80, 0x7e, 0xa8, 0x2a, 0xd3, 0xfb, 0xc2, 0x96, 0x22, 0x59, 0x25,
	0x9b, 0xa9, 0x9c, 0x47, 0xe6, 0xb6, 0xe4, 0x02, 0x52, 0xad, 0x50, 0x9b, 0x5a, 0x9c, 0xac, 0x92,
	0x0d, 0xdb, 0x4e, 0x64, 0xc4, 0x3c, 0x83, 0x59, 0xd9, 0xb9, 0x56, 0x4c, 0x57, 0xc9, 0xe6, 0xfc,
	0xad, 0xc8, 0x1e, 0x37, 0xce, 0xc6, 0xe2, 0x37, 0x9d, 0x6b, 0xb7, 0x13, 0x49, 0x3a, 0x7e, 0x05,
	0x33, 0xad, 0xea, 0x5a, 0xcc, 0x48, 0xff, 0xe2, 0xa8, 0xa7, 0xb1, 0x6e, 0xb1, 0xf7, 0xa1, 0x6c,
	0x1f, 0xe4, 0x41, 0xc6, 0x39, 0x4c, 0xef, 0x2d, 0x8a, 0xd3, 0xd8, 0x35, 0x00, 0xfe, 0x3c, 0xb4,
	0x54, 0xf7, 0x22, 0x8d, 0x24, 0x21, 0xbe, 0x84, 0x53, 0x74, 0xde, 0xa0, 0x38, 0x8b, 0xf4, 0x08,
	0xf9, 0x1b, 0xb8, 0xd0, 0x0e, 0xbd, 0xc5, 0xc1, 0x14, 0x61, 0x28, 0xc1, 0x62, 0x7c, 0xf1, 0x40,
	0x7f, 0x55, 0x8d, 0xe1, 0xef, 0x81, 0xf5, 0xc6, 0x7b, 0x8b, 0x55, 0x2f, 0xe6, 0x34, 0xdb, 0xf2,
	0xe9, 0x2e, 0xdf, 0x62, 0x74, 0x3b, 0x91, 0x47, 0xe5, 0xf5, 0x33, 0x58, 0x38, 0x34, 0x6e, 0x5f,
	0x8c, 0xee, 0xaf, 0x3f, 0x01, 0xfc, 0xdf, 0x99, 0x5f, 0x02, 0xb3, 0x71, 0xa3, 0x68, 0xe9, 0x11,
	0xf3, 0x25, 0xa4, 0x9d, 0xb5, 0xfa, 0x60, 0x47, 0x47, 0x65, 0x44, 0xeb, 0x3f, 0x09, 0xb0, 0x87,
	0x56, 0xe1, 0x2a, 0xc1, 0xb4, 0xc2, 0xf7, 0x43, 0xe3, 0xa8, 0x04, 0x93, 0xf3, 0xc0, 0xdc, 0x05,
	0x82, 0xbf, 0x04, 0xa6, 0x06, 0xef, 0x8a, 0xe0, 0xd0, 0x58, 0xe5, 0x2c, 0xe0, 0xef, 0x16, 0x43,
	0x26, 0xba, 0xa2, 0x34, 0xba, 0x56, 0x9d, 0xa1, 0xe3, 0x30, 0x39, 0x47, 0x77, 0x33, 0x12, 0xfc,
	0x35, 0x9c, 0x53, 0xa6, 0x37, 0x4d, 0xab, 0x2c, 0x1d, 0x83, 0x49, 0x08, 0xd4, 0x1d, 0x31, 0x7c,
	0x0d, 0x17, 0x24, 0x30, 0x58, 0x8e, 0xae, 0xd1, 0x05, 0x24, 0x65, 0x7d, 0xc1, 0x32, 0x58, 0x76,
	0xfd, 0xf1, 0xc7, 0x87, 0xca, 0xfa, 0xc3, 0xb0, 0xcb, 0xb4, 0x6b, 0xf2, 0x12, 0xdd, 0x6f, 0xfb,
	0xcb, 0xed, 0xf7, 0x79, 0xb4, 0xed, 0x4a, 0xb5, 0x36, 0xaf, 0x0c, 0xd2, 0x5f, 0xcb, 0x1f, 0x7d,
	0xd5, 0x5d, 0x4a, 0xd4, 0xbb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x42, 0x0c, 0x4f, 0xc0,
	0x02, 0x00, 0x00,
}
