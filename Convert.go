package maps

func Convert[TSource ~map[TKey]TValue, TKey comparable, TValue any, TResultKey comparable, TResultValue any](source TSource, convert convert[TKey, TValue, TResultKey, TResultValue]) (result map[TResultKey]TResultValue) {
	result = make(map[TResultKey]TResultValue, len(source))
	for key, value := range source {
		newKey, newValue := convert(key, value)
		result[newKey] = newValue
	}
	return result
}
