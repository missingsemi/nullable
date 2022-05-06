package nullable

import (
	"encoding/json"
	"testing"
)

// Tests don't rely on other functions, but error messages will mention the equivalent function call.

func TestFrom(t *testing.T) {
	{
		got := From(10)
		tmp := 10
		want := Nullable[int]{&tmp, true}
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Wanted false.")
		} else if *got.ptr != *want.ptr {
			t.Errorf("got.Value() = %v. Wanted %v.", *got.ptr, *want.ptr)
		}
	}
	{
		got := From(true)
		tmp := true
		want := Nullable[bool]{&tmp, true}
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Wanted false.")
		} else if *got.ptr != *want.ptr {
			t.Errorf("got.Value() = %v. Wanted %v.", *got.ptr, *want.ptr)
		}
	}
	{
		got := From("hello")
		tmp := "hello"
		want := Nullable[string]{&tmp, true}
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Wanted false.")
		} else if *got.ptr != *want.ptr {
			t.Errorf("got.Value() = %v. Wanted %v.", *got.ptr, *want.ptr)
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Wanted true.")
		}
	}
}

func TestNull(t *testing.T) {
	{
		got := Null[int]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Wanted true.")
		}
	}
	{
		got := Null[bool]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Wanted true.")
		}
	}
	{
		got := Null[string]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Wanted true.")
		}
	}
}

func TestAbsent(t *testing.T) {
	{
		got := Absent[int]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == true {
			t.Error("got.IsPresent() = true. Wanted false.")
		}
	}
	{
		got := Absent[bool]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == true {
			t.Error("got.IsPresent() = true. Wanted false.")
		}
	}
	{
		got := Absent[string]()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Wanted true.")
		}
		if got.present == true {
			t.Error("got.IsPresent() = true. Wanted false.")
		}
	}
}

