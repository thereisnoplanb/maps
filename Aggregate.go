package maps

// Applies an accumulate function over a sequence. The specified seed value is used as the initial accumulator value.
//
// # Parameters
//
//	source map[TKey]TValue
//
// A map to aggregate over.
//
//	seed TAccumulator
//
// The initial accumulator value.
//
//	accumulate accumulate[TAccumulator, TKey, TValue]
//
// An accumulate function to be invoked on each element.
//
// # Returns
//
//	result TAccumulator
//
// The final accumulator value.
func Aggreagate[TSource ~map[TKey]TValue, TKey comparable, TValue any, TAccumulator any](source TSource, seed TAccumulator, accumulate accumulate[TAccumulator, TKey, TValue]) (result TAccumulator) {
	result = seed
	for key, value := range source {
		result = accumulate(result, key, value)
	}
	return result
}
