package random

import "testing"

func TestGenerate(t *testing.T) {
	values := make(map[string]bool)
	numTests := 10

	for i := 0; i < numTests; i++ {
		randString := Generate(10)
		values[randString] = true
	}

	if len(values) != numTests {
		t.Error("Strings are not generating randomly")
	}
}
