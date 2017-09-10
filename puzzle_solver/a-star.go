package puzzle_solver

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver/internal/heap"
)

// Solve returns the moves to get to the goal state the fastest.
//
// It returns nil if there is no solution.
// It uses the A* star algorithm to achieve that goal.
func Solve(start State) []Move {
	states := make(map[string]State)
	h := heap.New()
	h.Push(start.key(), 0)
	states[start.key()] = start
	for !h.IsEmpty() {
		currentKey := h.Pop()
		current := states[currentKey]
		if current.isGoal() {
			return append(current.moves(), Goal)
		}
		for _, next := range current.neighbors() {
			key := next.key()
			// update if a better way to a state is found.
			if old, exists := states[key]; !exists || next.cost < old.cost {
				states[key] = next
				priority := next.cost + next.heuristic()
				h.Push(next.key(), priority)
			}
		}
	}
	return nil
}
