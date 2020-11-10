// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user_balance_log.proto

package user_balance_log

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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserBalanceLog struct {
	BalanceId            int64    `protobuf:"varint,1,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Source               int64    `protobuf:"varint,3,opt,name=source,proto3" json:"source,omitempty"`
	SourceSn             int64    `protobuf:"varint,4,opt,name=source_sn,json=sourceSn,proto3" json:"source_sn,omitempty"`
	CreateTime           string   `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Amount               float64  `protobuf:"fixed64,6,opt,name=amount,proto3" json:"amount,omitempty"`
	UserName             string   `protobuf:"bytes,7,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserBalanceLog) Reset()         { *m = UserBalanceLog{} }
func (m *UserBalanceLog) String() string { return proto.CompactTextString(m) }
func (*UserBalanceLog) ProtoMessage()    {}
func (*UserBalanceLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{1}
}

func (m *UserBalanceLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserBalanceLog.Unmarshal(m, b)
}
func (m *UserBalanceLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserBalanceLog.Marshal(b, m, deterministic)
}
func (m *UserBalanceLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBalanceLog.Merge(m, src)
}
func (m *UserBalanceLog) XXX_Size() int {
	return xxx_messageInfo_UserBalanceLog.Size(m)
}
func (m *UserBalanceLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBalanceLog.DiscardUnknown(m)
}

var xxx_messageInfo_UserBalanceLog proto.InternalMessageInfo

func (m *UserBalanceLog) GetBalanceId() int64 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

func (m *UserBalanceLog) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserBalanceLog) GetSource() int64 {
	if m != nil {
		return m.Source
	}
	return 0
}

func (m *UserBalanceLog) GetSourceSn() int64 {
	if m != nil {
		return m.SourceSn
	}
	return 0
}

func (m *UserBalanceLog) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *UserBalanceLog) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *UserBalanceLog) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type In_GetUserBalanceInfoById struct {
	//通过id
	BalanceId            int64    `protobuf:"varint,1,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetUserBalanceInfoById) Reset()         { *m = In_GetUserBalanceInfoById{} }
func (m *In_GetUserBalanceInfoById) String() string { return proto.CompactTextString(m) }
func (*In_GetUserBalanceInfoById) ProtoMessage()    {}
func (*In_GetUserBalanceInfoById) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{2}
}

func (m *In_GetUserBalanceInfoById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetUserBalanceInfoById.Unmarshal(m, b)
}
func (m *In_GetUserBalanceInfoById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetUserBalanceInfoById.Marshal(b, m, deterministic)
}
func (m *In_GetUserBalanceInfoById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetUserBalanceInfoById.Merge(m, src)
}
func (m *In_GetUserBalanceInfoById) XXX_Size() int {
	return xxx_messageInfo_In_GetUserBalanceInfoById.Size(m)
}
func (m *In_GetUserBalanceInfoById) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetUserBalanceInfoById.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetUserBalanceInfoById proto.InternalMessageInfo

func (m *In_GetUserBalanceInfoById) GetBalanceId() int64 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

