package puzzle_solver

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver/internal/heap"
)

// Solve returns the moves and the different board configurations to get to the goal state the fastest.
//
// It returns nil if there is no solution.
// It uses the A* star algorithm to achieve that goal.
func Solve(start State) ([]Move, []string) {
	configs := make([]string, 1)
	configs[0] = start.board.String()
	if !start.board.solvable() {
		return nil, configs
	}
	states := make(map[string]State)
	h := heap.New()
	h.Push(start.key(), 0)
	states[start.key()] = start
	for !h.IsEmpty() {
		currentKey := h.Pop()
		current := states[currentKey]
		if current.isGoal() {
			moves, temp := current.moves()
			return append(moves, Goal), append(configs, temp...)
		}
		for _, next := range current.neighbors() {
			key := next.key()
			// update if a better way to a state is found.
			if old, exists := states[key]; !exists || next.cost < old.cost {
				states[key] = next
				priority := next.cost + next.heuristic()
				h.Push(key, priority)
			}
		}
	}
	return nil, configs
}
