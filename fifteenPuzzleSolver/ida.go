package fifteenPuzzleSolver

import (
	"math"

	"github.com/Bastiantheone/8-Puzzle-Solver/game"
)

// max is the maximum number of moves required to solve the fifteen puzzle problem.
const max = 80

func Solve(start game.State) ([]game.Move, []string) {
	configs := make([]string, 1)
	configs[0] = start.Board().String()
	if !start.Board().Solvable() {
		return nil, configs
	}
	threshold := start.Heuristic()
	for {
		node, nrSteps := search(start, threshold)
		if node.IsGoal() {
			moves, temp := node.Moves()
			return append(moves, game.Goal), append(configs, temp...)
		}
		if nrSteps > max {
			return nil, configs
		}
		threshold = nrSteps
	}
}

func search(state game.State, threshold int) (game.State, int) {
	if state.IsGoal() {
		return state, 0
	}
	f := state.Cost() + state.Heuristic()
	if f > threshold {
		return state, f
	}
	min := math.MaxInt16
	var next game.State
	for _, neighbor := range state.Neighbors() {
		// avoid cycles
		path := neighbor.Path()
		if path[neighbor.Key()] {
			continue
		}
		s, h := search(neighbor, threshold)
		if s.IsGoal() {
			return s, h
		}
		if h < min {
			min = h
			next = s
		}
	}
	return next, min
}
