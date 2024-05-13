package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("test add", func(t *testing.T) {
		t.Parallel()

		a, b, expected := 1, 3, 4
		res := add(a, b)

		if(res != expected) {
			t.Fatalf("Expected: %d Got: %d", expected, res)
		}
	})
}
