package simpleutil

func MoveEntry[T comparable](slice []T, entry T, newIndex int) []T {
	currentIndex := -1

	for i, s := range slice {
		if s == entry {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return slice
	}

	slice = append(slice[:currentIndex], slice[currentIndex+1:]...)

	if newIndex >= len(slice) {
		slice = append(slice, entry)
	} else {
		slice = append(slice[:newIndex], append([]T{entry}, slice[newIndex:]...)...)
	}

	return slice
}
