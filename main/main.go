package main

import (
	"fmt"
	"os"

	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver"
)

// main reads the file given as command line argument. It extracts the board from the file
// and solves it using the A* algorithm with Manhattan distance as heuristic. The program
// assumes the input is correct. The result is printed on the command line.
func main() {
	args := os.Args
	if len(args) != 2 {
		panic(fmt.Errorf("puzzle_solver: got %d arguments, need 2", len(args)))
	}
	board, err := puzzle_solver.Read(args[1])
	if err != nil {
		panic(err)
	}
	start := puzzle_solver.NewState(board)
	moves := puzzle_solver.Solve(start)
	if moves == nil {
		fmt.Println("No Solution")
		return
	}
	output := "Solution: "
	for _, move := range moves {
		output += move.String()
	}
	fmt.Println(output)
}
