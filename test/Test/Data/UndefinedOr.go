package Test_Data_UndefinedOr

func Undefined() any {
	return nil
}

func Defined(x any) any {
	return x
}

func EqUndefinedOrImpl(eq func(any) func(any) any, a any, b any) any {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return eq(a)(b)
}

func CompareUndefinedOrImpl(lt any, eq any, gt any, compare func(any) func(any) any, a any, b any) any {
	if a == nil && b == nil {
		return eq
	}
	if a == nil {
		return lt
	}
	if b == nil {
		return gt
	}
	return compare(a)(b)
}
