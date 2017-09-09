package puzzle_solver

import (
	"math/rand"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		start State
		want  []Move
	}{
		{State{board: board{1, 0, 2, 3, 4, 5, 6, 7, 8}}, []Move{Left}},
		{State{board: board{1, 4, 2, 3, 0, 5, 6, 7, 8}}, []Move{Up, Left}},
		{State{board: board{0, 1, 2, 3, 4, 5, 6, 8, 7}}, nil},
	}
	for i, test := range tests {
		got := Solve(test.start)
		if len(got) != len(test.want) {
			t.Fatalf("test %d: got = %d moves, want = %d", i, len(got), len(test.want))
		}
		for j := range got {
			if !got[j].equals(test.want[j]) {
				t.Errorf("test %d: got[%d] = %s, want[%d] = %s", i, j, got[j], j, test.want[j])
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// create a random board
		var board board
	outer:
		for c := 0; c < 9; {
			n := rand.Intn(9)
			for _, old := range board {
				if n == old {
					continue outer
				}
			}
			board = append(board, n)
			c++
		}
		start := State{board: board}
		b.StartTimer()
		Solve(start)
	}
}
