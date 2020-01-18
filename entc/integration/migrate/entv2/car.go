// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/car"
)

// Car is the model entity for the Car schema.
type Car struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges struct {
		// Owner holds the value of the owner edge.
		Owner *User
	} `json:"edges,omitempty"`
	owner_id *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Car) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // owner_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(values ...interface{}) error {
	if m, n := len(values), len(car.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	values = values[0:]
	if len(values) == len(car.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field owner_id", value)
		} else if value.Valid {
			c.owner_id = new(int)
			*c.owner_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Car.
func (c *Car) QueryOwner() *UserQuery {
	return (&CarClient{c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Car.
// Note that, you need to call Car.Unwrap() before calling this method, if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return (&CarClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Car is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

func (c Cars) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
