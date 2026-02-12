package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(5, 3)
	expected := 8.0

	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2.0

	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(5, 3)
	expected :=15.0

	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(9, 3)
	expected :=3.0

	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}

func TestPower(t *testing.T) {
	result := Power(5, 3)
	expected :=125.0

	if result != expected {
		t.Errorf("expected %f but got %f", expected, result)
	}
}