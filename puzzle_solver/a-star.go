package puzzle_solver

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver/internal/heap"
)

func Solve(start State) []Move {
	var states map[string]State
	h := heap.New()
	h.Push(start.key(), 0)
	states[start.key()] = start
	for !h.IsEmpty() {
		currentKey := h.Pop()
		current := states[currentKey]
		if current.isGoal() {
			return current.Moves()
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
