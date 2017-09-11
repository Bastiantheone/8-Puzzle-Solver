package puzzle_solver

import (
	"strings"
	"testing"
)

func TestIsGoal(t *testing.T) {
	tests := []struct {
		board board
		want  bool
	}{
		{board{0, 1, 2, 3, 4, 5, 6, 7, 8}, false},
		{board{0, 1, 2, 6, 7, 8, 3, 4, 5}, true},
		{board{8, 7, 6, 5, 4, 3, 2, 1, 0}, false},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		s := State{board: test.board}
		got := s.isGoal()
		if got != test.want {
			t.Errorf("test %d: got = %v, want = %v", i, got, test.want)
		}
	}
}

func TestMoves(t *testing.T) {
	s0 := State{}
	s1 := State{origin: &s0, board: board{2}, move: Left, cost: 1}
	s2 := State{origin: &s1, board: board{1}, move: Right, cost: 2}
	s3 := State{origin: &s2, board: board{0}, move: Down, cost: 3}
	wantMoves := []Move{Left, Right, Down}
	wantConfigs := []string{"2 ", "1 ", "0 "}
	gotMoves, gotConfigs := s3.moves()
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
		{State{board: board{0, 1, 2, 3, 4, 5, 6, 7, 8}, cost: 0}, []State{State{board: board{1, 0, 2, 3, 4, 5, 6, 7, 8}, move: Right, cost: 1},
			State{board: board{3, 1, 2, 0, 4, 5, 6, 7, 8}, move: Down, cost: 1}}},
		{State{board: board{4, 1, 2, 3, 0, 5, 6, 7, 8}, cost: 1}, []State{State{board: board{4, 1, 2, 0, 3, 5, 6, 7, 8}, move: Left, cost: 2},
			State{board: board{4, 1, 2, 3, 5, 0, 6, 7, 8}, move: Right, cost: 2}, State{board: board{4, 0, 2, 3, 1, 5, 6, 7, 8}, move: Up, cost: 2},
			State{board: board{4, 1, 2, 3, 7, 5, 6, 0, 8}, move: Down, cost: 2}}},
	}
	for i, test := range tests {
		got := test.s.neighbors()
		if len(got) != len(test.want) {
			t.Fatalf("test %d: got = %d states, want = %d", i, len(got), len(test.want))
		}
		for j, n := range got {
			if strings.Compare(n.key(), test.want[j].key()) != 0 {
				t.Errorf("test %d: board differs got[%d] = %s, want[%d] = %s", i, j, n.key(), j, test.want[j].key())
			}
			if n.cost != test.want[j].cost {
				t.Errorf("test %d: cost differs got[%d] = %d, want[%d] = %d", i, j, n.cost, j, test.want[j].cost)
			}
		}
	}
}

func TestHeuristic(t *testing.T) {
	tests := []struct {
		s    State
		want int
	}{
		{State{board: board{0, 2, 1, 3, 4, 5, 6, 7, 8}}, 8},
		{State{board: board{1, 0, 2, 3, 4, 5, 6, 8, 7}}, 9},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got := test.s.heuristic()
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
		{State{board: board{0, 1, 2, 3, 4, 5, 6, 7, 8}}, "012345678"},
		{State{board: board{1, 0, 2, 3, 4, 5, 6, 8, 7}}, "102345687"},
	}
	for i, test := range tests {
		got := test.s.key()
		if strings.Compare(got, test.want) != 0 {
			t.Errorf("test %d: got = %s, want = %s", i, got, test.want)
		}
	}
}

func TestSolvable(t *testing.T) {
	tests := []struct {
		b    board
		want bool
	}{
		{board{0, 1, 2, 6, 7, 8, 3, 5, 4}, false},
		{board{1, 0, 2, 6, 7, 8, 3, 4, 5}, true},
		{board{7, 2, 4, 5, 0, 6, 8, 1, 3}, true},
	}
	SetGoal([]int{0, 1, 2, 6, 7, 8, 3, 4, 5})
	for i, test := range tests {
		got := test.b.solvable()
		if got != test.want {
			t.Errorf("test %d: got = %v, want = %v", i, got, test.want)
		}
	}
}
