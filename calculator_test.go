package main

import (
	"testing"
	"math"
)

func TestAdd(t *testing.T) {
	result, err := Add(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 8.0
	if math.Abs(result-expected) > 1e-9 {
	t.Errorf("expected %f but got %f", expected, result)
}
}

func TestSubtract(t *testing.T) {
	result, err  := Subtract(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 2.0
	if math.Abs(result-expected) > 1e-9 {
	t.Errorf("expected %f but got %f", expected, result)
}
}

func TestMultiply(t *testing.T) {
	result, err  := Multiply(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 15.0
	if math.Abs(result-expected) > 1e-9 {
	t.Errorf("expected %f but got %f", expected, result)
}
}

func TestDivide(t *testing.T) {
	result, err  := Divide(9, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 3.0
	if math.Abs(result-expected) > 1e-9 {
	t.Errorf("expected %f but got %f", expected, result)
}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(5, 0)
	if err == nil {
		t.Errorf("expected error but got nil")
	}
}

func TestPower(t *testing.T) {
	result, err  := Power(5, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := 125.0
	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("expected %f but got %f", expected, result)
	}
}