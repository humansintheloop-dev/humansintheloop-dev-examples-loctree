package state

// MoveUp moves the selection up, clamped to valid range
func MoveUp(currentIndex, maxIndex int) int {
	if currentIndex > 0 {
		return currentIndex - 1
	}
	return 0
}

// MoveDown moves the selection down, clamped to valid range
func MoveDown(currentIndex, maxIndex int) int {
	if currentIndex < maxIndex {
		return currentIndex + 1
	}
	return maxIndex
}