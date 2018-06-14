// Copyright 2018 GAUSS All Rights Reserved.
// Request Validator

package application

import (
	 "gopkg.in/go-playground/validator.v9"
)

var valIns = newValidator()

// Validator Validator
type Validator struct {
	validator *validator.Validate
}

// Validate Validate
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// New コンストラクタ
func newValidator() *Validator {
	ins := new(Validator)
	ins.validator = validator.New()

	return ins
}

// GetValidator GetValidator.
func GetValidator() *Validator {
	return valIns
}
