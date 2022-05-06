package nullable

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"
)

func TestValidatePresentPass(t *testing.T) {

	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:"required,min=5"`
	}

	v := A{From("hello")}
	if err := validate.Struct(v); err != nil {
		t.Errorf("err = %v; Expected nil", err)
	}
}

func TestValidatePresentFail(t *testing.T) {

	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:"required,min=5"`
	}

	// Fail on min

	v1 := A{From("a")}
	if err := validate.Struct(v1); err == nil {
		t.Error("err = nil; Expected err")
	}

	// Fail on required

	v2 := A{From("")}
	if err := validate.Struct(v2); err == nil {
		t.Error("err = nil; Expected err")
	}
}

func TestValidateAbsentPass(t *testing.T) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:""`
	}

	v1 := A{Absent[string]()}
	if err := validate.Struct(v1); err != nil {
		t.Errorf("err = %v; Expected nil", err)
	}
}

func TestValidateAbsentFail(t *testing.T) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:"required"`
	}

	v1 := A{Absent[string]()}
	if err := validate.Struct(v1); err == nil {
		t.Error("err = nil; Expected err")
	}
}

func TestValidateNullPass(t *testing.T) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:""`
	}

	v1 := A{Null[string]()}
	if err := validate.Struct(v1); err != nil {
		t.Errorf("err = %v; Expected nil", err)
	}
}

func TestValidateNullFail(t *testing.T) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[string]{})

	type A struct {
		A Nullable[string] `validate:"required"`
	}

	v1 := A{Null[string]()}
	if err := validate.Struct(v1); err == nil {
		t.Error("err = nil; Expected err")
	}
}
