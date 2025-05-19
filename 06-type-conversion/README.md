# 06 - 타입 변환 (Type Conversion)

이 예제는 Go 언어에서 **기본 타입 간의 변환**과 `strconv` 패키지를 사용한 **문자열 ↔ 숫자 변환**을 실습합니다.  
Go는 정적 타입 언어이므로, 타입 변환은 **명시적으로** 해야 합니다.

<br>

## 주요 변환 방법

| 변환 방향 | 방법                                         | 예시 |
|-----------|--------------------------------------------|------|
| `int` → `float64` | `float64(i)`                               | `float64(3)` → `3.0` |
| `float64` → `int` | `int(f)`                                   | `int(3.9)` → `3` |
| `int` → `string` | `strconv.Itoa(i)`                          | `strconv.Itoa(42)` → `"42"` |
| `string` → `int` | `strconv.Atoi(s)`                          | `strconv.Atoi("42")` → `42` |
| `float64` → `string` | `strconv.FormatFloat(f, 'f', 소수점 자릿수, 64)` | `"3.14"` |
| `string` → `float64` | `strconv.ParseFloat(s, 64)`                | `"2.718"` → `2.718` |

<br>

## 실행
```bash
go run .
```
출력
```bash
int → float64: 42
int → string: 42
string → int: 123
float64 → string: 3.14
string → float64: 2.718
```