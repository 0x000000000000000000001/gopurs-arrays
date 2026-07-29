func New_() func() interface{} {
	return func() interface{} {
		arr := make([]interface{}, 0)
		return &arr
	}
}

func PeekImpl(just func(interface{}) interface{}, nothing interface{}, i int64, arr interface{}) func() interface{} {
	return func() interface{} {
		a := arr.(*[]interface{})
		if i >= 0 && i < int64(len(*a)) {
			return just((*a)[i])
		}
		return nothing
	}
}

func PokeImpl(i int64, a interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if i >= 0 && i < int64(len(*ptr)) {
			(*ptr)[i] = a
			return true
		}
		return false
	}
}

func LengthImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		return int64(len(*ptr))
	}
}

func PopImpl(just func(interface{}) interface{}, nothing interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if len(*ptr) > 0 {
			last := (*ptr)[len(*ptr)-1]
			*ptr = (*ptr)[:len(*ptr)-1]
			return just(last)
		}
		return nothing
	}
}

func PushAllImpl(xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(*ptr, xs...)
		return int64(len(*ptr))
	}
}

func PushImpl(x interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(*ptr, x)
		return int64(len(*ptr))
	}
}

func ShiftImpl(just func(interface{}) interface{}, nothing interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		if len(*ptr) > 0 {
			first := (*ptr)[0]
			*ptr = (*ptr)[1:]
			return just(first)
		}
		return nothing
	}
}

func UnshiftAllImpl(xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append(xs, *ptr...)
		return int64(len(*ptr))
	}
}

func UnshiftImpl(x interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		*ptr = append([]interface{}{x}, *ptr...)
		return int64(len(*ptr))
	}
}

func SpliceImpl(start int64, count int64, xs []interface{}, arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		removed := make([]interface{}, count)
		copy(removed, (*ptr)[start:start+count])
		
		newArr := make([]interface{}, 0, len(*ptr) - int(count) + len(xs))
		newArr = append(newArr, (*ptr)[:start]...)
		newArr = append(newArr, xs...)
		newArr = append(newArr, (*ptr)[start+count:]...)
		*ptr = newArr
		return removed
	}
}

func UnsafeFreezeImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		return *(arr.(*[]interface{}))
	}
}

func UnsafeThawImpl(xs []interface{}) func() interface{} {
	return func() interface{} {
		return &xs
	}
}

func FreezeImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		res := make([]interface{}, len(*ptr))
		copy(res, *ptr)
		return res
	}
}

func ThawImpl(xs []interface{}) func() interface{} {
	return func() interface{} {
		res := make([]interface{}, len(xs))
		copy(res, xs)
		return &res
	}
}

func CloneImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		ptr := arr.(*[]interface{})
		res := make([]interface{}, len(*ptr))
		copy(res, *ptr)
		return &res
	}
}

func SortByImpl(f func(interface{}, interface{}) interface{}, toInt func(interface{}) int64, arr interface{}) func() interface{} {
	return func() interface{} {
		panic("Not implemented: sortByImpl")
	}
}

func ToAssocArrayImpl(arr interface{}) func() interface{} {
	return func() interface{} {
		panic("Not implemented: toAssocArrayImpl")
	}
}
