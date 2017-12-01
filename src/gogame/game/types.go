package game

import validator "gopkg.in/validator.v2"

var id int

// Object contains all basic and common attributes for mos object in the
// game.
type Object struct {
	ID   int    `validate:"nonzero"`
	Name string `validate:"nonzero"`
	Desc string
}

// NewObject function creates a new Object.
func NewObject(name string) Object {
	id++
	obj := Object{
		ID:   id,
		Name: name,
	}
	if errs := validator.Validate(obj); errs != nil {
		panic(errs)
	}
	return obj
}
