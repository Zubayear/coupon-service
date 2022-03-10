// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"Coupon/ent/migrate"

	"Coupon/ent/coupon"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Coupon is the client for interacting with the Coupon builders.
	Coupon *CouponClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Coupon = NewCouponClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Coupon: NewCouponClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Coupon: NewCouponClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Coupon.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Coupon.Use(hooks...)
}

// CouponClient is a client for the Coupon schema.
type CouponClient struct {
	config
}

// NewCouponClient returns a client for the Coupon from the given config.
func NewCouponClient(c config) *CouponClient {
	return &CouponClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coupon.Hooks(f(g(h())))`.
func (c *CouponClient) Use(hooks ...Hook) {
	c.hooks.Coupon = append(c.hooks.Coupon, hooks...)
}

// Create returns a create builder for Coupon.
func (c *CouponClient) Create() *CouponCreate {
	mutation := newCouponMutation(c.config, OpCreate)
	return &CouponCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Coupon entities.
func (c *CouponClient) CreateBulk(builders ...*CouponCreate) *CouponCreateBulk {
	return &CouponCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Coupon.
func (c *CouponClient) Update() *CouponUpdate {
	mutation := newCouponMutation(c.config, OpUpdate)
	return &CouponUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponClient) UpdateOne(co *Coupon) *CouponUpdateOne {
	mutation := newCouponMutation(c.config, OpUpdateOne, withCoupon(co))
	return &CouponUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponClient) UpdateOneID(id uuid.UUID) *CouponUpdateOne {
	mutation := newCouponMutation(c.config, OpUpdateOne, withCouponID(id))
	return &CouponUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Coupon.
func (c *CouponClient) Delete() *CouponDelete {
	mutation := newCouponMutation(c.config, OpDelete)
	return &CouponDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CouponClient) DeleteOne(co *Coupon) *CouponDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CouponClient) DeleteOneID(id uuid.UUID) *CouponDeleteOne {
	builder := c.Delete().Where(coupon.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponDeleteOne{builder}
}

// Query returns a query builder for Coupon.
func (c *CouponClient) Query() *CouponQuery {
	return &CouponQuery{
		config: c.config,
	}
}

// Get returns a Coupon entity by its id.
func (c *CouponClient) Get(ctx context.Context, id uuid.UUID) (*Coupon, error) {
	return c.Query().Where(coupon.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponClient) GetX(ctx context.Context, id uuid.UUID) *Coupon {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponClient) Hooks() []Hook {
	return c.hooks.Coupon
}
