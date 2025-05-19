package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	f := float64(i)
	fmt.Println("int → float64:", f)

	s := strconv.Itoa(i)
	fmt.Println("int → string:", s)

	s2 := "123"
	i2, _ := strconv.Atoi(s2)
	fmt.Println("string → int:", i2)

	sf := strconv.FormatFloat(3.14, 'f', 2, 64)
	fmt.Println("float64 → string:", sf)

	f2, _ := strconv.ParseFloat("2.718", 64)
	fmt.Println("string → float64:", f2)
}
