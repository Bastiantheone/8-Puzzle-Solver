package eightPuzzleSolver

import (
	"math/rand"
	"testing"

	"github.com/Bastiantheone/8-Puzzle-Solver/game"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		start game.State
		want  []game.Move
	}{
		{game.NewState([]int{1, 0, 2, 6, 7, 8, 3, 4, 5}), []game.Move{game.Left, game.Goal}},
		{game.NewState([]int{1, 7, 2, 6, 0, 8, 3, 4, 5}), []game.Move{game.Up, game.Left, game.Goal}},
		{game.NewState([]int{0, 1, 2, 6, 8, 7, 3, 4, 5}), nil},
		{game.NewState([]int{7, 2, 4, 5, 0, 6, 8, 1, 3}), []game.Move{game.Right, game.Up, game.Left, game.Down, game.Left, game.Down,
			game.Right, game.Right, game.Up, game.Left, game.Left, game.Down, game.Right, game.Right, game.Up, game.Left, game.Left, game.Up,
			game.Right, game.Down, game.Left, game.Up, game.Goal}},
	}
	game.SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got, _ := Solve(test.start)
		if len(got) != len(test.want) {
			t.Fatalf("test %d: got = %d moves, want = %d", i, len(got), len(test.want))
		}
		for j := range got {
			if !got[j].Equals(test.want[j]) {
				t.Errorf("test %d: got[%d] = %s, want[%d] = %s", i, j, got[j], j, test.want[j])
			}
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		// create a random board
		var board []int
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
		start := game.NewState(board)
		b.StartTimer()
		game.SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
		Solve(start)
	}
}
