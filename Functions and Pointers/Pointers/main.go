package main

import "fmt"

func increment(a int) {
	a = a + 1
}

func incrementPtr(b *int) {
	*b = *b + 1
}

func main() {
	var5 := 5
	fmt.Printf("My value is %d\n", var5)
	increment(var5)
	fmt.Printf("My value is %d\n", var5)
	incrementPtr(&var5)
	fmt.Printf("My address is %X, value is %d\n", &var5, var5)

	var var2 *int
	fmt.Printf("My value is %v\n", var2)

	var2 = &var5
	*var2 = 132
	fmt.Printf("My value is %v\n", var5)

	var2 = new(int)
	*var2 = 3
	fmt.Printf("My value is %v\n", *var2)
}
