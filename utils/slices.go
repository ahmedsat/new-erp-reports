package utils

func FindF[T any](slice []T, predicate func(T) bool) T {
	for _, item := range slice {
		if predicate(item) {
			return item
		}
	}
	var zero T
	return zero
}
