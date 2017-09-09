package puzzle_solver

type State struct {
	origin State
	cost   int
}

func (s *State) isGoal() bool {
	panic("not implemented")
}

func (s *State) Moves() []Move {
	panic("not implemented")
}

func (s *State) neighbors() []State {
	panic("not implemented")
}

func (s *State) heuristic() int {
	panic("not implemented")
}

func (s *State) key() string {
	panic("not implemented")
}

type Move struct {
}
