package nullable

import (
	"encoding/json"
	"fmt"
	"os"
)

// Struct representing a value that can either exist or be null.
// Implements Marshaler and Unmarshaler, falling back on default behavior after doing a null check.
type Nullable[T any] struct {
	ptr *T
}

// Returns true if the Nullable represents a null value.
func (n Nullable[T]) IsNull() bool {
	return n.ptr == nil
}

// Returns true if the Nullable represents a non null value.
func (n Nullable[T]) IsSome() bool {
	return n.ptr != nil
}

// Unwraps the Nullable object, calling os.Exit if the Nullable is null.
func (n Nullable[T]) Unwrap() T {
	if n.ptr == nil {
		var tmp T
		outStr := fmt.Sprintf("Unwrapped on a null %T", tmp)
		os.Stderr.WriteString(outStr)
		os.Exit(1)
	}
	return *n.ptr
}

// Unwraps the Nullable, printing failStr to Stderr then calling os.Exit if the nullable is null.
func (n Nullable[T]) Expect(failStr string) T {
	if n.ptr == nil {
		os.Stderr.WriteString(failStr)
		os.Exit(1)
	}
	return *n.ptr
}

// Unwraps the Nullable or returns the fallback value if the Nullable is null.
func (n Nullable[T]) Fallback(fallback T) T {
	if n.ptr == nil {
		return fallback
	}
	return *n.ptr
}

// Sets the Nullable's value to the passed value.
// No cleanup is done on the previously stored value if it existed.
// A pointer to the stored value is returned.
func (n *Nullable[T]) Set(value T) *T {
	n.ptr = &value
	return n.ptr
}

// Sets the Nullable's value to null.
// No cleanup is done on the previously stored value if it existed.
func (n *Nullable[T]) Clear() {
	n.ptr = nil
}

// Checks if the json data is equal to null. Otherwise falls back onto default behavior.
func (n *Nullable[T]) UnmarshalJSON(raw []byte) error {
	if string(raw) == "null" {
		n.ptr = nil
		return nil
	}

	var res T
	err := json.Unmarshal(raw, &res)

	if err != nil {
		n.ptr = nil
		return err
	}

	n.ptr = &res
	return nil
}

// Converts the Nullable to null or falls back on default behavior if the value exists.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.ptr == nil {
		return []byte{'n', 'u', 'l', 'l'}, nil
	}
	return json.Marshal(*n.ptr)
}

// Creates a Nullable object from the passed value.
func From[T any](val T) Nullable[T] {
	return Nullable[T]{&val}
}

// Creates a null Nullable object.
func Null[T any]() Nullable[T] {
	return Nullable[T]{nil}
}
