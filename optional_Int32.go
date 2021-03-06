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

var _Int32 = time.Time{}
var __Int32 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type int32
type Int32 optionalInt32

type optionalInt32 []int32

const (
	valueKeyInt32 = iota
)

func scanValueInt32(input string) (val int32, err error) {
	v, err := optional_scanner.ScanInt(input)
	return int32(v), err
}

// Of wraps the value in an optional.
func OfInt32(value int32) Int32 {
	return Int32{valueKeyInt32: value}
}

func OfInt32Ptr(ptr *int32) Int32 {
	if ptr == nil {
		return EmptyInt32()
	} else {
		return OfInt32(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyInt32() Int32 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Int32) Get() (value int32, ok bool) {
	o.If(func(v int32) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Int32) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Int32) If(f func(value int32)) {
	if o.IsPresent() {
		f(o[valueKeyInt32])
	}
}

func (o Int32) ElseFunc(f func() int32) (value int32) {
	if o.IsPresent() {
		o.If(func(v int32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Int32) Else(elseValue int32) (value int32) {
	return o.ElseFunc(func() int32 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Int32) V() (value int32) {
	var zero int32
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Int32) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int32) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyInt32])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Int32) UnmarshalJSON(data []byte) error {
	var v int32
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfInt32(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Int32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt32(v)
	return nil
}

func (c Int32) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Int32) Scan(input interface{}) (err error) {
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
		val, err := scanValueInt32(vv)
		if err != nil {
			return err
		}
		*c = OfInt32(val)
	}
	return
}
