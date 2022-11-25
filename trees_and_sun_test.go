package bluebird

import (
	"testing"
)

func TestGrabTrees(t *testing.T) {
	want1 := 3
	got1 := GrabTrees([]int{7, 4, 8, 2, 9}, 5)

	if got1 != want1 {
		t.Errorf("GrabTrees() got %v, want %v", got1, want1)
	}
}
