// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CouponsColumns holds the columns for the "coupons" table.
	CouponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "code", Type: field.TypeString, Unique: true, Size: 6},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(6,2)"}},
		{Name: "already_used", Type: field.TypeBool, Default: false},
		{Name: "expire_at", Type: field.TypeInt64, Default: 1647376068},
	}
	// CouponsTable holds the schema information for the "coupons" table.
	CouponsTable = &schema.Table{
		Name:       "coupons",
		Columns:    CouponsColumns,
		PrimaryKey: []*schema.Column{CouponsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CouponsTable,
	}
)

func init() {
}
