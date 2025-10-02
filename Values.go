package maps

// Gets all values from map and returns them as a slice.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A map from which all values are retrieved.
//
// # Returns
//
//	result []TValue
//
// A slice that contains all values of given map. The values in the returned slice are in an indeterminate order.
func Values[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource) (result []TValue) {
	result = make([]TValue, len(source))
	i := 0
	for _, value := range source {
		result[i] = value
		i++
	}
	return result
}
