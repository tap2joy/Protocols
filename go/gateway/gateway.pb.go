// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: gateway/gateway.proto

package gateway

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

// 登陆
type CLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Channel uint32 `protobuf:"varint,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *CLogin) Reset() {
	*x = CLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLogin) ProtoMessage() {}

func (x *CLogin) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLogin.ProtoReflect.Descriptor instead.
func (*CLogin) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *CLogin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CLogin) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type SLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Channel uint32 `protobuf:"varint,4,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *SLogin) Reset() {
	*x = SLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SLogin) ProtoMessage() {}

func (x *SLogin) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SLogin.ProtoReflect.Descriptor instead.
func (*SLogin) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *SLogin) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SLogin) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SLogin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SLogin) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

// 登出
type CLogout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CLogout) Reset() {
	*x = CLogout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CLogout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CLogout) ProtoMessage() {}

func (x *CLogout) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CLogout.ProtoReflect.Descriptor instead.
func (*CLogout) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *CLogout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SLogout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SLogout) Reset() {
	*x = SLogout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SLogout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SLogout) ProtoMessage() {}

func (x *SLogout) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SLogout.ProtoReflect.Descriptor instead.
func (*SLogout) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{3}
}

func (x *SLogout) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SLogout) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SLogout) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 进入指定聊天室
type CChangeChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel uint32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CChangeChannel) Reset() {
	*x = CChangeChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CChangeChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CChangeChannel) ProtoMessage() {}

func (x *CChangeChannel) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CChangeChannel.ProtoReflect.Descriptor instead.
func (*CChangeChannel) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{4}
}

func (x *CChangeChannel) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *CChangeChannel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SChangeChannel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Channel uint32     `protobuf:"varint,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Logs    []*ChatLog `protobuf:"bytes,4,rep,name=logs,proto3" json:"logs,omitempty"` // 该频道的聊天记录
}

func (x *SChangeChannel) Reset() {
	*x = SChangeChannel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SChangeChannel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SChangeChannel) ProtoMessage() {}

func (x *SChangeChannel) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SChangeChannel.ProtoReflect.Descriptor instead.
func (*SChangeChannel) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{5}
}

func (x *SChangeChannel) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SChangeChannel) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SChangeChannel) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *SChangeChannel) GetLogs() []*ChatLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

// 发送消息
type CSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderUid  uint32 `protobuf:"varint,1,opt,name=sender_uid,json=senderUid,proto3" json:"sender_uid,omitempty"`   // 发送者uid
	SenderName string `protobuf:"bytes,2,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"` // 发送者名字
	Channel    uint32 `protobuf:"varint,3,opt,name=channel,proto3" json:"channel,omitempty"`                        // 发送频道
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`                         // 聊天内容
}

func (x *CSend) Reset() {
	*x = CSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSend) ProtoMessage() {}

func (x *CSend) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSend.ProtoReflect.Descriptor instead.
func (*CSend) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{6}
}

func (x *CSend) GetSenderUid() uint32 {
	if x != nil {
		return x.SenderUid
	}
	return 0
}

func (x *CSend) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *CSend) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

func (x *CSend) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Result string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SSend) Reset() {
	*x = SSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSend) ProtoMessage() {}

func (x *SSend) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSend.ProtoReflect.Descriptor instead.
func (*SSend) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{7}
}

func (x *SSend) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SSend) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SSend) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ChatLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderName string `protobuf:"bytes,1,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"` // 发送者名字
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                         // 聊天内容
	Timestamp  uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                    // 时间戳
}

func (x *ChatLog) Reset() {
	*x = ChatLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatLog) ProtoMessage() {}

func (x *ChatLog) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatLog.ProtoReflect.Descriptor instead.
func (*ChatLog) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{8}
}

func (x *ChatLog) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *ChatLog) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatLog) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 获取聊天记录
type CGetLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel uint32 `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *CGetLog) Reset() {
	*x = CGetLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGetLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGetLog) ProtoMessage() {}

func (x *CGetLog) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGetLog.ProtoReflect.Descriptor instead.
func (*CGetLog) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{9}
}

func (x *CGetLog) GetChannel() uint32 {
	if x != nil {
		return x.Channel
	}
	return 0
}

type SGetLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Logs []*ChatLog `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *SGetLog) Reset() {
	*x = SGetLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SGetLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SGetLog) ProtoMessage() {}

func (x *SGetLog) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SGetLog.ProtoReflect.Descriptor instead.
func (*SGetLog) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{10}
}

func (x *SGetLog) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SGetLog) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SGetLog) GetLogs() []*ChatLog {
	if x != nil {
		return x.Logs
	}
	return nil
}

// 推送消息
type SPushMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderName string `protobuf:"bytes,1,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"` // 发送者名字
	Content    string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`                         // 聊天内容
	Timestamp  uint64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                    // 时间戳
}

func (x *SPushMessage) Reset() {
	*x = SPushMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPushMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPushMessage) ProtoMessage() {}

func (x *SPushMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPushMessage.ProtoReflect.Descriptor instead.
func (*SPushMessage) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{11}
}

func (x *SPushMessage) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *SPushMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SPushMessage) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 踢人
type SKickUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"` // 踢出原因
}