func TestIsNull(t *testing.T) {
	{
		got := Nullable[int]{nil, true}
		if got.IsNull() == false {
			t.Error("got.IsNull() = false. Wanted true.")
		}
	}
	{
		got := Nullable[bool]{nil, true}
		if got.IsNull() == false {
			t.Error("got.IsNull() = false. Wanted true.")
		}
	}
	{
		got := Nullable[string]{nil, true}
		if got.IsNull() == false {
			t.Error("got.IsNull() = false. Wanted true.")
		}
	}
	{
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.IsNull() == true {
			t.Error("got.IsNull() = true. Wanted false.")
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.IsNull() == true {
			t.Error("got.IsNull() = true. Wanted false.")
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.IsNull() == true {
			t.Error("got.IsNull() = true. Wanted false.")
		}
	}
}

func TestHasValue(t *testing.T) {
	{
		got := Nullable[int]{nil, true}
		if got.HasValue() == true {
			t.Error("got.HasValue() = true. Wanted false.")
		}
	}
	{
		got := Nullable[bool]{nil, true}
		if got.HasValue() == true {
			t.Error("got.HasValue() = true. Wanted false.")
		}
	}
	{
		got := Nullable[string]{nil, true}
		if got.HasValue() == true {
			t.Error("got.HasValue() = true. Wanted false.")
		}
	}
	{
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.HasValue() == false {
			t.Error("got.HasValue() = false. Wanted true.")
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.HasValue() == false {
			t.Error("got.HasValue() = false. Wanted true.")
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.HasValue() == false {
			t.Error("got.HasValue() = false. Wanted true.")
		}
	}
}

func TestIsPresent(t *testing.T) {
	{
		got := Nullable[int]{nil, false}
		if got.IsPresent() == true {
			t.Error("got.IsPresent() = true. Wanted false.")
		}
	}
	{
		got := Nullable[int]{nil, true}
		if got.IsPresent() == false {
			t.Error("got.IsPresent() = false. Wanted true.")
		}
	}
}

func TestIsAbsent(t *testing.T) {
	{
		got := Nullable[int]{nil, false}
		if got.IsAbsent() == false {
			t.Error("got.IsAbsent() = false. Wanted true.")
		}
	}
	{
		got := Nullable[int]{nil, true}
		if got.IsAbsent() == true {
			t.Error("got.IsAbsent() = true. Wanted false.")
		}
	}
}

func TestValue(t *testing.T) {
	// Calls to Value can panic so all calls are wrapped in closures with a recover call.
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Value() panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.Value() != tmp {
			t.Errorf("got.Value() = %v. Wanted %v", got.Value(), tmp)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Value() panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.Value() != tmp {
			t.Errorf("got.Value() = %v. Wanted %v", got.Value(), tmp)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Value() panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.Value() != tmp {
			t.Errorf("got.Value() = %v. Wanted %v", got.Value(), tmp)
		}
	}()
	// These next calls are supposed to panic
	func() {
		defer func() { recover() }()
		got := Nullable[int]{nil, true}
		got.Value()
		t.Errorf("got.Value() failed to panic.")
	}()
	func() {
		defer func() { recover() }()
		got := Nullable[bool]{nil, true}
		got.Value()
		t.Errorf("got.Value() failed to panic.")
	}()
	func() {
		defer func() { recover() }()
		got := Nullable[string]{nil, true}
		got.Value()
		t.Errorf("got.Value() failed to panic.")
	}()
}

func TestExpect(t *testing.T) {
	// Calls to Expect can panic so all calls are wrapped in closures with a recover call.
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Expect(\"hello\") panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.Expect("hello") != tmp {
			t.Errorf("got.Expect(\"hello\") = %v. Wanted %v", got.Expect("hello"), tmp)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Expect(\"hello\") panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.Expect("hello") != tmp {
			t.Errorf("got.Expect(\"hello\") = %v. Wanted %v", got.Expect("hello"), tmp)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("got.Expect(\"hello\") panicked with value %v. Call should not panic.", r)
			}
		}()
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.Expect("hello") != tmp {
			t.Errorf("got.Expect(\"hello\") = %v. Wanted %v", got.Expect("hello"), tmp)
		}
	}()
	// These next calls are supposed to panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				if str, ok := r.(string); !ok || str != "hello" {
					t.Errorf("got.Expect(\"hello\") panicked with value %v. Expected panic with value %v.", r, "hello")
				}
			}
		}()
		got := Nullable[int]{nil, true}
		got.Expect("hello")
		t.Errorf("got.Expect(\"hello\") failed to panic.")
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				if str, ok := r.(string); !ok || str != "hello" {
					t.Errorf("got.Expect(\"hello\") panicked with value %v. Expected panic with value %v.", r, "hello")
				}
			}
		}()
		got := Nullable[bool]{nil, true}
		got.Expect("hello")
		t.Errorf("got.Expect(\"hello\") failed to panic.")
	}()
	func() {
		defer func() {
			if r := recover(); r != nil {
				if str, ok := r.(string); !ok || str != "hello" {
					t.Errorf("got.Expect(\"hello\") panicked with value %v. Expected panic with value %v.", r, "hello")
				}
			}
		}()
		got := Nullable[string]{nil, true}
		got.Expect("hello")
		t.Errorf("got.Expect(\"hello\") failed to panic.")
	}()
}

