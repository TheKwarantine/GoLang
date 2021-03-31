package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T){
	t.Parallel()
	type testCase struct{
		fn func(float64, float64) float64
		a, b float64
		want float64
	}
	tcs := []testCase{
		{ fn: calculator.Add, a: 2, b: 2, want: 4},
		{ fn: calculator.Add, a: 3, b: 3, want: 6},
		{ fn: calculator.Subtract, a: -2, b: -2, want: 0},
		{ fn: calculator.Multiply, a: 2, b: 4, want: 8},
	}
	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%.1f, %.1f): Want %.1f, got %.1f", tc.a, tc.b, tc.want, got)
		}
	}
	
}

func TestSubtract(t *testing.T){
	t.Parallel()
	var a, b float64
	a = 4
	b = 2
	var want float64 = 2
	got := calculator.Subtract(a, b)
	if want != got {
		t.Errorf("Subtract(%.1f, %.1f): Want %.1f, got %.1f", a, b, want, got)
	}
}

func TestMultiply(t *testing.T){
	t.Parallel()
	var a, b float64
	a = 2
	b = 2
	var want float64 = 4
	got := calculator.Multiply(a, b)
	if want != got {
		t.Errorf("Multiply(%.1f, %.1f): Want %.1f, got %.1f", a, b, want, got)
	}
}

func TestDivide(t *testing.T){
	t.Parallel()

	tcs := []struct{
		a, b float64
		want float64
		errExpected bool
	}{
		{ a: 3, b:0, want: 999, errExpected: true},
		{ a: 3, b:3, want: 1, errExpected: false},
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)
		if errReceived != tc.errExpected {
			t.Fatalf("Divide(%.1f, %.1f): unexpected error status: %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%.1f, %.1f): Want %.1f, got %.1f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T){
	t.Parallel()

	tcs := []struct{
		a, b float64
		want float64
		errExpected bool
	}{
		{ a: 9, want: 3, errExpected: false},
		{ a: 81, want: 9, errExpected: false},
		{ a: -5, want: 999, errExpected: true},
	}
	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.a)
		errReceived := (err != nil)
		if errReceived != tc.errExpected {
			t.Fatalf("Square Root(%.1f): unexpected error status: %v", tc.a, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Square Root(%.1f): Want %.1f, got %.1f", tc.a, tc.want, got)
		}
	}
}