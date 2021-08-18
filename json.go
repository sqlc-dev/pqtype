package hoare

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// NullRawMessage represents a json.RawMessage that may be null.
// NullRawMessage implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullRawMessage struct {
	RawMessage json.RawMessage
	Valid      bool // Valid is true if RawMessage is not NULL
}

// Scan implements the Scanner interface.
func (n *NullRawMessage) Scan(value interface{}) error {
	if value == nil {
		n.Valid = false
		return nil
	}
	buf, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type %T", value, []byte{})
	}
	n.RawMessage, n.Valid = buf, true
	return nil
}

// Value implements the driver Valuer interface.
func (n NullRawMessage) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return []byte(n.RawMessage), nil
}
