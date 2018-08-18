/* Attest is a very lightweight testing library aimed at improving the
 * intuitiveness and aesthetics of the go standard testing library, as well as
 * reducing the amount of keystrokes per test, hence improving developer
 * efficiency by a marginal amount with very minimal overhead and risk. Any
 * given testing.T function can create an attest.Test object, whose methods
 * can then perform tests.
 *
 * A brief example:
 *
 *     package main
 *
 *        import (
 *            "testing"
 *
 *            "github.com/dscottboggs/attest"
 *        )
 *
 *        func TestExample(t \*testing.T) {
 *            test := attest.Test{t}
 *            test.Attest(fmt.Sprintf("%T", "that something is true") == "string", "or %s a message", "log")
 *            const unchanging = 0
 *            var variable int
 *            test.Equals(unchanging, variable)
 *            variable = 1
 *            test.Greater(variable, unchanging)
 *        }
 *
 * That same test function with the default testing library might be written
 * like:
 *
 *        func TestExample(t \*testing.T) {
 *            if fmt.Sprintf("%T", "something is true") != "string" {
 *                t.Errorf("or %s a message", "log")
 *            }
 *            const unchanging = 0
 *            var variable int
 *            if fmt.Sprintf("%T", unchanging) != fmt.Sprintf("%T", variable) {
 *                t.Errorf("Value 1 had a different type (%T) than value 2 (%T)", unchanging, variable)
 *            }
 *            if unchanging != variable {
 *                t.Errorf("Value 1 (%d) didn't equal value 2 (%d).")
 *            }
 *            variable = 1
 *            if fmt.Sprintf("%T", unchanging) != fmt.Sprintf("%T", variable) {
 *                t.Errorf("Value 1 had a different type (%T) than value 2 (%T)", unchanging, variable)
 *            }
 *            if variable <= unchanging {
 *                t.Errorf("Value 1 was less than or equal to value 2")
 *            }
 *	  }
 *
 * As you can see, this provides minimal benefit besides a reduced number of
 * keystrokes when *writing*, but when reading back, the attest way is much
 * easier to understand. Of course, you can mix-and-match:
 *
 *     func TestExample(t \*testing.T){
 *         test := attest.Test{t}
 *         if fmt.Sprintf("%T", "something is true") != "string" {
 *             test.Errorf("or %s a message", "log")
 *         }
 *         const unchanging = 0
 *         var variable int
 *         test.Equals(unchanging, variable)
 *         variable = 1
 *         test.GreaterThan(unchanging, variable)
 *     }
 *
 * ### Logging a custom message
 * All tests allow for an optional (or in the case of the few strictly boolean
 * tests, required) message string and formatters to be forwarded to
 * fmt.Sprintf()
 *
 * # Available test functions
 * The following tests are available:
 *
 *  - **Attest**: the first argument must equal the boolean value true.
 *  - **AttestNot** and **Not**: the first argument must equal the boolean value false.
 *  - **AttestOrDo**: takes a callback function and arguments to forward to the callback in case of a failure
 *  - **Nil** and **NotNil**: the first argument must be nil or not nil, respectively.
 *  - **Equals** and **NotEqual**: the second argument must equal (or not equal, respectively) the first argument. Both require that the arguments be the same type
 *  - **GreaterThan** and **LessThan**: like Equals, but checks for the second value to be greater or less than the first argument.
 *  - **Positive** and **Negative**: are shortcuts for test.LessThan(0, ...) and test.GreaterThan(0, ...)
 *  - **TypeIs** and **TypeIsNot**: check the type of a value
 *
 * In addition there are the following ways of handling error types and panics:
 *  - **Handle**: Log and fail if any of the arguments to this are non-nil errors.
 *  - **AttestPanics** and **AttestNoPanic**: ensure the given function panics or doesn't.
 *  - **StopIf**: Log and fail a fatal non-nil error
 *  - **MessageHandle**: Like handle but instead of accepting an arbitrary number of errors, it accepts a message and formatters
 *  - **EatError**: Logs and fails an error message if the second argument is a non-nil error, and returns the first argument. For handling function calls that return a value and an error in a single line.
 *
 * The following APIs are works in progress and are subject to change:
 *  - **In** and **NotIn**: Currently doesn't work as expected due to go's typing system.
 */

package attest

import (
	"fmt"
	"log"
	"testing"
)

// Test -- A structure for containing methods and data for asserting and
// testing assertion validity
type Test struct {
	*testing.T
}

func typeOf(val interface{}) string {
	return fmt.Sprintf("%T", val)
}

