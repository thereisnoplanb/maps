package maps

func Values[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource) (result []TValue) {
	result = make([]TValue, len(source))
	i := 0
	for _, value := range source {
		result[i] = value
		i++
	}
	return result
}
