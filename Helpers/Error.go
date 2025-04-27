/*
 * Vueghost Tech, FZE © 2018- 2019. Unauthorized copying of this file,
 * via any medium is strictly prohibited proprietary and confidential,
 * can not be copied and/or distributed without the express permission of Vueghost
 */

package Helpers

import "strings"

// Errors is an array of multiple errors and conforms to the error interface.
type Errors []error

// Errors returns itself.
func (es Errors) Errors() []error {
	return es
}

func (es Errors) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	return strings.Join(errs, ";")
}

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error struct {
	Name                     string
	Err                      error
	CustomErrorMessageExists bool

	// Validator indicates the name of the validator that failed
	Validator string
	Path      []string
}

func (e Error) Error() string {
	if e.CustomErrorMessageExists {
		return e.Err.Error()
	}

	errName := e.Name
	if len(e.Path) > 0 {
		errName = strings.Join(append(e.Path, e.Name), ".")
	}

	return errName + ": " + e.Err.Error()
}