// Equals that var1 is deeply equal to var2. Optionally, you can pass an
// additional string and additional string formatters to be passed to
// Test.Attest. If no message is specified, a message will be logged simply
// stating that the two values weren't equal.
func (t *Test) Equals(
	var1, var2 interface{}, msgAndFormatters ...interface{},
) {
	if len(msgAndFormatters) > 0 {
		t.Attest(
			typeOf(var1) == typeOf(var2),
			msgAndFormatters[0].(string),
			msgAndFormatters[1:]...)
		t.Attest(
			var1 == var2,
			msgAndFormatters[0].(string),
			msgAndFormatters[1:]...)
	} else {
		t.Attest(
			typeOf(var1) == typeOf(var2),
			fmt.Sprintf(
				"%#v of type %T didn't match the type of %#v, %T; so they can't be"+
					" compared. ",
				var1,
				var1,
				var2,
				var2))
		t.Attest(
			var1 == var2,
			fmt.Sprintf(
				"Expected %#v (%v) was actually %#v (%v)",
				var1,
				var1,
				var2,
				var2))
	}
}

// NotEqual fails the test if var1 doesn't equal var2, with the given message
// and formatting.
func (t *Test) NotEqual(var1, var2 interface{},
	msg string,
	fmt ...interface{},
) {
	if typeOf(var1) != typeOf(var2) {
		// types don't match, not equal by default.
		return
	}
	t.Attest(var1 != var2, msg, fmt...)
}

// Attest that `that` is true, or log `message` and fail the test.
func (t *Test) Attest(that bool, message string, formatters ...interface{}) {
	if !that {
		if len(formatters) == 0 {
			fmt.Println(message)
		} else {
			fmt.Printf(message+"\n", formatters...)
		}
		t.Fail()
	}
}

// AttestNot -- assert that `that` is false. It just calls t.Attest(!that...
func (t *Test) AttestNot(that bool, message string, formatters ...interface{}) {
	t.Attest(!that, message, formatters...)
}

// Not does exactly the same thing that AttestNot does.
func (t *Test) Not(that bool, message string, formatters ...interface{}) {
	t.AttestNot(that, message, formatters...)
}

// AttestOrDo -- call `callback` with the Test as a parameter and fail the test
// should `that` be false.
func (t *Test) AttestOrDo(that bool,
	callback func(*Test, ...interface{}),
	cbArgs ...interface{},
) {
	if !that {
		callback(t, cbArgs...)
		t.Fail()
	}
}

