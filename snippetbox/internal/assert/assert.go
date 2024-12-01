package assert

import "testing"

func Equal[T comparable](t *testing.T, actual, expected T) {
	// indicates to go test runner that Equal function is a test helper
	// go test runner will report the filename and line number of the code which called our Equal function
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
