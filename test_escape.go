package main

import (
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

func main() {}
