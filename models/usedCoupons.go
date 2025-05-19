package models

import "github.com/google/btree"

var bt = btree.NewOrderedG[string](2)

type CouponCode struct {
	Code string
}
