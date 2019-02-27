// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/server.proto

package api

import (
	fmt "fmt"
	base "github.com/dnovikoff/mahjong-api/genproto/base"
	log "github.com/dnovikoff/mahjong-api/genproto/log"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Server initiated message
type Server struct {
	// If server message contains suggset, the client is expected
	// to answer with Client message, with same "suggest_id" set.
	// The exception is when suggest.canceled = true. With such message
	// server tells that on of previous messages (selected by suggest_id) is
	// out of date (ex. by timeout) and the response is no longer expected.
	Suggest *Suggest `protobuf:"bytes,1,opt,name=suggest,proto3" json:"suggest,omitempty"`
	// Types that are valid to be assigned to OneofEvents:
	//	*Server_Take
	//	*Server_Drop
	//	*Server_Changes
	//	*Server_Say
	//	*Server_Declare
	//	*Server_Win
	//	*Server_Draw
	//	*Server_Indicator
	//	*Server_RoundStart
	//	*Server_GameStart
	//	*Server_GameEnd
	//	*Server_Furiten
	//	*Server_Recover
	//	*Server_PlayerStatus
	//	*Server_Settings
	OneofEvents          isServer_OneofEvents `protobuf_oneof:"oneof_events"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b13ee64afa9929, []int{0}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetSuggest() *Suggest {
	if m != nil {
		return m.Suggest
	}
	return nil
}

type isServer_OneofEvents interface {
	isServer_OneofEvents()
}

type Server_Take struct {
	Take *log.TakeEvent `protobuf:"bytes,3,opt,name=take,proto3,oneof"`
}

type Server_Drop struct {
	Drop *log.DropEvent `protobuf:"bytes,4,opt,name=drop,proto3,oneof"`
}

type Server_Changes struct {
	Changes *log.ScoreChangesEvent `protobuf:"bytes,5,opt,name=changes,proto3,oneof"`
}

type Server_Say struct {
	Say *log.SayEvent `protobuf:"bytes,6,opt,name=say,proto3,oneof"`
}

type Server_Declare struct {
	Declare *log.DeclareEvent `protobuf:"bytes,7,opt,name=declare,proto3,oneof"`
}

type Server_Win struct {
	Win *log.WinEvent `protobuf:"bytes,8,opt,name=win,proto3,oneof"`
}

type Server_Draw struct {
	Draw *log.DrawEvent `protobuf:"bytes,9,opt,name=draw,proto3,oneof"`
}

type Server_Indicator struct {
	Indicator *log.IndicatorEvent `protobuf:"bytes,10,opt,name=indicator,proto3,oneof"`
}

type Server_RoundStart struct {
	RoundStart *log.RoundInfo `protobuf:"bytes,11,opt,name=round_start,json=roundStart,proto3,oneof"`
}

type Server_GameStart struct {
	GameStart *GameStartEvent `protobuf:"bytes,12,opt,name=game_start,json=gameStart,proto3,oneof"`
}

type Server_GameEnd struct {
	GameEnd *GameEndEvent `protobuf:"bytes,13,opt,name=game_end,json=gameEnd,proto3,oneof"`
}

type Server_Furiten struct {
	Furiten log.Furiten `protobuf:"varint,14,opt,name=furiten,proto3,enum=log.Furiten,oneof"`
}

type Server_Recover struct {
	Recover *RoundRecover `protobuf:"bytes,15,opt,name=recover,proto3,oneof"`
}

type Server_PlayerStatus struct {
	PlayerStatus *log.PlayerStatusEvent `protobuf:"bytes,16,opt,name=player_status,json=playerStatus,proto3,oneof"`
}

type Server_Settings struct {
	Settings *Settings `protobuf:"bytes,17,opt,name=settings,proto3,oneof"`
}

func (*Server_Take) isServer_OneofEvents() {}

func (*Server_Drop) isServer_OneofEvents() {}

func (*Server_Changes) isServer_OneofEvents() {}

func (*Server_Say) isServer_OneofEvents() {}

func (*Server_Declare) isServer_OneofEvents() {}

func (*Server_Win) isServer_OneofEvents() {}

func (*Server_Draw) isServer_OneofEvents() {}

func (*Server_Indicator) isServer_OneofEvents() {}

func (*Server_RoundStart) isServer_OneofEvents() {}

func (*Server_GameStart) isServer_OneofEvents() {}

func (*Server_GameEnd) isServer_OneofEvents() {}

func (*Server_Furiten) isServer_OneofEvents() {}

func (*Server_Recover) isServer_OneofEvents() {}

func (*Server_PlayerStatus) isServer_OneofEvents() {}

func (*Server_Settings) isServer_OneofEvents() {}

func (m *Server) GetOneofEvents() isServer_OneofEvents {
	if m != nil {
		return m.OneofEvents
	}
	return nil
}

func (m *Server) GetTake() *log.TakeEvent {
	if x, ok := m.GetOneofEvents().(*Server_Take); ok {
		return x.Take
	}
	return nil
}

func (m *Server) GetDrop() *log.DropEvent {
	if x, ok := m.GetOneofEvents().(*Server_Drop); ok {
		return x.Drop
	}
	return nil
}

func (m *Server) GetChanges() *log.ScoreChangesEvent {
	if x, ok := m.GetOneofEvents().(*Server_Changes); ok {
		return x.Changes
	}
	return nil
}

func (m *Server) GetSay() *log.SayEvent {
	if x, ok := m.GetOneofEvents().(*Server_Say); ok {
		return x.Say
	}
	return nil
}

func (m *Server) GetDeclare() *log.DeclareEvent {
	if x, ok := m.GetOneofEvents().(*Server_Declare); ok {
		return x.Declare
	}
	return nil
}

func (m *Server) GetWin() *log.WinEvent {
	if x, ok := m.GetOneofEvents().(*Server_Win); ok {
		return x.Win
	}
	return nil
}

func (m *Server) GetDraw() *log.DrawEvent {
	if x, ok := m.GetOneofEvents().(*Server_Draw); ok {
		return x.Draw
	}
	return nil
}

func (m *Server) GetIndicator() *log.IndicatorEvent {
	if x, ok := m.GetOneofEvents().(*Server_Indicator); ok {
		return x.Indicator
	}
	return nil
}

func (m *Server) GetRoundStart() *log.RoundInfo {
	if x, ok := m.GetOneofEvents().(*Server_RoundStart); ok {
		return x.RoundStart
	}
	return nil
}

func (m *Server) GetGameStart() *GameStartEvent {
	if x, ok := m.GetOneofEvents().(*Server_GameStart); ok {
		return x.GameStart
	}
	return nil
}

func (m *Server) GetGameEnd() *GameEndEvent {
	if x, ok := m.GetOneofEvents().(*Server_GameEnd); ok {
		return x.GameEnd
	}
	return nil
}

func (m *Server) GetFuriten() log.Furiten {
	if x, ok := m.GetOneofEvents().(*Server_Furiten); ok {
		return x.Furiten
	}
	return log.Furiten_FURITEN_UNDEFINED
}

func (m *Server) GetRecover() *RoundRecover {
	if x, ok := m.GetOneofEvents().(*Server_Recover); ok {
		return x.Recover
	}
	return nil
}

func (m *Server) GetPlayerStatus() *log.PlayerStatusEvent {
	if x, ok := m.GetOneofEvents().(*Server_PlayerStatus); ok {
		return x.PlayerStatus
	}
	return nil
}

func (m *Server) GetSettings() *Settings {
	if x, ok := m.GetOneofEvents().(*Server_Settings); ok {
		return x.Settings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Server) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Server_Take)(nil),
		(*Server_Drop)(nil),
		(*Server_Changes)(nil),
		(*Server_Say)(nil),
		(*Server_Declare)(nil),
		(*Server_Win)(nil),
		(*Server_Draw)(nil),
		(*Server_Indicator)(nil),
		(*Server_RoundStart)(nil),
		(*Server_GameStart)(nil),
		(*Server_GameEnd)(nil),
		(*Server_Furiten)(nil),
		(*Server_Recover)(nil),
		(*Server_PlayerStatus)(nil),
		(*Server_Settings)(nil),
	}
}

// Send after reconnect
type RoundRecover struct {
	Events               []*Server `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoundRecover) Reset()         { *m = RoundRecover{} }
func (m *RoundRecover) String() string { return proto.CompactTextString(m) }
func (*RoundRecover) ProtoMessage()    {}
func (*RoundRecover) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b13ee64afa9929, []int{1}
}

