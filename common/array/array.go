package array

// Check in an slice if an element is duplicated (exists-twice)
func FindDuplicates[T comparable](points []T) bool {
	if len(points) == 0 {
		return false
	}

	// create a map to store the element if any
	elements := make(map[T]bool)

	for _, point := range points {
		// this means an element is duplicate
		if elements[point] == true {
			return true
		}

		elements[point] = true
	}

	return false
}
