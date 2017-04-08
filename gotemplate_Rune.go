package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Rune struct {
	value *rune
}

// For wraps the value in an Optional.
func ForRune(value rune) Rune {
	return Rune{&value}
}

func ForRunePtr(ptr *rune) Rune {
	return Rune{ptr}
}

// Empty returns an empty Optional.
func EmptyRune() Rune {
	return Rune{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Rune) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Rune) IfPresent(f func(value rune)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Rune) OrElse(value rune) rune {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
