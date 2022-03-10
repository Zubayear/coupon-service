// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Coupon/ent/coupon"
	"Coupon/ent/schema"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	couponFields := schema.Coupon{}.Fields()
	_ = couponFields
	// couponDescCode is the schema descriptor for Code field.
	couponDescCode := couponFields[1].Descriptor()
	// coupon.CodeValidator is a validator for the "Code" field. It is called by the builders before save.
	coupon.CodeValidator = couponDescCode.Validators[0].(func(string) error)
	// couponDescAlreadyUsed is the schema descriptor for AlreadyUsed field.
	couponDescAlreadyUsed := couponFields[3].Descriptor()
	// coupon.DefaultAlreadyUsed holds the default value on creation for the AlreadyUsed field.
	coupon.DefaultAlreadyUsed = couponDescAlreadyUsed.Default.(bool)
	// couponDescExpireAt is the schema descriptor for ExpireAt field.
	couponDescExpireAt := couponFields[4].Descriptor()
	// coupon.DefaultExpireAt holds the default value on creation for the ExpireAt field.
	coupon.DefaultExpireAt = couponDescExpireAt.Default.(int64)
	// couponDescID is the schema descriptor for id field.
	couponDescID := couponFields[0].Descriptor()
	// coupon.DefaultID holds the default value on creation for the id field.
	coupon.DefaultID = couponDescID.Default.(func() uuid.UUID)
}