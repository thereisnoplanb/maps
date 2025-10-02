package maps

func Where[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource, predicate predicate[TKey, TValue]) (result TSource) {
	result = make(TSource, 0)
	for key, value := range source {
		if predicate(key, value) {
			result[key] = value
		}
	}
	return result
}
