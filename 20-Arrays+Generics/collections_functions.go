package main

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}

	return result
}

func Find[A any](collection []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range collection {
		if predicate(v) {
			return v, true
		}
	}
	return
}
