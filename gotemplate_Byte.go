// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Byte = time.Time{}

// template type Optional(T,scan)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Byte optionalByte

type optionalByte []byte

const (
	valueKeyByte = iota
)

func scanValueByte(input string) (val byte, err error) {
	v, err := ScanByte(input)
	return byte(v), err
}

// Of wraps the value in an optional.
func OfByte(value byte) Byte {
	return Byte{valueKeyByte: value}
}

func OfBytePtr(ptr *byte) Byte {
	if ptr == nil {
		return EmptyByte()
	} else {
		return OfByte(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyByte() Byte {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Byte) Get() (value byte, ok bool) {
	o.If(func(v byte) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Byte) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Byte) If(f func(value byte)) {
	if o.IsPresent() {
		f(o[valueKeyByte])
	}
}

func (o Byte) ElseFunc(f func() byte) (value byte) {
	if o.IsPresent() {
		o.If(func(v byte) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Byte) Else(elseValue byte) (value byte) {
	return o.ElseFunc(func() byte { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Byte) V() (value byte) {
	var zero byte
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Byte) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Byte) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyByte])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Byte) UnmarshalJSON(data []byte) error {
	var v byte
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfByte(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Byte) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Byte) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v byte
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfByte(v)
	return nil
}

func (c Byte) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Byte) Scan(input interface{}) (err error) {
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
		val, err := scanValueByte(vv)
		if err != nil {
			return err
		}
		*c = OfByte(val)
	}
	return
}
