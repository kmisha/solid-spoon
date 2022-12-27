package rotate

import "testing"

func TestRotate(t *testing.T) {
	t.Run("the first example", func(t *testing.T) {
		example := []int{3, 8, 9, 7, 6}
		want := []int{9, 7, 6, 3, 8}
		K := 3

		got := Solution(example, K)

		for i, n := range want {
			if n != got[i] {
				t.Errorf("wront answer; got %v; want %v", got, want)
			}
		}
	})
	t.Run("the first example", func(t *testing.T) {
		example := []int{1, 2, 3, 4}
		want := []int{1, 2, 3, 4}
		K := 4

		got := Solution(example, K)

		for i, n := range want {
			if n != got[i] {
				t.Errorf("wront answer; got %v; want %v", got, want)
			}
		}
	})

}
