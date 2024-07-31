package basemath

import (
	"fmt"
	"testing"
)

func TestHypergeom(t *testing.T) {
	h := HypGeomPMF(4, 50, 5, 10)

	fmt.Printf("%f", h)
}

func TestHypergeomCDF(t *testing.T) {
	h := HypGeomCDF(4, 50, 5, 10)

	fmt.Printf("%f", h)
}
