func PeekImpl(i int, xs []any) func() any { return func() any { return xs[i] } }
func PokeImpl(i int, a any, xs []any) func() bool { return func() bool { xs[i] = a; return true } }
