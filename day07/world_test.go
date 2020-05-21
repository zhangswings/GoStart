package day07

import "testing"

/**
https://stackoverflow.com/questions/28240489/golang-testing-no-test-files/28240537

Files containing tests should be called name_test, with the _test suffix. They should be alongside the code that they are testing.

To run the tests recursively call go test -v ./...

From How to Write Go Code:

You write a test by creating a file with a name ending in _test.go that contains functions named TestXXX with signature func (t *testing.T).
The test framework runs each such function; if the function calls a failure function such as t.Error or t.Fail, the test is considered to have failed.
*/
func TestAllocsPerRun(t *testing.T) {
	t.Log("hello")
}
