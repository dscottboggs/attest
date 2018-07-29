# attest
A small library to make go tests more readable.

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
