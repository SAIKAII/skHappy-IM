// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cli-interface.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Sex                  int32    `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	Birthday             int64    `protobuf:"varint,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	PhoneNum             string   `protobuf:"bytes,7,opt,name=phone_num,json=phoneNum,proto3" json:"phone_num,omitempty"`
	CreateTime           int64    `protobuf:"varint,8,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           int64    `protobuf:"varint,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *User) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *User) GetBirthday() int64 {
	if m != nil {
		return m.Birthday
	}
	return 0
}

func (m *User) GetPhoneNum() string {
	if m != nil {
		return m.PhoneNum
	}
	return ""
}

func (m *User) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *User) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type RegisterReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{1}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type RegisterResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{2}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

type AddFriendReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId             string   `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFriendReq) Reset()         { *m = AddFriendReq{} }
func (m *AddFriendReq) String() string { return proto.CompactTextString(m) }
func (*AddFriendReq) ProtoMessage()    {}
func (*AddFriendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{3}
}

func (m *AddFriendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFriendReq.Unmarshal(m, b)
}
func (m *AddFriendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFriendReq.Marshal(b, m, deterministic)
}
func (m *AddFriendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFriendReq.Merge(m, src)
}
func (m *AddFriendReq) XXX_Size() int {
	return xxx_messageInfo_AddFriendReq.Size(m)
}
func (m *AddFriendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFriendReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddFriendReq proto.InternalMessageInfo

func (m *AddFriendReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *AddFriendReq) GetFriendId() string {
	if m != nil {
		return m.FriendId
	}
	return ""
}

type AddFriendResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddFriendResp) Reset()         { *m = AddFriendResp{} }
func (m *AddFriendResp) String() string { return proto.CompactTextString(m) }
func (*AddFriendResp) ProtoMessage()    {}
func (*AddFriendResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{4}
}

func (m *AddFriendResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddFriendResp.Unmarshal(m, b)
}
func (m *AddFriendResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddFriendResp.Marshal(b, m, deterministic)
}
func (m *AddFriendResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddFriendResp.Merge(m, src)
}
func (m *AddFriendResp) XXX_Size() int {
	return xxx_messageInfo_AddFriendResp.Size(m)
}
func (m *AddFriendResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddFriendResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddFriendResp proto.InternalMessageInfo

type GetUserReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{5}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

func (m *GetUserReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserResp struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{6}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ListUsersReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersReq) Reset()         { *m = ListUsersReq{} }
func (m *ListUsersReq) String() string { return proto.CompactTextString(m) }
func (*ListUsersReq) ProtoMessage()    {}
func (*ListUsersReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{7}
}

func (m *ListUsersReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersReq.Unmarshal(m, b)
}
func (m *ListUsersReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersReq.Marshal(b, m, deterministic)
}
func (m *ListUsersReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersReq.Merge(m, src)
}
func (m *ListUsersReq) XXX_Size() int {
	return xxx_messageInfo_ListUsersReq.Size(m)
}
func (m *ListUsersReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersReq proto.InternalMessageInfo

func (m *ListUsersReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ListUsersResp struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResp) Reset()         { *m = ListUsersResp{} }
func (m *ListUsersResp) String() string { return proto.CompactTextString(m) }
func (*ListUsersResp) ProtoMessage()    {}
func (*ListUsersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{8}
}

func (m *ListUsersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResp.Unmarshal(m, b)
}
func (m *ListUsersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResp.Marshal(b, m, deterministic)
}
func (m *ListUsersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResp.Merge(m, src)
}
func (m *ListUsersResp) XXX_Size() int {
	return xxx_messageInfo_ListUsersResp.Size(m)
}
func (m *ListUsersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResp proto.InternalMessageInfo

func (m *ListUsersResp) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type SendMessageReq struct {
	Item                 *MessageItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SendMessageReq) Reset()         { *m = SendMessageReq{} }
func (m *SendMessageReq) String() string { return proto.CompactTextString(m) }
func (*SendMessageReq) ProtoMessage()    {}
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{9}
}

func (m *SendMessageReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReq.Unmarshal(m, b)
}
func (m *SendMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReq.Marshal(b, m, deterministic)
}
func (m *SendMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReq.Merge(m, src)
}
func (m *SendMessageReq) XXX_Size() int {
	return xxx_messageInfo_SendMessageReq.Size(m)
}
func (m *SendMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReq proto.InternalMessageInfo

func (m *SendMessageReq) GetItem() *MessageItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type SendMessageResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageResp) Reset()         { *m = SendMessageResp{} }
func (m *SendMessageResp) String() string { return proto.CompactTextString(m) }
func (*SendMessageResp) ProtoMessage()    {}
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{10}
}

func (m *SendMessageResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResp.Unmarshal(m, b)
}
func (m *SendMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResp.Marshal(b, m, deterministic)
}
func (m *SendMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResp.Merge(m, src)
}
func (m *SendMessageResp) XXX_Size() int {
	return xxx_messageInfo_SendMessageResp.Size(m)
}
func (m *SendMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResp proto.InternalMessageInfo

type DelFriendReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FriendName           string   `protobuf:"bytes,2,opt,name=friendName,proto3" json:"friendName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelFriendReq) Reset()         { *m = DelFriendReq{} }
func (m *DelFriendReq) String() string { return proto.CompactTextString(m) }
func (*DelFriendReq) ProtoMessage()    {}
func (*DelFriendReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{11}
}

