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
		{State{board: board{1, 0, 2, 6, 7, 8, 3, 4, 5}}, []Move{Left, Goal}},
		{State{board: board{1, 7, 2, 6, 0, 8, 3, 4, 5}}, []Move{Up, Left, Goal}},
		{State{board: board{0, 1, 2, 6, 8, 7, 3, 4, 5}}, nil},
		{State{board: board{7, 2, 4, 5, 0, 6, 8, 1, 3}}, []Move{Right, Up, Left, Down, Left, Down,
			Right, Right, Up, Left, Left, Down, Right, Right, Up, Left, Left, Up, Right, Down, Left, Up, Goal}},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got, _ := Solve(test.start)
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
		SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
		Solve(start)
	}
}
