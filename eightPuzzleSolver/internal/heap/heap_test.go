package heap

import (
	"strings"
	"testing"
)

func TestHeapPop(t *testing.T) {
	h := Heap{items: []node{{}, {key: "a", p: 0}, {key: "b", p: 2}, {key: "d", p: 4}, {key: "c", p: 3}, {key: "e", p: 5}}, n: 5}
	want := []string{"a", "b", "c", "d", "e"}
	for i := 0; !h.IsEmpty(); i++ {
		s := h.Pop()
		if i >= len(want) {
			t.Fatal("no item expected, but heap is not empty")
		}
		if strings.Compare(s, want[i]) != 0 {
			t.Errorf("got = %s, want[%d] = %s", s, i, want[i])
		}
	}
}

func TestHeapPush(t *testing.T) {
	items := []node{{key: "e", p: 5}, {key: "b", p: 2}, {key: "a", p: 0}, {key: "c", p: 3}, {key: "d", p: 4}}
	want := []node{{}, {key: "a", p: 0}, {key: "c", p: 3}, {key: "b", p: 2}, {key: "e", p: 5}, {key: "d", p: 4}}
	h := New()
	for _, item := range items {
		h.Push(item.key, item.p)
	}
	if len(want) != len(h.items) {
		t.Fatalf("got = %d items, want = %d", len(h.items), len(want))
	}
	for i, item := range h.items {
		if i == 0 {
			continue
		}
		if strings.Compare(item.key, want[i].key) != 0 {
			t.Errorf("key differs, got = %s, want = %s", item.key, want[i].key)
		}
		if item.p != want[i].p {
			t.Errorf("priority differs, got = %d, want = %d", item.p, want[i].p)
		}
	}
}
