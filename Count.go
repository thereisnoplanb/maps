package maps

// Returns a number that represents how many elements in the specified sequence satisfy a condition.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A sequence that contains elements to be tested and counted.
//
//	predicate predicate[TKey, Tavlue] [OPTIONAL]
//
// A function to test each element for a condition.
//
// # Returns
//
//	count int
//
// A number that represents how many elements in the sequence satisfy the condition in the predicate function or
// the number of elements in the input sequence if predicate is ommited.
func Count[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource, predicate ...predicate[TKey, TValue]) (count int) {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for key, value := range source {
			if Predicate(key, value) {
				count++
			}
		}
		return count
	}
	return len(source)
}
