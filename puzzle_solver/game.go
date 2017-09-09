package puzzle_solver

import (
	"fmt"
	"strconv"
)

type State struct {
	// board is the 3x3 puzzle board.
	board  board
	origin *State
	cost   int
	move   Move
}

// isGoal return whether the state is the goal state.
func (s State) isGoal() bool {
	for i, n := range s.board {
		if i != n {
			return false
		}
	}
	return true
}

// Moves returns the moves made to get to the state.
func (s State) Moves() []Move {
	moves := make([]Move, s.cost)
	if s.cost < 1 {
		return moves
	}
	moves[s.cost-1] = s.move
	for i := s.cost - 2; ; i-- {
		s = *s.origin
		if s.origin == nil {
			return moves
		}
		if i < 0 {
			panic(fmt.Errorf("puzzle_solver: more moves than cost, for state %v", s))
		}
		moves[i] = s.move
	}
}

// neighbors returns the States that can be reached by doing
// one move from this State.
func (s State) neighbors() []State {
	// find 0
	var movable int
	for i, n := range s.board {
		if n == 0 {
			movable = i
		}
	}
	// find potential moves
	neighbors := make([]State, 0, 4)
	col := movable % 3
	row := movable / 3
	if col > 0 {
		// move left
		nBoard := make(board, 9)
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable-1)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Left})
	}
	if col < 2 {
		// move right
		nBoard := make(board, 9)
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable+1)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Right})
	}
	if row > 0 {
		// move up
		nBoard := make(board, 9)
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable-3)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Up})
	}
	if row < 2 {
		// move down
		nBoard := make(board, 9)
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable+3)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Down})
	}
	return neighbors
}

// heuristic returns the manhattan distance to the goal state.
//
// It is calculated by adding the distance from each number to its goal.
// This is an underestimate of the actual cost and that makes the algorithm
// complete.
func (s State) heuristic() int {
	sum := 0
	for i, n := range s.board {
		if n == 0 {
			// skip the movable piece
			continue
		}
		goalRow := n % 3
		goalCol := n / 3
		actualRow := i % 3
		actualCol := i / 3
		sum += abs(actualRow-goalRow) + abs(actualCol-goalCol)
	}
	return sum
}

// key returns a string representation of the state's board.
func (s State) key() string {
	key := ""
	for _, n := range s.board {
		key += strconv.Itoa(n)
	}
	return key
}

type Move string

const (
	Up    Move = "Up"
	Down  Move = "Down"
	Left  Move = "Left"
	Right Move = "Right"
)

// board is the 3x3 puzzle board.
type board []int

// swap swaps the items at index i and j and returns the new board.
func (b board) swap(i, j int) board {
	b[i], b[j] = b[j], b[i]
	return b
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
