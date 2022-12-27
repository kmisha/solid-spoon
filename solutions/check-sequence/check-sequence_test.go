package checksequence

import "testing"

func TestCheckSequence(t *testing.T) {
	t.Run("it is a sequence", func(t *testing.T) {
		example := []int{4, 1, 3, 2}
		want := 1

		got := Solution(example)

		if got != want {
			t.Errorf("wrong answer; got %d; want %d", got, want)
		}
	})

	t.Run("it is not a sequence", func(t *testing.T) {
		example := []int{4, 1, 3}
		want := 0

		got := Solution(example)

		if got != want {
			t.Errorf("wrong answer; got %d; want %d", got, want)
		}
	})
}
