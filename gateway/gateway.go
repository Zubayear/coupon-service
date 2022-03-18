package main

import (
	Coupon "Coupon/proto"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-micro.dev/v4/client"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	gateway = flag.Int("gateway", 8080, "gateway")
	port    = flag.Int("port", 60009, "coupon address")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	Coupon.RegisterCouponServiceServer(s, &Proxy{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		return err
	}

	go func() {
		err := s.Serve(listen)
		if err != nil {
			return
		}
	}()

	mux := runtime.NewServeMux()
	conn, err := grpc.Dial("localhost:60009", grpc.WithInsecure())
	if err != nil {
		return err
	}

	err = Coupon.RegisterCouponServiceUT(ctx, mux, conn)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *gateway),
		Handler: mux,
	}
	return gwServer.ListenAndServe()
}

func main() {
	flag.Parse()

	defer glog.Flush()
	err := run()
	if err != nil {
		return
	}
}

type Proxy struct {
}

func (p *Proxy) GetCouponByCode(ctx context.Context, request *Coupon.GetCouponByCodeRequest) (*Coupon.GetCouponByCodeResponse, error) {
	couponClient := Coupon.NewCouponService("coupon", client.DefaultClient)
	couponToReturn, err := couponClient.GetCouponByCode(ctx, request)
	if err != nil {
		return nil, err
	}
	return couponToReturn, nil
}

func (p *Proxy) GetCouponById(ctx context.Context, request *Coupon.GetCouponByIdRequest) (*Coupon.GetCouponByIdResponse, error) {
	couponClient := Coupon.NewCouponService("coupon", client.DefaultClient)
	couponToReturn, err := couponClient.GetCouponById(ctx, request)
	if err != nil {
		return nil, err
	}
	return couponToReturn, nil
}

func (p *Proxy) UseCoupon(ctx context.Context, request *Coupon.UseCouponRequest) (*Coupon.UseCouponResponse, error) {
	couponClient := Coupon.NewCouponService("coupon", client.DefaultClient)
	couponToReturn, err := couponClient.UseCoupon(ctx, request)
	if err != nil {
		return nil, err
	}
	return couponToReturn, nil
}

func (p *Proxy) CreateCoupon(ctx context.Context, request *Coupon.CreateCouponRequest) (*Coupon.CreateCouponResponse, error) {
	couponClient := Coupon.NewCouponService("coupon", client.DefaultClient)
	couponToReturn, err := couponClient.CreateCoupon(ctx, request)
	if err != nil {
		return nil, err
	}
	return couponToReturn, nil
}
