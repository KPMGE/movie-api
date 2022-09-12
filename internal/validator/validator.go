package validator

import "regexp"

// var EmailRegex = regexp.MustCompile("^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$")

type Validator struct {
	Errors map[string]string
}

// creates a new validator, with no Errors
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Checks if the input is valid
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// add an error to the Errors map
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// checks if a condition is satisfied, if not add the message to the Errors map
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// checks if a value is in a given list
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

// checks if a value matches the provided regex
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// checks if the entry has only Unique values
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
