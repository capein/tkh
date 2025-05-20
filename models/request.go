package models

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"
	"tkh/coupons"
	"tkh/logger"
)

type IReqModel interface {
	Load(ctx context.Context, reader io.ReadCloser) error
	Validate(ctx context.Context) error
}

type OrderReq struct {
	CouponCode string `json:"couponCode" validate:"greater=7,lesser=11"`
	Items      []Item `json:"items" validate:"greater=0"`
}

func (o *OrderReq) Load(ctx context.Context, r io.ReadCloser) error {
	err := json.NewDecoder(r).Decode(o)
	defer r.Close()
	return err
}

func (o *OrderReq) Validate(ctx context.Context) error {
	if o.CouponCode != "" {
		if bt.Has(o.CouponCode) {
			return errors.New("coupon code already used")
		}
		ok, err := coupons.FindCoupon(ctx, o.CouponCode)
		if err != nil {
			logger.Println("error while validating the coupon", err)
			return err
		}
		if !ok {
			return errors.New("coupon code invalid")
		}
		bt.ReplaceOrInsert(o.CouponCode)
	}
	var temp int64
	p := Products{}
	err := p.GetData(ctx)
	if err != nil {
		logger.Println("error while getting the product data", err)
		return errors.New("error getting the product data")
	}
	for _, item := range o.Items {
		temp, err = strconv.ParseInt(item.ProductId, 10, 64)
		if err != nil {
			logger.Println("error while parsing the product id", err)
			return errors.New("invalid product specified")
		}
		if temp < 1 || temp > int64(len(p.Products)) {
			logger.Println("error while getting the product data", err)
			return errors.New("invalid product specified")
		}
	}
	return nil
}

type Item struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,greater=0"`
}