type Out_GetUserBalanceInfoById struct {
	Error                *Error          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserBalanceLog       *UserBalanceLog `protobuf:"bytes,2,opt,name=user_balance_log,json=userBalanceLog,proto3" json:"user_balance_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Out_GetUserBalanceInfoById) Reset()         { *m = Out_GetUserBalanceInfoById{} }
func (m *Out_GetUserBalanceInfoById) String() string { return proto.CompactTextString(m) }
func (*Out_GetUserBalanceInfoById) ProtoMessage()    {}
func (*Out_GetUserBalanceInfoById) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{3}
}

func (m *Out_GetUserBalanceInfoById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetUserBalanceInfoById.Unmarshal(m, b)
}
func (m *Out_GetUserBalanceInfoById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetUserBalanceInfoById.Marshal(b, m, deterministic)
}
func (m *Out_GetUserBalanceInfoById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetUserBalanceInfoById.Merge(m, src)
}
func (m *Out_GetUserBalanceInfoById) XXX_Size() int {
	return xxx_messageInfo_Out_GetUserBalanceInfoById.Size(m)
}
func (m *Out_GetUserBalanceInfoById) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetUserBalanceInfoById.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetUserBalanceInfoById proto.InternalMessageInfo

func (m *Out_GetUserBalanceInfoById) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetUserBalanceInfoById) GetUserBalanceLog() *UserBalanceLog {
	if m != nil {
		return m.UserBalanceLog
	}
	return nil
}

type Out_UpdateUserBalanceInfo struct {
	Error                *Error          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserBalanceLog       *UserBalanceLog `protobuf:"bytes,2,opt,name=user_balance_log,json=userBalanceLog,proto3" json:"user_balance_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Out_UpdateUserBalanceInfo) Reset()         { *m = Out_UpdateUserBalanceInfo{} }
func (m *Out_UpdateUserBalanceInfo) String() string { return proto.CompactTextString(m) }
func (*Out_UpdateUserBalanceInfo) ProtoMessage()    {}
func (*Out_UpdateUserBalanceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{4}
}

func (m *Out_UpdateUserBalanceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_UpdateUserBalanceInfo.Unmarshal(m, b)
}
func (m *Out_UpdateUserBalanceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_UpdateUserBalanceInfo.Marshal(b, m, deterministic)
}
func (m *Out_UpdateUserBalanceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_UpdateUserBalanceInfo.Merge(m, src)
}
func (m *Out_UpdateUserBalanceInfo) XXX_Size() int {
	return xxx_messageInfo_Out_UpdateUserBalanceInfo.Size(m)
}
func (m *Out_UpdateUserBalanceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_UpdateUserBalanceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Out_UpdateUserBalanceInfo proto.InternalMessageInfo

func (m *Out_UpdateUserBalanceInfo) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_UpdateUserBalanceInfo) GetUserBalanceLog() *UserBalanceLog {
	if m != nil {
		return m.UserBalanceLog
	}
	return nil
}

type In_UpdateUserBalanceInfo struct {
	UserBalanceLog       *UserBalanceLog `protobuf:"bytes,1,opt,name=user_balance_log,json=userBalanceLog,proto3" json:"user_balance_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *In_UpdateUserBalanceInfo) Reset()         { *m = In_UpdateUserBalanceInfo{} }
func (m *In_UpdateUserBalanceInfo) String() string { return proto.CompactTextString(m) }
func (*In_UpdateUserBalanceInfo) ProtoMessage()    {}
func (*In_UpdateUserBalanceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{5}
}

func (m *In_UpdateUserBalanceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_UpdateUserBalanceInfo.Unmarshal(m, b)
}
func (m *In_UpdateUserBalanceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_UpdateUserBalanceInfo.Marshal(b, m, deterministic)
}
func (m *In_UpdateUserBalanceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_UpdateUserBalanceInfo.Merge(m, src)
}
func (m *In_UpdateUserBalanceInfo) XXX_Size() int {
	return xxx_messageInfo_In_UpdateUserBalanceInfo.Size(m)
}
func (m *In_UpdateUserBalanceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_In_UpdateUserBalanceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_In_UpdateUserBalanceInfo proto.InternalMessageInfo

func (m *In_UpdateUserBalanceInfo) GetUserBalanceLog() *UserBalanceLog {
	if m != nil {
		return m.UserBalanceLog
	}
	return nil
}

type In_GetUserBalances struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64    `protobuf:"varint,2,opt,name=pages,proto3" json:"pages,omitempty"`
	SearchKey            string   `protobuf:"bytes,3,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	StartTime            string   `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,5,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_GetUserBalances) Reset()         { *m = In_GetUserBalances{} }
func (m *In_GetUserBalances) String() string { return proto.CompactTextString(m) }
func (*In_GetUserBalances) ProtoMessage()    {}
func (*In_GetUserBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{6}
}

func (m *In_GetUserBalances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_GetUserBalances.Unmarshal(m, b)
}
func (m *In_GetUserBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_GetUserBalances.Marshal(b, m, deterministic)
}
func (m *In_GetUserBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_GetUserBalances.Merge(m, src)
}
func (m *In_GetUserBalances) XXX_Size() int {
	return xxx_messageInfo_In_GetUserBalances.Size(m)
}
func (m *In_GetUserBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_In_GetUserBalances.DiscardUnknown(m)
}

var xxx_messageInfo_In_GetUserBalances proto.InternalMessageInfo

func (m *In_GetUserBalances) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *In_GetUserBalances) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *In_GetUserBalances) GetSearchKey() string {
	if m != nil {
		return m.SearchKey
	}
	return ""
}

