package main

import "fmt"

func main() {
	var x, y, z, a, b float32
	fmt.Scan(&a)
	z = a - 5
	y = z * 5
	b = 2 - y
	x = b / z
	fmt.Println(x)
}
