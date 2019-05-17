package nilsql

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// Int64 is an alias for sql.NullInt64 data type
type Int64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *Int64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = Int64{i.Int64, false}
	} else {
		*ni = Int64{i.Int64, true}
	}
	return nil
}

// MarshalJSON is to convert NullInt64 into JSON data
func (ni *Int64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ni.Int64)
}

// Bool is an alias for sql.NullBool data type
type Bool sql.NullBool

// Scan implements the Scanner interface for NullBool
func (nb *Bool) Scan(value interface{}) error {
	var b sql.NullBool
	if err := b.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nb = Bool{b.Bool, false}
	} else {
		*nb = Bool{b.Bool, true}
	}

	return nil
}

// MarshalJSON is to convert NullBool into JSON data
func (nb *Bool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nb.Bool)
}

// Float64 is an alias for sql.NullFloat64 data type
type Float64 sql.NullFloat64

// Scan implements the Scanner interface for NullFloat64
func (nf *Float64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = Float64{f.Float64, false}
	} else {
		*nf = Float64{f.Float64, true}
	}

	return nil
}

// MarshalJSON is to convert NullFloat64 into JSON data
func (nf *Float64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(nf.Float64)
}

// String is an alias for sql.NullString data type
type String sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *String) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = String{s.String, false}
	} else {
		*ns = String{s.String, true}
	}

	return nil
}

// MarshalJSON is to convert NullString into JSON data
func (ns *String) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(ns.String)
}
