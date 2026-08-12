package main

import (
	"fmt"
	"unsafe"
)

type Value struct {
	Type      int
	IntVal    int64
	UnsafePtr unsafe.Pointer
}

//go:noinline
func Func(f func(Value) Value) Value {
	return Value{Type: 1, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&f))}
}

//go:noinline
func Apply(f Value, arg Value) Value {
    fn := *(*func(Value, Value) Value)(unsafe.Pointer(&f.UnsafePtr))
    return Func(func(a Value) Value { return fn(arg, a) })
}

func main() {
	var a int
	fmt.Printf("Stack address approx: %p\n", &a)
    
    dummy := func(x Value, y Value) Value { return x }
    val := Value{Type: 2, UnsafePtr: *(*unsafe.Pointer)(unsafe.Pointer(&dummy))}
    
    res := Apply(val, Value{})
    fmt.Printf("Closure address: %p\n", res.UnsafePtr)
}
