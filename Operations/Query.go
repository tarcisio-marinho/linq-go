package Operations

type Query[T any] struct {
	Data []T
}

func (q Query[T]) ToSlice() []T {
	return q.Data
}
