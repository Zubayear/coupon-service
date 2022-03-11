package repository

import (
	"Coupon/ent"
	"Coupon/ent/coupon"
	"Coupon/ent/migrate"
	"Coupon/external"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	log "go-micro.dev/v4/logger"
)

type ICouponRepository interface {
	GetCouponByCode(ctx context.Context, couponCode string) (*ent.Coupon, error)
	UseCoupon(ctx context.Context, couponId uuid.UUID) (string, error)
	GetCouponById(ctx context.Context, couponId uuid.UUID) (*ent.Coupon, error)
	CreateCoupon(ctx context.Context, coupon *ent.Coupon) (*ent.Coupon, error)
}

type CouponRepository struct {
	couponClient *ent.Client
}

func (c *CouponRepository) GetCouponByCode(ctx context.Context, couponCode string) (*ent.Coupon, error) {
	couponFromRepo, err := c.couponClient.Coupon.Query().Where(coupon.Code(couponCode)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting coupon by code: %w", err)
	}
	return couponFromRepo, nil
}

func (c *CouponRepository) UseCoupon(ctx context.Context, couponId uuid.UUID) (string, error) {
	err := c.couponClient.Coupon.UpdateOneID(couponId).SetAlreadyUsed(true).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("failed updating coupon: %w", err)
	}
	return "coupon used", nil
}

func (c *CouponRepository) GetCouponById(ctx context.Context, couponId uuid.UUID) (*ent.Coupon, error) {
	couponFromRepo, err := c.couponClient.Coupon.Get(ctx, couponId)
	if err != nil {
		return nil, fmt.Errorf("failed getting coupon by id: %w", err)
	}
	return couponFromRepo, nil
}

func (c *CouponRepository) CreateCoupon(ctx context.Context, coupon *ent.Coupon) (*ent.Coupon, error) {
	savedCoupon, err := c.couponClient.Coupon.Create().
		SetCode(coupon.Code).
		SetAmount(coupon.Amount).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed saving coupon: %w", err)
	}
	return savedCoupon, nil
}

func CouponRepositoryProvider(h *external.Host) (*CouponRepository, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", h.User, h.Password, h.Address, h.Port, h.Name)
	client, err := ent.Open(h.Type, connString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &CouponRepository{couponClient: client}, nil
}

func NewCouponRepository() ICouponRepository {
	return &CouponRepository{}
}
