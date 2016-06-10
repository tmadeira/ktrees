package renyi

import (
	"reflect"
	"testing"

	"ktree"
)

func TestFromTk(t *testing.T) {
	fig1A := ktree.Ktree{
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

	wantRenyi := RenyiKtree{
		&ktree.Ktree{
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
		},
		[]int{1, 2, 8},
	}

	invalidKtree := ktree.Ktree{
		[][]int{
			{1, 2, 3},
			{0, 2, 3},
			{0, 1, 3},
			{0, 1, 2},
		},
		2,
	}

	tests := []struct {
		t       *ktree.Ktree
		wantErr bool
		want    *RenyiKtree
	}{
		{&fig1A, false, &wantRenyi},
		{&invalidKtree, true, nil},
	}

	for _, test := range tests {
		got, err := FromTk(test.t)
		if (err != nil) != test.wantErr {
			want := "<nil>"
			if test.wantErr {
				want = "<error>"
			}
			t.Errorf("FromTk(%v) = %v, %v; want _, %v", test.t, got, err, want)
		} else if !reflect.DeepEqual(got, test.want) {
			t.Errorf("FromTk(%v) = %v, %v; want %v, _", test.t, got, err, test.want)
		}
	}
}
