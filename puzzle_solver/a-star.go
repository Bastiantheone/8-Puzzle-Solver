package puzzle_solver

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver/internal/heap"
)

func Solve(start *State) []Move {
	var states map[string]*State
	h := heap.New()
	h.Push(start, 0)
	states[start.key()] = start
	for !h.IsEmpty() {
		current := h.Pop()
		if current.isGoal() {
			return current.Moves()
		}
		for _, next := range current.neighbors() {
			newCost := current.cost + 1
			key := next.key()
			// update if a better way to a state is found.
			if old, exists := states[key]; !exists || newCost < old.cost {
				next.cost = newCost
				next.origin = current
				states[key] = next
				priority := newCost + next.heuristic()
				h.Push(next, priority)
			}
		}
	}
	return nil
}
