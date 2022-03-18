package proxy

import (
	"Coupon/handler"
	pb "Coupon/proto"
	"context"
)

type Proxy struct {
	coupon *handler.Coupon
}

func ProxyProvider(c *handler.Coupon) (*Proxy, error) {
	return &Proxy{coupon: c}, nil
}

func (p *Proxy) GetCouponByCode(ctx context.Context, request *pb.GetCouponByCodeRequest) (*pb.GetCouponByCodeResponse, error) {
	response := &pb.GetCouponByCodeResponse{}
	err := p.coupon.GetCouponByCode(ctx, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Proxy) GetCouponById(ctx context.Context, request *pb.GetCouponByIdRequest) (*pb.GetCouponByIdResponse, error) {
	response := &pb.GetCouponByIdResponse{}
	err := p.coupon.GetCouponById(ctx, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Proxy) UseCoupon(ctx context.Context, request *pb.UseCouponRequest) (*pb.UseCouponResponse, error) {
	response := &pb.UseCouponResponse{}
	err := p.coupon.UseCoupon(ctx, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *Proxy) CreateCoupon(ctx context.Context, request *pb.CreateCouponRequest) (*pb.CreateCouponResponse, error) {
	response := &pb.CreateCouponResponse{}
	err := p.coupon.CreateCoupon(ctx, request, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewProxy() pb.CouponServiceServer {
	return &Proxy{}
}
