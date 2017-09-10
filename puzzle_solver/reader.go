package puzzle_solver

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// Read reads the file and returns all integers seperated by white space
// of the file as a slice.
func Read(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return readInts(io.Reader(f))
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, n)
	}
	return result, scanner.Err()
}
