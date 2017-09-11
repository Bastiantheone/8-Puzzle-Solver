package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver"
)

// main reads the file given as command line argument. It extracts the board from the file
// given as the second command line argument and solves it using the A* algorithm with Manhattan
// distance as heuristic. The third command line argument is optional and can be a file that
// contains the goal state. The default goal configuration is:
// 	|0 1 2|
// 	|6 7 8|
//	|3 4 5|
// The program assumes the input is correct. The result is printed on the command line.
func main() {
	args := os.Args
	if len(args) != 2 && len(args) != 3 {
		panic(fmt.Errorf("puzzle_solver: got %d arguments, need 2 or 3", len(args)))
	}
	board, err := puzzle_solver.Read(args[1])
	if err != nil {
		panic(err)
	}
	if len(args) == 3 {
		goal, err := puzzle_solver.Read(args[2])
		if err != nil {
			panic(err)
		}
		puzzle_solver.SetGoalBoard(goal)
	} else {
		puzzle_solver.SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	}
	start := puzzle_solver.NewState(board)
	moves, configs := puzzle_solver.Solve(start)
	if moves == nil {
		fmt.Println("No Solution")
		fmt.Println(configs[0])
		return
	}
	output := "Solution: \n"
	for i, move := range moves {
		output += move.String() + "\n" + configs[i] + "\n"
	}
	output += strconv.Itoa(len(moves)-1) + " Steps"
	fmt.Println(output)
}
