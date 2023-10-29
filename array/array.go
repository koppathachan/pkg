package array

type GenericArray[T, R any] []T

func (arr GenericArray[T, R]) Map(mapper func(T) R) []R {
	result := make([]R, len(arr))

	for i, item := range arr {
		result[i] = mapper(item)
	}

	return result
}

func (arr GenericArray[T, R]) Reduce(reducer func(R, T) R, initialValue R) R {
	accumulator := initialValue

	for _, item := range arr {
		accumulator = reducer(accumulator, item)
	}

	return accumulator
}

func New[T, R any](arr []T) GenericArray[T, R] {
	return arr
}
