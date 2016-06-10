package ktree

import (
	"reflect"
	"testing"
)

// fig1A is the k-Tree seen in Fig. 1(a) from Caminiti et al, but 0-indexed.
var fig1A Ktree = Ktree{
	[][]int{
		{1, 4, 6, 7},
		{0, 2, 4, 5, 7, 8, 9, 10},
		{1, 3, 4, 7, 9, 10},
		{2, 8, 10},
		{0, 1, 2, 6, 7},
		{1, 7, 8},
		{0, 4, 7},
		{0, 1, 2, 4, 5, 6, 8},
		{1, 2, 3, 5, 7, 9, 10},
		{1, 2, 8},
		{1, 2, 3, 8},
	},
	3,
}

func TestGetQ(t *testing.T) {
	invalidKtree := Ktree{
		[][]int{
			{1, 2, 3},
			{0, 2, 3},
			{0, 1, 3},
			{0, 1, 2},
		},
		2,
	}

	tests := []struct {
		t       *Ktree
		wantErr bool
		want    []int
	}{
		{&fig1A, false, []int{1, 2, 8}},
		{&invalidKtree, true, nil},
	}

	for _, test := range tests {
		got, err := GetQ(test.t)
		if (err != nil) != test.wantErr {
			want := "<nil>"
			if test.wantErr {
				want = "<error>"
			}
			t.Errorf("GetQ(%v) = %v, %v; want _, %v", test.t, got, err, want)
		} else if !reflect.DeepEqual(got, test.want) {
			t.Errorf("GetQ(%v) = %v, %v; want %v, _", test.t, got, err, test.want)
		}
	}
}

func TestComputePhi(t *testing.T) {
	// This is n, k, Q from the k-Tree in Fig. 1(a) in Caminiti et al.
	n, k := 11, 3
	Q := []int{1, 2, 8}
	want := []int{0, 8, 9, 3, 4, 5, 6, 7, 10, 2, 1}

	got := ComputePhi(n, k, Q)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ComputePhi(%v, %v, %v) = %v; want %v", n, k, Q, got, want)
	}
}

func TestRelabel(t *testing.T) {
	phi := []int{0, 8, 9, 3, 4, 5, 6, 7, 10, 2, 1}

	// This is the what we want after relabeling Fig. 1(a) from Caminiti et al.
	want := &Ktree{
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
			{1, 2, 3, 4, 7, 8},
			{1, 2, 3, 5, 7, 8, 9},
		},
		3,
	}

	got := Relabel(&fig1A, phi)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Relabel(%v, %v) = %v; want %v", fig1A, phi, got, want)
	}
}
