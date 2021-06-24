package unitconv

import "testing"

func TestCtoK(t *testing.T) {
	x := AbsoluteZeroC
	got := CtoK(x)
	want := Kelvin(0)
	if got != want {
		t.Errorf("CtoK(%s) = %s, want %s", x, got, want)
	}
}

func TestKtoC(t *testing.T) {
	x := Kelvin(0)
	got := KtoC(x)
	want := AbsoluteZeroC
	if got != want {
		t.Errorf("KtoC(%s) = %s, want %s", x, got, want)
	}
}
