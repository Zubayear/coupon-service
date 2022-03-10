package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Coupon holds the schema definition for the Coupon entity.
type Coupon struct {
	ent.Schema
}

// Fields of the Coupon.
func (Coupon) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("Code").MaxLen(6).Unique(),
		field.Float("Amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(6,2)",
		}),
		field.Bool("AlreadyUsed").Default(false),
		field.Int64("ExpireAt").Default(time.Now().AddDate(0, 0, 5).Unix()),
	}
}

// Edges of the Coupon.
func (Coupon) Edges() []ent.Edge {
	return nil
}
