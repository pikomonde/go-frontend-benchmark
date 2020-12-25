package main

import (
	"fmt"
	"syscall/js"
	"time"
)

// Factorial
func factorialRecursiveInt64(input int64) int64 {
	if (input == 0) || (input == 1) {
		return 1
	}
	return input * factorialRecursiveInt64(input-1)
}

func factorialRecursiveInt(input int) int {
	if (input == 0) || (input == 1) {
		return 1
	}
	return input * factorialRecursiveInt(input-1)
}

func factorialRecursiveInt8(input int8) int8 {
	if (input == 0) || (input == 1) {
		return 1
	}
	return input * factorialRecursiveInt8(input-1)
}

// Fibonacci
func fibonacciRecursiveInt64(input int64) int64 {
	if (input == 0) || (input == 1) {
		return 1
	}
	return fibonacciRecursiveInt64(input-2) + fibonacciRecursiveInt64(input-1)
}
func fibonacciRecursiveInt(input int) int {
	if (input == 0) || (input == 1) {
		return 1
	}
	return fibonacciRecursiveInt(input-2) + fibonacciRecursiveInt(input-1)
}
func fibonacciRecursiveInt8(input int8) int8 {
	if (input == 0) || (input == 1) {
		return 1
	}
	return fibonacciRecursiveInt8(input-2) + fibonacciRecursiveInt8(input-1)
}

func main() {
	nTest := 1000
	t0 := time.Now()
	for i := 0; i < nTest; i++ {
		factorialRecursiveInt64(20)
	}
	t1 := time.Now()
	for i := 0; i < nTest; i++ {
		factorialRecursiveInt(20)
	}
	t2 := time.Now()
	for i := 0; i < nTest; i++ {
		factorialRecursiveInt8(20)
	}
	t3 := time.Now()
	for i := 0; i < nTest; i++ {
		fibonacciRecursiveInt64(20)
	}
	t4 := time.Now()
	for i := 0; i < nTest; i++ {
		fibonacciRecursiveInt(20)
	}
	t5 := time.Now()
	for i := 0; i < nTest; i++ {
		fibonacciRecursiveInt8(20)
	}
	t6 := time.Now()

	js.Global().Set("factorialRecursiveInt64WASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return factorialRecursiveInt64(int64(args[0].Int()))
	}))
	js.Global().Set("factorialRecursiveIntWASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return factorialRecursiveInt(args[0].Int())
	}))
	js.Global().Set("factorialRecursiveInt8WASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return factorialRecursiveInt8(int8(args[0].Int()))
	}))
	js.Global().Set("fibonacciRecursiveInt64WASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fibonacciRecursiveInt64(int64(args[0].Int()))
	}))
	js.Global().Set("fibonacciRecursiveIntWASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fibonacciRecursiveInt(args[0].Int())
	}))
	js.Global().Set("fibonacciRecursiveInt8WASM", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return fibonacciRecursiveInt8(int8(args[0].Int()))
	}))

	fmt.Printf("factorialRecursiveInt64: %v\n", t1.Sub(t0))
	fmt.Printf("factorialRecursiveInt: %v\n", t2.Sub(t1))
	fmt.Printf("factorialRecursiveInt8: %v\n", t3.Sub(t2))
	fmt.Printf("fibonacciRecursiveInt64: %v\n", t4.Sub(t3))
	fmt.Printf("fibonacciRecursiveInt: %v\n", t5.Sub(t4))
	fmt.Printf("fibonacciRecursiveInt8: %v\n", t6.Sub(t5))
	c := make(chan bool)
	<-c
}
