package main

import "testing"

func TestSum(t *testing.T) {
	if sum(2, 3) != 5 {
		t.Error("Expected 5")
	}
}

func TestSub(t *testing.T) {
	if sub(5, 3) != 2 {
		t.Error("Expected 2")
	}
}

func TestDiv(t *testing.T) {
	if div(6, 2) != 3 {
		t.Error("Expected 3")
	}
	if div(1, 0) != 0 {
		t.Error("Expected 0 for division by zero")
	}
}

func TestMod(t *testing.T) {
	if mod(5, 2) != 1 {
		t.Error("Expected 1")
	}
	if mod(5, 0) != 0 {
		t.Error("Expected 0 for modulo by zero")
	}
}

func TestPow(t *testing.T) {
	if pow(2, 3) != 8 {
		t.Error("Expected 8")
	}
}

func TestSqrt(t *testing.T) {
	if sqrt(9) != 3 {
		t.Error("Expected 3")
	}
}

func TestCube(t *testing.T) {
	if cube(2) != 8 {
		t.Error("Expected 8")
	}
}

func TestSquare(t *testing.T) {
	if square(4) != 16 {
		t.Error("Expected 16")
	}
}

func TestAdd(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Error("Expected 4")
	}
}

func TestSubFunc(t *testing.T) {
	if Sub(5, 3) != 2 {
		t.Error("Expected 2")
	}
}

func TestMul(t *testing.T) {
	if Mul(3, 3) != 9 {
		t.Error("Expected 9")
	}
}

func TestDivFunc(t *testing.T) {
	if Div(4, 2) != 2 {
		t.Error("Expected 2")
	}
	if Div(1, 0) != 0 {
		t.Error("Expected 0 for division by zero")
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(4) {
		t.Error("Expected true")
	}
	if IsEven(3) {
		t.Error("Expected false")
	}
}

func TestMax(t *testing.T) {
	if Max(5, 3) != 5 {
		t.Error("Expected 5")
	}
	if Max(2, 7) != 7 {
		t.Error("Expected 7")
	}
}
