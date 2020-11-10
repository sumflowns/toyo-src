// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user_info.proto

package user_info

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserInfo service

type UserInfoService interface {
	//用户登陆
	DoneUserLogin(ctx context.Context, in *InDoneUserLogin, opts ...client.CallOption) (*OutDoneUserLogin, error)
	//获取用户详细信息
	GetUserInfo(ctx context.Context, in *InGetUserInfo, opts ...client.CallOption) (*OutGetUserInfo, error)
	//获取用户列表
	GetUserInfoList(ctx context.Context, in *InGetUserInfoList, opts ...client.CallOption) (*OutGetUserInfoList, error)
	//修改用户信息
	UpdateUserInfo(ctx context.Context, in *InUpdateUserInfo, opts ...client.CallOption) (*OutUpdateUserInfo, error)
	//用户注册
	DoneUserRegister(ctx context.Context, in *InDoneUserRegister, opts ...client.CallOption) (*OutDoneUserRegister, error)
	//获取验证码
	GetVerificationCode(ctx context.Context, in *InGetVerificationCode, opts ...client.CallOption) (*OutGetVerificationCode, error)
}

type userInfoService struct {
	c    client.Client
	name string
}

func NewUserInfoService(name string, c client.Client) UserInfoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user_info"
	}
	return &userInfoService{
		c:    c,
		name: name,
	}
}

func (c *userInfoService) DoneUserLogin(ctx context.Context, in *InDoneUserLogin, opts ...client.CallOption) (*OutDoneUserLogin, error) {
	req := c.c.NewRequest(c.name, "UserInfo.DoneUserLogin", in)
	out := new(OutDoneUserLogin)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoService) GetUserInfo(ctx context.Context, in *InGetUserInfo, opts ...client.CallOption) (*OutGetUserInfo, error) {
	req := c.c.NewRequest(c.name, "UserInfo.GetUserInfo", in)
	out := new(OutGetUserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoService) GetUserInfoList(ctx context.Context, in *InGetUserInfoList, opts ...client.CallOption) (*OutGetUserInfoList, error) {
	req := c.c.NewRequest(c.name, "UserInfo.GetUserInfoList", in)
	out := new(OutGetUserInfoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoService) UpdateUserInfo(ctx context.Context, in *InUpdateUserInfo, opts ...client.CallOption) (*OutUpdateUserInfo, error) {
	req := c.c.NewRequest(c.name, "UserInfo.UpdateUserInfo", in)
	out := new(OutUpdateUserInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoService) DoneUserRegister(ctx context.Context, in *InDoneUserRegister, opts ...client.CallOption) (*OutDoneUserRegister, error) {
	req := c.c.NewRequest(c.name, "UserInfo.DoneUserRegister", in)
	out := new(OutDoneUserRegister)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoService) GetVerificationCode(ctx context.Context, in *InGetVerificationCode, opts ...client.CallOption) (*OutGetVerificationCode, error) {
	req := c.c.NewRequest(c.name, "UserInfo.GetVerificationCode", in)
	out := new(OutGetVerificationCode)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInfo service

type UserInfoHandler interface {
	//用户登陆
	DoneUserLogin(context.Context, *InDoneUserLogin, *OutDoneUserLogin) error
	//获取用户详细信息
	GetUserInfo(context.Context, *InGetUserInfo, *OutGetUserInfo) error
	//获取用户列表
	GetUserInfoList(context.Context, *InGetUserInfoList, *OutGetUserInfoList) error
	//修改用户信息
	UpdateUserInfo(context.Context, *InUpdateUserInfo, *OutUpdateUserInfo) error
	//用户注册
	DoneUserRegister(context.Context, *InDoneUserRegister, *OutDoneUserRegister) error
	//获取验证码
	GetVerificationCode(context.Context, *InGetVerificationCode, *OutGetVerificationCode) error
}

func RegisterUserInfoHandler(s server.Server, hdlr UserInfoHandler, opts ...server.HandlerOption) error {
	type userInfo interface {
		DoneUserLogin(ctx context.Context, in *InDoneUserLogin, out *OutDoneUserLogin) error
		GetUserInfo(ctx context.Context, in *InGetUserInfo, out *OutGetUserInfo) error
		GetUserInfoList(ctx context.Context, in *InGetUserInfoList, out *OutGetUserInfoList) error
		UpdateUserInfo(ctx context.Context, in *InUpdateUserInfo, out *OutUpdateUserInfo) error
		DoneUserRegister(ctx context.Context, in *InDoneUserRegister, out *OutDoneUserRegister) error
		GetVerificationCode(ctx context.Context, in *InGetVerificationCode, out *OutGetVerificationCode) error
	}
	type UserInfo struct {
		userInfo
	}
	h := &userInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&UserInfo{h}, opts...))
}

type userInfoHandler struct {
	UserInfoHandler
}

func (h *userInfoHandler) DoneUserLogin(ctx context.Context, in *InDoneUserLogin, out *OutDoneUserLogin) error {
	return h.UserInfoHandler.DoneUserLogin(ctx, in, out)
}

func (h *userInfoHandler) GetUserInfo(ctx context.Context, in *InGetUserInfo, out *OutGetUserInfo) error {
	return h.UserInfoHandler.GetUserInfo(ctx, in, out)
}

func (h *userInfoHandler) GetUserInfoList(ctx context.Context, in *InGetUserInfoList, out *OutGetUserInfoList) error {
	return h.UserInfoHandler.GetUserInfoList(ctx, in, out)
}

func (h *userInfoHandler) UpdateUserInfo(ctx context.Context, in *InUpdateUserInfo, out *OutUpdateUserInfo) error {
	return h.UserInfoHandler.UpdateUserInfo(ctx, in, out)
}

func (h *userInfoHandler) DoneUserRegister(ctx context.Context, in *InDoneUserRegister, out *OutDoneUserRegister) error {
	return h.UserInfoHandler.DoneUserRegister(ctx, in, out)
}

func (h *userInfoHandler) GetVerificationCode(ctx context.Context, in *InGetVerificationCode, out *OutGetVerificationCode) error {
	return h.UserInfoHandler.GetVerificationCode(ctx, in, out)
}
