package Operations

// From return the query of the collection
func From[T any](data []T) Query[T] {
	return Query[T]{Data: data}
}
