package main

import (
	"Coupon/di"
	pb "Coupon/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "coupon"
	version = "latest"
	port    = ":60009"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(port),
	)
	srv.Init()

	// wire up all the dependency
	coupon, err := di.DependencyProvider()
	if err != nil {
		return
	}

	// Register handler
	err = pb.RegisterCouponServiceHandler(srv.Server(), coupon)
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
