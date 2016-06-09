package dandelion

import (
	"characteristic"
	"reflect"
	"testing"
)

func TestCodeFig2C(t *testing.T) {
	e := characteristic.Epsilon

	// This is the tree shown in Fig. 2(c) of Caminiti et al.
	characteristic := &characteristic.Tree{
		[]int{e, 5, 0, 0, 2, 8, 8, 1, 0},
		[]int{0, 3, e, e, 1, 3, 2, 3, e},
	}

	want := DandelionCode{
		[]int{0, 0, 2, 8, 8, 1, 5},
		[]int{e, e, 1, 3, 2, 3, 3},
	}
	got := Code(characteristic)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Code(Fig. 2c) = %v; want %v", characteristic, got, want)
	}
}

func TestDecodeFig2C(t *testing.T) {
	e := characteristic.Epsilon

	// This is the code shown in Fig. 2(c) of Caminiti et al.
	code := &DandelionCode{
		[]int{0, 0, 2, 8, 8, 1, 5},
		[]int{e, e, 1, 3, 2, 3, 3},
	}

	want := characteristic.Tree{
		[]int{e, 5, 0, 0, 2, 8, 8, 1, 0},
		[]int{e, 3, e, e, 1, 3, 2, 3, e},
	}
	got := Decode(code)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Decode(Fig. 2c) = %v; want %v", code, got, want)
	}
}

func CodeFig3(t *testing.T) {
	e := characteristic.Epsilon

	// This is the tree shown in Fig. 2(c) of Caminiti et al.
	characteristic := &characteristic.Tree{
		[]int{e, 9, 3, 0, 6, 2, 10, 1, 10, 6, 5, 1, 8, 3, 0},
		[]int{e, 2, 2, e, 3, 1, 3, 2, 3, 4, 4, 3, 1, 3, e},
	}

	want := DandelionCode{
		[]int{3, 2, 6, 5, 10, 1, 10, 6, 9, 1, 8, 3, 0},
		[]int{2, 1, 3, 4, 3, 2, 3, 4, 2, 3, 1, 3, e},
	}
	got := Code(characteristic)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Code(Fig. 3) = %v; want %v", characteristic, got, want)
	}
}

func TestDecodeFig3(t *testing.T) {
	e := characteristic.Epsilon

	// This is the code shown in Fig. 3 of Caminiti et al.
	code := &DandelionCode{
		[]int{3, 2, 6, 5, 10, 1, 10, 6, 9, 1, 8, 3, 0},
		[]int{2, 1, 3, 4, 3, 2, 3, 4, 2, 3, 1, 3, e},
	}

	want := characteristic.Tree{
		[]int{e, 9, 3, 0, 6, 2, 10, 1, 10, 6, 5, 1, 8, 3, 0},
		[]int{e, 2, 2, e, 3, 1, 3, 2, 3, 4, 4, 3, 1, 3, e},
	}
	got := Decode(code)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Decode(Fig. 3) = %v; want %v", code, got, want)
	}
}