// Nil -- Log a message and fail if the variable is not nil
func (t *Test) Nil(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Attest(
			variable == nil,
			"%#v was expected to be nil, but was not!",
			variable)
	} else if len(msgAndFmt) == 1 {
		t.Attest(
			variable == nil,
			msgAndFmt[0].(string))
	} else {
		t.Attest(
			variable == nil,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
}

// NotNil --  Log a message and fail if the variable is nil.
func (t *Test) NotNil(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Attest(
			variable != nil,
			"%#v was expected to have a value but instead was nil",
			variable)
	} else if len(msgAndFmt) == 1 {
		t.Attest(
			variable != nil,
			msgAndFmt[0].(string))
	} else {
		t.Attest(
			variable != nil,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
}

// GreaterThan -- log a message and fail if the variable is less than the
// expected value
func (t *Test) GreaterThan(expected,
	variable interface{},
	msgAndFmt ...interface{},
) {
	defMsg := fmt.Sprintf(
		"Value (%#v) was less than expected (%#v).",
		variable,
		expected)
	msg := func() string {
		if len(msgAndFmt) == 0 {
			return defMsg
		}
		if len(msgAndFmt) == 1 {
			return msgAndFmt[0].(string)
		}
		return fmt.Sprintf(msgAndFmt[0].(string), msgAndFmt[1:]...)
	}
	switch variable.(type) {
	default:
		log.Printf(
			"When trying check that %v was greater than %v, found non-numeric "+
				"types %T and %T",
			expected,
			variable,
			expected,
			variable)
		t.Fail()
	case int:
		t.Attest(variable.(int) > expected.(int), msg())
	case int8:
		t.Attest(variable.(int8) > expected.(int8), msg())
	case int16:
		t.Attest(variable.(int16) > expected.(int16), msg())
	case int32:
		t.Attest(variable.(int32) > expected.(int32), msg())
	case int64:
		t.Attest(variable.(int64) > expected.(int64), msg())
	case float32:
		t.Attest(variable.(float32) > expected.(float32), msg())
	case float64:
		t.Attest(variable.(float64) > expected.(float64), msg())
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// LessThan -- log a message and fail if variable is negative.
func (t *Test) LessThan(expected,
	variable interface{},
	msgAndFmt ...interface{},
) {
	defMsg := fmt.Sprintf(
		"Value (%#v) was greater than expected (%#v).",
		variable,
		expected)
	msg := func() string {
		if len(msgAndFmt) == 0 {
			return defMsg
		}
		if len(msgAndFmt) == 1 {
			return msgAndFmt[0].(string)
		}
		return fmt.Sprintf(msgAndFmt[0].(string), msgAndFmt[1:]...)
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check value of %#v: check isn't implemented for type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(variable.(int) < expected.(int), msg())
	case int8:
		t.Attest(variable.(int8) < expected.(int8), msg())
	case int16:
		t.Attest(variable.(int16) < expected.(int16), msg())
	case int32:
		t.Attest(variable.(int32) < expected.(int32), msg())
	case int64:
		t.Attest(variable.(int64) < expected.(int64), msg())
	case float32:
		t.Attest(variable.(float32) < expected.(float32), msg())
	case float64:
		t.Attest(variable.(float64) < expected.(float64), msg())
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// Positive -- log a message and fail if variable is negative or zero.
func (t *Test) Positive(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		msgAndFmt = []interface{}{"%#v was not positive", variable}
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is positive: check isn't implemented for "+
				"type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(
			variable.(int) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int8:
		t.Attest(
			variable.(int8) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int16:
		t.Attest(
			variable.(int16) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int32:
		t.Attest(
			variable.(int32) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int64:
		t.Attest(
			variable.(int64) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float32:
		t.Attest(
			variable.(float32) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float64:
		t.Attest(
			variable.(float64) > 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// Negative -- log a message and fail if variable is positive or zero.
func (t *Test) Negative(variable interface{}, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		msgAndFmt = []interface{}{"%#v was not positive", variable}
	}
	switch variable.(type) {
	default:
		log.Printf(
			"Can't check that %#v is negative: check isn't implemented for "+
				"type %T",
			variable,
			variable)
		t.Fail()
	case int:
		t.Attest(
			variable.(int) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int8:
		t.Attest(
			variable.(int8) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int16:
		t.Attest(
			variable.(int16) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int32:
		t.Attest(
			variable.(int32) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case int64:
		t.Attest(
			variable.(int64) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float32:
		t.Attest(
			variable.(float32) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	case float64:
		t.Attest(
			variable.(float64) < 0,
			msgAndFmt[0].(string),
			msgAndFmt[1:]...)
	}
	// can't use > on complex numbers for some reason.
	// FIXME: implement GT/LT for complex64 and complex128
}

// TypeIs fails the test if the type of the value does not match the typestring,
// as determined by fmt.Sprintf("%T"). For example, a "Test" struct from the
// "attest" package (this one), would have the type "attest.Test".
func (t *Test) TypeIs(typestring string, value interface{}) {
	if fmt.Sprintf("%T", value) == typestring {
		return
	}
	t.Errorf("Type of %#v wasn't %s.", value, typestring)
}

// TypeIsNot is the inverse of TypeIs; it fails the test if the type of value
// matches the typestring.
func (t *Test) TypeIsNot(typestring string, value interface{}) {
	if fmt.Sprintf("%T", value) == typestring {
		t.Errorf("Type of %#v was %s.", value, typestring)
	}
}

/* 						ERROR HANDLING TESTS

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
func (t *Test) Handle(e ...error) {
	for _, err := range e {
		if err != nil {
			log.Println(err)
			t.Fail()
		}
	}
}

// MessageHandle -- handle an error with a custom message.
func (t *Test) MessageHandle(err error, msgAndFmt ...interface{}) {
	if len(msgAndFmt) == 0 {
		t.Handle(err)
	} else {
		if err != nil {
			log.Printf(msgAndFmt[0].(string), msgAndFmt[1:]...)
			t.Fail()
		}
	}
}

// StopIf -- Fail the test and stop running it if an error is present, with
// optional message.
func (t *Test) StopIf(err error, msgAndFmt ...interface{}) {
	if err != nil {
		if len(msgAndFmt) == 0 {
			msgAndFmt = []interface{}{"Fatal error: %#v", err}
		}
		log.Printf(msgAndFmt[0].(string), msgAndFmt[1:]...)
		t.FailNow()
	}

}

// EatError accepts two values, the latter of which is a nillable error. If the
// error is not nil, the test is failed. Regardless, the first value is
// returned through the function. The returned value must be asserted to have
// its type in order to convert it from the generic interface{} type.
func (t *Test) EatError(value interface{}, err error) interface{} {
	if err != nil {
		t.Errorf("When aquiring value %#v, got error %#v", value, err)
	}
	return value
}
