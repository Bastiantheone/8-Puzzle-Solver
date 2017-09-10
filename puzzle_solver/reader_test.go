package puzzle_solver

import (
	"testing"
)

func TestRead(t *testing.T) {
	path := "testdata\\puzzle1.txt"
	got, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	if len(want) != len(got) {
		t.Fatalf("got = %d ints, want = %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %d, want[%d] = %d", i, got[i], i, want[i])
		}
	}
}
