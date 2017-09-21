package eightPuzzleSolver

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/game"
	"github.com/Bastiantheone/8-Puzzle-Solver/heap"
)

// Solve returns the moves and the different board configurations to get to the goal state the fastest.
//
// It returns nil if there is no solution.
// It uses the A* star algorithm to achieve that goal.
func Solve(start game.State) ([]game.Move, []string) {
	configs := make([]string, 1)
	configs[0] = start.Board().String()
	if !start.Board().Solvable() {
		return nil, configs
	}
	states := make(map[string]game.State)
	h := heap.New()
	h.Push(start.Key(), 0)
	states[start.Key()] = start
	for !h.IsEmpty() {
		currentKey := h.Pop()
		current := states[currentKey]
		if current.IsGoal() {
			moves, temp := current.Moves()
			return append(moves, game.Goal), append(configs, temp...)
		}
		for _, next := range current.Neighbors() {
			key := next.Key()
			// update if a better way to a state is found.
			if old, exists := states[key]; !exists || next.Cost() < old.Cost() {
				states[key] = next
				priority := next.Cost() + next.Heuristic()
				h.Push(key, priority)
			}
		}
	}
	return nil, configs
}