func (m *RoundRecover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundRecover.Unmarshal(m, b)
}
func (m *RoundRecover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundRecover.Marshal(b, m, deterministic)
}
func (m *RoundRecover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundRecover.Merge(m, src)
}
func (m *RoundRecover) XXX_Size() int {
	return xxx_messageInfo_RoundRecover.Size(m)
}
func (m *RoundRecover) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundRecover.DiscardUnknown(m)
}

var xxx_messageInfo_RoundRecover proto.InternalMessageInfo

func (m *RoundRecover) GetEvents() []*Server {
	if m != nil {
		return m.Events
	}
	return nil
}

// Final screen
type GameEndEvent struct {
	Changes              *log.Changes      `protobuf:"bytes,1,opt,name=changes,proto3" json:"changes,omitempty"`
	EndReason            log.GameEndReason `protobuf:"varint,2,opt,name=end_reason,json=endReason,proto3,enum=log.GameEndReason" json:"end_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GameEndEvent) Reset()         { *m = GameEndEvent{} }
func (m *GameEndEvent) String() string { return proto.CompactTextString(m) }
func (*GameEndEvent) ProtoMessage()    {}
func (*GameEndEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b13ee64afa9929, []int{2}
}

func (m *GameEndEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEndEvent.Unmarshal(m, b)
}
func (m *GameEndEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEndEvent.Marshal(b, m, deterministic)
}
func (m *GameEndEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEndEvent.Merge(m, src)
}
func (m *GameEndEvent) XXX_Size() int {
	return xxx_messageInfo_GameEndEvent.Size(m)
}
func (m *GameEndEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEndEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GameEndEvent proto.InternalMessageInfo

func (m *GameEndEvent) GetChanges() *log.Changes {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *GameEndEvent) GetEndReason() log.GameEndReason {
	if m != nil {
		return m.EndReason
	}
	return log.GameEndReason_GAME_END_REASON_UNSPECIFIED
}

type GameStartEvent struct {
	ClientIndex          int64         `protobuf:"varint,1,opt,name=client_index,json=clientIndex,proto3" json:"client_index,omitempty"`
	Info                 *log.GameInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameStartEvent) Reset()         { *m = GameStartEvent{} }
func (m *GameStartEvent) String() string { return proto.CompactTextString(m) }
func (*GameStartEvent) ProtoMessage()    {}
func (*GameStartEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b13ee64afa9929, []int{3}
}

func (m *GameStartEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartEvent.Unmarshal(m, b)
}
func (m *GameStartEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartEvent.Marshal(b, m, deterministic)
}
func (m *GameStartEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartEvent.Merge(m, src)
}
func (m *GameStartEvent) XXX_Size() int {
	return xxx_messageInfo_GameStartEvent.Size(m)
}
func (m *GameStartEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartEvent.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartEvent proto.InternalMessageInfo

func (m *GameStartEvent) GetClientIndex() int64 {
	if m != nil {
		return m.ClientIndex
	}
	return 0
}

func (m *GameStartEvent) GetInfo() *log.GameInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type Suggest struct {
	// Should response with this id
	SuggestId        int64           `protobuf:"varint,1,opt,name=suggest_id,json=suggestId,proto3" json:"suggest_id,omitempty"`
	Pon              bool            `protobuf:"varint,2,opt,name=pon,proto3" json:"pon,omitempty"`
	Kan              bool            `protobuf:"varint,3,opt,name=kan,proto3" json:"kan,omitempty"`
	ClosedKanMask    int64           `protobuf:"varint,4,opt,name=closed_kan_mask,json=closedKanMask,proto3" json:"closed_kan_mask,omitempty"`
	UpgradeInstances *base.Instances `protobuf:"bytes,5,opt,name=upgrade_instances,json=upgradeInstances,proto3" json:"upgrade_instances,omitempty"`
	// 1+23
	ChiLeft bool `protobuf:"varint,6,opt,name=chi_left,json=chiLeft,proto3" json:"chi_left,omitempty"`
	// 2+13
	ChiCenter bool `protobuf:"varint,7,opt,name=chi_center,json=chiCenter,proto3" json:"chi_center,omitempty"`
	// 3+12
	ChiRight   bool  `protobuf:"varint,8,opt,name=chi_right,json=chiRight,proto3" json:"chi_right,omitempty"`
	DropMask   int64 `protobuf:"varint,9,opt,name=drop_mask,json=dropMask,proto3" json:"drop_mask,omitempty"`
	RiichiMask int64 `protobuf:"varint,10,opt,name=riichi_mask,json=riichiMask,proto3" json:"riichi_mask,omitempty"`
	Win        bool  `protobuf:"varint,11,opt,name=win,proto3" json:"win,omitempty"`
	Draw       bool  `protobuf:"varint,12,opt,name=draw,proto3" json:"draw,omitempty"`
	// Suggested to show noten (could show tempai)
	// Could be used for some rulesets
	Noten bool `protobuf:"varint,13,opt,name=noten,proto3" json:"noten,omitempty"`
	// Specific for some rules, where leader could continue game (agari yame).
	ContinueGame bool `protobuf:"varint,14,opt,name=continue_game,json=continueGame,proto3" json:"continue_game,omitempty"`
	// Time to answer. Should be used to display timer to player.
	// Client should not take any action if player does not select an answer in given time.
	// After the timeout server will discard the event.
	// Server could add additional time on its side to gurantee player will have at least
	// timeout time.
	Timeout *duration.Duration `protobuf:"bytes,15,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Means all suggest are canceled. Must not be answered.
	Canceled             bool     `protobuf:"varint,16,opt,name=canceled,proto3" json:"canceled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suggest) Reset()         { *m = Suggest{} }
func (m *Suggest) String() string { return proto.CompactTextString(m) }
func (*Suggest) ProtoMessage()    {}
func (*Suggest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19b13ee64afa9929, []int{4}
}

func (m *Suggest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suggest.Unmarshal(m, b)
}
func (m *Suggest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suggest.Marshal(b, m, deterministic)
}
func (m *Suggest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suggest.Merge(m, src)
}
func (m *Suggest) XXX_Size() int {
	return xxx_messageInfo_Suggest.Size(m)
}
func (m *Suggest) XXX_DiscardUnknown() {
	xxx_messageInfo_Suggest.DiscardUnknown(m)
}

var xxx_messageInfo_Suggest proto.InternalMessageInfo

func (m *Suggest) GetSuggestId() int64 {
	if m != nil {
		return m.SuggestId
	}
	return 0
}

func (m *Suggest) GetPon() bool {
	if m != nil {
		return m.Pon
	}
	return false
}

func (m *Suggest) GetKan() bool {
	if m != nil {
		return m.Kan
	}
	return false
}

func (m *Suggest) GetClosedKanMask() int64 {
	if m != nil {
		return m.ClosedKanMask
	}
	return 0
}

func (m *Suggest) GetUpgradeInstances() *base.Instances {
	if m != nil {
		return m.UpgradeInstances
	}
	return nil
}

func (m *Suggest) GetChiLeft() bool {
	if m != nil {
		return m.ChiLeft
	}
	return false
}

func (m *Suggest) GetChiCenter() bool {
	if m != nil {
		return m.ChiCenter
	}
	return false
}

func (m *Suggest) GetChiRight() bool {
	if m != nil {
		return m.ChiRight
	}
	return false
}

func (m *Suggest) GetDropMask() int64 {
	if m != nil {
		return m.DropMask
	}
	return 0
}

func (m *Suggest) GetRiichiMask() int64 {
	if m != nil {
		return m.RiichiMask
	}
	return 0
}

func (m *Suggest) GetWin() bool {
	if m != nil {
		return m.Win
	}
	return false
}

func (m *Suggest) GetDraw() bool {
	if m != nil {
		return m.Draw
	}
	return false
}

func (m *Suggest) GetNoten() bool {
	if m != nil {
		return m.Noten
	}
	return false
}

func (m *Suggest) GetContinueGame() bool {
	if m != nil {
		return m.ContinueGame
	}
	return false
}

func (m *Suggest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *Suggest) GetCanceled() bool {
	if m != nil {
		return m.Canceled
	}
	return false
}

func init() {
	proto.RegisterType((*Server)(nil), "api.Server")
	proto.RegisterType((*RoundRecover)(nil), "api.RoundRecover")
	proto.RegisterType((*GameEndEvent)(nil), "api.GameEndEvent")
	proto.RegisterType((*GameStartEvent)(nil), "api.GameStartEvent")
	proto.RegisterType((*Suggest)(nil), "api.Suggest")
}

func init() { proto.RegisterFile("api/server.proto", fileDescriptor_19b13ee64afa9929) }

var fileDescriptor_19b13ee64afa9929 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x55, 0xdf, 0x6f, 0xeb, 0x34,
	0x14, 0xde, 0xe8, 0xb6, 0xa6, 0xa7, 0x69, 0xd7, 0x19, 0x84, 0x72, 0x87, 0x80, 0xad, 0x17, 0x5d,
	0x4d, 0x42, 0x4b, 0x75, 0x37, 0x1e, 0xe1, 0xe5, 0xfe, 0x80, 0x55, 0x80, 0x84, 0x5c, 0x04, 0x12,
	0x2f, 0x91, 0x97, 0xb8, 0xa9, 0x69, 0x6b, 0x47, 0x8e, 0xb3, 0xb2, 0x7f, 0x9a, 0x77, 0xde, 0xd0,
	0x39, 0x76, 0x72, 0x73, 0x5f, 0xa2, 0xf8, 0xfb, 0x3e, 0x9f, 0x73, 0x7c, 0xec, 0xcf, 0x86, 0x99,
	0xa8, 0xd4, 0xa2, 0x96, 0xf6, 0x49, 0xda, 0xb4, 0xb2, 0xc6, 0x19, 0x36, 0x10, 0x95, 0xba, 0xfc,
	0xaa, 0x34, 0xa6, 0xdc, 0xc9, 0x05, 0x41, 0x8f, 0xcd, 0x7a, 0x51, 0x34, 0x56, 0x38, 0x65, 0xb4,
	0x17, 0x5d, 0xce, 0x76, 0xa6, 0x5c, 0xc8, 0x27, 0xa9, 0x5d, 0xdd, 0x47, 0x72, 0xb3, 0xdf, 0x77,
	0x9a, 0x09, 0x22, 0x3b, 0x53, 0x86, 0xe1, 0xf9, 0xa3, 0xa8, 0xe5, 0x02, 0x3f, 0xed, 0x0c, 0x4c,
	0x9d, 0xef, 0x94, 0xd4, 0xce, 0x23, 0xf3, 0x7f, 0x4f, 0xe1, 0x6c, 0x45, 0xb5, 0xb0, 0x57, 0x30,
	0xac, 0x9b, 0xb2, 0x94, 0xb5, 0x4b, 0x8e, 0xaf, 0x8e, 0x6f, 0xc6, 0x77, 0x71, 0x2a, 0x2a, 0x95,
	0xae, 0x3c, 0xc6, 0x5b, 0x92, 0x7d, 0x03, 0x27, 0x4e, 0x6c, 0x65, 0x32, 0x20, 0xd1, 0x34, 0xc5,
	0x7c, 0xbf, 0x8b, 0xad, 0x7c, 0x8f, 0xb5, 0x3d, 0x1c, 0x71, 0x62, 0x51, 0x55, 0x58, 0x53, 0x25,
	0x27, 0x3d, 0xd5, 0x3b, 0x6b, 0xaa, 0x4e, 0x85, 0x2c, 0xbb, 0x83, 0x61, 0xbe, 0x11, 0xba, 0x94,
	0x75, 0x72, 0x4a, 0xc2, 0xcf, 0x49, 0xb8, 0xca, 0x8d, 0x95, 0x6f, 0x3d, 0xd1, 0x4e, 0x68, 0x85,
	0xec, 0x1a, 0x06, 0xb5, 0x78, 0x4e, 0xce, 0x48, 0x3f, 0xf1, 0x7a, 0xf1, 0xdc, 0xca, 0x90, 0x63,
	0xb7, 0x30, 0x2c, 0x64, 0xbe, 0x13, 0x56, 0x26, 0x43, 0x92, 0x5d, 0xf8, 0xfc, 0x1e, 0xeb, 0x22,
	0x06, 0x0d, 0x46, 0x3c, 0x28, 0x9d, 0x44, 0xbd, 0x88, 0x7f, 0x2a, 0xdd, 0x45, 0x3c, 0x28, 0xed,
	0x97, 0x23, 0x0e, 0xc9, 0xe8, 0xa3, 0xe5, 0x88, 0x43, 0x6f, 0x39, 0xe2, 0xc0, 0xee, 0x61, 0xa4,
	0x74, 0xa1, 0x72, 0xe1, 0x8c, 0x4d, 0x80, 0xa4, 0x9f, 0x92, 0x74, 0xd9, 0xa2, 0xad, 0xfe, 0x83,
	0x8e, 0xbd, 0x86, 0xb1, 0x35, 0x8d, 0x2e, 0xb2, 0xda, 0x09, 0xeb, 0x92, 0x71, 0x2f, 0x03, 0x47,
	0x7c, 0xa9, 0xd7, 0xe6, 0xe1, 0x88, 0x03, 0x89, 0x56, 0xa8, 0x61, 0xdf, 0x01, 0x94, 0x62, 0x2f,
	0xc3, 0x8c, 0x38, 0x24, 0xc2, 0xdd, 0xfa, 0x49, 0xec, 0x25, 0x69, 0xba, 0x44, 0x65, 0x8b, 0xb0,
	0x14, 0x22, 0x9a, 0x25, 0x75, 0x91, 0x4c, 0x42, 0x5b, 0xda, 0x39, 0xef, 0x75, 0xd1, 0xb5, 0xa5,
	0xf4, 0x63, 0x76, 0x03, 0xc3, 0x75, 0x63, 0x95, 0x93, 0x3a, 0x99, 0x5e, 0x1d, 0xdf, 0x4c, 0xef,
	0x62, 0x2a, 0xea, 0x47, 0x8f, 0xa1, 0x32, 0xd0, 0xd8, 0x6f, 0x2b, 0x73, 0xf3, 0x24, 0x6d, 0x72,
	0xde, 0x0b, 0x4c, 0xe5, 0x73, 0x4f, 0xa0, 0x3c, 0x68, 0xd8, 0x0f, 0x30, 0xa9, 0x76, 0xe2, 0x59,
	0x5a, 0x5c, 0x80, 0x6b, 0xea, 0x64, 0xd6, 0xdb, 0xfb, 0xdf, 0x88, 0x59, 0x11, 0xd1, 0x96, 0x14,
	0x57, 0x3d, 0x90, 0x7d, 0x0b, 0x51, 0x2d, 0x9d, 0x53, 0xba, 0xac, 0x93, 0x8b, 0xb0, 0x67, 0x74,
	0x52, 0x03, 0xf8, 0x70, 0xc4, 0x3b, 0xc1, 0x9b, 0x29, 0xc4, 0x46, 0x4b, 0xb3, 0xce, 0xbc, 0x75,
	0xe6, 0xf7, 0x10, 0xf7, 0xcb, 0x62, 0x2f, 0xe1, 0xcc, 0x33, 0xc9, 0xf1, 0xd5, 0xe0, 0x66, 0x7c,
	0x37, 0x0e, 0xa1, 0xd0, 0x12, 0x3c, 0x50, 0x73, 0x05, 0x71, 0xbf, 0x49, 0x68, 0x95, 0xf6, 0xd8,
	0xb6, 0x56, 0xc1, 0xd2, 0xc3, 0x89, 0xfd, 0x70, 0x54, 0x5f, 0x03, 0x48, 0x5d, 0x64, 0x56, 0x8a,
	0xda, 0xe8, 0xe4, 0x13, 0x6a, 0x22, 0x23, 0x69, 0x08, 0xc7, 0x89, 0xe1, 0x23, 0xd9, 0xfe, 0xce,
	0xff, 0x80, 0xe9, 0xc7, 0x7b, 0xc8, 0xae, 0x21, 0xf6, 0x96, 0xcd, 0x94, 0x2e, 0xe4, 0x3f, 0x94,
	0x71, 0xc0, 0xc7, 0x1e, 0x5b, 0x22, 0xc4, 0xae, 0xe1, 0x44, 0xe9, 0xb5, 0xa1, 0x0c, 0xed, 0x09,
	0xc6, 0x28, 0x78, 0x74, 0x38, 0x51, 0xf3, 0xff, 0x06, 0x30, 0x0c, 0x56, 0x66, 0x5f, 0x02, 0x04,
	0x33, 0x67, 0xaa, 0x08, 0xf1, 0x46, 0x01, 0x59, 0x16, 0x6c, 0x06, 0x83, 0x2a, 0x94, 0x1b, 0x71,
	0xfc, 0x45, 0x64, 0x2b, 0x34, 0x39, 0x3e, 0xe2, 0xf8, 0xcb, 0x5e, 0xc1, 0x79, 0xbe, 0x33, 0xb5,
	0x2c, 0xb2, 0xad, 0xd0, 0xd9, 0x5e, 0xd4, 0x5b, 0x72, 0xfa, 0x80, 0x4f, 0x3c, 0xfc, 0xb3, 0xd0,
	0xbf, 0x8a, 0x7a, 0xcb, 0xbe, 0x87, 0x8b, 0xa6, 0x2a, 0xad, 0x28, 0x64, 0xa6, 0x74, 0xed, 0x84,
	0xce, 0x3b, 0xab, 0x9f, 0xa7, 0x74, 0x33, 0x2d, 0x5b, 0x98, 0xcf, 0x82, 0xb2, 0x43, 0xd8, 0x0b,
	0x88, 0xf2, 0x8d, 0xca, 0x76, 0x72, 0xed, 0xc8, 0xef, 0x11, 0xb6, 0x56, 0xfd, 0x22, 0xd7, 0xb4,
	0x06, 0xa4, 0x72, 0xa9, 0x9d, 0xb4, 0xe4, 0xf2, 0x88, 0x8f, 0xf2, 0x8d, 0x7a, 0x4b, 0x00, 0xfb,
	0x02, 0x70, 0x90, 0x59, 0x55, 0x6e, 0x1c, 0x19, 0x3b, 0xe2, 0x18, 0x8a, 0xe3, 0x18, 0x49, 0xbc,
	0x7d, 0x7c, 0xd9, 0x23, 0x2a, 0x3b, 0x42, 0x80, 0x2a, 0xfe, 0x1a, 0xc6, 0x56, 0x29, 0x9c, 0x4c,
	0x34, 0x10, 0x0d, 0x1e, 0x22, 0xc1, 0xcc, 0xdf, 0x16, 0x63, 0xdf, 0x0c, 0xbc, 0x1c, 0x58, 0xb8,
	0x1c, 0x62, 0x82, 0xfc, 0x55, 0xf0, 0x19, 0x9c, 0x6a, 0x83, 0xd6, 0x99, 0x10, 0xe8, 0x07, 0xec,
	0x25, 0x4c, 0x72, 0xa3, 0x9d, 0xd2, 0x8d, 0xcc, 0xd0, 0x66, 0x64, 0xac, 0x88, 0xc7, 0x2d, 0x88,
	0x9b, 0xc6, 0xee, 0x61, 0xe8, 0xd4, 0x5e, 0x9a, 0xc6, 0x05, 0x37, 0xbd, 0x48, 0xfd, 0xdb, 0x90,
	0xb6, 0x6f, 0x43, 0xfa, 0x2e, 0xbc, 0x0d, 0xbc, 0x55, 0xb2, 0x4b, 0x88, 0x72, 0x6c, 0xda, 0x4e,
	0x16, 0x64, 0x27, 0x5c, 0x6f, 0x18, 0xbf, 0x59, 0xfc, 0x75, 0x5b, 0x2a, 0xb7, 0x69, 0x1e, 0xd3,
	0xdc, 0xec, 0x17, 0x85, 0x36, 0x4f, 0x6a, 0x6b, 0xd6, 0xeb, 0xc5, 0x5e, 0x6c, 0xfe, 0x36, 0xba,
	0xbc, 0xc5, 0x57, 0xa1, 0x94, 0x9a, 0xa2, 0x2f, 0x44, 0xa5, 0x1e, 0xcf, 0xe8, 0xf7, 0xfe, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xbe, 0x86, 0x19, 0xab, 0x06, 0x00, 0x00,
}