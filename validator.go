package nullable

import "reflect"

// Non-generic version of Nullable that is usable in validator
type interfaced struct {
	ptr     *interface{}
	present bool
}

// Interface that Nullable implements
// Allows us to get an interfaced without ever knowing what type of Nullable was passed in.
type interfaceable interface {
	toInterfaced() interfaced
}

// Essentially converts a Nullable[T] to a Nullable[interface{}]
// Unfortunately seems to be necessary for validator support :(
func (n Nullable[T]) toInterfaced() interfaced {
	if n.ptr == nil {
		return interfaced{
			ptr:     nil,
			present: n.present,
		}
	}

	tmp := interface{}(*n.ptr)

	return interfaced{
		ptr:     &tmp,
		present: n.present,
	}
}

// Function to be registered with github.com/go-playground/validator
// Unfortunately, it must be registered against every instantiation of Nullable that needs to be validated.
func ValidateNullable(field reflect.Value) interface{} {
	if converted, ok := field.Interface().(interfaceable); ok {
		interfaced := converted.toInterfaced()

		if !interfaced.present || interfaced.ptr == nil {
			return nil
		}
		return *interfaced.ptr
	}

	return nil
}
