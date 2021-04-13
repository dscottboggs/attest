/**
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package attest

import (
	"testing"
)

func TestAttestPanics(t *testing.T) {
	test := New(t)
	test.AttestPanics(
		func(a ...interface{}) { panic(a[0].(string)) },
		"test panic",
	)
}

func TestAttestNoPanic(t *testing.T) {
	test := New(t)
	test.AttestNoPanic(
		func(a ...interface{}) { a[0] = a[1] },
		"args for",
		"callback func",
	)
}

// this function is called by EatError, it returns a value and a nil error
func returnsNilError() (string, error) {
	return "success", nil
}

func TestEatError(t *testing.T) {
	test := New(t)
	test.Equals("success", test.EatError(returnsNilError()).(string))
}

func TestFailOnError(t *testing.T) {
	test := New(t)
	test.Equals("success", test.FailOnError(returnsNilError()).(string))
}

// the following are explicit tests on the implementation, not implicit tests
// like the others.
func TestPanicCheckImplementationWithPanic(t *testing.T) {
	test := NewTest(t)
	defer func() {
		r := recover()
		test.Attest(r != nil, "Test error function didn't panic!")
	}()
	panic("test panic")
}
func TestPanicCheckImplementationWithNoPanic(t *testing.T) {
	test := NewTest(t)
	defer func() {
		r := recover()
		test.Attest(r == nil, "Basic math caused a panic.")
	}()
	_ = 2 + 2
}
