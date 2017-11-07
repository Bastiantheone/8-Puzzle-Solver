package eightPuzzleSolver

import (
	"math/rand"
	"testing"

	"github.com/Bastiantheone/8-Puzzle-Solver/game"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		start game.State
		goal  []int
		want  []game.Move
	}{
		{game.NewState([]int{1, 0, 2, 6, 7, 8, 3, 4, 5}), []int{0, 1, 2, 6, 7, 8, 3, 4, 5}, []game.Move{game.Left, game.Goal}},
		{game.NewState([]int{1, 7, 2, 6, 0, 8, 3, 4, 5}), []int{0, 1, 2, 6, 7, 8, 3, 4, 5}, []game.Move{game.Up, game.Left, game.Goal}},
		{game.NewState([]int{0, 1, 2, 6, 8, 7, 3, 4, 5}), []int{0, 1, 2, 6, 7, 8, 3, 4, 5}, nil},
		{game.NewState([]int{7, 2, 4, 5, 0, 6, 8, 1, 3}), []int{0, 1, 2, 6, 7, 8, 3, 4, 5}, []game.Move{game.Right, game.Up, game.Left, game.Down, game.Left, game.Down,
			game.Right, game.Right, game.Up, game.Left, game.Left, game.Down, game.Right, game.Right, game.Up, game.Left, game.Left, game.Up,
			game.Right, game.Down, game.Left, game.Up, game.Goal}},
		{game.NewState([]int{2, 8, 3, 1, 6, 4, 7, 0, 5}), []int{1, 2, 3, 8, 0, 4, 7, 6, 5}, []game.Move{game.Up, game.Up, game.Left, game.Down, game.Right, game.Goal}},
		{game.NewState([]int{3, 7, 8, 1, 6, 4, 0, 2, 5}), []int{1, 2, 3, 4, 5, 6, 7, 8, 0}, []game.Move{game.Right, game.Up, game.Up, game.Left, game.Down, game.Right,
			game.Right, game.Down, game.Left, game.Left, game.Up, game.Right, game.Down, game.Right, game.Up, game.Up, game.Left, game.Down, game.Down, game.Right,
			game.Up, game.Left, game.Down, game.Right, game.Goal}},
		{game.NewState([]int{1, 8, 0, 3, 2, 5, 7, 4, 6}), []int{1, 2, 3, 4, 5, 6, 7, 8, 0}, []game.Move{game.Left, game.Down, game.Left, game.Up, game.Right, game.Down,
			game.Down, game.Right, game.Up, game.Up, game.Left, game.Left, game.Down, game.Right, game.Down, game.Right, game.Up, game.Left, game.Down, game.Right, game.Goal}},
		{game.NewState([]int{7, 6, 5, 8, 2, 3, 4, 1, 0}), []int{1, 2, 3, 4, 5, 6, 7, 8, 0}, []game.Move{game.Up, game.Left, game.Down, game.Left, game.Up, game.Right, game.Up,
			game.Left, game.Down, game.Down, game.Right, game.Up, game.Up, game.Right, game.Down, game.Down, game.Left, game.Up, game.Left, game.Down, game.Right, game.Up,
			game.Up, game.Right, game.Down, game.Down, game.Goal}},
		{game.NewState([]int{5, 8, 4, 7, 1, 2, 0, 6, 3}), []int{1, 2, 3, 4, 5, 6, 7, 8, 0}, []game.Move{game.Up, game.Right, game.Up, game.Right, game.Down, game.Down, game.Left,
			game.Up, game.Up, game.Left, game.Down, game.Right, game.Up, game.Right, game.Down, game.Down, game.Goal}},
	}
	for i, test := range tests {
		game.SetGoalBoard(test.goal)
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
		game.SetGoalBoard([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
		Solve(start)
	}
}
