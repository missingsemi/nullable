package nullable

import (
	"encoding/json"
	"os"
)

type Nullable[T any] struct {
	ptr *T
}

func (n Nullable[T]) IsNull() bool {
	return n.ptr == nil
}

func (n Nullable[T]) IsSome() bool {
	return n.ptr != nil
}

func (n Nullable[T]) Unwrap() T {
	if n.ptr == nil {
		os.Exit(1)
	}
	return *n.ptr
}

func (n Nullable[T]) Expect(failStr string) T {
	if n.ptr == nil {
		os.Stderr.WriteString(failStr)
		os.Exit(1)
	}
	return *n.ptr
}

func (n Nullable[T]) Fallback(fallback T) T {
	if n.ptr == nil {
		return fallback
	}
	return *n.ptr
}

func (n *Nullable[T]) Set(value T) {
	n.ptr = &value
}

func (n *Nullable[T]) Clear() {
	n.ptr = nil
}

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

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if n.ptr == nil {
		return []byte{'n', 'u', 'l', 'l'}, nil
	}
	return json.Marshal(*n.ptr)
}

func From[T any](val T) Nullable[T] {
	return Nullable[T]{&val}
}
