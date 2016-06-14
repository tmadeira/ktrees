package generator

import (
	"io/ioutil"
	"log"

	"testing"
)

func setup() func() int64 {
	log.SetOutput(ioutil.Discard)
	oldSeed := seed
	seed = func() int64 {
		return 0
	}
	return oldSeed
}

func tearDown(oldSeed func() int64) {
	seed = oldSeed
}

func TestRandomKtree(t *testing.T) {
	oldSeed := setup()
	defer tearDown(oldSeed)

	tests := []struct {
		n       int
		k       int
		wantErr bool
	}{
		{5, 0, true},
		{4, 3, true},
		{5, 3, false},
		{500, 5, false},
		{2000, 8, false},
	}

	for _, test := range tests {
		_, err := RandomKtree(test.n, test.k)
		if (err != nil) != test.wantErr {
			want := "<nil>"
			if test.wantErr {
				want = "<error>"
			}
			t.Errorf("RandomKtree(%d, %d) = %v; want %v", test.n, test.k, err, want)
		}
	}
}
