package maps

// Determines whether any element of a sequence satisfies a condition.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A map whose elements to apply the predicate to.
//
// 	predicate predicate[TKey, TValue]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result bool
//
// True if the source sequence is not empty or nil and at least one of its elements passes the test in the specified predicate; otherwise, false.
//
// # Remarks
//
// If predicate parameter is ommited function returns true if the source sequence contains any elements; otherwise, false.
func Any[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource, predicate ...predicate[TKey, TValue]) bool {
	if len(predicate) > 0 {
		Predicate := predicate[0]
		for key, value := range source {
			if Predicate(key, value) {
				return true
			}
		}
		return false
	}
	return len(source) > 0
}
