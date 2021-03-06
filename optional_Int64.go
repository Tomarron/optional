// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/imiskolee/optional/optional_scanner"
)

var _Int64 = time.Time{}
var __Int64 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type int64
type Int64 optionalInt64

type optionalInt64 []int64

const (
	valueKeyInt64 = iota
)

func scanValueInt64(input string) (val int64, err error) {
	v, err := optional_scanner.ScanInt(input)
	return int64(v), err
}

// Of wraps the value in an optional.
func OfInt64(value int64) Int64 {
	return Int64{valueKeyInt64: value}
}

func OfInt64Ptr(ptr *int64) Int64 {
	if ptr == nil {
		return EmptyInt64()
	} else {
		return OfInt64(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyInt64() Int64 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Int64) Get() (value int64, ok bool) {
	o.If(func(v int64) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Int64) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Int64) If(f func(value int64)) {
	if o.IsPresent() {
		f(o[valueKeyInt64])
	}
}

func (o Int64) ElseFunc(f func() int64) (value int64) {
	if o.IsPresent() {
		o.If(func(v int64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Int64) Else(elseValue int64) (value int64) {
	return o.ElseFunc(func() int64 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Int64) V() (value int64) {
	var zero int64
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Int64) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int64) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyInt64])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Int64) UnmarshalJSON(data []byte) error {
	var v int64
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfInt64(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Int64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt64(v)
	return nil
}

func (c Int64) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Int64) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true

	if input == nil {
		isvalid = false
	}

	if isvalid {
		switch value := input.(type) {
		case string:
			vv = value
		case []byte:
			if value != nil {
				vv = string(value)
			} else {
				isvalid = false
			}
		default:
			vv = fmt.Sprint(input)
		}
	}

	if isvalid {
		val, err := scanValueInt64(vv)
		if err != nil {
			return err
		}
		*c = OfInt64(val)
	}
	return
}
