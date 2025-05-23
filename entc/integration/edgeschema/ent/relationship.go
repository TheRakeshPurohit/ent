// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// Relationship is the model entity for the Relationship schema.
type Relationship struct {
	config `json:"-"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// RelativeID holds the value of the "relative_id" field.
	RelativeID int `json:"relative_id,omitempty"`
	// InfoID holds the value of the "info_id" field.
	InfoID int `json:"info_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RelationshipQuery when eager-loading is set.
	Edges        RelationshipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RelationshipEdges holds the relations/edges for other nodes in the graph.
type RelationshipEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Relative holds the value of the relative edge.
	Relative *User `json:"relative,omitempty"`
	// Info holds the value of the info edge.
	Info *RelationshipInfo `json:"info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RelativeOrErr returns the Relative value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) RelativeOrErr() (*User, error) {
	if e.Relative != nil {
		return e.Relative, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "relative"}
}

// InfoOrErr returns the Info value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RelationshipEdges) InfoOrErr() (*RelationshipInfo, error) {
	if e.Info != nil {
		return e.Info, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: relationshipinfo.Label}
	}
	return nil, &NotLoadedError{edge: "info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Relationship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case relationship.FieldWeight, relationship.FieldUserID, relationship.FieldRelativeID, relationship.FieldInfoID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Relationship fields.
func (_m *Relationship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case relationship.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				_m.Weight = int(value.Int64)
			}
		case relationship.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				_m.UserID = int(value.Int64)
			}
		case relationship.FieldRelativeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relative_id", values[i])
			} else if value.Valid {
				_m.RelativeID = int(value.Int64)
			}
		case relationship.FieldInfoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field info_id", values[i])
			} else if value.Valid {
				_m.InfoID = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Relationship.
// This includes values selected through modifiers, order, etc.
func (_m *Relationship) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Relationship entity.
func (_m *Relationship) QueryUser() *UserQuery {
	return NewRelationshipClient(_m.config).QueryUser(_m)
}

// QueryRelative queries the "relative" edge of the Relationship entity.
func (_m *Relationship) QueryRelative() *UserQuery {
	return NewRelationshipClient(_m.config).QueryRelative(_m)
}

// QueryInfo queries the "info" edge of the Relationship entity.
func (_m *Relationship) QueryInfo() *RelationshipInfoQuery {
	return NewRelationshipClient(_m.config).QueryInfo(_m)
}

// Update returns a builder for updating this Relationship.
// Note that you need to call Relationship.Unwrap() before calling this method if this Relationship
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Relationship) Update() *RelationshipUpdateOne {
	return NewRelationshipClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Relationship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Relationship) Unwrap() *Relationship {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Relationship is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Relationship) String() string {
	var builder strings.Builder
	builder.WriteString("Relationship(")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", _m.Weight))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.UserID))
	builder.WriteString(", ")
	builder.WriteString("relative_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.RelativeID))
	builder.WriteString(", ")
	builder.WriteString("info_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.InfoID))
	builder.WriteByte(')')
	return builder.String()
}

// Relationships is a parsable slice of Relationship.
type Relationships []*Relationship
