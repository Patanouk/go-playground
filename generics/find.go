package generics

func Find[T any](items []T, predicate func(T) bool) (item T, found bool) {
	var zero T

	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}

	return zero, false
}