func (m *DelFriendReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelFriendReq.Unmarshal(m, b)
}
func (m *DelFriendReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelFriendReq.Marshal(b, m, deterministic)
}
func (m *DelFriendReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelFriendReq.Merge(m, src)
}
func (m *DelFriendReq) XXX_Size() int {
	return xxx_messageInfo_DelFriendReq.Size(m)
}
func (m *DelFriendReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelFriendReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelFriendReq proto.InternalMessageInfo

func (m *DelFriendReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DelFriendReq) GetFriendName() string {
	if m != nil {
		return m.FriendName
	}
	return ""
}

type DelFriendResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelFriendResp) Reset()         { *m = DelFriendResp{} }
func (m *DelFriendResp) String() string { return proto.CompactTextString(m) }
func (*DelFriendResp) ProtoMessage()    {}
func (*DelFriendResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{12}
}

func (m *DelFriendResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelFriendResp.Unmarshal(m, b)
}
func (m *DelFriendResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelFriendResp.Marshal(b, m, deterministic)
}
func (m *DelFriendResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelFriendResp.Merge(m, src)
}
func (m *DelFriendResp) XXX_Size() int {
	return xxx_messageInfo_DelFriendResp.Size(m)
}
func (m *DelFriendResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelFriendResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelFriendResp proto.InternalMessageInfo

type UpdateProfileReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileReq) Reset()         { *m = UpdateProfileReq{} }
func (m *UpdateProfileReq) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileReq) ProtoMessage()    {}
func (*UpdateProfileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{13}
}

func (m *UpdateProfileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileReq.Unmarshal(m, b)
}
func (m *UpdateProfileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileReq.Marshal(b, m, deterministic)
}
func (m *UpdateProfileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileReq.Merge(m, src)
}
func (m *UpdateProfileReq) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileReq.Size(m)
}
func (m *UpdateProfileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileReq proto.InternalMessageInfo

func (m *UpdateProfileReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateProfileResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateProfileResp) Reset()         { *m = UpdateProfileResp{} }
func (m *UpdateProfileResp) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileResp) ProtoMessage()    {}
func (*UpdateProfileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eb1df2ca22d1b82, []int{14}
}

func (m *UpdateProfileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileResp.Unmarshal(m, b)
}
func (m *UpdateProfileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileResp.Marshal(b, m, deterministic)
}
func (m *UpdateProfileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileResp.Merge(m, src)
}
func (m *UpdateProfileResp) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileResp.Size(m)
}
func (m *UpdateProfileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*RegisterReq)(nil), "pb.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "pb.RegisterResp")
	proto.RegisterType((*AddFriendReq)(nil), "pb.AddFriendReq")
	proto.RegisterType((*AddFriendResp)(nil), "pb.AddFriendResp")
	proto.RegisterType((*GetUserReq)(nil), "pb.GetUserReq")
	proto.RegisterType((*GetUserResp)(nil), "pb.GetUserResp")
	proto.RegisterType((*ListUsersReq)(nil), "pb.ListUsersReq")
	proto.RegisterType((*ListUsersResp)(nil), "pb.ListUsersResp")
	proto.RegisterType((*SendMessageReq)(nil), "pb.SendMessageReq")
	proto.RegisterType((*SendMessageResp)(nil), "pb.SendMessageResp")
	proto.RegisterType((*DelFriendReq)(nil), "pb.DelFriendReq")
	proto.RegisterType((*DelFriendResp)(nil), "pb.DelFriendResp")
	proto.RegisterType((*UpdateProfileReq)(nil), "pb.UpdateProfileReq")
	proto.RegisterType((*UpdateProfileResp)(nil), "pb.UpdateProfileResp")
}

