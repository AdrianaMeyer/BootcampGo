package main

import (
	"fmt"
)

func Fibonacci(num int) int{
	if num <= 1 {
		return num
	}

	fib := make([]int, num+1)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i <= num; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[num]
}

func main() {
	num := 30
	fmt.Printf("Fibonacci(%d) = %d\n", num, Fibonacci(num))
}
