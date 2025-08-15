package state

import (
	"testing"
)

func TestMoveUp(t *testing.T) {
	// Test moving up from middle
	current := 5
	max := 10
	result := MoveUp(current, max)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestMoveUp_AtTop(t *testing.T) {
	// Should stay at top
	current := 0
	max := 10
	result := MoveUp(current, max)
	if result != 0 {
		t.Errorf("Expected to stay at 0, got %d", result)
	}
}

func TestMoveDown(t *testing.T) {
	// Test moving down from middle
	current := 5
	max := 10
	result := MoveDown(current, max)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestMoveDown_AtBottom(t *testing.T) {
	// Should stay at bottom
	current := 10
	max := 10
	result := MoveDown(current, max)
	if result != 10 {
		t.Errorf("Expected to stay at 10, got %d", result)
	}
}