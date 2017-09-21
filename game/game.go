package game

import (
	"fmt"
	"strconv"
	"strings"
)

// goal is the goal board configuration.
// See SetGoal for further explanation.
var goal Board

// n represents the n*n dimensions of the board.
var n int

// SetGoal sets the goal configuration.
// The value indicates the position the number at the index
// should hold.
//
// e.g: nGoal[1] = 6 means that one should be at index six.
func SetGoal(nGoal []int) {
	goal = nGoal
	setN()
}

// SetGoalBoard takes the input board and converts it to the goal configuration.
// Then it sets the goal configuration.
//
// e.g: goalBoard[1] = 6 means that six should be at index one.
func SetGoalBoard(goalBoard []int) {
	goal = make(Board, 9)
	for i, n := range goalBoard {
		goal[n] = i
	}
	setN()
}

func setN() {
	switch len(goal) {
	case 9:
		n = 3
	case 16:
		n = 4
	default:
		panic(fmt.Errorf("game: Board of length %d not implemented", len(goal)))
	}
}

// State is a state of the 8-puzzle game.
type State struct {
	// board is the 3x3 puzzle board.
	board  Board
	origin *State
	cost   int
	move   Move
}

// NewState returns a State pointer for the given board. It should be used to
// create the start state.
func NewState(board []int) State {
	return State{board: board}
}

// Board returns the board at the current state.
func (s State) Board() Board {
	return s.board
}

// Cost returns the cost to get to the current state.
func (s State) Cost() int {
	return s.cost
}

// IsGoal returns whether the state is the goal state.
func (s State) IsGoal() bool {
	for i, v := range s.board {
		if i != goal[v] {
			return false
		}
	}
	return true
}

// Moves returns the moves made to get to the state. It also returns
// each board configuration.
func (s State) Moves() ([]Move, []string) {
	moves := make([]Move, s.cost)
	configs := make([]string, s.cost)
	if s.cost < 1 {
		return moves, configs
	}
	moves[s.cost-1] = s.move
	configs[s.cost-1] = s.board.String()
	for i := s.cost - 2; ; i-- {
		s = *s.origin
		if s.origin == nil {
			return moves, configs
		}
		if i < 0 {
			panic(fmt.Errorf("game: more moves than cost, for state %v", s))
		}
		moves[i] = s.move
		configs[i] = s.board.String()
	}
}

// Neighbors returns the States that can be reached by doing
// one move from this State.
func (s State) Neighbors() []State {
	// find 0
	var movable int
	for i, n := range s.board {
		if n == 0 {
			movable = i
		}
	}
	// find potential moves
	neighbors := make([]State, 0, 4)
	col := movable % n
	row := movable / n
	if col > 0 {
		// move left
		nBoard := make(Board, len(goal))
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable-1)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Left})
	}
	if col < n-1 {
		// move right
		nBoard := make(Board, len(goal))
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable+1)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Right})
	}
	if row > 0 {
		// move up
		nBoard := make(Board, len(goal))
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable-n)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Up})
	}
	if row < n-1 {
		// move down
		nBoard := make(Board, len(goal))
		copy(nBoard, s.board)
		nBoard = nBoard.swap(movable, movable+n)
		neighbors = append(neighbors, State{board: nBoard, cost: s.cost + 1, origin: &s, move: Down})
	}
	return neighbors
}

// Heuristic returns the manhattan distance to the goal state.
//
// It is calculated by adding the distance from each number to its goal.
// This is an underestimate of the actual cost and that makes the algorithm
// complete.
func (s State) Heuristic() int {
	sum := 0
	for i, v := range s.board {
		if v == 0 {
			// skip the movable piece
			continue
		}
		goalRow := goal[v] % n
		goalCol := goal[v] / n
		actualRow := i % n
		actualCol := i / n
		sum += abs(actualRow-goalRow) + abs(actualCol-goalCol)
	}
	return sum
}

// Key returns a string representation of the state's board.
func (s State) Key() string {
	key := ""
	for _, n := range s.board {
		key += strconv.Itoa(n)
	}
	return key
}

// Move is the possible move of the zero puzzle tile
type Move string

const (
	Up    Move = "Up"
	Down  Move = "Down"
	Left  Move = "Left"
	Right Move = "Right"
	Goal  Move = "Goal"
)

// Equals returns whether m and m2 are the same.
func (m Move) Equals(m2 Move) bool {
	return strings.Compare(string(m), string(m2)) == 0
}

func (m Move) String() string {
	return string(m)
}

// Board is the puzzle board.
type Board []int

// swap swaps the items at index i and j and returns the new board.
func (b Board) swap(i, j int) Board {
	b[i], b[j] = b[j], b[i]
	return b
}

// Solvable returns whether a board is solvable.
// It counts the number of inversions and if it
// is even the board is solvable.
func (b Board) Solvable() bool {
	invGoal := 0
	for i := range goal {
		if goal[i] == 0 {
			continue
		}
		for j := i + 1; j < len(goal); j++ {
			if goal[i] > goal[j] && goal[j] != 0 {
				invGoal++
			}
		}
	}
	invBoard := 0
	for i := range b {
		if b[i] == 0 {
			continue
		}
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] && b[j] != 0 {
				invBoard++
			}
		}
	}
	return invGoal%2 == invBoard%2
}

func (b Board) String() string {
	str := ""
	for i, v := range b {
		if i != 0 && i%n == 0 {
			str += "\n"
		}
		str += strconv.Itoa(v) + " "
	}
	return str
}

// abs returns the absolute value of x.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
