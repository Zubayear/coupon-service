// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/Coupon.proto

package Coupon

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CouponService service

func NewCouponServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CouponService service

type CouponService interface {
	GetCouponByCode(ctx context.Context, in *GetCouponByCodeRequest, opts ...client.CallOption) (*GetCouponByCodeResponse, error)
	GetCouponById(ctx context.Context, in *GetCouponByIdRequest, opts ...client.CallOption) (*GetCouponByIdResponse, error)
	UseCoupon(ctx context.Context, in *UseCouponRequest, opts ...client.CallOption) (*UseCouponResponse, error)
	CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...client.CallOption) (*CreateCouponResponse, error)
}

type couponService struct {
	c    client.Client
	name string
}

func NewCouponService(name string, c client.Client) CouponService {
	return &couponService{
		c:    c,
		name: name,
	}
}

func (c *couponService) GetCouponByCode(ctx context.Context, in *GetCouponByCodeRequest, opts ...client.CallOption) (*GetCouponByCodeResponse, error) {
	req := c.c.NewRequest(c.name, "CouponService.GetCouponByCode", in)
	out := new(GetCouponByCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponService) GetCouponById(ctx context.Context, in *GetCouponByIdRequest, opts ...client.CallOption) (*GetCouponByIdResponse, error) {
	req := c.c.NewRequest(c.name, "CouponService.GetCouponById", in)
	out := new(GetCouponByIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponService) UseCoupon(ctx context.Context, in *UseCouponRequest, opts ...client.CallOption) (*UseCouponResponse, error) {
	req := c.c.NewRequest(c.name, "CouponService.UseCoupon", in)
	out := new(UseCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *couponService) CreateCoupon(ctx context.Context, in *CreateCouponRequest, opts ...client.CallOption) (*CreateCouponResponse, error) {
	req := c.c.NewRequest(c.name, "CouponService.CreateCoupon", in)
	out := new(CreateCouponResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CouponService service

type CouponServiceHandler interface {
	GetCouponByCode(context.Context, *GetCouponByCodeRequest, *GetCouponByCodeResponse) error
	GetCouponById(context.Context, *GetCouponByIdRequest, *GetCouponByIdResponse) error
	UseCoupon(context.Context, *UseCouponRequest, *UseCouponResponse) error
	CreateCoupon(context.Context, *CreateCouponRequest, *CreateCouponResponse) error
}

func RegisterCouponServiceHandler(s server.Server, hdlr CouponServiceHandler, opts ...server.HandlerOption) error {
	type couponService interface {
		GetCouponByCode(ctx context.Context, in *GetCouponByCodeRequest, out *GetCouponByCodeResponse) error
		GetCouponById(ctx context.Context, in *GetCouponByIdRequest, out *GetCouponByIdResponse) error
		UseCoupon(ctx context.Context, in *UseCouponRequest, out *UseCouponResponse) error
		CreateCoupon(ctx context.Context, in *CreateCouponRequest, out *CreateCouponResponse) error
	}
	type CouponService struct {
		couponService
	}
	h := &couponServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CouponService{h}, opts...))
}

type couponServiceHandler struct {
	CouponServiceHandler
}

func (h *couponServiceHandler) GetCouponByCode(ctx context.Context, in *GetCouponByCodeRequest, out *GetCouponByCodeResponse) error {
	return h.CouponServiceHandler.GetCouponByCode(ctx, in, out)
}

func (h *couponServiceHandler) GetCouponById(ctx context.Context, in *GetCouponByIdRequest, out *GetCouponByIdResponse) error {
	return h.CouponServiceHandler.GetCouponById(ctx, in, out)
}

func (h *couponServiceHandler) UseCoupon(ctx context.Context, in *UseCouponRequest, out *UseCouponResponse) error {
	return h.CouponServiceHandler.UseCoupon(ctx, in, out)
}

func (h *couponServiceHandler) CreateCoupon(ctx context.Context, in *CreateCouponRequest, out *CreateCouponResponse) error {
	return h.CouponServiceHandler.CreateCoupon(ctx, in, out)
}
