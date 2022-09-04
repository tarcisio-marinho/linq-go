package Operations

func (q Query[T]) Select(function func(T) any) Query[any] {
	var newQuery Query[any]
	for _, obj := range q.Data {
		response := function(obj)
		newQuery.Data = append(newQuery.Data, response)
	}
	return newQuery
}
