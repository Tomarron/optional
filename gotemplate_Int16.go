// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Int16 = time.Time{}

// template type Optional(T,scan)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int16 optionalInt16

type optionalInt16 []int16

const (
	valueKeyInt16 = iota
)

func scanValueInt16(input string) (val int16, err error) {
	v, err := ScanInt(input)
	return int16(v), err
}

// Of wraps the value in an optional.
func OfInt16(value int16) Int16 {
	return Int16{valueKeyInt16: value}
}

func OfInt16Ptr(ptr *int16) Int16 {
	if ptr == nil {
		return EmptyInt16()
	} else {
		return OfInt16(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyInt16() Int16 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Int16) Get() (value int16, ok bool) {
	o.If(func(v int16) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Int16) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Int16) If(f func(value int16)) {
	if o.IsPresent() {
		f(o[valueKeyInt16])
	}
}

func (o Int16) ElseFunc(f func() int16) (value int16) {
	if o.IsPresent() {
		o.If(func(v int16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Int16) Else(elseValue int16) (value int16) {
	return o.ElseFunc(func() int16 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Int16) V() (value int16) {
	var zero int16
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Int16) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int16) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyInt16])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Int16) UnmarshalJSON(data []byte) error {
	var v int16
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfInt16(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Int16) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Int16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int16
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt16(v)
	return nil
}

func (c Int16) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Int16) Scan(input interface{}) (err error) {
	var vv string
	var isvalid = true
	switch value := input.(type) {
	case string:
		if len(value) > 0 {
			vv = value
		} else {
			isvalid = false
		}
	case []byte:
		if value != nil && len(value) > 0 {
			vv = string(value)
		} else {
			isvalid = false
		}
	default:
		isvalid = false
	}
	if isvalid {
		val, err := scanValueInt16(vv)
		if err != nil {
			return err
		}
		*c = OfInt16(val)
	}
	return
}
