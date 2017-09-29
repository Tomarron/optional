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

var _Uint32 = time.Time{}
var __Uint32 = optional_scanner.ScanBool

// template type Optional(T,scan)

//swagger:type uint32
type Uint32 optionalUint32

type optionalUint32 []uint32

const (
	valueKeyUint32 = iota
)

func scanValueUint32(input string) (val uint32, err error) {
	v, err := optional_scanner.ScanUint(input)
	return uint32(v), err
}

// Of wraps the value in an optional.
func OfUint32(value uint32) Uint32 {
	return Uint32{valueKeyUint32: value}
}

func OfUint32Ptr(ptr *uint32) Uint32 {
	if ptr == nil {
		return EmptyUint32()
	} else {
		return OfUint32(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyUint32() Uint32 {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Uint32) Get() (value uint32, ok bool) {
	o.If(func(v uint32) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Uint32) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Uint32) If(f func(value uint32)) {
	if o.IsPresent() {
		f(o[valueKeyUint32])
	}
}

func (o Uint32) ElseFunc(f func() uint32) (value uint32) {
	if o.IsPresent() {
		o.If(func(v uint32) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Uint32) Else(elseValue uint32) (value uint32) {
	return o.ElseFunc(func() uint32 { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Uint32) V() (value uint32) {
	var zero uint32
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Uint32) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint32) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyUint32])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Uint32) UnmarshalJSON(data []byte) error {
	var v uint32
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfUint32(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Uint32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Uint32) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint32
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint32(v)
	return nil
}

func (c Uint32) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *Uint32) Scan(input interface{}) (err error) {
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
		val, err := scanValueUint32(vv)
		if err != nil {
			return err
		}
		*c = OfUint32(val)
	}
	return
}
