package maps

//type valueSelector[TObject any, TResult any] = func(object TObject) TResult
//type accumulate[TAccumulator any, TObject any] = func(accumulator TAccumulator, object TObject) TAccumulator

type predicate[TKey comparable, TValue any] = func(key TKey, value TValue) bool

type convert[TKey comparable, TValue any, TResultKey comparable, TResultValue any] = func(key TKey, value TValue) (TResultKey, TResultValue)

//type keySelector[TObject any, TKey comparable] = func(object TObject) TKey
