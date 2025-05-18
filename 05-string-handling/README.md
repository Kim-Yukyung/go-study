# 05 - 문자열 처리

이 예제는 Go의 `strings` 패키지를 활용해 문자열을 다루는 다양한 방법을 실습합니다.  

<br>

## 사용 패키지

```go
import "strings"
```

<br>

## 주요 함수 정리

| 함수 | 설명 | 예시 입력 | 예시 출력 |
|------|------|------------|------------|
| `strings.Contains(s, substr)` | 문자열 포함 여부 확인 | `"Hello, Go"`, `"Go"` | `true` |
| `strings.HasPrefix(s, prefix)` | 접두사(prefix) 확인 | `"Hello, Go"`, `"Hello"` | `true` |
| `strings.HasSuffix(s, suffix)` | 접미사(suffix) 확인 | `"go.go"`, `".go"` | `true` |
| `strings.ToUpper(s)` | 문자열을 대문자로 변환 | `"go"` | `"GO"` |
| `strings.ToLower(s)` | 문자열을 소문자로 변환 | `"GO"` | `"go"` |
| `strings.Replace(s, old, new, n)` | 문자열 일부 치환 (`n`: 바꿀 횟수) | `"a-b-c"`, `"-"`, `":"`, `2` | `"a:b:c"` |
| `strings.Split(s, sep)` | 문자열을 구분자로 나눠 슬라이스로 반환 | `"a,b,c"`, `","` | `["a", "b", "c"]` |
| `strings.Join([]string, sep)` | 문자열 슬라이스를 구분자로 연결 | `["a", "b", "c"]`, `"-"` | `"a-b-c"` |
| `strings.Trim(s, cutset)` | 문자열 양 끝의 문자를 제거 | `"!!hi!!"`, `"!"` | `"hi"` |

<br>

## 실행
```bash
go run .
```
출력
```bash
원본: Hello, Go Language!
포함 여부(Go): true
접두사 확인: true
접미사 확인: true
대문자: HELLO, GO LANGUAGE!
소문자: hello, go language!
치환: Hello, Golang Language!
분할: [Hello, Go Language!]
다시 합침: Hello,-Go-Language!
양쪽 문자 제거: Hello
```