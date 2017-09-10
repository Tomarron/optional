// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uint16 = time.Time{}

// template type Optional(T,scan)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint16 optionalUint16

type optionalUint16 []uint16

const (
	valueKeyUint16 = iota
)

func scanValueUint16(input string) (val uint16, err error) {
	v, err := ScanUint(input)
	return uint16(v), err
}

// Of wraps the value in an optional.
func OfUint16(value uint16) Uint16 {
	return Uint16{valueKeyUint16: value}
}

func OfUint16Ptr(ptr *uint16) Uint16 {
	if ptr == nil {
		return EmptyUint16()
	} else {
		return OfUint16(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint16() Uint16 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint16) Get() (value uint16, ok bool) {
	o.If(func(v uint16) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Uint16) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint16) If(f func(value uint16)) {
	if o.IsPresent() {
		f(o[valueKeyUint16])
	}
}

func (o Uint16) ElseFunc(f func() uint16) (value uint16) {
	if o.IsPresent() {
		o.If(func(v uint16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint16) Else(elseValue uint16) (value uint16) {
	return o.ElseFunc(func() uint16 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint16) V() (value uint16) {
	var zero uint16
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint16) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint16) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyUint16])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint16) UnmarshalJSON(data []byte) error {
	var v uint16
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfUint16(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint16) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint16
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint16(v)
	return nil
}

func (c Uint16) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint16) Scan(input interface{}) (err error) {
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
		val, err := scanValueUint16(vv)
		if err != nil {
			return err
		}
		*c = OfUint16(val)
	}
	return
}
