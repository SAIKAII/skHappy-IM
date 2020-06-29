// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tcp-conn.proto

package pb

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

// 请求类型
type PackageType int32

const (
	PackageType_PT_UNKNOWN      PackageType = 0
	PackageType_PT_SIGN_IN      PackageType = 1
	PackageType_PT_SYNC_MESSAGE PackageType = 2
	PackageType_PT_HEART_BEAT   PackageType = 3
	PackageType_PT_MESSAGE      PackageType = 4
)

var PackageType_name = map[int32]string{
	0: "PT_UNKNOWN",
	1: "PT_SIGN_IN",
	2: "PT_SYNC_MESSAGE",
	3: "PT_HEART_BEAT",
	4: "PT_MESSAGE",
}

var PackageType_value = map[string]int32{
	"PT_UNKNOWN":      0,
	"PT_SIGN_IN":      1,
	"PT_SYNC_MESSAGE": 2,
	"PT_HEART_BEAT":   3,
	"PT_MESSAGE":      4,
}

func (x PackageType) String() string {
	return proto.EnumName(PackageType_name, int32(x))
}

func (PackageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{0}
}

// 消息类型
type MessageType int32

const (
	MessageType_MT_UNKNOWN MessageType = 0
	MessageType_MT_TEXT    MessageType = 1
	MessageType_MT_IMAGE   MessageType = 2
)

var MessageType_name = map[int32]string{
	0: "MT_UNKNOWN",
	1: "MT_TEXT",
	2: "MT_IMAGE",
}

var MessageType_value = map[string]int32{
	"MT_UNKNOWN": 0,
	"MT_TEXT":    1,
	"MT_IMAGE":   2,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{1}
}

type SenderType int32

const (
	SenderType_ST_UNKNOWN  SenderType = 0
	SenderType_ST_USER     SenderType = 1
	SenderType_ST_SYSTEM   SenderType = 2
	SenderType_ST_BUSINESS SenderType = 3
)

var SenderType_name = map[int32]string{
	0: "ST_UNKNOWN",
	1: "ST_USER",
	2: "ST_SYSTEM",
	3: "ST_BUSINESS",
}

var SenderType_value = map[string]int32{
	"ST_UNKNOWN":  0,
	"ST_USER":     1,
	"ST_SYSTEM":   2,
	"ST_BUSINESS": 3,
}

func (x SenderType) String() string {
	return proto.EnumName(SenderType_name, int32(x))
}

func (SenderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{2}
}

type ReceiverType int32

const (
	ReceiverType_RT_UNKNOWN ReceiverType = 0
	ReceiverType_RT_USER    ReceiverType = 1
	ReceiverType_RT_GROUP   ReceiverType = 2
)

var ReceiverType_name = map[int32]string{
	0: "RT_UNKNOWN",
	1: "RT_USER",
	2: "RT_GROUP",
}

var ReceiverType_value = map[string]int32{
	"RT_UNKNOWN": 0,
	"RT_USER":    1,
	"RT_GROUP":   2,
}

func (x ReceiverType) String() string {
	return proto.EnumName(ReceiverType_name, int32(x))
}

func (ReceiverType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{3}
}

type ConnInput struct {
	PackageType          PackageType `protobuf:"varint,1,opt,name=package_type,json=packageType,proto3,enum=pb.PackageType" json:"package_type,omitempty"`
	Data                 []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConnInput) Reset()         { *m = ConnInput{} }
func (m *ConnInput) String() string { return proto.CompactTextString(m) }
func (*ConnInput) ProtoMessage()    {}
func (*ConnInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{0}
}

func (m *ConnInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnInput.Unmarshal(m, b)
}
func (m *ConnInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnInput.Marshal(b, m, deterministic)
}
func (m *ConnInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnInput.Merge(m, src)
}
func (m *ConnInput) XXX_Size() int {
	return xxx_messageInfo_ConnInput.Size(m)
}
func (m *ConnInput) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnInput.DiscardUnknown(m)
}

var xxx_messageInfo_ConnInput proto.InternalMessageInfo

func (m *ConnInput) GetPackageType() PackageType {
	if m != nil {
		return m.PackageType
	}
	return PackageType_PT_UNKNOWN
}

