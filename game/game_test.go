package game

import (
	"strings"
	"testing"
)

func TestIsGoal(t *testing.T) {
	tests := []struct {
		board Board
		want  bool
	}{
		{Board{0, 1, 2, 3, 4, 5, 6, 7, 8}, false},
		{Board{0, 1, 2, 6, 7, 8, 3, 4, 5}, true},
		{Board{8, 7, 6, 5, 4, 3, 2, 1, 0}, false},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		s := State{board: test.board}
		got := s.IsGoal()
		if got != test.want {
			t.Errorf("test %d: got = %v, want = %v", i, got, test.want)
		}
	}
}

func TestMoves(t *testing.T) {
	s0 := State{}
	s1 := State{origin: &s0, board: Board{2}, move: Left, cost: 1}
	s2 := State{origin: &s1, board: Board{1}, move: Right, cost: 2}
	s3 := State{origin: &s2, board: Board{0}, move: Down, cost: 3}
	wantMoves := []Move{Left, Right, Down}
	wantConfigs := []string{"2 ", "1 ", "0 "}
	gotMoves, gotConfigs := s3.Moves()
	if len(gotMoves) != len(wantMoves) {
		t.Fatalf("got = %d moves, want = %d", len(gotMoves), len(wantMoves))
	}
	for i := range gotMoves {
		if gotMoves[i] != wantMoves[i] {
			t.Errorf("got[%d] = %s, want[%d] = %s", i, gotMoves[i], i, wantMoves[i])
		}
	}
	if len(gotConfigs) != len(wantConfigs) {
		t.Fatalf("got = %d configurations, want = %d", len(gotConfigs), len(wantConfigs))
	}
	for i := range gotConfigs {
		if strings.Compare(gotConfigs[i], wantConfigs[i]) != 0 {
			t.Errorf("got[%d] = %s, want[%d] = %s", i, gotConfigs[i], i, wantConfigs[i])
		}
	}
}

func TestNeighbors(t *testing.T) {
	tests := []struct {
		s    State
		want []State
	}{
		{State{board: Board{0, 1, 2, 3, 4, 5, 6, 7, 8}, cost: 0}, []State{State{board: Board{1, 0, 2, 3, 4, 5, 6, 7, 8}, move: Right, cost: 1},
			State{board: Board{3, 1, 2, 0, 4, 5, 6, 7, 8}, move: Down, cost: 1}}},
		{State{board: Board{4, 1, 2, 3, 0, 5, 6, 7, 8}, cost: 1}, []State{State{board: Board{4, 1, 2, 0, 3, 5, 6, 7, 8}, move: Left, cost: 2},
			State{board: Board{4, 1, 2, 3, 5, 0, 6, 7, 8}, move: Right, cost: 2}, State{board: Board{4, 0, 2, 3, 1, 5, 6, 7, 8}, move: Up, cost: 2},
			State{board: Board{4, 1, 2, 3, 7, 5, 6, 0, 8}, move: Down, cost: 2}}},
	}
	for i, test := range tests {
		got := test.s.Neighbors()
		if len(got) != len(test.want) {
			t.Fatalf("test %d: got = %d states, want = %d", i, len(got), len(test.want))
		}
		for j, n := range got {
			if strings.Compare(n.Key(), test.want[j].Key()) != 0 {
				t.Errorf("test %d: board differs got[%d] = %s, want[%d] = %s", i, j, n.Key(), j, test.want[j].Key())
			}
			if n.cost != test.want[j].cost {
				t.Errorf("test %d: cost differs got[%d] = %d, want[%d] = %d", i, j, n.cost, j, test.want[j].cost)
			}
		}
	}
}

func TestManhattan(t *testing.T) {
	tests := []struct {
		s    State
		want int
	}{
		{State{board: Board{0, 2, 1, 3, 4, 5, 6, 7, 8}}, 8},
		{State{board: Board{1, 0, 2, 3, 4, 5, 6, 8, 7}}, 9},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got := test.s.manhattan()
		if got != test.want {
			t.Errorf("test %d: got = %d, want = %d", i, got, test.want)
		}
	}
}

func TestLinearConflict(t *testing.T) {
	tests := []struct {
		s    State
		want int
	}{
		{State{board: Board{0, 2, 1, 3, 4, 5, 6, 7, 8}}, 8},
		{State{board: Board{1, 0, 2, 3, 4, 5, 6, 8, 7}}, 2},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got := test.s.linearConflict()
		if got != test.want {
			t.Errorf("test %d: got = %d, want = %d", i, got, test.want)
		}
	}
}

func TestKey(t *testing.T) {
	tests := []struct {
		s    State
		want string
	}{
		{State{board: Board{0, 1, 2, 3, 4, 5, 6, 7, 8}}, "012345678"},
		{State{board: Board{1, 0, 2, 3, 4, 5, 6, 8, 7}}, "102345687"},
	}
	for i, test := range tests {
		got := test.s.Key()
		if strings.Compare(got, test.want) != 0 {
			t.Errorf("test %d: got = %s, want = %s", i, got, test.want)
		}
	}
}

func TestSolvable(t *testing.T) {
	tests := []struct {
		b    Board
		want bool
	}{
		{Board{0, 1, 2, 6, 7, 8, 3, 5, 4}, false},
		{Board{1, 0, 2, 6, 7, 8, 3, 4, 5}, true},
		{Board{7, 2, 4, 5, 0, 6, 8, 1, 3}, true},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got := test.b.Solvable()
		if got != test.want {
			t.Errorf("test %d: got = %v, want = %v", i, got, test.want)
		}
	}
}
