package maps

func Select[TSource ~map[TKey]TValue, TKey comparable, TValue any, TResult any](source TSource, valueSelector func(key TKey, item TValue) TResult) (result map[TKey]TResult) {
	result = make(map[TKey]TResult)
	for key, value := range source {
		result[key] = valueSelector(key, value)
	}
	return result
}
