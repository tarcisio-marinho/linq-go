package Operations

func (q Query[T]) Where(function func(T) bool) Query[T] {
	newQuery := Query[T]{}
	for _, obj := range q.Data {
		if function(obj) {
			newQuery.Data = append(newQuery.Data, obj)
		}
	}

	return newQuery
}
