package Operations

func (q Query[T]) Distinct() Query[T] {

	set := make(map[any]bool)
	newQuery := Query[T]{}
	for _, value := range q.Data {
		got := set[value]
		if !got {
			set[value] = true
			newQuery.Data = append(newQuery.Data, value)
		}
	}
	return newQuery
}
