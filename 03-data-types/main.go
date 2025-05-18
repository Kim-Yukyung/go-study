package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a int     = 10
		b float64 = 3.14
		c bool    = true
		d string  = "Go"
		e rune    = '한' // 유니코드 문자 -> int32와 동일
		f byte    = 'A'
		g         = []int{1, 2, 3} // 슬라이스
		h         = map[string]int{"x": 1, "y": 2}
	)

	fmt.Println("a:", a, reflect.TypeOf(a))
	fmt.Println("b:", b, reflect.TypeOf(b))
	fmt.Println("c:", c, reflect.TypeOf(c))
	fmt.Println("d:", d, reflect.TypeOf(d))
	fmt.Println("e:", e, string(e), reflect.TypeOf(e))
	fmt.Println("f:", f, string(f), reflect.TypeOf(f))
	fmt.Println("g:", g, reflect.TypeOf(g))
	fmt.Println("h:", h, reflect.TypeOf(h))
}