func TestValueOr(t *testing.T) {
	{
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.ValueOr(100) != tmp {
			t.Errorf("got.ValueOr(100) = %v. Wanted %v.", got.ValueOr(100), tmp)
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.ValueOr(false) != tmp {
			t.Errorf("got.ValueOr(false) = %v. Wanted %v.", got.ValueOr(false), tmp)
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.ValueOr("bye") != tmp {
			t.Errorf("got.ValueOr(\"bye\") = %v. Wanted %v.", got.ValueOr("bye"), tmp)
		}
	}
	{
		got := Nullable[int]{nil, true}
		if got.ValueOr(100) != 100 {
			t.Errorf("got.ValueOr(100) = %v. Wanted %v.", got.ValueOr(100), 100)
		}
	}
	{
		got := Nullable[bool]{nil, true}
		if got.ValueOr(false) != false {
			t.Errorf("got.ValueOr(false) = %v. Wanted %v.", got.ValueOr(false), false)
		}
	}
	{
		got := Nullable[string]{nil, true}
		if got.ValueOr("bye") != "bye" {
			t.Errorf("got.ValueOr(\"bye\") = %v. Wanted %v.", got.ValueOr("bye"), "bye")
		}
	}
}

func TestValueOrElse(t *testing.T) {
	{
		tmp := 10
		got := Nullable[int]{&tmp, true}
		if got.ValueOrElse(func() int { return 100 }) != tmp {
			t.Errorf("got.ValueOrElse(func() int { return 100 }) = %v. Wanted %v.", got.ValueOrElse(func() int { return 100 }), tmp)
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, true}
		if got.ValueOrElse(func() bool { return false }) != tmp {
			t.Errorf("got.ValueOrElse(func() bool { return false }) = %v. Wanted %v.", got.ValueOrElse(func() bool { return false }), tmp)
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		if got.ValueOrElse(func() string { return "bye" }) != tmp {
			t.Errorf("got.ValueOrElse(func() string { return \"bye\" }) = %v. Wanted %v.", got.ValueOrElse(func() string { return "bye" }), tmp)
		}
	}
	{
		got := Nullable[int]{nil, true}
		if got.ValueOrElse(func() int { return 100 }) != 100 {
			t.Errorf("got.ValueOrElse(func() int { return 100 }) = %v. Wanted %v.", got.ValueOrElse(func() int { return 100 }), 100)
		}
	}
	{
		got := Nullable[bool]{nil, true}
		if got.ValueOrElse(func() bool { return false }) != false {
			t.Errorf("got.ValueOrElse(func() bool { return false }) = %v. Wanted %v.", got.ValueOrElse(func() bool { return false }), false)
		}
	}
	{
		got := Nullable[string]{nil, true}
		if got.ValueOrElse(func() string { return "bye" }) != "bye" {
			t.Errorf("got.ValueOrElse(func() string { return \"bye\" }) = %v. Wanted %v.", got.ValueOrElse(func() string { return "bye" }), "bye")
		}
	}
}

func TestSet(t *testing.T) {
	{
		got := Nullable[int]{nil, false}
		ptr := got.Set(10)
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Expected false.")
		} else if *got.ptr != 10 {
			t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, 10)
		}

		if ptr == nil {
			t.Errorf("ptr = %v. Expected valid pointer.", ptr)
		} else {
			*ptr = 100
			if got.ptr == nil {
				t.Error("got.IsNull() = true. Expected false.")
			} else if *got.ptr != 100 {
				t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, 100)
			}
		}

		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
	{
		got := Nullable[bool]{nil, false}
		ptr := got.Set(true)
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Expected false.")
		} else if *got.ptr != true {
			t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, true)
		}

		if ptr == nil {
			t.Errorf("ptr = %v. Expected valid pointer.", ptr)
		} else {
			*ptr = false
			if got.ptr == nil {
				t.Error("got.IsNull() = true. Expected false.")
			} else if *got.ptr != false {
				t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, false)
			}
		}

		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
	{
		got := Nullable[string]{nil, false}
		ptr := got.Set("hello")
		if got.ptr == nil {
			t.Error("got.IsNull() = true. Expected false.")
		} else if *got.ptr != "hello" {
			t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, "hello")
		}

		if ptr == nil {
			t.Errorf("ptr = %v. Expected valid pointer.", ptr)
		} else {
			*ptr = "bye"
			if got.ptr == nil {
				t.Error("got.IsNull() = true. Expected false.")
			} else if *got.ptr != "bye" {
				t.Errorf("got.Value() = %v. Expected %v.", *got.ptr, "bye")
			}
		}

		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
}

