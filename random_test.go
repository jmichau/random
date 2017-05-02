package random

import "testing"

func TestGenerate(t *testing.T) {
	values := make(map[string]bool)
	num_tests := 10

	for i := 0; i < num_tests; i++ {
		randString := Generate(10)
		values[randString] = true
	}

	if len(values) != num_tests {
		t.Error("Strings are not generating randomly")
	}
}
