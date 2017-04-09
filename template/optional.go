package template

import (
	"fmt"
)

// template type Optional(T)

type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional map[key]T

type key int

const (
	valueKey key = iota
)

// Of wraps the value in an Optional.
func Of(value T) Optional {
	return Optional{valueKey: value}
}

func OfOptionalPtr(ptr *T) Optional {
	if ptr == nil {
		return Empty()
	} else {
		return Of(*ptr)
	}
}

// Empty returns an empty Optional.
func Empty() Optional {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Optional) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Optional) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Optional) If(f func(value T)) {
	if o.IsPresent() {
		f(o[valueKey])
	}
}

func (o Optional) ElseFunc(f func() T) (value T) {
	if o.IsPresent() {
		o.If(func(v T) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Optional) Else(elseValue T) (value T) {
	return o.ElseFunc(func() T { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Optional) String() string {
	if o.IsPresent() {
		var value T
		o.If(func(v T) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
