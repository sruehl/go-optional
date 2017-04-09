package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Byte map[keyByte]byte

type keyByte int

const (
	valueKeyByte keyByte = iota
)

// Of wraps the value in an Optional.
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

// Empty returns an empty Optional.
func EmptyByte() Byte {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Byte) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Byte) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
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

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Byte) Else(elseValue byte) (value byte) {
	return o.ElseFunc(func() byte { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o Byte) MarshalText() (text []byte, err error) {
	o.If(func(v byte) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *Byte) UnmarshalText(text []byte) error {
	*o = EmptyByte()
	return nil
}
