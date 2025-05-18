package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, Go Language!"

	fmt.Println("원본:", str)
	fmt.Println("포함 여부(Go):", strings.Contains(str, "Go"))
	fmt.Println("접두사 확인:", strings.HasPrefix(str, "Hello"))
	fmt.Println("접미사 확인:", strings.HasSuffix(str, "!"))
	fmt.Println("대문자:", strings.ToUpper(str))
	fmt.Println("소문자:", strings.ToLower(str))

	replaced := strings.Replace(str, "Go", "Golang", 1)
	fmt.Println("치환:", replaced)

	split := strings.Split(str, " ")
	fmt.Println("분할:", split)

	joined := strings.Join(split, "-")
	fmt.Println("다시 합침:", joined)

	trimmed := strings.Trim("...Hello...", ".")
	fmt.Println("양쪽 문자 제거:", trimmed)
}
