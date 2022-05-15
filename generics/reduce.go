package generics

func Reduce[A, B any](values []A, f func(B, A) B, zeroValue B) B {
	var result B = zeroValue
	for _, value := range values {
		result = f(result, value)
	}
	return result
}
