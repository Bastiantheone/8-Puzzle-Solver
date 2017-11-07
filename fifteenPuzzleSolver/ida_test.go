package fifteenPuzzleSolver

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
		{game.NewState([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}), []game.Move{game.Right, game.Goal}},
		{game.NewState([]int{13, 10, 11, 6, 5, 7, 4, 8, 1, 12, 14, 9, 3, 15, 2, 0}), nil},
		{game.NewState([]int{1, 2, 3, 4, 5, 6, 0, 8, 9, 10, 7, 11, 13, 14, 15, 12}), []game.Move{game.Down, game.Right, game.Down, game.Goal}},
		/*{game.NewState([]int{0, 11, 5, 12, 2, 3, 6, 4, 7, 10, 14, 8, 9, 1, 15, 13}), []game.Move{game.Right, game.Right, game.Right, game.Down, game.Down, game.Left,
		game.Down, game.Right, game.Up, game.Left, game.Left, game.Down, game.Right, game.Up, game.Up, game.Up, game.Up, game.Left, game.Left, game.Down, game.Down,
		game.Right, game.Right, game.Up, game.Left, game.Up, game.Right, game.Down, game.Left, game.Left, game.Down, game.Down, game.Right, game.Up, game.Left, game.Up,
		game.Left, game.Down, game.Right, game.Down, game.Down, game.Right, game.Right, game.Goal}},*/
		/*{game.NewState([]int{15, 14, 1, 6, 9, 11, 4, 12, 0, 10, 7, 3, 13, 8, 5, 2}),[]game.Move{game.Right, game.Up, game.Up, game.Right, game.Down, game.Down, game.Right,
		game.Down, game.Left, game.Left, game.Up, game.Up, game.Right, game.Down, game.Right, game.Down, game.Left, game.Left, game.Up, game.Right,game.Right,
		game.Up, game.Up, game.Left, game.Down, game.Left, game.Down, game.Left, game.Up, game.Up, game.Right, game.Down, game.Left, game.Down, game.Right, game.Up,
		game.Right, game.Down, game.Right, game.Down, game.Left, game.Up, game.Left, game.Up, game.Right, game.Right, game.Down, game.Down, game.Left, game.Up, game.Left,
		game.Left, game.Up, game.Right, game.Right, game.Down, game.Right, game.Down, game.Goal}},*/
		/*{game.NewState([]int{8,3,2,4,13,5,7,11,9,15,1,12,6,14,10,0}), []game.Move{game.Up, game.Up, game.Up, game.Left, game.Left, game.Left, game.Down, game.Down, game.Down,
		game.Right, game.Up, game.Right, game.Up, game.Up, game.Left, game.Down, game.Down, game.Left, game.Up, game.Up, game.Right, game.Down, game.Right, game.Up,
		game.Right, game.Down, game.Left, game.Left, game.Down, game.Left, game.Up, game.Up, game.Right, game.Right, game.Down, game.Down, game.Down, game.Left, game.Left,
		game.Up, game.Right, game.Right, game.Right, game.Down, game.Goal}},*/
	}
	game.SetGoalBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0})
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
		for c := 0; c < 16; {
			n := rand.Intn(16)
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
		game.SetGoalBoard([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0})
		Solve(start)
	}
}
