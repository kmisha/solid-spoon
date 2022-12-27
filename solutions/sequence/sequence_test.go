package sequence

import "testing"

func TestSolution(t *testing.T) {
	example := []int{2, 3, 1, 5}
	want := 4

	got := Solution(example)

	if got != want {
		t.Errorf("wrong anwser; got %d; want %d", got, want)
	}
}