func (x *SKickUser) Reset() {
	*x = SKickUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKickUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKickUser) ProtoMessage() {}

func (x *SKickUser) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKickUser.ProtoReflect.Descriptor instead.
func (*SKickUser) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{12}
}

func (x *SKickUser) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

// 获取频道列表
type CGetChannelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CGetChannelList) Reset() {
	*x = CGetChannelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGetChannelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGetChannelList) ProtoMessage() {}

func (x *CGetChannelList) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGetChannelList.ProtoReflect.Descriptor instead.
func (*CGetChannelList) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{13}
}

func (x *CGetChannelList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ChannelData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *ChannelData) Reset() {
	*x = ChannelData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelData) ProtoMessage() {}

func (x *ChannelData) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelData.ProtoReflect.Descriptor instead.
func (*ChannelData) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{14}
}

func (x *ChannelData) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChannelData) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type SGetChannelList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	List []*ChannelData `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *SGetChannelList) Reset() {
	*x = SGetChannelList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SGetChannelList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SGetChannelList) ProtoMessage() {}

func (x *SGetChannelList) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SGetChannelList.ProtoReflect.Descriptor instead.
func (*SGetChannelList) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{15}
}

func (x *SGetChannelList) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SGetChannelList) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SGetChannelList) GetList() []*ChannelData {
	if x != nil {
		return x.List
	}
	return nil
}

var File_gateway_gateway_proto protoreflect.FileDescriptor

var file_gateway_gateway_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x06, 0x43, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0x5c, 0x0a, 0x06, 0x53, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x1d, 0x0a,
	0x07, 0x43, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x43, 0x0a, 0x07,
	0x53, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x6e, 0x0a, 0x0e, 0x53, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x22, 0x7b, 0x0a, 0x05, 0x43, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x45,
	0x0a, 0x05, 0x53, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x62, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x4d,
	0x0a, 0x07, 0x53, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x1c, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x67, 0x0a,
	0x0c, 0x53, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x23, 0x0a, 0x09, 0x53, 0x4b, 0x69, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0f, 0x43,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x31, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x59, 0x0a, 0x0f, 0x53, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x20,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x61, 0x70, 0x32, 0x6a, 0x6f, 0x79, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_gateway_gateway_proto_rawDescOnce sync.Once
	file_gateway_gateway_proto_rawDescData = file_gateway_gateway_proto_rawDesc
)

func file_gateway_gateway_proto_rawDescGZIP() []byte {
	file_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_proto_rawDescData)
	})
	return file_gateway_gateway_proto_rawDescData
}

var file_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_gateway_gateway_proto_goTypes = []interface{}{
	(*CLogin)(nil),          // 0: CLogin
	(*SLogin)(nil),          // 1: SLogin
	(*CLogout)(nil),         // 2: CLogout
	(*SLogout)(nil),         // 3: SLogout
	(*CChangeChannel)(nil),  // 4: CChangeChannel
	(*SChangeChannel)(nil),  // 5: SChangeChannel
	(*CSend)(nil),           // 6: CSend
	(*SSend)(nil),           // 7: SSend
	(*ChatLog)(nil),         // 8: ChatLog
	(*CGetLog)(nil),         // 9: CGetLog
	(*SGetLog)(nil),         // 10: SGetLog
	(*SPushMessage)(nil),    // 11: SPushMessage
	(*SKickUser)(nil),       // 12: SKickUser
	(*CGetChannelList)(nil), // 13: CGetChannelList
	(*ChannelData)(nil),     // 14: ChannelData
	(*SGetChannelList)(nil), // 15: SGetChannelList
}
var file_gateway_gateway_proto_depIdxs = []int32{
	8,  // 0: SChangeChannel.logs:type_name -> ChatLog
	8,  // 1: SGetLog.logs:type_name -> ChatLog
	14, // 2: SGetChannelList.list:type_name -> ChannelData
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_gateway_gateway_proto_init() }
func file_gateway_gateway_proto_init() {
	if File_gateway_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLogin); i {
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
		file_gateway_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SLogin); i {
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
		file_gateway_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CLogout); i {
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
		file_gateway_gateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SLogout); i {
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
		file_gateway_gateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CChangeChannel); i {
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
		file_gateway_gateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SChangeChannel); i {
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
		file_gateway_gateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CSend); i {
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
		file_gateway_gateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSend); i {
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
		file_gateway_gateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatLog); i {
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
		file_gateway_gateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CGetLog); i {
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
		file_gateway_gateway_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SGetLog); i {
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
		file_gateway_gateway_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPushMessage); i {
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
		file_gateway_gateway_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKickUser); i {
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
		file_gateway_gateway_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CGetChannelList); i {
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
		file_gateway_gateway_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelData); i {
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
		file_gateway_gateway_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SGetChannelList); i {
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
			RawDescriptor: file_gateway_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_proto_msgTypes,
	}.Build()
	File_gateway_gateway_proto = out.File
	file_gateway_gateway_proto_rawDesc = nil
	file_gateway_gateway_proto_goTypes = nil
	file_gateway_gateway_proto_depIdxs = nil
}
