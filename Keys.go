package maps

// Gets all keys from map and returns them as a slice.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A map from which all keys are retrieved.
//
// # Returns
//
//	result []TKey
//
// A slice that contains all keys of given map. The keys in the returned slice are in an indeterminate order.
func Keys[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource) (result []TKey) {
	result = make([]TKey, len(source))
	i := 0
	for key, _ := range source {
		result[i] = key
		i++
	}
	return result
}
