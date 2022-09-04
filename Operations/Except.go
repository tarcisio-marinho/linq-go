package Operations

// Except return the set difference of two queries: A - B
func (q Query[T]) Except(secondQuery Query[T]) Query[T] {
	set := make(map[any]bool)
	newQuery := Query[T]{}
	for _, value := range secondQuery.Data {
		set[value] = true
	}

	for _, value := range q.Data {
		got := set[value]
		if !got {
			newQuery.Data = append(newQuery.Data, value)
		}
	}
	return newQuery
}
