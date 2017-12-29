package main

import "errors"

//function to convert array of errors into single error
func ErrorsConv(errorsArr []error) error {
	var error1 error
	var errString string

	for i := range errorsArr {
		errString += errorsArr[i].Error()
	}
	error1 = errors.New(errString)
	return error1
}
