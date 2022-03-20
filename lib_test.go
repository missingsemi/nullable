package nullable

import (
	"encoding/json"
	"testing"
)

func TestFrom(t *testing.T) {
	v1 := From(10)
	i := 10
	v2 := Nullable[int]{&i}
	if *v1.ptr != *v2.ptr {
		t.Errorf("v.From(10) != Nullable{&10}; Expected equal")
	}
}

func TestIsNull(t *testing.T) {
	v := From(10)
	if v.IsNull() {
		t.Errorf("v.IsNull() = true; Expected false")
	}
	v = Nullable[int]{}
	if !v.IsNull() {
		t.Errorf("v.IsNull() = false; Expected true")
	}
}

func TestIsSome(t *testing.T) {
	v := Nullable[int]{}
	if v.IsSome() {
		t.Errorf("v.IsSome() = true; Expected false")
	}
	v = From(10)
	if !v.IsSome() {
		t.Errorf("v.IsSome() = false; Expected true")
	}
}

func TestNullableUnwrapExpect(t *testing.T) {
	v := From(10)
	got := v.Unwrap()
	if got != 10 {
		t.Errorf("v.Unwrap() = %d; Wanted 10", got)
	}

	got = v.Expect("")
	if got != 10 {
		t.Errorf("v.Expect(\"\") = %d; Wanted 10", got)
	}
}

func TestNullableFallbackNull(t *testing.T) {
	v := Nullable[int]{}
	got := v.Fallback(20)
	if got != 20 {
		t.Errorf("v.Fallback(20) = %d; Wanted 20", got)
	}
}

func TestNullableFallbackSome(t *testing.T) {
	v := From(10)
	got := v.Fallback(20)
	if got != 10 {
		t.Errorf("v.Fallback(20) = %d; Wanted 10", got)
	}
}

func TestNullableSet(t *testing.T) {
	v := From(10)
	v.Set(20)
	got := v.Unwrap()
	if got != 20 {
		t.Errorf("v.Set(20) = %d; Wanted 20", got)
	}
}

func TestNullableClear(t *testing.T) {
	v := From(10)
	v.Clear()
	got := v.ptr == nil
	if got != true {
		t.Errorf("v.ptr == nil = %t; Wanted true", got)
	}
}
func TestUnmarshalNull(t *testing.T) {
	var v Nullable[int]
	input := []byte("null")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %t; Wanted nil", err)
	}
	got := v.Fallback(20)
	if got != 20 {
		t.Errorf("v.Fallback(20) = %d; Wanted 20", got)
	}
}

func TestUnmarshalSome(t *testing.T) {
	var v Nullable[int]
	input := []byte("10")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %t; Wanted nil", err)
	}
	got := v.Fallback(20)
	if got != 10 {
		t.Errorf("v.Fallback(20) = %d; Wanted 10", got)
	}
}

func TestUnmarshalFail(t *testing.T) {
	var v Nullable[int]
	input := []byte("what")
	err := json.Unmarshal(input, &v)
	if err == nil {
		t.Errorf("err = %t; Wanted error", err)
	}
}

func TestMarshalSome(t *testing.T) {
	v := From(10)
	out, err := json.Marshal(v)
	if err != nil {
		t.Errorf("err = %t; Wanted nil", err)
	}
	if string(out) != "10" {
		t.Errorf("json.Marshal(v) = %s; Wanted \"10\"", string(out))
	}
}

func TestMarshalNull(t *testing.T) {
	v := Nullable[int]{}
	out, err := json.Marshal(v)
	if err != nil {
		t.Errorf("err = %t; Wanted nil", err)
	}
	if string(out) != "null" {
		t.Errorf("json.Marshal(v) = %s; Wanted \"null\"", string(out))
	}
}
