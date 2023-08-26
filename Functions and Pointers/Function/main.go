package main

import "fmt"

func greet(name string) string {
	text := "Hello " + name
	return text
}

func add(x, y int) int {
	return x + y
}

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

func mulDiv(x, y float32) (mulRes float32, divRes float32) {
	mulRes = x * y
	divRes = x / y
	return
}

func main() {
	fmt.Println(greet("Yessir"))
	fmt.Println(add(2, 5))
	fmt.Println(addSub(2, 5))
	fmt.Println(mulDiv(2, 5))
}
