package nilsql

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NilInt64 is an alias for sql.NullInt64 data type
type NilInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NilInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NilInt64{i.Int64, false}
	} else {
		*ni = NilInt64{i.Int64, true}
	}
	return nil
}

// MarshalJSON is to convert NullInt64 into JSON data
func (ni *NilInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

// NilBool is an alias for sql.NullBool data type
type NilBool sql.NullBool

// Scan implements the Scanner interface for NullBool
func (nb *NilBool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = NilBool{b.Bool, false}
	} else {
		*nb = NilBool{b.Bool, true}
	}

	return nil
}

// MarshalJSON is to convert NullBool into JSON data
func (nb *NilBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

// NilFloat64 is an alias for sql.NullFloat64 data type
type NilFloat64 sql.NullFloat64

// Scan implements the Scanner interface for NullFloat64
func (nf *NilFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = NilFloat64{f.Float64, false}
	} else {
		*nf = NilFloat64{f.Float64, true}
	}

	return nil
}

// MarshalJSON is to convert NullFloat64 into JSON data
func (nf *NilFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float64)
}

// NilString is an alias for sql.NullString data type
type NilString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NilString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NilString{s.String, false}
	} else {
		*ns = NilString{s.String, true}
	}

	return nil
}

// MarshalJSON is to convert NullString into JSON data
func (ns *NilString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.String)
}
