package optional

import (
	"encoding/json"
	"testing"
	"time"
)

func TestIsPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional          Optional[string]
		ExpectedIsPresent bool
	}{
		{Empty[string](), false},
		{Of(""), true},
		{Of("string"), true},
		{OfPtr((*string)(nil)), false},
		{OfPtr((*string)(&s)), true},
	}

	for _, test := range tests {
		isPresent := test.Optional.IsPresent()

		if isPresent != test.ExpectedIsPresent {
			t.Errorf("%#v IsPresent got %#v, want %#v", test.Optional, isPresent, test.ExpectedIsPresent)
		}
	}
}

func TestGet(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional      Optional[string]
		ExpectedValue string
		ExpectedOk    bool
	}{
		{Empty[string](), "", false},
		{Of(""), "", true},
		{Of("string"), "string", true},
		{OfPtr((*string)(nil)), "", false},
		{OfPtr((*string)(&s)), "ptr to string", true},
	}

	for _, test := range tests {
		value, ok := test.Optional.Get()

		if value != test.ExpectedValue || ok != test.ExpectedOk {
			t.Errorf("%#v Get got %#v, %#v, want %#v, %#v", test.Optional, ok, test.ExpectedOk, value, test.ExpectedValue)
		}
	}
}

func TestIfPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional       Optional[string]
		ExpectedCalled bool
		IfCalledValue  string
	}{
		{Empty[string](), false, ""},
		{Of(""), true, ""},
		{Of("string"), true, "string"},
		{OfPtr((*string)(nil)), false, ""},
		{OfPtr((*string)(&s)), true, "ptr to string"},
	}

	for _, test := range tests {
		called := false
		test.Optional.If(func(v string) {
			called = true
			if v != test.IfCalledValue {
				t.Errorf("%#v IfPresent got %#v, want #%v", test.Optional, v, test.IfCalledValue)
			}
		})

		if called != test.ExpectedCalled {
			t.Errorf("%#v IfPresent called %#v, want %#v", test.Optional, called, test.ExpectedCalled)
		}
	}
}

func TestElse(t *testing.T) {
	s := "ptr to string"
	const orElse = "orelse"
	tests := []struct {
		Optional       Optional[string]
		ExpectedResult string
	}{
		{Empty[string](), orElse},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfPtr((*string)(nil)), orElse},
		{OfPtr((*string)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.Else(orElse)

		if result != test.ExpectedResult {
			t.Errorf("%#v OrElse(%#v) got %#v, want %#v", test.Optional, orElse, result, test.ExpectedResult)
		}
	}
}

func TestElseFunc(t *testing.T) {
	s := "ptr to string"
	const orElse = "orelse"
	tests := []struct {
		Optional       Optional[string]
		ExpectedResult string
	}{
		{Empty[string](), orElse},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfPtr((*string)(nil)), orElse},
		{OfPtr((*string)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.ElseFunc(func() string { return orElse })

		if result != test.ExpectedResult {
			t.Errorf("%#v OrElse(%#v) got %#v, want %#v", test.Optional, orElse, result, test.ExpectedResult)
		}
	}
}

func TestElseZero(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional       Optional[string]
		ExpectedResult string
	}{
		{Empty[string](), ""},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfPtr((*string)(nil)), ""},
		{OfPtr((*string)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.ElseZero()

		if result != test.ExpectedResult {
			t.Errorf("%#v ElseZero() got %#v, want %#v", test.Optional, result, test.ExpectedResult)
		}
	}
}

func TestUnmarshalNullValuesToEmpty(t *testing.T) {
	s := struct {
		Bool    Optional[bool]      `json:"bool"`
		Byte    Optional[byte]      `json:"byte"`
		Uint    Optional[uint]      `json:"uint"`
		Uint8   Optional[uint]      `json:"uint8"`
		Uint16  Optional[uint16]    `json:"uint16"`
		Uint32  Optional[uint32]    `json:"uint32"`
		Uint64  Optional[uint64]    `json:"uint64"`
		Uintptr Optional[uintptr]   `json:"uintptr"`
		Int     Optional[int]       `json:"int"`
		Int8    Optional[int16]     `json:"int8"`
		Int16   Optional[int16]     `json:"int16"`
		Int32   Optional[int32]     `json:"int32"`
		Int64   Optional[int64]     `json:"int64"`
		Float32 Optional[float32]   `json:"float32"`
		Float64 Optional[float64]   `json:"float64"`
		Rune    Optional[rune]      `json:"rune"`
		String  Optional[string]    `json:"string"`
		Time    Optional[time.Time] `json:"time"`
	}{}

	x := `{
		  "bool": null,
		  "byte": null,
		  "uint8": null,
		  "uint16": null,
		  "uint32": null,
		  "uint64": null,
		  "uintptr": null,
		  "int8": null,
		  "int16": null,
		  "int32": null,
		  "int64": null,
		  "float32": null,
		  "float64": null,
		  "rune": null,
		  "string": null,
		  "time": null
		}`
	err := json.Unmarshal([]byte(x), &s)
	if err != nil {
		t.Error(err, "error unmarshalling")
		t.Fail()
	}
	failIfPresent := func(f func() bool) {
		if f() {
			t.Errorf("value present")
		}
	}
	failIfPresent(s.Bool.IsPresent)
	failIfPresent(s.Byte.IsPresent)
	failIfPresent(s.Float32.IsPresent)
	failIfPresent(s.Float64.IsPresent)
	failIfPresent(s.Int.IsPresent)
	failIfPresent(s.Int8.IsPresent)
	failIfPresent(s.Int16.IsPresent)
	failIfPresent(s.Int32.IsPresent)
	failIfPresent(s.Int64.IsPresent)
	failIfPresent(s.Rune.IsPresent)
	failIfPresent(s.String.IsPresent)
	failIfPresent(s.Time.IsPresent)
	failIfPresent(s.Uint.IsPresent)
	failIfPresent(s.Uint8.IsPresent)
	failIfPresent(s.Uint16.IsPresent)
	failIfPresent(s.Uint32.IsPresent)
	failIfPresent(s.Uint64.IsPresent)
	failIfPresent(s.Uintptr.IsPresent)
}
