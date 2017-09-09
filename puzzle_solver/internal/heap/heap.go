package heap

import (
	"github.com/Bastiantheone/8-Puzzle-Solver/puzzle_solver"
)

// node is an item in the heap. It contains a game state and its priority.
type node struct {
	state *puzzle_solver.State
	// p is the priority of the node.
	p int
}

// Heap contains a slice of items and the last index.
type Heap struct {
	n     int // n is the last index of the slice
	items []node
}

// New returns a heap pointer.
func New() *Heap {
	h := &Heap{n: 0}
	h.items = make([]node, 1)
	return h
}

// Push inserts the state and its priority into the heap.
func (h *Heap) Push(state *puzzle_solver.State, priority int) {
	item := node{state: state, p: priority}
	h.n++
	if h.n >= len(h.items) {
		h.items = append(h.items, item)
	} else {
		h.items[h.n] = item
	}
	h.moveUp(h.n)
}

// Pop returns the state with the lowest priority and
// removes it from the heap.
//
// It returns nil, 0 if the heap is empty.
func (h *Heap) Pop() *puzzle_solver.State {
	if h.n == 0 {
		return nil, 0
	}
	item := h.items[1]
	h.n--
	if h.n == 0 {
		return item.state
	}
	h.swap(1, h.n+1)
	h.moveDown(1)
	return item.state
}

// IsEmpty returns whether a heap is empty.
func (h *Heap) IsEmpty() bool {
	return h.n == 0
}

// moveUp compares an item with its parent and moves it up if it has
// a smaller priority. It will keep moving up until the parent's priority is smaller
// or if it reaches the top.
func (h *Heap) moveUp(i int) {
	for {
		parent := int(i / 2)
		if parent == 0 {
			return
		}
		if h.less(i, parent) {
			h.swap(i, parent)
			i = parent
			continue
		}
		return
	}
}

// moveDown compares an item to its children and moves it down
// if one of the children has a smaller priority. It will keep
// moving down until it reaches a leaf or both children have
// a bigger priority.
func (h *Heap) moveDown(i int) {
	for {
		left := 2 * i
		if left > h.n {
			return
		}
		right := 2*i + 1
		var smallestChild int
		switch {
		case left == h.n:
			smallestChild = left
		case h.less(left, right):
			smallestChild = left
		default:
			smallestChild = right
		}
		if h.less(smallestChild, i) {
			h.swap(i, smallestChild)
			i = smallestChild
			continue
		}
		return
	}
}

// less returns whether item at index i has a smaller priority
// than the item at index j.
func (h *Heap) less(i, j int) bool {
	return h.items[i].p < h.items[j].p
}

// swap swaps the two items at the indices i and j.
func (h *Heap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}
