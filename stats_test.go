package basemath

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const tolerance = .0001

func TestHypergeom(t *testing.T) {
	h := HypGeomPMF(4, 50, 5, 10)

	fmt.Printf("%f", h)
}

func TestHypergeomCDF(t *testing.T) {
	h := HypGeomCDF(4, 50, 5, 10)

	fmt.Printf("%f", h)
}

func TestHypergeomCDF2(t *testing.T) {

	got := HypGeomCDF(2, 100, 20, 10)
	want := 0.6812

	opt := cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)

		return diff < tolerance
	})

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

	fmt.Printf("%f", got)
}