func init() {
	proto.RegisterFile("cli-interface.proto", fileDescriptor_6eb1df2ca22d1b82)
}

var fileDescriptor_6eb1df2ca22d1b82 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6f, 0xd3, 0x4e,
	0x10, 0x55, 0x3e, 0x1b, 0x4f, 0xbe, 0x37, 0xbf, 0x9f, 0xb0, 0x0c, 0x94, 0xc8, 0x5c, 0x22, 0x20,
	0xa1, 0x0a, 0x82, 0x13, 0x17, 0x44, 0x45, 0x15, 0x04, 0x15, 0x72, 0xc9, 0x39, 0x72, 0xec, 0x49,
	0xbb, 0xc2, 0x1f, 0xdb, 0xdd, 0x4d, 0x81, 0x1b, 0x27, 0xfe, 0x6e, 0xb4, 0xbb, 0x8e, 0x63, 0x07,
	0xd4, 0xdc, 0x3a, 0xef, 0xbd, 0x99, 0xf1, 0xbe, 0x37, 0x0d, 0x8c, 0x82, 0x88, 0x4e, 0x69, 0x22,
	0x91, 0x6f, 0xfc, 0x00, 0x67, 0x8c, 0xa7, 0x32, 0x25, 0x55, 0xb6, 0x76, 0x7a, 0x32, 0x60, 0xd3,
	0x20, 0x4d, 0x12, 0x83, 0xb9, 0xbf, 0xab, 0x50, 0x5f, 0x0a, 0xe4, 0xc4, 0x81, 0xd6, 0x56, 0x20,
	0x4f, 0xfc, 0x18, 0xed, 0xca, 0xb8, 0x32, 0xb1, 0xbc, 0xbc, 0x56, 0x5c, 0x42, 0x83, 0x6f, 0x9a,
	0xab, 0x1a, 0x6e, 0x57, 0x2b, 0x8e, 0xf9, 0x42, 0x7c, 0x4f, 0x79, 0x68, 0xd7, 0x0c, 0xb7, 0xab,
	0xc9, 0x63, 0x00, 0xff, 0xce, 0x97, 0x3e, 0x5f, 0x6d, 0x79, 0x64, 0xd7, 0x35, 0x6b, 0x19, 0x64,
	0xc9, 0x23, 0x32, 0x80, 0x9a, 0xc0, 0x1f, 0x76, 0x63, 0x5c, 0x99, 0x34, 0x3c, 0xf5, 0xa7, 0x1a,
	0xb6, 0xa6, 0x5c, 0xde, 0x84, 0xfe, 0x4f, 0xbb, 0x39, 0xae, 0x4c, 0x6a, 0x5e, 0x5e, 0x93, 0x87,
	0x60, 0xb1, 0x9b, 0x34, 0xc1, 0x55, 0xb2, 0x8d, 0xed, 0x93, 0x6c, 0x93, 0x02, 0x2e, 0xb7, 0x31,
	0x79, 0x02, 0xed, 0x80, 0xa3, 0x2f, 0x71, 0x25, 0x69, 0x8c, 0x76, 0x4b, 0xf7, 0x82, 0x81, 0xbe,
	0xd2, 0x18, 0x95, 0x60, 0xcb, 0xc2, 0x5c, 0x60, 0x19, 0x81, 0x81, 0x94, 0xc0, 0x7d, 0x0e, 0x6d,
	0x0f, 0xaf, 0xa9, 0x90, 0xc8, 0x3d, 0xbc, 0x25, 0x8f, 0xa0, 0xae, 0x9e, 0xaf, 0xad, 0x68, 0xcf,
	0x5b, 0x33, 0xb6, 0x9e, 0x29, 0x9b, 0x3c, 0x8d, 0xba, 0x3d, 0xe8, 0xec, 0xc5, 0x82, 0xb9, 0xe7,
	0xd0, 0x79, 0x17, 0x86, 0x1f, 0x38, 0xc5, 0x24, 0x54, 0xdd, 0x0f, 0xe0, 0x44, 0xe9, 0x56, 0x34,
	0xcc, 0xbc, 0x6c, 0xaa, 0x72, 0x11, 0xaa, 0x47, 0x6c, 0xb4, 0x4a, 0x51, 0x99, 0x95, 0x06, 0x58,
	0x84, 0x6e, 0x1f, 0xba, 0x85, 0x29, 0x82, 0xb9, 0x13, 0x80, 0x0b, 0x94, 0x7a, 0x2f, 0xde, 0xde,
	0x97, 0x90, 0xfa, 0xfa, 0x5c, 0x29, 0xd8, 0x91, 0xaf, 0x7f, 0x06, 0x9d, 0x4f, 0x54, 0x68, 0xb5,
	0x38, 0x36, 0xf8, 0x25, 0x74, 0x0b, 0x5a, 0xc1, 0xc8, 0x29, 0x34, 0x14, 0x29, 0xec, 0xca, 0xb8,
	0x56, 0x9a, 0x6d, 0x60, 0xf7, 0x35, 0xf4, 0xae, 0x30, 0x09, 0x3f, 0xa3, 0x10, 0xfe, 0x35, 0xaa,
	0xf1, 0x4f, 0xa1, 0x4e, 0x25, 0xc6, 0xd9, 0xc7, 0xf4, 0x55, 0x43, 0xc6, 0x2e, 0x24, 0xc6, 0x9e,
	0x26, 0xdd, 0x21, 0xf4, 0x4b, 0x6d, 0x82, 0xb9, 0x1f, 0xa1, 0x73, 0x8e, 0xd1, 0xde, 0xd4, 0xfb,
	0x2e, 0xf4, 0x14, 0xc0, 0xd8, 0x78, 0xb9, 0xbf, 0xd1, 0x02, 0xa2, 0xac, 0x2d, 0xcc, 0x12, 0xcc,
	0x3d, 0x83, 0xc1, 0x52, 0x87, 0xff, 0x85, 0xa7, 0x1b, 0x1a, 0xe1, 0xf1, 0xcc, 0x47, 0x30, 0x3c,
	0xe8, 0x10, 0x6c, 0xfe, 0xab, 0x06, 0xa3, 0xf7, 0x11, 0x5d, 0xec, 0xfe, 0xd3, 0xae, 0x90, 0xdf,
	0xd1, 0x00, 0xc9, 0x14, 0x5a, 0xbb, 0x03, 0x21, 0xfa, 0xc5, 0x85, 0xdb, 0x72, 0x06, 0x65, 0x40,
	0x30, 0x72, 0x06, 0x56, 0x9e, 0x3c, 0xd1, 0x74, 0xf1, 0x9c, 0x9c, 0xe1, 0x01, 0x22, 0x18, 0x79,
	0x01, 0xd6, 0x05, 0xca, 0xac, 0xa3, 0xa7, 0xf8, 0xfd, 0xa5, 0x38, 0xfd, 0x52, 0x2d, 0x18, 0x99,
	0x43, 0x5b, 0xa5, 0x68, 0xe4, 0xc2, 0x6c, 0x28, 0x9e, 0x80, 0xd9, 0x50, 0x0e, 0xfa, 0x0d, 0xb4,
	0x0b, 0x89, 0x10, 0xa2, 0x14, 0xe5, 0x64, 0x9d, 0xd1, 0x5f, 0x98, 0x79, 0x4b, 0x6e, 0xb5, 0xd9,
	0x54, 0x4c, 0xd1, 0x6c, 0x2a, 0x65, 0x41, 0xde, 0x42, 0xb7, 0xe4, 0x2c, 0xf9, 0x4f, 0x5b, 0x7f,
	0x10, 0x8f, 0xf3, 0xff, 0x3f, 0x50, 0xc1, 0xd6, 0x4d, 0xfd, 0x43, 0xf6, 0xea, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf3, 0xe9, 0x69, 0xac, 0xf3, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CliInterfaceServiceClient is the client API for CliInterfaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliInterfaceServiceClient interface {
	// Register 用户注册
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	// AddFriend 添加好友
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error)
	// GetFriend 获取指定好友详细信息
	GetFriend(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	// ListFriends 列出指定用户的所有好友
	ListFriends(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error)
	// SendMessage 发送消息给好友
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
	// DelFriend 删除指定好友
	DelFriend(ctx context.Context, in *DelFriendReq, opts ...grpc.CallOption) (*DelFriendResp, error)
	// UpdateProfile 更新个人信息
	UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*UpdateProfileResp, error)
}

type cliInterfaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCliInterfaceServiceClient(cc grpc.ClientConnInterface) CliInterfaceServiceClient {
	return &cliInterfaceServiceClient{cc}
}

func (c *cliInterfaceServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*AddFriendResp, error) {
	out := new(AddFriendResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) GetFriend(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/GetFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) ListFriends(ctx context.Context, in *ListUsersReq, opts ...grpc.CallOption) (*ListUsersResp, error) {
	out := new(ListUsersResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/ListFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) DelFriend(ctx context.Context, in *DelFriendReq, opts ...grpc.CallOption) (*DelFriendResp, error) {
	out := new(DelFriendResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/DelFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliInterfaceServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileReq, opts ...grpc.CallOption) (*UpdateProfileResp, error) {
	out := new(UpdateProfileResp)
	err := c.cc.Invoke(ctx, "/pb.CliInterfaceService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CliInterfaceServiceServer is the server API for CliInterfaceService service.
type CliInterfaceServiceServer interface {
	// Register 用户注册
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	// AddFriend 添加好友
	AddFriend(context.Context, *AddFriendReq) (*AddFriendResp, error)
	// GetFriend 获取指定好友详细信息
	GetFriend(context.Context, *GetUserReq) (*GetUserResp, error)
	// ListFriends 列出指定用户的所有好友
	ListFriends(context.Context, *ListUsersReq) (*ListUsersResp, error)
	// SendMessage 发送消息给好友
	SendMessage(context.Context, *SendMessageReq) (*SendMessageResp, error)
	// DelFriend 删除指定好友
	DelFriend(context.Context, *DelFriendReq) (*DelFriendResp, error)
	// UpdateProfile 更新个人信息
	UpdateProfile(context.Context, *UpdateProfileReq) (*UpdateProfileResp, error)
}

// UnimplementedCliInterfaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCliInterfaceServiceServer struct {
}

func (*UnimplementedCliInterfaceServiceServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) AddFriend(ctx context.Context, req *AddFriendReq) (*AddFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) GetFriend(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriend not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) ListFriends(ctx context.Context, req *ListUsersReq) (*ListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFriends not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) SendMessage(ctx context.Context, req *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) DelFriend(ctx context.Context, req *DelFriendReq) (*DelFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelFriend not implemented")
}
func (*UnimplementedCliInterfaceServiceServer) UpdateProfile(ctx context.Context, req *UpdateProfileReq) (*UpdateProfileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}

func RegisterCliInterfaceServiceServer(s *grpc.Server, srv CliInterfaceServiceServer) {
	s.RegisterService(&_CliInterfaceService_serviceDesc, srv)
}

func _CliInterfaceService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_GetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).GetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/GetFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).GetFriend(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_ListFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).ListFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/ListFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).ListFriends(ctx, req.(*ListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).SendMessage(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_DelFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).DelFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/DelFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).DelFriend(ctx, req.(*DelFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliInterfaceService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliInterfaceServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CliInterfaceService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliInterfaceServiceServer).UpdateProfile(ctx, req.(*UpdateProfileReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CliInterfaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CliInterfaceService",
	HandlerType: (*CliInterfaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _CliInterfaceService_Register_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _CliInterfaceService_AddFriend_Handler,
		},
		{
			MethodName: "GetFriend",
			Handler:    _CliInterfaceService_GetFriend_Handler,
		},
		{
			MethodName: "ListFriends",
			Handler:    _CliInterfaceService_ListFriends_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _CliInterfaceService_SendMessage_Handler,
		},
		{
			MethodName: "DelFriend",
			Handler:    _CliInterfaceService_DelFriend_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _CliInterfaceService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli-interface.proto",
}
