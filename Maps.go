package maps

//type valueSelector[TObject any, TResult any] = func(object TObject) TResult
type accumulate[TAccumulator any, TKey comparable, TValue any] = func(accumulator TAccumulator, key TKey, value TValue) TAccumulator

type predicate[TKey comparable, TValue any] = func(key TKey, value TValue) bool

type convert[TKey comparable, TValue any, TResultKey comparable, TResultValue any] = func(key TKey, value TValue) (TResultKey, TResultValue)

//type keySelector[TObject any, TKey comparable] = func(object TObject) TKey
