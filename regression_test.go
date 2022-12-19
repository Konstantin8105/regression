package regression

import (
	"fmt"
	"os"
	"testing"
)

func ExampleLinear() {
	a, b, R2, err := Linear([][2]float64{
		{2, 3},
		{4, 7},
		{6, 5},
		{8, 10},
	})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "y   = %.4f*x+%.4f\n", a, b)
	fmt.Fprintf(os.Stdout, "R^2 = %.4f\n", R2)
	// Output:
	// y   = 1.5000*x+0.9500
	// R^2 = 0.6748
}

func ExampleQuadratic() {
	a, b, c, R2, err := Quadratic([][2]float64{
		{-3.00, 7.50},
		{-2.00, 3.00},
		{-1.00, 0.50},
		{+0.00, 1.00},
		{+1.00, 3.00},
		{+2.00, 6.00},
		{+3.00, 14.0},
	})
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stdout, "y   = %.4f*x^2+%.4f*x+%.4f\n", a, b, c)
	fmt.Fprintf(os.Stdout, "R^2 = %.4f\n", R2)
	// Output:
	// y   = 1.1071*x^2+1.0000*x+0.5714
	// R^2 = 0.9884
}

func TestQuadratic2points(t *testing.T) {
	a, b, c, R2, err := Quadratic([][2]float64{
		{-3.00, 7.50},
		{-2.00, 3.00},
	})
	t.Logf("%v %v %v %v", a, b, c, R2)
	if err == nil {
		t.Errorf("strange data")
	}
}

func TestQuadraticLine(t *testing.T) {
	a, b, c, R2, err := Quadratic([][2]float64{
		{1, 1},
		{2, 2},
		{4, 4},
		{8, 8},
		{16, 16},
	})
	t.Logf("%v %v %v %v", a, b, c, R2)
	if err != nil {
		t.Error(err)
	}
}
