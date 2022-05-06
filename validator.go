package nullable

import "reflect"

/*
interfaceNullable is an Interface-based Nullable that is used for validator support.
Because validator doesn't support generics, a Nullable[T] has to be converted into an interfaceNullable before being validated.
*/
type interfaceNullable struct {
	ptr     *interface{}
	present bool
}

/*
interfaceable requires a method that generates an interfaceNullable.
*/
type interfaceable interface {
	toInterfaceNullable() interfaceNullable
}

/*
toInterfaceNullable implements interfaceable for the Nullable type.
*/
func (n Nullable[T]) toInterfaceNullable() interfaceNullable {
	if n.ptr == nil {
		return interfaceNullable{
			ptr:     nil,
			present: n.present,
		}
	}

	tmp := interface{}(*n.ptr)

	return interfaceNullable{
		ptr:     &tmp,
		present: n.present,
	}
}

/*
Handler to be registered with validator.
Due to how go handles generics, each instantiated type of Nullable must be registered with validator.
*/
func ValidateNullable(field reflect.Value) interface{} {
	if converted, ok := field.Interface().(interfaceable); ok {
		interfaced := converted.toInterfaceNullable()

		if !interfaced.present || interfaced.ptr == nil {
			return nil
		}
		return *interfaced.ptr
	}

	return nil
}
