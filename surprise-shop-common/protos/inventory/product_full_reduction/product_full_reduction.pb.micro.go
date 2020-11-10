// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product_full_reduction.proto

package product_full_reduction

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

// Client API for ProductFullReductionHandler service

type ProductFullReductionHandlerService interface {
	//获取信息
	GetProductFullReductionById(ctx context.Context, in *In_GetProductFullReductionById, opts ...client.CallOption) (*Out_GetProductFullReductionById, error)
	//修改信息
	UpdateProductFullReductionInfo(ctx context.Context, in *In_UpdateProductFullReductionInfo, opts ...client.CallOption) (*Out_UpdateProductFullReductionInfo, error)
	//获取列表
	GetProductFullReductions(ctx context.Context, in *In_GetProductFullReductions, opts ...client.CallOption) (*Out_GetProductFullReductions, error)
	//删除列表
	DeleteProductFullReductions(ctx context.Context, in *In_DeleteProductFullReductions, opts ...client.CallOption) (*Out_DeleteProductFullReductions, error)
	//新建信息
	CreateProductFullReduction(ctx context.Context, in *In_CreateProductFullReduction, opts ...client.CallOption) (*Out_CreateProductFullReduction, error)
}

type productFullReductionHandlerService struct {
	c    client.Client
	name string
}

func NewProductFullReductionHandlerService(name string, c client.Client) ProductFullReductionHandlerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "product_full_reduction"
	}
	return &productFullReductionHandlerService{
		c:    c,
		name: name,
	}
}

func (c *productFullReductionHandlerService) GetProductFullReductionById(ctx context.Context, in *In_GetProductFullReductionById, opts ...client.CallOption) (*Out_GetProductFullReductionById, error) {
	req := c.c.NewRequest(c.name, "ProductFullReductionHandler.GetProductFullReductionById", in)
	out := new(Out_GetProductFullReductionById)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFullReductionHandlerService) UpdateProductFullReductionInfo(ctx context.Context, in *In_UpdateProductFullReductionInfo, opts ...client.CallOption) (*Out_UpdateProductFullReductionInfo, error) {
	req := c.c.NewRequest(c.name, "ProductFullReductionHandler.UpdateProductFullReductionInfo", in)
	out := new(Out_UpdateProductFullReductionInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFullReductionHandlerService) GetProductFullReductions(ctx context.Context, in *In_GetProductFullReductions, opts ...client.CallOption) (*Out_GetProductFullReductions, error) {
	req := c.c.NewRequest(c.name, "ProductFullReductionHandler.GetProductFullReductions", in)
	out := new(Out_GetProductFullReductions)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFullReductionHandlerService) DeleteProductFullReductions(ctx context.Context, in *In_DeleteProductFullReductions, opts ...client.CallOption) (*Out_DeleteProductFullReductions, error) {
	req := c.c.NewRequest(c.name, "ProductFullReductionHandler.DeleteProductFullReductions", in)
	out := new(Out_DeleteProductFullReductions)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productFullReductionHandlerService) CreateProductFullReduction(ctx context.Context, in *In_CreateProductFullReduction, opts ...client.CallOption) (*Out_CreateProductFullReduction, error) {
	req := c.c.NewRequest(c.name, "ProductFullReductionHandler.CreateProductFullReduction", in)
	out := new(Out_CreateProductFullReduction)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductFullReductionHandler service

type ProductFullReductionHandlerHandler interface {
	//获取信息
	GetProductFullReductionById(context.Context, *In_GetProductFullReductionById, *Out_GetProductFullReductionById) error
	//修改信息
	UpdateProductFullReductionInfo(context.Context, *In_UpdateProductFullReductionInfo, *Out_UpdateProductFullReductionInfo) error
	//获取列表
	GetProductFullReductions(context.Context, *In_GetProductFullReductions, *Out_GetProductFullReductions) error
	//删除列表
	DeleteProductFullReductions(context.Context, *In_DeleteProductFullReductions, *Out_DeleteProductFullReductions) error
	//新建信息
	CreateProductFullReduction(context.Context, *In_CreateProductFullReduction, *Out_CreateProductFullReduction) error
}

func RegisterProductFullReductionHandlerHandler(s server.Server, hdlr ProductFullReductionHandlerHandler, opts ...server.HandlerOption) error {
	type productFullReductionHandler interface {
		GetProductFullReductionById(ctx context.Context, in *In_GetProductFullReductionById, out *Out_GetProductFullReductionById) error
		UpdateProductFullReductionInfo(ctx context.Context, in *In_UpdateProductFullReductionInfo, out *Out_UpdateProductFullReductionInfo) error
		GetProductFullReductions(ctx context.Context, in *In_GetProductFullReductions, out *Out_GetProductFullReductions) error
		DeleteProductFullReductions(ctx context.Context, in *In_DeleteProductFullReductions, out *Out_DeleteProductFullReductions) error
		CreateProductFullReduction(ctx context.Context, in *In_CreateProductFullReduction, out *Out_CreateProductFullReduction) error
	}
	type ProductFullReductionHandler struct {
		productFullReductionHandler
	}
	h := &productFullReductionHandlerHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductFullReductionHandler{h}, opts...))
}

type productFullReductionHandlerHandler struct {
	ProductFullReductionHandlerHandler
}

func (h *productFullReductionHandlerHandler) GetProductFullReductionById(ctx context.Context, in *In_GetProductFullReductionById, out *Out_GetProductFullReductionById) error {
	return h.ProductFullReductionHandlerHandler.GetProductFullReductionById(ctx, in, out)
}

func (h *productFullReductionHandlerHandler) UpdateProductFullReductionInfo(ctx context.Context, in *In_UpdateProductFullReductionInfo, out *Out_UpdateProductFullReductionInfo) error {
	return h.ProductFullReductionHandlerHandler.UpdateProductFullReductionInfo(ctx, in, out)
}

func (h *productFullReductionHandlerHandler) GetProductFullReductions(ctx context.Context, in *In_GetProductFullReductions, out *Out_GetProductFullReductions) error {
	return h.ProductFullReductionHandlerHandler.GetProductFullReductions(ctx, in, out)
}

func (h *productFullReductionHandlerHandler) DeleteProductFullReductions(ctx context.Context, in *In_DeleteProductFullReductions, out *Out_DeleteProductFullReductions) error {
	return h.ProductFullReductionHandlerHandler.DeleteProductFullReductions(ctx, in, out)
}

func (h *productFullReductionHandlerHandler) CreateProductFullReduction(ctx context.Context, in *In_CreateProductFullReduction, out *Out_CreateProductFullReduction) error {
	return h.ProductFullReductionHandlerHandler.CreateProductFullReduction(ctx, in, out)
}