func TestClear(t *testing.T) {
	{
		tmp := 10
		got := Nullable[int]{&tmp, false}
		got.Clear()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Expected true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, false}
		got.Clear()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Expected true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, false}
		got.Clear()
		if got.ptr != nil {
			t.Error("got.IsNull() = false. Expected true.")
		}
		if got.present == false {
			t.Error("got.IsPresent() = false. Expected true.")
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	{
		type S struct {
			Got Nullable[int] `json:"got"`
		}
		var s S
		j := []byte(`{"got": 10}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr == nil {
				t.Error("s.Got.IsNull() = true. Expected false.")
			} else if *s.Got.ptr != 10 {
				t.Errorf("s.Got.Value() = %v. Expected %v.", *s.Got.ptr, 10)
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[bool] `json:"got"`
		}
		var s S
		j := []byte(`{"got": true}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr == nil {
				t.Error("s.Got.IsNull() = true. Expected false.")
			} else if *s.Got.ptr != true {
				t.Errorf("s.Got.Value() = %v. Expected %v.", *s.Got.ptr, true)
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[string] `json:"got"`
		}
		var s S
		j := []byte(`{"got": "hello"}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr == nil {
				t.Error("s.Got.IsNull() = true. Expected false.")
			} else if *s.Got.ptr != "hello" {
				t.Errorf("s.Got.Value() = %v. Expected %v.", *s.Got.ptr, "hello")
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[int] `json:"got"`
		}
		var s S
		j := []byte(`{"got": null}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[bool] `json:"got"`
		}
		var s S
		j := []byte(`{"got": null}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[string] `json:"got"`
		}
		var s S
		j := []byte(`{"got": null}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == false {
				t.Error("s.Got.IsPresent() = false. Expected true.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[int] `json:"got"`
		}
		var s S
		j := []byte(`{}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == true {
				t.Error("s.Got.IsPresent() = true. Expected false.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[bool] `json:"got"`
		}
		var s S
		j := []byte(`{}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == true {
				t.Error("s.Got.IsPresent() = true. Expected false.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[string] `json:"got"`
		}
		var s S
		j := []byte(`{}`)
		err := json.Unmarshal(j, &s)
		if err != nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected %v.", err, nil)
		} else {
			if s.Got.ptr != nil {
				t.Error("s.Got.IsNull() = false. Expected true.")
			}
			if s.Got.present == true {
				t.Error("s.Got.IsPresent() = true. Expected false.")
			}
		}
	}
	{
		type S struct {
			Got Nullable[int] `json:"got"`
		}
		var s S
		j := []byte(`{"got": "hello"}`)
		err := json.Unmarshal(j, &s)
		if err == nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			Got Nullable[bool] `json:"got"`
		}
		var s S
		j := []byte(`{"got": 10}`)
		err := json.Unmarshal(j, &s)
		if err == nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			Got Nullable[string] `json:"got"`
		}
		var s S
		j := []byte(`{"got": true}`)
		err := json.Unmarshal(j, &s)
		if err == nil {
			t.Errorf("json.Unmarshal(j, &s) = %v. Expected error.", err)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	{
		tmp := 10
		got := Nullable[int]{&tmp, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "10" {
			t.Errorf("j = %v. Expected %v.", string(j), "10")
		}
	}
	{
		tmp := true
		got := Nullable[bool]{&tmp, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "true" {
			t.Errorf("j = %v. Expected %v.", string(j), "true")
		}
	}
	{
		tmp := "hello"
		got := Nullable[string]{&tmp, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "\"hello\"" {
			t.Errorf("j = %v. Expected %v.", string(j), "\"hello\"")
		}
	}
	{
		got := Nullable[int]{nil, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "null" {
			t.Errorf("j = %v. Expected %v.", string(j), "null")
		}
	}
	{
		got := Nullable[bool]{nil, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "null" {
			t.Errorf("j = %v. Expected %v.", string(j), "null")
		}
	}
	{
		got := Nullable[string]{nil, true}
		j, err := json.Marshal(got)
		if err != nil {
			t.Errorf("json.Marshal(got) err = %v. Expected nil.", err)
		} else if string(j) != "null" {
			t.Errorf("j = %v. Expected %v.", string(j), "null")
		}
	}
}