func (m *In_GetUserBalances) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *In_GetUserBalances) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type Out_GetUserBalances struct {
	Error                *Error            `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Limit                int64             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Pages                int64             `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Total                int64             `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	UserBalanceLogList   []*UserBalanceLog `protobuf:"bytes,5,rep,name=user_balance_log_list,json=userBalanceLogList,proto3" json:"user_balance_log_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Out_GetUserBalances) Reset()         { *m = Out_GetUserBalances{} }
func (m *Out_GetUserBalances) String() string { return proto.CompactTextString(m) }
func (*Out_GetUserBalances) ProtoMessage()    {}
func (*Out_GetUserBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{7}
}

func (m *Out_GetUserBalances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_GetUserBalances.Unmarshal(m, b)
}
func (m *Out_GetUserBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_GetUserBalances.Marshal(b, m, deterministic)
}
func (m *Out_GetUserBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_GetUserBalances.Merge(m, src)
}
func (m *Out_GetUserBalances) XXX_Size() int {
	return xxx_messageInfo_Out_GetUserBalances.Size(m)
}
func (m *Out_GetUserBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_GetUserBalances.DiscardUnknown(m)
}

var xxx_messageInfo_Out_GetUserBalances proto.InternalMessageInfo

func (m *Out_GetUserBalances) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_GetUserBalances) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Out_GetUserBalances) GetPages() int64 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *Out_GetUserBalances) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Out_GetUserBalances) GetUserBalanceLogList() []*UserBalanceLog {
	if m != nil {
		return m.UserBalanceLogList
	}
	return nil
}

