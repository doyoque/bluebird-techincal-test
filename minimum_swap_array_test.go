package bluebird

import "testing"

func TestMinSwapArray(t *testing.T) {
	want1 := 5
	got1 := MinSwapArray([]int{7, 1, 3, 2, 4, 5, 6})

	if got1 != want1 {
		t.Errorf("MinSwapArray() got %v, want %v", got1, want1)
	}
}
