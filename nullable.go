/*
Package nullable provides a Nullable type that is able to represent null values in addition to differentiating them from absent json keys.
*/
package nullable

import (
	"encoding/json"
	"errors"
	"fmt"
)

/*
Nullable represents a value that can either exist or be null.
This type was designed to be especially useful for receiving input from JSON APIs.
As such, it implements the Marhsaler and Unmarshaler interfaces.
It is also possible to use this type with the validator package.
*/
type Nullable[T any] struct {
	ptr     *T
	present bool
}

/*
IsNull returns true if the Nullable is null, false otherwise.
*/
func (n Nullable[T]) IsNull() bool {
	return n.ptr == nil
}

/*
HasValue returns true if the Nullable holds a value, false otherwise.
*/
func (n Nullable[T]) HasValue() bool {
	return n.ptr != nil
}

/*
IsPresent returns true if the Nullable holds a value or was explicitly set to null.
This is useful after unmarshalling json to differentiate between keys that were absent or directly set to null.

	type Example struct {
		Key nullable.Nullable[string] `json:"key"`
	}

	var input Example
	json.Unmarshal(
		[]byte("{\"key\": null}"),
		&input,
	)
	input.Key.IsPresent() // true

	json.Unmarshal(
		[]byte("{}"),
		&input,
	)
	input.Key.IsPresent() // false

*/
func (n Nullable[T]) IsPresent() bool {
	return n.present
}

/*
IsAbsent returns false if the Nullable holds a value or was explicitly set to null.
See the documentation for IsPresent to see the diference between absent and present nulls.
*/
func (n Nullable[T]) IsAbsent() bool {
	return !n.present
}

/*
Value returns the value held by the Nullable.
If the Nullable is null, Value panics with a default message.
*/
func (n Nullable[T]) Value() T {
	if n.ptr == nil {
		var tmp T
		outStr := fmt.Sprintf("Value() called on a null %T", tmp)
		panic(outStr)
	}
	return *n.ptr
}

/*
Expect returns the value held by the Nullable.
If the Nullable is null, Expect panics with the provided message.
*/
func (n Nullable[T]) Expect(msg string) T {
	if n.ptr == nil {
		panic(msg)
	}
	return *n.ptr
}

/*
ValueOr returns the value held by the Nullable.
If the Nullable is null, ValueOr returns the provided fallback.
*/
func (n Nullable[T]) ValueOr(fallback T) T {
	if n.ptr == nil {
		return fallback
	}
	return *n.ptr
}

/*
ValueOrElse returns the value held by the Nullable.
If the Nullable is null, ValueOrElse calls the provided callback and returns its return value.
*/
func (n Nullable[T]) ValueOrElse(callback func() T) T {
	if n.ptr == nil {
		return callback()
	}
	return *n.ptr
}

/*
ValueOrDefault returns the value held by the Nullable.
If the Nullable is null, ValueOrDefault returns the zero value of the type T.
*/
func (n Nullable[T]) ValueOrDefault() T {
	if n.ptr == nil {
		var tmp T
		return tmp
	}
	return *n.ptr
}

/*
TryValue returns the value held by the Nullable.
If the Nullable is null, TryValue returns a non-nil error.
*/
func (n Nullable[T]) TryValue() (T, error) {
	if n.ptr == nil {
		var tmp T
		return tmp, errors.New(fmt.Sprintf("Value() called on a null %T", tmp))
	}
	return *n.ptr, nil
}

/*
Set stores the provided value in the Nullable and marks the Nullable as present.
A pointer to the held value is returned.
*/
func (n *Nullable[T]) Set(value T) *T {
	n.ptr = &value
	n.present = true
	return n.ptr
}

/*
Clear removes the stored value and marks the Nullable as present.
*/
func (n *Nullable[T]) Clear() {
	n.ptr = nil
	n.present = true
}

/*
UnmarshalJSON implements the json.Unmarshaler interface.
Calls to UnmarshalJSON always mark the Nullable as present.
*/
func (n *Nullable[T]) UnmarshalJSON(raw []byte) error {
	n.present = true
	err := json.Unmarshal(raw, &n.ptr)
	if err != nil {
		n.ptr = nil
		return err
	}
	return nil
}

/*
MarshalJSON implements the json.Marshaler interface.
Whether or not the Nullable is marked as present has no effect on MarshalJSON.
*/
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.ptr == nil {
		return []byte("null"), nil
	}
	return json.Marshal(*n.ptr)
}

/*
From creates a new Nullable that holds the provided value.
*/
func From[T any](val T) Nullable[T] {
	return Nullable[T]{&val, true}
}

/*
Null creates a new Nullable that is marked present and holds no value.
*/
func Null[T any]() Nullable[T] {
	return Nullable[T]{nil, true}
}

/*
Absent creates a new Nullable that is marked absent and holds no value.
*/
func Absent[T any]() Nullable[T] {
	return Nullable[T]{nil, false}
}
