//go:build wireinject
// +build wireinject

package di

import (
	"Coupon/external"
	"Coupon/handler"
	"Coupon/repository"
	"github.com/google/wire"
)

func DependencyProvider() (*handler.Coupon, error) {
	wire.Build(
		external.HostProvider, repository.CouponRepositoryProvider, handler.NewCouponProvider,
		wire.Bind(new(repository.ICouponRepository), new(*repository.CouponRepository)))
	return &handler.Coupon{}, nil
}
