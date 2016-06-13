package codec

import (
	"reflect"
	"testing"

	"github.com/tmadeira/tcc/characteristic"
	"github.com/tmadeira/tcc/dandelion"
	"github.com/tmadeira/tcc/ktree"
)

const e = characteristic.E

var fig1A ktree.Ktree = ktree.Ktree{
	[][]int{
		{1, 4, 6, 7},
		{0, 2, 4, 5, 7, 8, 9, 10},
		{1, 3, 4, 7, 8, 9, 10},
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

var codedFig1A Code = Code{
	[]int{1, 2, 8},
	&dandelion.DandelionCode{
		[]int{0, 2, 8, 8, 1, 5},
		[]int{e, 0, 2, 1, 2, 2},
	},
}

var invalidKtree ktree.Ktree = ktree.Ktree{
	[][]int{
		{1, 2, 3},
		{0, 2, 3},
		{0, 1, 3},
		{0, 1, 2},
	},
	2,
}

func TestCodingAlgorithm(t *testing.T) {
	tests := []struct {
		Tk      *ktree.Ktree
		wantErr bool
		want    *Code
	}{
		{&fig1A, false, &codedFig1A},
		{&invalidKtree, true, nil},
	}

	for _, test := range tests {
		got, err := CodingAlgorithm(test.Tk)
		if (err != nil) != test.wantErr {
			want := "<nil>"
			if test.wantErr {
				want = "<error>"
			}
			t.Errorf("Coding(%v) = %v, %v; want _, %v", test.Tk, got, err, want)
		} else if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Coding(%v) = %v, %v; want %v, _", test.Tk, got, err, test.want)
		}
	}
}

/*
func TestDecodingAlgorithm(t *testing.T) {
	got, err := DecodingAlgorithm(&codedFig1A)
	want := &fig1A

	if err != nil {
		t.Errorf("Decoding(%v) = _, %v; want %v, <nil>", codedFig1A, err, want)
	} else if !reflect.DeepEqual(got, want) {
		t.Errorf("Decoding(%v) = %v, _; want %v, <nil>", codedFig1A, got, want)
	}
}
*/
