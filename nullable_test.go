package nullable

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
)

func TestFrom(t *testing.T) {
	v1 := From(10)
	i := 10
	v2 := Nullable[int]{&i, true}
	if *v1.ptr != *v2.ptr {
		t.Errorf("v.From(10) != Nullable{&10}; Expected equal")
	}
}

func TestNull(t *testing.T) {
	v := Null[int]()
	if v.ptr != nil {
		t.Errorf("v.ptr != nil; Expected equal")
	}
}

func TestEmpty(t *testing.T) {
	v := Absent[int]()
	if v.ptr != nil || v.present != false {
		t.Errorf("v = %v; Wanted {nil, false}", v)
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

func TestHasValue(t *testing.T) {
	v := Nullable[int]{nil, true}
	if v.HasValue() {
		t.Errorf("v.HasValue() = true; Expected false")
	}
	v = From(10)
	if !v.HasValue() {
		t.Errorf("v.HasValue() = false; Expected true")
	}
}

func TestIsPresent(t *testing.T) {
	v := Nullable[int]{nil, true}
	if !v.IsPresent() {
		t.Errorf("v.IsPresent() = false; Expected true")
	}
	v = Nullable[int]{nil, false}
	if v.IsPresent() {
		t.Errorf("v.IsPresent() = true; Expected false")
	}
}

func TestIsAbsent(t *testing.T) {
	v := Nullable[int]{nil, true}
	if v.IsAbsent() {
		t.Errorf("v.IsAbsent() = true; Expected false")
	}
	v = Nullable[int]{nil, false}
	if !v.IsAbsent() {
		t.Errorf("v.IsAbsent() = false; Expected true")
	}
}
func TestNullableValueNull(t *testing.T) {
	if os.Getenv("OS_EXIT_TEST") == "1" {
		v := Null[int]()
		_ = v.Value()
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestNullableValueNull")
	cmd.Env = append(os.Environ(), "OS_EXIT_TEST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("v.Value() ran with error %v; Wanted exit status 1", err)
}

func TestNullableValueSome(t *testing.T) {
	v := From(10)
	got := v.Value()
	if got != 10 {
		t.Errorf("v.Value() = %d; Wanted 10", got)
	}

	got = v.Expect("")
	if got != 10 {
		t.Errorf("v.Expect(\"\") = %d; Wanted 10", got)
	}
}

func TestNullableExpectNull(t *testing.T) {
	if os.Getenv("OS_EXIT_TEST") == "1" {
		v := Null[int]()
		_ = v.Expect("")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestNullableExpectNull")
	cmd.Env = append(os.Environ(), "OS_EXIT_TEST=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("v.Expect() ran with error %v; Wanted exit status 1", err)
}

func TestNullableExpectSome(t *testing.T) {
	v := From(10)
	got := v.Expect("")
	if got != 10 {
		t.Errorf("v.Expect(\"\") = %d; Wanted 10", got)
	}
}

func TestNullableValueOrNull(t *testing.T) {
	v := Nullable[int]{}
	got := v.ValueOr(20)
	if got != 20 {
		t.Errorf("v.ValueOr(20) = %d; Wanted 20", got)
	}
}

func TestNullableValueOrSome(t *testing.T) {
	v := From(10)
	got := v.ValueOr(20)
	if got != 10 {
		t.Errorf("v.ValueOr(20) = %d; Wanted 10", got)
	}
}

func TestNullableValueOrElseNull(t *testing.T) {
	v := Nullable[int]{}
	f := func() int {
		return 20
	}
	got := v.ValueOrElse(f)
	if got != 20 {
		t.Errorf("v.ValueOrElse(f) = %d; Wanted 20", got)
	}
}

func TestNullableValueOrElseSome(t *testing.T) {
	v := From(10)
	f := func() int {
		return 20
	}
	got := v.ValueOrElse(f)
	if got != 10 {
		t.Errorf("v.ValueOrElse(f) = %d; Wanted 10", got)
	}
}

func TestNullableSet(t *testing.T) {
	v := From(10)
	v.Set(20)
	got := v.Value()
	if got != 20 {
		t.Errorf("v.Set(20) = %d; Wanted 20", got)
	}
}

func TestNullableClear(t *testing.T) {
	v := From(10)
	v.Clear()
	got := v.ptr == nil
	if got != true {
		t.Errorf("v.ptr == nil = %v; Wanted true", got)
	}
}
func TestUnmarshalNull(t *testing.T) {
	var v Nullable[int]
	input := []byte("null")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	got := v.ValueOr(20)
	if got != 20 {
		t.Errorf("v.ValueOr(20) = %d; Wanted 20", got)
	}
}

func TestUnmarshalSome(t *testing.T) {
	var v Nullable[int]
	input := []byte("10")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	got := v.ValueOr(20)
	if got != 10 {
		t.Errorf("v.ValueOr(20) = %d; Wanted 10", got)
	}
}

func TestUnmarshalFail(t *testing.T) {
	var v Nullable[int]
	input := []byte("\"what\"")
	err := json.Unmarshal(input, &v)
	if err == nil {
		t.Errorf("err = nil; Wanted error")
	}
}

func TestUnmarshalAbsent(t *testing.T) {
	type S struct {
		A Nullable[int]
	}
	var v S
	input := []byte("{}")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	if v.A.ptr != nil || v.A.present != false {
		t.Errorf("v = %v; Wanted {nil, false}", v)
	}
}

func TestMarshalSome(t *testing.T) {
	v := From(10)
	out, err := json.Marshal(v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	if string(out) != "10" {
		t.Errorf("json.Marshal(v) = %s; Wanted \"10\"", string(out))
	}
}

func TestMarshalNull(t *testing.T) {
	v := Nullable[int]{}
	out, err := json.Marshal(v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	if string(out) != "null" {
		t.Errorf("json.Marshal(v) = %s; Wanted \"null\"", string(out))
	}
}

func TestUnmarshalSliceSome(t *testing.T) {
	type S struct {
		A Nullable[[]int] `json:"a"`
	}
	var v S
	input := []byte("{\"a\": [1, 2, 3]}")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	if v.A.ptr == nil || v.A.present == false {
		t.Errorf("v = %v; Wanted {nil, false}", v)
	}
	a := v.A.Value()
	if len(a) != 3 || a[0] != 1 || a[1] != 2 || a[2] != 3 {
		t.Errorf("v.A = %v; Wanted [1, 2, 3]", a)
	}
}

func TestUnmarshalSliceAbsent(t *testing.T) {
	type S struct {
		A Nullable[[]int] `json:"a"`
	}
	var v S
	input := []byte("{}")
	err := json.Unmarshal(input, &v)
	if err != nil {
		t.Errorf("err = %v; Wanted nil", err)
	}
	if v.A.ptr != nil || v.A.present != false {
		t.Errorf("v = %v; Wanted {nil, false}", v)
	}
}
