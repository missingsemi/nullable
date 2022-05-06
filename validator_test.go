package nullable

import (
	"testing"

	"gopkg.in/go-playground/validator.v9"
)

func TestValidate(t *testing.T) {
	validate := validator.New()
	validate.RegisterCustomTypeFunc(ValidateNullable, Nullable[int]{}, Nullable[string]{})
	{
		type S struct {
			S Nullable[int] `validate:"required,min=5"`
		}
		tmp := 10
		got := S{Nullable[int]{&tmp, true}}
		if err := validate.Struct(got); err != nil {
			t.Errorf("validate.Struct(got) = %v. Expected %v.", err, nil)
		}
	}
	{
		type S struct {
			S Nullable[string] `validate:"required,min=5"`
		}
		tmp := "hello"
		got := S{Nullable[string]{&tmp, true}}
		if err := validate.Struct(got); err != nil {
			t.Errorf("validate.Struct(got) = %v. Expected %v.", err, nil)
		}
	}
	{
		type S struct {
			S Nullable[int] `validate:"required,min=5"`
		}
		tmp := 1
		got := S{Nullable[int]{&tmp, true}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[string] `validate:"required,min=5"`
		}
		tmp := "hi"
		got := S{Nullable[string]{&tmp, true}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[int] `validate:"required,min=5"`
		}
		tmp := 10
		got := S{Nullable[int]{&tmp, false}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[string] `validate:"required,min=5"`
		}
		tmp := "hello"
		got := S{Nullable[string]{&tmp, false}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[int] `validate:"required,min=5"`
		}
		got := S{Nullable[int]{nil, true}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[string] `validate:"required,min=5"`
		}
		got := S{Nullable[string]{nil, true}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[int] `validate:"required,min=5"`
		}
		got := S{Nullable[int]{nil, false}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type S struct {
			S Nullable[string] `validate:"required,min=5"`
		}
		got := S{Nullable[string]{nil, false}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
	{
		type Invalid struct {
			Invalid int
		}
		type S struct {
			S Invalid `validate:"required"`
		}
		validate.RegisterCustomTypeFunc(ValidateNullable, Invalid{})
		got := S{Invalid{10}}
		if err := validate.Struct(got); err == nil {
			t.Errorf("validate.Struct(got) = %v. Expected error.", err)
		}
	}
}
