package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(10, 20)
	if got != 30 {
		t.Errorf("Add(10,20) = %d; want 30", got)
	}
}

func TestDivide1(t *testing.T) {
	got, err := Divide("100", "20")
	if err != nil {
		t.Errorf("Divide(100,20) returned an error: %v", err)
	}
	if got != 5 {
		t.Errorf("Divide(100,20) = %d; want 5", got)
	}
}

/*func TestDivide2(t *testing.T) {
	_, err := Divide("100", "0")
	if err == nil {
		t.Errorf("Expected an error for division by zero, but got nil")
	}
}*/
