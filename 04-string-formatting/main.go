package main

import "fmt"

func main() {
	name := "유경"
	age := 24
	height := 160.55

	fmt.Printf("이름: %s\n", name)
	fmt.Printf("나이: %d살\n", age)
	fmt.Printf("키: %.1fcm\n", height) // 반올림

	// 여러 값 한 번에 출력
	fmt.Printf("안녕하세요:) 제 이름은 %s이고, 나이는 %d살, 키는 %.1fcm입니다.\n", name, age, height)

	// 문자열 조합 후 저장
	intro := fmt.Sprintf("소개: %s (%d세)", name, age)
	fmt.Println(intro)
}
