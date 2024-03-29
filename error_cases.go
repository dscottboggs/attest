/**
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package attest

import "log"

/*
These tests are passed (possibly nil) errors. The test fails if the error is
not nil, and logs the error and, in some cases, an optional custom message.
*/

// AttestPanics -- Attest that when fun is called with args, it causes a panic.
// e.g.
//	t.AttestPanics(func(){log.Printf("Panics, passes test."); panic()})
//	t.AttestPanics(func(){log.Printf("Doesn't panic, fails test.")})
func (t *Test) AttestPanics(fun func(...interface{}), args ...interface{}) {
	defer func() {
		r := recover()
		t.Attest(r != nil, "Function %v didn't cause a panic!", fun)
	}()
	fun(args...)
}

// AttestNoPanic -- the inverse of AttestPanics
func (t *Test) AttestNoPanic(fun func(...interface{}), args ...interface{}) {
	defer func() {
		r := recover()
		t.Attest(r == nil, "Function %v caused a panic!", fun)
	}()
	fun(args...)
}

// Handle -- log and fail for an arbitrary number of errors.
func (t *Test) HandleMultiple(e ...error) {
	for _, err := range e {
		if err != nil {
			t.Error(err)
		}
	}
}

// Handle -- handle an error with an optional custom message.
func (t *Test) Handle(err error, msgAndFmt ...interface{}) {
	if err == nil && len(msgAndFmt) == 0 {
		return
	}
	if len(msgAndFmt) == 0 {
		t.Error(err)
		return
	}
	if err != nil {
		switch msgAndFmt[0].(type) {
		case string:
			t.errorf(msgAndFmt[0].(string), msgAndFmt[1:]...)
		case error:
			t.errorf(
				"WARNING! starting at attest version 1.0, use HandleMultiple to handle" +
					"multiple error cases.")
		default:
			t.errorf(
				"Got type %T for second argument to Test.Handle(). If more than one"+
					"argument is specified, the second one MUST be a string.",
				msgAndFmt[0])
		}
	}
}

// StopIf -- Fail the test and stop running it if an error is present, with
// optional message.
func (t *Test) StopIf(err error, msgAndFmt ...interface{}) {
	if err != nil {
		if len(msgAndFmt) == 0 {
			msgAndFmt = []interface{}{"Fatal error: %s (%#+v)", err.Error(), err}
		}
		log.Printf(msgAndFmt[0].(string), msgAndFmt[1:]...)
		t.FailNow()
	}
}

// EatError accepts two values, the latter of which is a nillable error. If the
// error is not nil, the test is failed. Regardless, the first value is
// returned through the function.
func (t *Test) EatError(value interface{}, err error) interface{} {
	if err != nil {
		t.errorf("When aquiring value %#v, got error %s (%#+v)", value, err.Error(), err)
	}
	return value
}

// FailOnError accepts two values, the latter of which is a nillable error. If the
// error is not nil, the test is failed immediately.
func (t *Test) FailOnError(value interface{}, err error, msgAndFormat ...interface{}) interface{} {
	t.StopIf(err, msgAndFormat...)
	return value
}
