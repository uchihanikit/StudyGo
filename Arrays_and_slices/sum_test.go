package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of different numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("want %q, want %q, given %v", got, want, numbers)
		}
	})
}
