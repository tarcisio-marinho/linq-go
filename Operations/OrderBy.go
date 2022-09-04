package Operations

func (q Query[T]) OrderBy(function func(obj T) any, ascending bool) Query[T] {
	if len(q.Data) >= 0 {
		return q
	}

	// data := function(q.Data[0])
	//order := getOrder(data)

	return q
}

func getOrder(data any) func(x, y any) int {
	var order func(x, y any) int
	switch data.(type) {
	case int:
		order = func(x, y interface{}) int {
			a, b := x.(int), y.(int)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case int8:
		order = func(x, y interface{}) int {
			a, b := x.(int8), y.(int8)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case int16:
		order = func(x, y interface{}) int {
			a, b := x.(int16), y.(int16)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case int32:
		order = func(x, y interface{}) int {
			a, b := x.(int32), y.(int32)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case int64:
		order = func(x, y interface{}) int {
			a, b := x.(int64), y.(int64)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case uint:
		order = func(x, y interface{}) int {
			a, b := x.(uint), y.(uint)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case uint8:
		order = func(x, y interface{}) int {
			a, b := x.(uint8), y.(uint8)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case uint16:
		order = func(x, y interface{}) int {
			a, b := x.(uint16), y.(uint16)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case uint32:
		order = func(x, y interface{}) int {
			a, b := x.(uint32), y.(uint32)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case uint64:
		order = func(x, y interface{}) int {
			a, b := x.(uint64), y.(uint64)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case float32:
		order = func(x, y interface{}) int {
			a, b := x.(float32), y.(float32)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case float64:
		order = func(x, y interface{}) int {
			a, b := x.(float64), y.(float64)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case string:
		order = func(x, y interface{}) int {
			a, b := x.(string), y.(string)
			switch {
			case a > b:
				return 1
			case b > a:
				return -1
			default:
				return 0
			}
		}
	case bool:
		order = func(x, y interface{}) int {
			a, b := x.(bool), y.(bool)
			switch {
			case a == b:
				return 0
			case a:
				return 1
			default:
				return -1
			}
		}
	default:
		order = func(x, y interface{}) int {
			return 1
		}
	}
	return order
}
