// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Coupon/ent/coupon"
	"Coupon/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCoupon = "Coupon"
)

// CouponMutation represents an operation that mutates the Coupon nodes in the graph.
type CouponMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	_Code         *string
	_Amount       *float64
	add_Amount    *float64
	_AlreadyUsed  *bool
	_ExpireAt     *int64
	add_ExpireAt  *int64
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Coupon, error)
	predicates    []predicate.Coupon
}

var _ ent.Mutation = (*CouponMutation)(nil)

// couponOption allows management of the mutation configuration using functional options.
type couponOption func(*CouponMutation)

// newCouponMutation creates new mutation for the Coupon entity.
func newCouponMutation(c config, op Op, opts ...couponOption) *CouponMutation {
	m := &CouponMutation{
		config:        c,
		op:            op,
		typ:           TypeCoupon,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCouponID sets the ID field of the mutation.
func withCouponID(id uuid.UUID) couponOption {
	return func(m *CouponMutation) {
		var (
			err   error
			once  sync.Once
			value *Coupon
		)
		m.oldValue = func(ctx context.Context) (*Coupon, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Coupon.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCoupon sets the old Coupon of the mutation.
func withCoupon(node *Coupon) couponOption {
	return func(m *CouponMutation) {
		m.oldValue = func(context.Context) (*Coupon, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CouponMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CouponMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Coupon entities.
func (m *CouponMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *CouponMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *CouponMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Coupon.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCode sets the "Code" field.
func (m *CouponMutation) SetCode(s string) {
	m._Code = &s
}

// Code returns the value of the "Code" field in the mutation.
func (m *CouponMutation) Code() (r string, exists bool) {
	v := m._Code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old "Code" field's value of the Coupon entity.
// If the Coupon object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CouponMutation) OldCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode resets all changes to the "Code" field.
func (m *CouponMutation) ResetCode() {
	m._Code = nil
}

// SetAmount sets the "Amount" field.
func (m *CouponMutation) SetAmount(f float64) {
	m._Amount = &f
	m.add_Amount = nil
}

// Amount returns the value of the "Amount" field in the mutation.
func (m *CouponMutation) Amount() (r float64, exists bool) {
	v := m._Amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "Amount" field's value of the Coupon entity.
// If the Coupon object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CouponMutation) OldAmount(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds f to the "Amount" field.
func (m *CouponMutation) AddAmount(f float64) {
	if m.add_Amount != nil {
		*m.add_Amount += f
	} else {
		m.add_Amount = &f
	}
}

// AddedAmount returns the value that was added to the "Amount" field in this mutation.
func (m *CouponMutation) AddedAmount() (r float64, exists bool) {
	v := m.add_Amount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "Amount" field.
func (m *CouponMutation) ResetAmount() {
	m._Amount = nil
	m.add_Amount = nil
}

// SetAlreadyUsed sets the "AlreadyUsed" field.
func (m *CouponMutation) SetAlreadyUsed(b bool) {
	m._AlreadyUsed = &b
}

// AlreadyUsed returns the value of the "AlreadyUsed" field in the mutation.
func (m *CouponMutation) AlreadyUsed() (r bool, exists bool) {
	v := m._AlreadyUsed
	if v == nil {
		return
	}
	return *v, true
}

// OldAlreadyUsed returns the old "AlreadyUsed" field's value of the Coupon entity.
// If the Coupon object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CouponMutation) OldAlreadyUsed(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAlreadyUsed is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAlreadyUsed requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAlreadyUsed: %w", err)
	}
	return oldValue.AlreadyUsed, nil
}

// ResetAlreadyUsed resets all changes to the "AlreadyUsed" field.
func (m *CouponMutation) ResetAlreadyUsed() {
	m._AlreadyUsed = nil
}

// SetExpireAt sets the "ExpireAt" field.
func (m *CouponMutation) SetExpireAt(i int64) {
	m._ExpireAt = &i
	m.add_ExpireAt = nil
}

// ExpireAt returns the value of the "ExpireAt" field in the mutation.
func (m *CouponMutation) ExpireAt() (r int64, exists bool) {
	v := m._ExpireAt
	if v == nil {
		return
	}
	return *v, true
}

// OldExpireAt returns the old "ExpireAt" field's value of the Coupon entity.
// If the Coupon object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *CouponMutation) OldExpireAt(ctx context.Context) (v int64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldExpireAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldExpireAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldExpireAt: %w", err)
	}
	return oldValue.ExpireAt, nil
}

// AddExpireAt adds i to the "ExpireAt" field.
func (m *CouponMutation) AddExpireAt(i int64) {
	if m.add_ExpireAt != nil {
		*m.add_ExpireAt += i
	} else {
		m.add_ExpireAt = &i
	}
}

// AddedExpireAt returns the value that was added to the "ExpireAt" field in this mutation.
func (m *CouponMutation) AddedExpireAt() (r int64, exists bool) {
	v := m.add_ExpireAt
	if v == nil {
		return
	}
	return *v, true
}

// ResetExpireAt resets all changes to the "ExpireAt" field.
func (m *CouponMutation) ResetExpireAt() {
	m._ExpireAt = nil
	m.add_ExpireAt = nil
}

// Where appends a list predicates to the CouponMutation builder.
func (m *CouponMutation) Where(ps ...predicate.Coupon) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *CouponMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Coupon).
func (m *CouponMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *CouponMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m._Code != nil {
		fields = append(fields, coupon.FieldCode)
	}
	if m._Amount != nil {
		fields = append(fields, coupon.FieldAmount)
	}
	if m._AlreadyUsed != nil {
		fields = append(fields, coupon.FieldAlreadyUsed)
	}
	if m._ExpireAt != nil {
		fields = append(fields, coupon.FieldExpireAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *CouponMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case coupon.FieldCode:
		return m.Code()
	case coupon.FieldAmount:
		return m.Amount()
	case coupon.FieldAlreadyUsed:
		return m.AlreadyUsed()
	case coupon.FieldExpireAt:
		return m.ExpireAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *CouponMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case coupon.FieldCode:
		return m.OldCode(ctx)
	case coupon.FieldAmount:
		return m.OldAmount(ctx)
	case coupon.FieldAlreadyUsed:
		return m.OldAlreadyUsed(ctx)
	case coupon.FieldExpireAt:
		return m.OldExpireAt(ctx)
	}
	return nil, fmt.Errorf("unknown Coupon field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CouponMutation) SetField(name string, value ent.Value) error {
	switch name {
	case coupon.FieldCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	case coupon.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	case coupon.FieldAlreadyUsed:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAlreadyUsed(v)
		return nil
	case coupon.FieldExpireAt:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetExpireAt(v)
		return nil
	}
	return fmt.Errorf("unknown Coupon field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *CouponMutation) AddedFields() []string {
	var fields []string
	if m.add_Amount != nil {
		fields = append(fields, coupon.FieldAmount)
	}
	if m.add_ExpireAt != nil {
		fields = append(fields, coupon.FieldExpireAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *CouponMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case coupon.FieldAmount:
		return m.AddedAmount()
	case coupon.FieldExpireAt:
		return m.AddedExpireAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *CouponMutation) AddField(name string, value ent.Value) error {
	switch name {
	case coupon.FieldAmount:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	case coupon.FieldExpireAt:
		v, ok := value.(int64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddExpireAt(v)
		return nil
	}
	return fmt.Errorf("unknown Coupon numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *CouponMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *CouponMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *CouponMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Coupon nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *CouponMutation) ResetField(name string) error {
	switch name {
	case coupon.FieldCode:
		m.ResetCode()
		return nil
	case coupon.FieldAmount:
		m.ResetAmount()
		return nil
	case coupon.FieldAlreadyUsed:
		m.ResetAlreadyUsed()
		return nil
	case coupon.FieldExpireAt:
		m.ResetExpireAt()
		return nil
	}
	return fmt.Errorf("unknown Coupon field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *CouponMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *CouponMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *CouponMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *CouponMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *CouponMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *CouponMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *CouponMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Coupon unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *CouponMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Coupon edge %s", name)
}
