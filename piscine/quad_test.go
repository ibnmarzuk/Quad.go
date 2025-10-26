package piscine

import "testing"

func TestQuadA(t *testing.T) {
	// This just verifies that QuadA runs without crashing.
	QuadA(5, 3)
	QuadA(1, 1)
	QuadA(5, 1)
}

func TestQuadB(t *testing.T) {
	QuadB(5, 3)
	QuadB(1, 1)
	QuadB(5, 1)
}

func TestQuadC(t *testing.T) {
	QuadC(5, 3)
	QuadC(1, 1)
	QuadC(5, 1)
}

func TestQuadD(t *testing.T) {
	QuadD(5, 3)
	QuadD(1, 1)
	QuadD(5, 1)
}

func TestQuadE(t *testing.T) {
	QuadE(5, 3)
	QuadE(1, 1)
	QuadE(5, 1)
}