type In_DeleteUserBalances struct {
	UserBalanceLogs      []int64  `protobuf:"varint,1,rep,packed,name=user_balance_logs,json=userBalanceLogs,proto3" json:"user_balance_logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In_DeleteUserBalances) Reset()         { *m = In_DeleteUserBalances{} }
func (m *In_DeleteUserBalances) String() string { return proto.CompactTextString(m) }
func (*In_DeleteUserBalances) ProtoMessage()    {}
func (*In_DeleteUserBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{8}
}

func (m *In_DeleteUserBalances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_DeleteUserBalances.Unmarshal(m, b)
}
func (m *In_DeleteUserBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_DeleteUserBalances.Marshal(b, m, deterministic)
}
func (m *In_DeleteUserBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_DeleteUserBalances.Merge(m, src)
}
func (m *In_DeleteUserBalances) XXX_Size() int {
	return xxx_messageInfo_In_DeleteUserBalances.Size(m)
}
func (m *In_DeleteUserBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_In_DeleteUserBalances.DiscardUnknown(m)
}

var xxx_messageInfo_In_DeleteUserBalances proto.InternalMessageInfo

func (m *In_DeleteUserBalances) GetUserBalanceLogs() []int64 {
	if m != nil {
		return m.UserBalanceLogs
	}
	return nil
}

type Out_DeleteUserBalances struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out_DeleteUserBalances) Reset()         { *m = Out_DeleteUserBalances{} }
func (m *Out_DeleteUserBalances) String() string { return proto.CompactTextString(m) }
func (*Out_DeleteUserBalances) ProtoMessage()    {}
func (*Out_DeleteUserBalances) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{9}
}

func (m *Out_DeleteUserBalances) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_DeleteUserBalances.Unmarshal(m, b)
}
func (m *Out_DeleteUserBalances) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_DeleteUserBalances.Marshal(b, m, deterministic)
}
func (m *Out_DeleteUserBalances) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_DeleteUserBalances.Merge(m, src)
}
func (m *Out_DeleteUserBalances) XXX_Size() int {
	return xxx_messageInfo_Out_DeleteUserBalances.Size(m)
}
func (m *Out_DeleteUserBalances) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_DeleteUserBalances.DiscardUnknown(m)
}

var xxx_messageInfo_Out_DeleteUserBalances proto.InternalMessageInfo

func (m *Out_DeleteUserBalances) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type In_CreateUserBalance struct {
	UserBalanceLog       *UserBalanceLog `protobuf:"bytes,1,opt,name=user_balance_log,json=userBalanceLog,proto3" json:"user_balance_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *In_CreateUserBalance) Reset()         { *m = In_CreateUserBalance{} }
func (m *In_CreateUserBalance) String() string { return proto.CompactTextString(m) }
func (*In_CreateUserBalance) ProtoMessage()    {}
func (*In_CreateUserBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{10}
}

func (m *In_CreateUserBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In_CreateUserBalance.Unmarshal(m, b)
}
func (m *In_CreateUserBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In_CreateUserBalance.Marshal(b, m, deterministic)
}
func (m *In_CreateUserBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In_CreateUserBalance.Merge(m, src)
}
func (m *In_CreateUserBalance) XXX_Size() int {
	return xxx_messageInfo_In_CreateUserBalance.Size(m)
}
func (m *In_CreateUserBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_In_CreateUserBalance.DiscardUnknown(m)
}

var xxx_messageInfo_In_CreateUserBalance proto.InternalMessageInfo

func (m *In_CreateUserBalance) GetUserBalanceLog() *UserBalanceLog {
	if m != nil {
		return m.UserBalanceLog
	}
	return nil
}

type Out_CreateUserBalance struct {
	Error                *Error          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	UserBalanceLog       *UserBalanceLog `protobuf:"bytes,2,opt,name=user_balance_log,json=userBalanceLog,proto3" json:"user_balance_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Out_CreateUserBalance) Reset()         { *m = Out_CreateUserBalance{} }
func (m *Out_CreateUserBalance) String() string { return proto.CompactTextString(m) }
func (*Out_CreateUserBalance) ProtoMessage()    {}
func (*Out_CreateUserBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5426dd8d173b8f1, []int{11}
}

func (m *Out_CreateUserBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out_CreateUserBalance.Unmarshal(m, b)
}
func (m *Out_CreateUserBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out_CreateUserBalance.Marshal(b, m, deterministic)
}
func (m *Out_CreateUserBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out_CreateUserBalance.Merge(m, src)
}
func (m *Out_CreateUserBalance) XXX_Size() int {
	return xxx_messageInfo_Out_CreateUserBalance.Size(m)
}
func (m *Out_CreateUserBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_Out_CreateUserBalance.DiscardUnknown(m)
}

var xxx_messageInfo_Out_CreateUserBalance proto.InternalMessageInfo

func (m *Out_CreateUserBalance) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Out_CreateUserBalance) GetUserBalanceLog() *UserBalanceLog {
	if m != nil {
		return m.UserBalanceLog
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "user_balance_log.Error")
	proto.RegisterType((*UserBalanceLog)(nil), "user_balance_log.User_balance_log")
	proto.RegisterType((*In_GetUserBalanceInfoById)(nil), "user_balance_log.In_GetUserBalanceInfoById")
	proto.RegisterType((*Out_GetUserBalanceInfoById)(nil), "user_balance_log.Out_GetUserBalanceInfoById")
	proto.RegisterType((*Out_UpdateUserBalanceInfo)(nil), "user_balance_log.Out_UpdateUserBalanceInfo")
	proto.RegisterType((*In_UpdateUserBalanceInfo)(nil), "user_balance_log.In_UpdateUserBalanceInfo")
	proto.RegisterType((*In_GetUserBalances)(nil), "user_balance_log.In_GetUserBalances")
	proto.RegisterType((*Out_GetUserBalances)(nil), "user_balance_log.Out_GetUserBalances")
	proto.RegisterType((*In_DeleteUserBalances)(nil), "user_balance_log.In_DeleteUserBalances")
	proto.RegisterType((*Out_DeleteUserBalances)(nil), "user_balance_log.Out_DeleteUserBalances")
	proto.RegisterType((*In_CreateUserBalance)(nil), "user_balance_log.In_CreateUserBalance")
	proto.RegisterType((*Out_CreateUserBalance)(nil), "user_balance_log.Out_CreateUserBalance")
}

func init() { proto.RegisterFile("user_balance_log.proto", fileDescriptor_e5426dd8d173b8f1) }

var fileDescriptor_e5426dd8d173b8f1 = []byte{
	// 632 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xce, 0xd6, 0x71, 0xd2, 0x4c, 0xa5, 0xbf, 0xed, 0xfe, 0x4d, 0xea, 0x06, 0x21, 0xa2, 0x15,
	0xd0, 0xa8, 0x85, 0x1e, 0x8a, 0xb8, 0x70, 0x6c, 0x41, 0x95, 0x45, 0x05, 0x92, 0x4b, 0xcf, 0x66,
	0x1b, 0x4f, 0x53, 0x8b, 0x78, 0x5d, 0x79, 0x37, 0x48, 0x7d, 0x0f, 0x0e, 0xf4, 0x3d, 0x78, 0x15,
	0x8e, 0xbc, 0x0b, 0xda, 0xb5, 0x0b, 0xa9, 0x77, 0x23, 0x52, 0x84, 0xd4, 0x9b, 0x67, 0x76, 0x76,
	0xbe, 0x6f, 0x66, 0xbe, 0x1d, 0x19, 0x7a, 0x53, 0x89, 0x45, 0x7c, 0xc6, 0x27, 0x5c, 0x8c, 0x30,
	0x9e, 0xe4, 0xe3, 0xbd, 0xcb, 0x22, 0x57, 0x39, 0x5d, 0xab, 0xfb, 0xd9, 0x4b, 0xf0, 0xdf, 0x14,
	0x45, 0x5e, 0x50, 0x0a, 0xcd, 0x51, 0x9e, 0x60, 0x40, 0x06, 0x64, 0xe8, 0x47, 0xe6, 0x9b, 0x06,
	0xd0, 0xce, 0x50, 0x4a, 0x3e, 0xc6, 0x60, 0x69, 0x40, 0x86, 0x9d, 0xe8, 0xc6, 0x64, 0xdf, 0x09,
	0xac, 0x9d, 0xd6, 0x72, 0xd1, 0x87, 0x00, 0x37, 0x66, 0x9a, 0x98, 0x44, 0x5e, 0xd4, 0xa9, 0x3c,
	0x61, 0x42, 0x37, 0xa1, 0x6d, 0xe0, 0xd3, 0xc4, 0x64, 0xf3, 0xa2, 0x96, 0x36, 0xc3, 0x84, 0xf6,
	0xa0, 0x25, 0xf3, 0x69, 0x31, 0xc2, 0xc0, 0x2b, 0xfd, 0xa5, 0x45, 0x1f, 0x40, 0xa7, 0xfc, 0x8a,
	0xa5, 0x08, 0x9a, 0xe6, 0x68, 0xb9, 0x74, 0x9c, 0x08, 0xfa, 0x08, 0x56, 0x46, 0x05, 0x72, 0x85,
	0xb1, 0x4a, 0x33, 0x0c, 0x7c, 0xc3, 0x0f, 0x4a, 0xd7, 0x87, 0x34, 0x43, 0x9d, 0x95, 0x67, 0xf9,
	0x54, 0xa8, 0xa0, 0x35, 0x20, 0x43, 0x12, 0x55, 0x96, 0xce, 0x6a, 0x68, 0x08, 0x9e, 0x61, 0xd0,
	0x36, 0xd7, 0x96, 0xb5, 0xe3, 0x1d, 0xcf, 0x90, 0xbd, 0x82, 0xad, 0x50, 0xc4, 0x47, 0xa8, 0x74,
	0x71, 0x07, 0x15, 0x75, 0x71, 0x9e, 0x1f, 0x5c, 0x85, 0xc9, 0x1f, 0xea, 0x63, 0xd7, 0x04, 0xfa,
	0xef, 0xa7, 0x6a, 0xde, 0xed, 0xe7, 0xe0, 0xa3, 0xee, 0xb4, 0xb9, 0xb8, 0xb2, 0xbf, 0xb9, 0x67,
	0xcd, 0xc8, 0x0c, 0x22, 0x2a, 0xa3, 0xe8, 0x31, 0x58, 0xc3, 0x32, 0x6d, 0x5b, 0xd9, 0x67, 0xf6,
	0xcd, 0xfa, 0x28, 0xa2, 0xff, 0xa6, 0xbf, 0x19, 0x1c, 0xe7, 0x63, 0xf6, 0x95, 0xc0, 0x96, 0xe6,
	0x76, 0x7a, 0x99, 0x70, 0x85, 0x35, 0x7a, 0xf7, 0x4b, 0xed, 0x02, 0x82, 0x50, 0xcc, 0x21, 0xe6,
	0x42, 0x22, 0x7f, 0x8d, 0x74, 0x4d, 0x80, 0x5a, 0xd3, 0x95, 0x74, 0x03, 0xfc, 0x49, 0x9a, 0xa5,
	0xaa, 0x9a, 0x68, 0x69, 0x68, 0xef, 0x25, 0x1f, 0xa3, 0xac, 0xb4, 0x5a, 0x1a, 0x5a, 0x02, 0x12,
	0x79, 0x31, 0xba, 0x88, 0x3f, 0xe1, 0x95, 0x91, 0x6b, 0x27, 0xea, 0x94, 0x9e, 0xb7, 0x78, 0x65,
	0x8e, 0x15, 0x2f, 0x54, 0xa9, 0xc9, 0x66, 0x75, 0xac, 0x3d, 0x46, 0x92, 0x5b, 0xb0, 0x8c, 0x22,
	0x99, 0x15, 0x6c, 0x1b, 0x45, 0xa2, 0x8f, 0xd8, 0x0f, 0x02, 0xff, 0xdb, 0xe2, 0x91, 0x77, 0x1d,
	0xcd, 0xaf, 0x5a, 0x96, 0x9c, 0xb5, 0x78, 0xb3, 0xb5, 0x6c, 0x80, 0xaf, 0x72, 0xc5, 0x27, 0xd5,
	0xd3, 0x2a, 0x0d, 0x7a, 0x0a, 0xdd, 0x3a, 0x44, 0x3c, 0x49, 0xa5, 0x0a, 0xfc, 0x81, 0xb7, 0x60,
	0xdf, 0xe9, 0xed, 0xbe, 0x1f, 0xa7, 0x52, 0xb1, 0x43, 0xe8, 0x86, 0x22, 0x7e, 0x8d, 0x13, 0xbc,
	0x35, 0x65, 0x49, 0x77, 0x60, 0xbd, 0x9e, 0x51, 0x06, 0x64, 0xe0, 0x0d, 0xbd, 0x68, 0xf5, 0x76,
	0x1e, 0xc9, 0x8e, 0xa0, 0xa7, 0x7b, 0xe4, 0xc8, 0x72, 0xb7, 0x36, 0xb1, 0x04, 0x36, 0x42, 0x11,
	0x1f, 0x9a, 0x65, 0x31, 0x93, 0xe7, 0x1f, 0xeb, 0xed, 0x0b, 0x81, 0xae, 0xe6, 0x6b, 0xe3, 0xdc,
	0xe7, 0x83, 0xdb, 0xff, 0xd6, 0x04, 0x3a, 0x43, 0xe6, 0x04, 0x8b, 0xcf, 0xe9, 0x08, 0xe9, 0x14,
	0x7a, 0x73, 0x36, 0xd7, 0xae, 0x0d, 0x32, 0x77, 0x49, 0xf6, 0x9f, 0xd9, 0xc1, 0xf3, 0x97, 0x22,
	0x6b, 0xd0, 0x02, 0xba, 0xee, 0xb7, 0xbf, 0xe3, 0x44, 0x75, 0xc6, 0xf6, 0x77, 0xdd, 0xa0, 0xce,
	0x60, 0xd6, 0xa0, 0x1f, 0x61, 0xb5, 0xfe, 0xce, 0x1e, 0x2f, 0x50, 0xa3, 0xec, 0x3f, 0x59, 0xa4,
	0x38, 0xc9, 0x1a, 0x34, 0x05, 0xea, 0x50, 0xe9, 0xb6, 0x13, 0xc4, 0x0e, 0xec, 0x0f, 0xdd, 0x38,
	0x76, 0x24, 0x6b, 0xd0, 0x73, 0x58, 0xb7, 0x05, 0xf6, 0xd4, 0x89, 0x64, 0xc5, 0xf5, 0xb7, 0xdd,
	0x40, 0x56, 0x20, 0x6b, 0x9c, 0xb5, 0xcc, 0x2f, 0xc4, 0x8b, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x3e, 0x72, 0x1f, 0x5c, 0x08, 0x00, 0x00,
}
