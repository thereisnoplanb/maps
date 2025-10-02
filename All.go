package maps

// Determines whether all elements of a sequence satisfy a condition.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A map that contains the elements to apply the predicate to.
//
// 	predicate predicate[TKey, TValue]
//
// A function to test each element for a condition.
//
// # Returns
//
//	result bool
//
// True if every element of the source sequence passes the test in the specified predicate, or if the sequence is empty or nil; otherwise, false.
func All[TSource ~map[TKey]TValue, TKey comparable, TValue any](source TSource, predicate predicate[TKey, TValue]) (result bool) {
	for key, value := range source {
		if !predicate(key, value) {
			return false
		}
	}
	return true
}
