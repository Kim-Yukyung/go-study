package main

import "fmt"

const (
	A = iota
	B
	C
)

func main() {
	// 명시적 선언
	var i1 int = 10
	var s1 string = "Go1"

	// 타입 생략
	var i2 = 20
	var s2 = "Go2"

	// 축약형
	i3 := 30
	s3 := "Go3"

	// 다중 선언
	var x, y = 1, 2
	a, b := "a", "b"

	// 블록 선언
	var (
		name  = "yukyung"
		age   = 24
		valid = true
	)

	fmt.Println(i1, s1)
	fmt.Println(i2, s2)
	fmt.Println(i3, s3)
	fmt.Println(x, y, a, b)
	fmt.Println(name, age, valid)

	// 상수 출력
	fmt.Println("iota constants:", A, B, C)
}
