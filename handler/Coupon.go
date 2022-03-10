package handler

import (
	"Coupon/ent"
	pb "Coupon/proto"
	"Coupon/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	log "go-micro.dev/v4/logger"
	"strconv"
	"time"
)

type Coupon struct {
	couponRepository repository.ICouponRepository
}

func (e *Coupon) CreateCoupon(ctx context.Context, request *pb.CreateCouponRequest, response *pb.CreateCouponResponse) error {
	log.Infof("Received Coupon.CreateCoupon request: %v", request)
	couponToSave := &ent.Coupon{
		Code:   request.Code,
		Amount: request.Amount,
	}
	couponFromRepo, err := e.couponRepository.CreateCoupon(ctx, couponToSave)
	if err != nil {
		return err
	}
	couponToReturn := &pb.Coupon{}
	err = mapCoupon(couponFromRepo, couponToReturn)
	if err != nil {
		return err
	}
	response.Coupon = couponToReturn
	return nil
}

func (e *Coupon) GetCouponByCode(ctx context.Context, request *pb.GetCouponByCodeRequest, response *pb.GetCouponByCodeResponse) error {
	log.Infof("Received Coupon.GetCouponByCode request: %v", request)
	couponFromRepo, err := e.couponRepository.GetCouponByCode(ctx, request.CouponCode)
	if err != nil {
		return err
	}
	couponToReturn := &pb.Coupon{}
	err = mapCoupon(couponFromRepo, couponToReturn)
	if err != nil {
		return err
	}
	response.Coupon = couponToReturn
	return nil
}

func (e *Coupon) GetCouponById(ctx context.Context, request *pb.GetCouponByIdRequest, response *pb.GetCouponByIdResponse) error {
	couponId, err := uuid.Parse(request.CouponId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	couponFromRepo, err := e.couponRepository.GetCouponById(ctx, couponId)
	couponToReturn := &pb.Coupon{}
	err = mapCoupon(couponFromRepo, couponToReturn)
	if err != nil {
		return err
	}
	response.Coupon = couponToReturn
	return nil
}

func (e *Coupon) UseCoupon(ctx context.Context, request *pb.UseCouponRequest, response *pb.UseCouponResponse) error {
	couponId, err := uuid.Parse(request.CouponId)
	if err != nil {
		return fmt.Errorf("failed parsing id: %w", err)
	}
	couponStatus, err := e.couponRepository.UseCoupon(ctx, couponId)
	if err != nil {
		return err
	}
	response.Msg = couponStatus
	return nil
}

func NewCoupon() pb.CouponServiceHandler {
	return &Coupon{}
}

func NewCouponProvider(couponRepository repository.ICouponRepository) (*Coupon, error) {
	return &Coupon{couponRepository: couponRepository}, nil
}

func convertToTime(expireAt int64) (time.Time, error) {
	i, err := strconv.ParseInt(strconv.FormatInt(expireAt, 10), 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed parsing expire_at: %w", err)
	}
	tm := time.Unix(i, 0)
	return tm, nil
}

func mapCoupon(src *ent.Coupon, dst *pb.Coupon) error {
	expireAt, err := convertToTime(src.ExpireAt)
	if err != nil {
		return err
	}
	dst.Code = src.Code
	dst.CouponId = src.ID.String()
	dst.Amount = src.Amount
	dst.AlreadyUsed = src.AlreadyUsed
	dst.ExpireAt = expireAt.String()
	return nil
}
