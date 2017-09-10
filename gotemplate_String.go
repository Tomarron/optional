// Code generated by gotemplate. DO NOT EDIT.

package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _String = time.Time{}

// template type Optional(T,scan)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String optionalString

type optionalString []string

const (
	valueKeyString = iota
)

func scanValueString(input string) (val string, err error) {
	v, err := ScanString(input)
	return string(v), err
}

// Of wraps the value in an optional.
func OfString(value string) String {
	return String{valueKeyString: value}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// Empty returns an empty optional.
func EmptyString() String {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o String) Get() (value string, ok bool) {
	o.If(func(v string) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o String) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o String) If(f func(value string)) {
	if o.IsPresent() {
		f(o[valueKeyString])
	}
}

func (o String) ElseFunc(f func() string) (value string) {
	if o.IsPresent() {
		o.If(func(v string) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o String) Else(elseValue string) (value string) {
	return o.ElseFunc(func() string { return elseValue })
}

// V returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o String) V() (value string) {
	var zero string
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o String) String() string {
	if o.IsPresent() {
		return fmt.Sprintf("%v", o.V())
	}
	return fmt.Sprintf("%v", nil)
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o String) MarshalJSON() (data []byte, err error) {
	if o.IsPresent() {
		return json.Marshal(o[valueKeyString])
	}
	return []byte("null"), nil
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *String) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	//Try unmarshal string numbers with quote
	if err != nil && len(data) > 2 {
		cpy := data[1 : len(data)-2]
		err = json.Unmarshal(cpy, &v)
	}
	if err != nil {
		return err
	}

	*o = OfString(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.V(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfString(v)
	return nil
}

func (c String) Value() (driver.Value, error) {
	v, ok := c.Get()
	if ok {
		return driver.DefaultParameterConverter.ConvertValue(v)
	}
	return driver.DefaultParameterConverter.ConvertValue(nil)
}

func (c *String) Scan(input interface{}) (err error) {
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
		val, err := scanValueString(vv)
		if err != nil {
			return err
		}
		*c = OfString(val)
	}
	return
}
