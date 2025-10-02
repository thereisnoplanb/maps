package maps

func Any[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource, predicate predicate[TKey, TValue]) bool {
	for key, value := range source {
		if predicate(key, value) {
			return true
		}
	}
	return false
}
