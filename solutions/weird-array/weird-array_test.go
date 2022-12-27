package weirdarray

import "testing"

func TestSolution(t *testing.T) {
	example := []int{9, 3, 9, 3, 9, 7, 9}
	want := 7

	got := Solution(example)

	if got != want {
		t.Errorf("wrong answer; got %d; want %d", got, want)
	}
}
