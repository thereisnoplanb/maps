package maps

func Keys[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource) (result []TKey) {
	result = make([]TKey, len(source))
	i := 0
	for key, _ := range source {
		result[i] = key
		i++
	}
	return result
}
