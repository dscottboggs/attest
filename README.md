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
        test.AttestEquals(val1, val2)
        val2 = 1
        test.AttestGreater(val2, val1)
    }

This package currently has 74.6% test coverage according to go test. The
implicit tests I use to perform testing on this package also serves as a great
set of examples its use.
