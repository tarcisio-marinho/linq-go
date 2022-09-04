package Operations

func (q Query[T]) Reverse() Query[T] {
	if len(q.Data) == 1 {
		return q
	}

	maxIndex := len(q.Data) - 1
	newQuery := Query[T]{}
	for i := maxIndex; i >= 0; i-- {
		elem := q.Data[i]
		newQuery.Data = append(newQuery.Data, elem)
	}
	return newQuery
}
