package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	var v, l float64
	fmt.Print("Masukkan jari-jari (r): ")
	fmt.Scan(&r)
	v = (4.0 / 3.0) * math.Pi * float64(r) * float64(r) * float64(r)
	l = 4.0 * math.Pi * float64(r) * float64(r)
	fmt.Printf("Volume: %.4f\nLuas Permukaan: %.4f\n", v, l)
}
