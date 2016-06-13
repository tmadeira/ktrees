package characteristic

import (
	"reflect"
	"testing"

	"github.com/tmadeira/tcc/ktree"
)

var relabeledFig1A ktree.Ktree = ktree.Ktree{
	[][]int{
		{4, 6, 7, 8},
		{3, 8, 9, 10},
		{8, 9, 10},
		{1, 9, 10},
		{0, 6, 7, 8, 9},
		{7, 8, 10},
		{0, 4, 7},
		{0, 4, 5, 6, 8, 9, 10},
		{0, 1, 2, 4, 5, 7, 9, 10},
		{1, 2, 3, 4, 7, 8, 10},
		{1, 2, 3, 5, 7, 8, 9},
	},
	3,
}

var Rk ktree.RenyiKtree = ktree.RenyiKtree{
	&relabeledFig1A,
	[]int{1, 2, 8},
}

var Tk Tree = Tree{
	[]int{-1, 5, 0, 0, 2, 8, 8, 1, 0},
	[]int{-1, 2, -1, -1, 0, 2, 1, 2, -1},
}

func TestTreeFrom(t *testing.T) {
	want := &Tk
	got := TreeFrom(&Rk)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TreeFrom(%v) = %v; want %v", Rk, Tk)
	}
}

// TODO: add unit tests to test pruneRk.
// TODO: add unit tests to test addEdges.
// TODO: add unit tests to test buildLabels.
