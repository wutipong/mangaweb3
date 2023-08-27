// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/wutipong/mangaweb3-backend/ent/meta"
)

// Meta is the model entity for the Meta schema.
type Meta struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Favorite holds the value of the "favorite" field.
	Favorite bool `json:"favorite,omitempty"`
	// FileIndices holds the value of the "file_indices" field.
	FileIndices []int `json:"file_indices,omitempty"`
	// Thumbnail holds the value of the "thumbnail" field.
	Thumbnail []byte `json:"thumbnail,omitempty"`
	// Read holds the value of the "read" field.
	Read bool `json:"read,omitempty"`
	// Tags holds the value of the "tags" field.
	Tags         []string `json:"tags,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Meta) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case meta.FieldFileIndices, meta.FieldThumbnail, meta.FieldTags:
			values[i] = new([]byte)
		case meta.FieldFavorite, meta.FieldRead:
			values[i] = new(sql.NullBool)
		case meta.FieldID:
			values[i] = new(sql.NullInt64)
		case meta.FieldName:
			values[i] = new(sql.NullString)
		case meta.FieldCreateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Meta fields.
func (m *Meta) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case meta.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case meta.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case meta.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				m.CreateTime = value.Time
			}
		case meta.FieldFavorite:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field favorite", values[i])
			} else if value.Valid {
				m.Favorite = value.Bool
			}
		case meta.FieldFileIndices:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field file_indices", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.FileIndices); err != nil {
					return fmt.Errorf("unmarshal field file_indices: %w", err)
				}
			}
		case meta.FieldThumbnail:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail", values[i])
			} else if value != nil {
				m.Thumbnail = *value
			}
		case meta.FieldRead:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field read", values[i])
			} else if value.Valid {
				m.Read = value.Bool
			}
		case meta.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Meta.
// This includes values selected through modifiers, order, etc.
func (m *Meta) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this Meta.
// Note that you need to call Meta.Unwrap() before calling this method if this Meta
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Meta) Update() *MetaUpdateOne {
	return NewMetaClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Meta entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Meta) Unwrap() *Meta {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Meta is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Meta) String() string {
	var builder strings.Builder
	builder.WriteString("Meta(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(m.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("favorite=")
	builder.WriteString(fmt.Sprintf("%v", m.Favorite))
	builder.WriteString(", ")
	builder.WriteString("file_indices=")
	builder.WriteString(fmt.Sprintf("%v", m.FileIndices))
	builder.WriteString(", ")
	builder.WriteString("thumbnail=")
	builder.WriteString(fmt.Sprintf("%v", m.Thumbnail))
	builder.WriteString(", ")
	builder.WriteString("read=")
	builder.WriteString(fmt.Sprintf("%v", m.Read))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", m.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// MetaSlice is a parsable slice of Meta.
type MetaSlice []*Meta
