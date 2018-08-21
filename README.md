# attest
An ever-growing list of assertions that make Go tests more readable and concise.

## Installation:
Install Go, from your operating system's package manager or from the [Golang website](https://golang.org/dl/).

Open a terminal/console window and run

`go get github.com/dscottboggs/attest`

done!

## Usage:
    package main

    import (
        "testing"

        "github.com/dscottboggs/attest"
    )

    func TestExample(t \*testing.T) {
        test := attest.Test{t}
        test.Attest(fmt.Sprintf("%T", "that something is true") == "string", "or %s a message", "log")
        var val1, val2 int
        test.Equals(val1, val2)
        val2 = 1
        test.Greater(val2, val1)
    }

The implicit tests I use to perform testing on this package also serves as a great
set of examples of its use. Unfortunately due to limitations in the Go testing
methods, it's not really possible to test for failure cases, as they would fail
the test. Recommendations are welcome for additional testing methodologies.

### Contributing
Have a clever function that makes testing easier in Go? Submit a pull request or open an issue and let's discuss it!

## About

Attest is a very lightweight testing library aimed at improving the
 intuitiveness and aesthetics of the go standard testing library, as well as
 reducing the amount of keystrokes per test, hence improving developer
 efficiency by a marginal amount with very minimal overhead and risk. Any
 given testing.T function can create an attest.Test object, whose methods
 can then perform tests.

 A brief example:

     package main

        import (
            "testing"

            "github.com/dscottboggs/attest"
        )

        func TestExample(t \*testing.T) {
            test := attest.NewTest(t)
            test.Attest(fmt.Sprintf("%T", "that something is true") == "string", "or %s a message", "log")
            const unchanging = 0
            var variable int
            test.Equals(unchanging, variable)
            variable = 1
            test.Greater(variable, unchanging)
        }

 That same test function with the default testing library might be written
 like:

        func TestExample(t \*testing.T) {
            if fmt.Sprintf("%T", "something is true") != "string" {
                t.Errorf("or %s a message", "log")
            }
            const unchanging = 0
            var variable int
            if fmt.Sprintf("%T", unchanging) != fmt.Sprintf("%T", variable) {
                t.Errorf("Value 1 had a different type (%T) than value 2 (%T)", unchanging, variable)
            }
            if unchanging != variable {
                t.Errorf("Value 1 (%d) didn't equal value 2 (%d).")
            }
            variable = 1
            if fmt.Sprintf("%T", unchanging) != fmt.Sprintf("%T", variable) {
                t.Errorf("Value 1 had a different type (%T) than value 2 (%T)", unchanging, variable)
            }
            if variable <= unchanging {
                t.Errorf("Value 1 was less than or equal to value 2")
            }
 	  }

 As you can see, this provides minimal benefit besides a reduced number of
 keystrokes when *writing*, but when reading back, the attest way is much
 easier to understand. Of course, you can mix-and-match:

     func TestExample(t \*testing.T){
         test := attest.New(t)
         if fmt.Sprintf("%T", "something is true") != "string" {
             test.Errorf("or %s a message", "log")
         }
         const unchanging = 0
         var variable int
         test.Equals(unchanging, variable)
         variable = 1
         test.GreaterThan(unchanging, variable)
     }

 ### Logging a custom message
 All tests allow for an optional (or in the case of the few strictly boolean
 tests, required) message string and formatters to be forwarded to
 fmt.Sprintf()

 # Available test functions
 The following tests are available:

  - **Attest** and **That**: the first argument must equal the boolean value true.
  - **AttestNot** and **Not**: the first argument must equal the boolean value false.
  - **AttestOrDo**: takes a callback function and arguments to forward to the callback in case of a failure
  - **Nil** and **NotNil**: the first argument must be nil or not nil, respectively.
  - **Equals** and **NotEqual**: the second argument must equal (or not equal, respectively) the first argument. Both require that the arguments be the same type
  - **Compares**, **SimilarTo**, **DoesNotCompare**, and **NotSimilarTo**: like Equals and NotEquals but the types don't have to be the same.
  - **GreaterThan** and **LessThan**: like Equals, but checks for the second value to be greater or less than the first argument.
  - **Positive** and **Negative**: are shortcuts for test.LessThan(0, ...) and test.GreaterThan(0, ...)
  - **TypeIs** and **TypeIsNot**: check the type of a value
  - **Matches** and **DoesNotMatch**: Check if the value matches a given regular expression.

 In addition there are the following ways of handling error types and panics:
  - **Handle**: Log and fail if the first argument is a non-nil error.
  - **HandleMultiple**: Log and fail if any of the arguments to this are non-nil errors. Does not accept a callback or message.
  - **AttestPanics** and **AttestNoPanic**: ensure the given function panics or doesn't.
  - **StopIf**: Log and fail a fatal non-nil error
  - **EatError**: Logs and fails an error message if the second argument is a non-nil error, and returns the first argument. For handling function calls that return a value and an error in a single line.
