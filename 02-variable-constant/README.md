# 02 - 변수와 상수 (Variables & Constants)

이 예제는 Go 언어에서 변수와 상수를 선언하고 사용하는 다양한 방법을 실습합니다.

<br>

## 주요 문법 요약

### 변수 선언 방식

| 방식 | 예시                                    | 특징 |
|------|---------------------------------------|------|
| 명시적 선언 | `var i int = 10`                      | 타입 명시 |
| 타입 생략 | `var i = 10`                          | 타입 자동 추론 |
| 축약형 선언 | `i := 10`                             | 함수 내부에서만 사용 가능 |
| 다중 선언 | `var x, y = 1, 2`<br>`a, b := "a", "b"` | 여러 변수 동시 선언 가능 |
| 블록 선언 | `var (`<br>`    name = "kim"`<br>`    age = 24`<br>`)`    | 선언 묶음 |

> **기본값**: int → 0, string → "", bool → false

<br>

### 상수 (`const`)

- 값이 고정되어 변경되지 않음
- 문자열, 숫자, boolean 등 모든 타입 사용 가능

```go
const pi = 3.14
const name = "Gopher"
```

<br>

### iota로 열거형 만들기

iota는 0부터 시작해서 자동으로 1씩 증가하는 상수 생성 도구입니다.
```
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

<br> 

## 실행
```bash
go run .
```
출력
```bash
10 Go1
20 Go2
30 Go3
1 2 a b
yukyung 24 true
iota constants: 0 1 2
```