func (m *ConnInput) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConnOutput struct {
	PackageType          PackageType `protobuf:"varint,1,opt,name=package_type,json=packageType,proto3,enum=pb.PackageType" json:"package_type,omitempty"`
	ErrCode              int32       `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	ErrMsg               string      `protobuf:"bytes,3,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	Data                 []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConnOutput) Reset()         { *m = ConnOutput{} }
func (m *ConnOutput) String() string { return proto.CompactTextString(m) }
func (*ConnOutput) ProtoMessage()    {}
func (*ConnOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{1}
}

func (m *ConnOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnOutput.Unmarshal(m, b)
}
func (m *ConnOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnOutput.Marshal(b, m, deterministic)
}
func (m *ConnOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnOutput.Merge(m, src)
}
func (m *ConnOutput) XXX_Size() int {
	return xxx_messageInfo_ConnOutput.Size(m)
}
func (m *ConnOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ConnOutput proto.InternalMessageInfo

func (m *ConnOutput) GetPackageType() PackageType {
	if m != nil {
		return m.PackageType
	}
	return PackageType_PT_UNKNOWN
}

func (m *ConnOutput) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ConnOutput) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *ConnOutput) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

//************************************************************************
type SignInReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}
func (*SignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{2}
}

func (m *SignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReq.Unmarshal(m, b)
}
func (m *SignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReq.Marshal(b, m, deterministic)
}
func (m *SignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReq.Merge(m, src)
}
func (m *SignInReq) XXX_Size() int {
	return xxx_messageInfo_SignInReq.Size(m)
}
func (m *SignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReq proto.InternalMessageInfo

func (m *SignInReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInResp struct {
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResp) Reset()         { *m = SignInResp{} }
func (m *SignInResp) String() string { return proto.CompactTextString(m) }
func (*SignInResp) ProtoMessage()    {}
func (*SignInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{3}
}

func (m *SignInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResp.Unmarshal(m, b)
}
func (m *SignInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResp.Marshal(b, m, deterministic)
}
func (m *SignInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResp.Merge(m, src)
}
func (m *SignInResp) XXX_Size() int {
	return xxx_messageInfo_SignInResp.Size(m)
}
func (m *SignInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResp proto.InternalMessageInfo

func (m *SignInResp) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

type SyncReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	LastSeqId            uint64   `protobuf:"varint,2,opt,name=last_seq_id,json=lastSeqId,proto3" json:"last_seq_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncReq) Reset()         { *m = SyncReq{} }
func (m *SyncReq) String() string { return proto.CompactTextString(m) }
func (*SyncReq) ProtoMessage()    {}
func (*SyncReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{4}
}

func (m *SyncReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncReq.Unmarshal(m, b)
}
func (m *SyncReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncReq.Marshal(b, m, deterministic)
}
func (m *SyncReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncReq.Merge(m, src)
}
func (m *SyncReq) XXX_Size() int {
	return xxx_messageInfo_SyncReq.Size(m)
}
func (m *SyncReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncReq.DiscardUnknown(m)
}

var xxx_messageInfo_SyncReq proto.InternalMessageInfo

func (m *SyncReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SyncReq) GetLastSeqId() uint64 {
	if m != nil {
		return m.LastSeqId
	}
	return 0
}

type SyncResp struct {
	Msg                  []*MessageItem `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
	HasMore              bool           `protobuf:"varint,2,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SyncResp) Reset()         { *m = SyncResp{} }
func (m *SyncResp) String() string { return proto.CompactTextString(m) }
func (*SyncResp) ProtoMessage()    {}
func (*SyncResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{5}
}

func (m *SyncResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResp.Unmarshal(m, b)
}
func (m *SyncResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResp.Marshal(b, m, deterministic)
}
func (m *SyncResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResp.Merge(m, src)
}
func (m *SyncResp) XXX_Size() int {
	return xxx_messageInfo_SyncResp.Size(m)
}
func (m *SyncResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResp.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResp proto.InternalMessageInfo

func (m *SyncResp) GetMsg() []*MessageItem {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *SyncResp) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

type MessageOutput struct {
	Item                 *MessageItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MessageOutput) Reset()         { *m = MessageOutput{} }
func (m *MessageOutput) String() string { return proto.CompactTextString(m) }
func (*MessageOutput) ProtoMessage()    {}
func (*MessageOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{6}
}

func (m *MessageOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageOutput.Unmarshal(m, b)
}
func (m *MessageOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageOutput.Marshal(b, m, deterministic)
}
func (m *MessageOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageOutput.Merge(m, src)
}
func (m *MessageOutput) XXX_Size() int {
	return xxx_messageInfo_MessageOutput.Size(m)
}
func (m *MessageOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageOutput.DiscardUnknown(m)
}

var xxx_messageInfo_MessageOutput proto.InternalMessageInfo

func (m *MessageOutput) GetItem() *MessageItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type MessageBody struct {
	Type                 MessageType     `protobuf:"varint,1,opt,name=type,proto3,enum=pb.MessageType" json:"type,omitempty"`
	Content              *MessageContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageBody) Reset()         { *m = MessageBody{} }
func (m *MessageBody) String() string { return proto.CompactTextString(m) }
func (*MessageBody) ProtoMessage()    {}
func (*MessageBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{7}
}

func (m *MessageBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBody.Unmarshal(m, b)
}
func (m *MessageBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBody.Marshal(b, m, deterministic)
}
func (m *MessageBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBody.Merge(m, src)
}
func (m *MessageBody) XXX_Size() int {
	return xxx_messageInfo_MessageBody.Size(m)
}
func (m *MessageBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBody.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBody proto.InternalMessageInfo

func (m *MessageBody) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_MT_UNKNOWN
}

func (m *MessageBody) GetContent() *MessageContent {
	if m != nil {
		return m.Content
	}
	return nil
}

type MessageContent struct {
	// Types that are valid to be assigned to Content:
	//	*MessageContent_Text
	//	*MessageContent_Image
	Content              isMessageContent_Content `protobuf_oneof:"Content"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MessageContent) Reset()         { *m = MessageContent{} }
func (m *MessageContent) String() string { return proto.CompactTextString(m) }
func (*MessageContent) ProtoMessage()    {}
func (*MessageContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{8}
}

func (m *MessageContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageContent.Unmarshal(m, b)
}
func (m *MessageContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageContent.Marshal(b, m, deterministic)
}
func (m *MessageContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageContent.Merge(m, src)
}
func (m *MessageContent) XXX_Size() int {
	return xxx_messageInfo_MessageContent.Size(m)
}
func (m *MessageContent) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageContent.DiscardUnknown(m)
}

var xxx_messageInfo_MessageContent proto.InternalMessageInfo

type isMessageContent_Content interface {
	isMessageContent_Content()
}

type MessageContent_Text struct {
	Text *Text `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type MessageContent_Image struct {
	Image *Image `protobuf:"bytes,2,opt,name=image,proto3,oneof"`
}

func (*MessageContent_Text) isMessageContent_Content() {}

func (*MessageContent_Image) isMessageContent_Content() {}

func (m *MessageContent) GetContent() isMessageContent_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MessageContent) GetText() *Text {
	if x, ok := m.GetContent().(*MessageContent_Text); ok {
		return x.Text
	}
	return nil
}

func (m *MessageContent) GetImage() *Image {
	if x, ok := m.GetContent().(*MessageContent_Image); ok {
		return x.Image
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MessageContent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MessageContent_Text)(nil),
		(*MessageContent_Image)(nil),
	}
}

type MessageItem struct {
	SenderName           string       `protobuf:"bytes,1,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
	SenderType           SenderType   `protobuf:"varint,2,opt,name=sender_type,json=senderType,proto3,enum=pb.SenderType" json:"sender_type,omitempty"`
	ReceiverName         string       `protobuf:"bytes,3,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	ReceiverType         ReceiverType `protobuf:"varint,4,opt,name=receiver_type,json=receiverType,proto3,enum=pb.ReceiverType" json:"receiver_type,omitempty"`
	MsgBody              *MessageBody `protobuf:"bytes,5,opt,name=msg_body,json=msgBody,proto3" json:"msg_body,omitempty"`
	SendTime             int64        `protobuf:"varint,6,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MessageItem) Reset()         { *m = MessageItem{} }
func (m *MessageItem) String() string { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()    {}
func (*MessageItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{9}
}

func (m *MessageItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageItem.Unmarshal(m, b)
}
func (m *MessageItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageItem.Marshal(b, m, deterministic)
}
func (m *MessageItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageItem.Merge(m, src)
}
func (m *MessageItem) XXX_Size() int {
	return xxx_messageInfo_MessageItem.Size(m)
}
func (m *MessageItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageItem.DiscardUnknown(m)
}

var xxx_messageInfo_MessageItem proto.InternalMessageInfo

func (m *MessageItem) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *MessageItem) GetSenderType() SenderType {
	if m != nil {
		return m.SenderType
	}
	return SenderType_ST_UNKNOWN
}

func (m *MessageItem) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *MessageItem) GetReceiverType() ReceiverType {
	if m != nil {
		return m.ReceiverType
	}
	return ReceiverType_RT_UNKNOWN
}

func (m *MessageItem) GetMsgBody() *MessageBody {
	if m != nil {
		return m.MsgBody
	}
	return nil
}

func (m *MessageItem) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{10}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Image struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ThumbnailUrl         string   `protobuf:"bytes,5,opt,name=thumbnail_url,json=thumbnailUrl,proto3" json:"thumbnail_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_e16fdb746331c880, []int{11}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetThumbnailUrl() string {
	if m != nil {
		return m.ThumbnailUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("pb.PackageType", PackageType_name, PackageType_value)
	proto.RegisterEnum("pb.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("pb.SenderType", SenderType_name, SenderType_value)
	proto.RegisterEnum("pb.ReceiverType", ReceiverType_name, ReceiverType_value)
	proto.RegisterType((*ConnInput)(nil), "pb.ConnInput")
	proto.RegisterType((*ConnOutput)(nil), "pb.ConnOutput")
	proto.RegisterType((*SignInReq)(nil), "pb.SignInReq")
	proto.RegisterType((*SignInResp)(nil), "pb.SignInResp")
	proto.RegisterType((*SyncReq)(nil), "pb.SyncReq")
	proto.RegisterType((*SyncResp)(nil), "pb.SyncResp")
	proto.RegisterType((*MessageOutput)(nil), "pb.MessageOutput")
	proto.RegisterType((*MessageBody)(nil), "pb.MessageBody")
	proto.RegisterType((*MessageContent)(nil), "pb.MessageContent")
	proto.RegisterType((*MessageItem)(nil), "pb.MessageItem")
	proto.RegisterType((*Text)(nil), "pb.Text")
	proto.RegisterType((*Image)(nil), "pb.Image")
}

func init() {
	proto.RegisterFile("tcp-conn.proto", fileDescriptor_e16fdb746331c880)
}

var fileDescriptor_e16fdb746331c880 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdf, 0x6f, 0xeb, 0x34,
	0x14, 0xc7, 0x97, 0x34, 0x5d, 0x93, 0x93, 0xae, 0x0b, 0x06, 0x41, 0xb9, 0x48, 0xa3, 0x37, 0xf7,
	0xa5, 0xaa, 0x60, 0x48, 0x03, 0x24, 0x78, 0x5c, 0xab, 0x68, 0x8d, 0x50, 0xb2, 0xca, 0x4e, 0x05,
	0xbc, 0x60, 0xd2, 0xc6, 0x4a, 0x03, 0xcd, 0x8f, 0xc5, 0x2e, 0xbb, 0x95, 0x78, 0xe4, 0x81, 0x3f,
	0x1b, 0xd9, 0x49, 0x73, 0xcb, 0x40, 0x42, 0xe2, 0xcd, 0xe7, 0x7c, 0xcf, 0xf9, 0xf8, 0xfc, 0x88,
	0x03, 0x23, 0xb1, 0xad, 0x3e, 0xdf, 0x96, 0x45, 0x71, 0x5b, 0xd5, 0xa5, 0x28, 0x91, 0x5e, 0x6d,
	0x5c, 0x02, 0xd6, 0xa2, 0x2c, 0x0a, 0xbf, 0xa8, 0x0e, 0x02, 0xdd, 0xc1, 0xb0, 0x8a, 0xb7, 0xbf,
	0xc6, 0x29, 0xa3, 0xe2, 0x58, 0xb1, 0xb1, 0x36, 0xd1, 0xa6, 0xa3, 0xbb, 0xeb, 0xdb, 0x6a, 0x73,
	0xbb, 0x6a, 0xfc, 0xd1, 0xb1, 0x62, 0xd8, 0xae, 0xde, 0x19, 0x08, 0x81, 0x91, 0xc4, 0x22, 0x1e,
	0xeb, 0x13, 0x6d, 0x3a, 0xc4, 0xea, 0xec, 0xfe, 0xa9, 0x01, 0x48, 0xea, 0xe3, 0x41, 0xfc, 0x5f,
	0xec, 0xc7, 0x60, 0xb2, 0xba, 0xa6, 0xdb, 0x32, 0x61, 0x0a, 0xdd, 0xc7, 0x03, 0x56, 0xd7, 0x8b,
	0x32, 0x61, 0xe8, 0x23, 0x90, 0x47, 0x9a, 0xf3, 0x74, 0xdc, 0x9b, 0x68, 0x53, 0x0b, 0x5f, 0xb2,
	0xba, 0x0e, 0x78, 0xda, 0x95, 0x62, 0x9c, 0x95, 0xb2, 0x00, 0x8b, 0x64, 0x69, 0xe1, 0x17, 0x98,
	0x3d, 0xa1, 0x57, 0x60, 0x1e, 0x38, 0xab, 0x8b, 0x38, 0x6f, 0x8a, 0xb0, 0x70, 0x67, 0x4b, 0xad,
	0x8a, 0x39, 0x7f, 0x2e, 0xeb, 0x44, 0x5d, 0x68, 0xe1, 0xce, 0x76, 0x6f, 0x00, 0x4e, 0x10, 0x5e,
	0x21, 0x07, 0x7a, 0xbf, 0x3c, 0x8b, 0x16, 0x20, 0x8f, 0xae, 0x07, 0x03, 0x72, 0x2c, 0xb6, 0xff,
	0x75, 0xc5, 0x0d, 0xd8, 0xfb, 0x98, 0x0b, 0xca, 0xd9, 0x13, 0xcd, 0x9a, 0x5b, 0x0c, 0x6c, 0x49,
	0x17, 0x61, 0x4f, 0x7e, 0xe2, 0x2e, 0xc1, 0x6c, 0x30, 0xbc, 0x42, 0xaf, 0xa1, 0x27, 0x1b, 0xd4,
	0x26, 0xbd, 0xa9, 0xdd, 0x8c, 0x2a, 0x60, 0x9c, 0xc7, 0x29, 0xf3, 0x05, 0xcb, 0xb1, 0xd4, 0xe4,
	0x88, 0x76, 0x31, 0xa7, 0x79, 0x59, 0x37, 0x23, 0x32, 0xf1, 0x60, 0x17, 0xf3, 0xa0, 0xac, 0x99,
	0xfb, 0x15, 0x5c, 0xb5, 0xe1, 0xed, 0x0a, 0xde, 0x80, 0x91, 0x09, 0x96, 0xab, 0x92, 0xfe, 0x85,
	0xa7, 0x44, 0xf7, 0x67, 0xb0, 0x5b, 0xe7, 0xbc, 0x4c, 0x8e, 0x32, 0xe7, 0xe5, 0xba, 0x5a, 0x59,
	0xad, 0x4b, 0x89, 0xe8, 0x33, 0x18, 0x6c, 0xcb, 0x42, 0xb0, 0x42, 0xa8, 0x1a, 0xec, 0x3b, 0x74,
	0x16, 0xb7, 0x68, 0x14, 0x7c, 0x0a, 0x71, 0x7f, 0x82, 0xd1, 0xdf, 0x25, 0x74, 0x03, 0x86, 0x60,
	0x6f, 0x45, 0x5b, 0x98, 0x29, 0x93, 0x23, 0xf6, 0x56, 0x2c, 0x2f, 0xb0, 0xf2, 0xa3, 0xd7, 0xd0,
	0xcf, 0xf2, 0x38, 0x65, 0x2d, 0xdd, 0x92, 0x01, 0xbe, 0x74, 0x2c, 0x2f, 0x70, 0xa3, 0xcc, 0x2d,
	0x18, 0xb4, 0x34, 0xf7, 0x0f, 0xbd, 0x6b, 0x41, 0xf6, 0x85, 0x3e, 0x05, 0x9b, 0xb3, 0x22, 0x61,
	0x35, 0x3d, 0x5b, 0x08, 0x34, 0xae, 0x50, 0xae, 0xe4, 0x8b, 0x2e, 0x40, 0xb5, 0xaa, 0xab, 0x56,
	0x47, 0xf2, 0x12, 0xa2, 0xdc, 0xaa, 0xd3, 0x36, 0x41, 0x7d, 0x97, 0x6f, 0xe0, 0xaa, 0x66, 0x5b,
	0x96, 0xfd, 0x76, 0x62, 0x36, 0x9f, 0xe0, 0xf0, 0xe4, 0x54, 0xd4, 0xaf, 0xcf, 0x82, 0x14, 0xd7,
	0x50, 0x5c, 0x47, 0x72, 0x71, 0x2b, 0x28, 0x72, 0x97, 0xa6, 0xd8, 0x33, 0x30, 0x73, 0x9e, 0xd2,
	0x4d, 0x99, 0x1c, 0xc7, 0xfd, 0x7f, 0x2c, 0x4a, 0xee, 0x04, 0x0f, 0x72, 0x9e, 0xaa, 0xe5, 0x7c,
	0x02, 0x96, 0xac, 0x8a, 0x8a, 0x2c, 0x67, 0xe3, 0xcb, 0x89, 0x36, 0xed, 0x61, 0x53, 0x3a, 0xa2,
	0x2c, 0x67, 0xee, 0x2b, 0x30, 0xe4, 0x10, 0xe5, 0x83, 0xe8, 0x86, 0x6b, 0x35, 0x03, 0x75, 0x7f,
	0x87, 0xbe, 0x9a, 0x1f, 0x1a, 0x81, 0x9e, 0x25, 0xad, 0xa4, 0x67, 0x09, 0xfa, 0x00, 0xfa, 0xcf,
	0x59, 0x22, 0x76, 0xed, 0x73, 0x6b, 0x0c, 0xf4, 0x21, 0x5c, 0xee, 0x58, 0x96, 0xee, 0x84, 0x6a,
	0xb4, 0x8f, 0x5b, 0x4b, 0x3e, 0x82, 0x43, 0xbd, 0x57, 0x8d, 0x59, 0x58, 0x1e, 0xe5, 0x64, 0xc4,
	0xee, 0x90, 0x6f, 0x8a, 0x38, 0xdb, 0x53, 0xa9, 0xf5, 0x9b, 0xc9, 0x74, 0xce, 0x75, 0xbd, 0x9f,
	0x31, 0xb0, 0xcf, 0x9e, 0x3c, 0x1a, 0x01, 0xac, 0x22, 0xba, 0x0e, 0xbf, 0x0b, 0x1f, 0xbf, 0x0f,
	0x9d, 0x8b, 0xd6, 0x26, 0xfe, 0x43, 0x48, 0xfd, 0xd0, 0xd1, 0xd0, 0xfb, 0x70, 0x2d, 0xed, 0x1f,
	0xc3, 0x05, 0x0d, 0x3c, 0x42, 0xee, 0x1f, 0x3c, 0x47, 0x47, 0xef, 0xc1, 0xd5, 0x2a, 0xa2, 0x4b,
	0xef, 0x1e, 0x47, 0x74, 0xee, 0xdd, 0x47, 0x4e, 0xaf, 0xcd, 0x3b, 0x85, 0x18, 0xb3, 0x6f, 0xba,
	0xcf, 0xe0, 0x74, 0x4d, 0x70, 0x7e, 0x8d, 0x0d, 0x83, 0x20, 0xa2, 0x91, 0xf7, 0x43, 0xe4, 0x68,
	0x68, 0x08, 0x66, 0x10, 0x51, 0x3f, 0x50, 0xf0, 0x99, 0x0f, 0xf0, 0x6e, 0xf3, 0x32, 0x91, 0xbc,
	0x48, 0x94, 0x36, 0xf1, 0xb0, 0xa3, 0xa1, 0x2b, 0xb0, 0x88, 0x2c, 0x8e, 0x44, 0x5e, 0xe0, 0xe8,
	0xe8, 0x1a, 0x6c, 0x12, 0xd1, 0xf9, 0x9a, 0xf8, 0xa1, 0x47, 0x88, 0xd3, 0x9b, 0x7d, 0x0b, 0xc3,
	0xf3, 0x65, 0x4b, 0x18, 0x7e, 0x01, 0xc3, 0x1d, 0x6c, 0x08, 0x26, 0x8e, 0xe8, 0x03, 0x7e, 0x5c,
	0xaf, 0x1c, 0x7d, 0x73, 0xa9, 0x7e, 0xd0, 0x5f, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0x97, 0x58,
	0x65, 0xfc, 0xb2, 0x05, 0x00, 0x00,
}
