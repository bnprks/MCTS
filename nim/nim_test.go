package nim

import "testing"

func TestAvailableMoves(t *testing.T) {
	var ns NimState
	var arr []string
	states := []NimState{
		NimState{10, 0},
		NimState{3, 0},
		NimState{2, 0},
		NimState{1, 0},
		NimState{0, 0},
	}
	answers := [][]string{
		[]string{"1", "2", "3"},
		[]string{"1", "2", "3"},
		[]string{"1", "2"},
		[]string{"1"},
		[]string{},
	}
	for i, _ := range states {
		ns = states[i]
		arr = ns.AvailableMoves()
		if !arrEquals(arr, answers[i]) {
			t.Error("Available Moves for ", ns, "got",
				arr, "expected", answers[i])
		}
	}
}

func arrEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMakeMove(t *testing.T) {
	ns := MakeNimState()
	moves := []string{
		"2",
		"3",
		"1",
		"2",
		"2",
	}
	states := []NimState{
		NimState{8, 1},
		NimState{5, 0},
		NimState{4, 1},
		NimState{2, 0},
		NimState{0, 1},
	}
	if ns.currentPlayer != 0 {
		t.Error("Initial player not set to 0")
	}
	for i, move := range moves {
		ns.MakeMove(move)
		if ns != states[i] {
			t.Error("After move", move, "got", ns, "expected", states[i])
		}
	}
}

func TestTerminalPlayout(t *testing.T) {
	ns := NimState{0, 0}
	result := ns.RandomPlayout()
	if (ns != NimState{0, 0}) {
		t.Error("RandomPlayout changed terminal state from {0 0} to", ns)
	}
	if result != false {
		t.Error("RandomPlayout returned wrong winner")
	}
}

func BenchmarkClonedPlayoutsPointers(b *testing.B) {
	ns := NimState{10, 0}
	p := &ns
	var cpy *NimState
	b.ResetTimer()
	cpy.Clone(ns)
	cpy.RandomPlayout()
}

func BenchmarkClonedPlayoutsValues(b *testing.B) {
	ns := NimState{10,0}
	var cpy NimState
	b.ResetTimer()
	cpy = ns.Clone()
	cpy.RandomPlayout()
}
