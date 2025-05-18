# 00 - Hello World

이 코드는 Go 언어에서 가장 기본적인 예제인 "Hello, World!"를 출력합니다.

## 코드 설명

- `package main`: 이 파일이 실행 가능한 **메인 프로그램**임을 의미합니다. Go에서 실행 파일을 만들려면 반드시 main 패키지여야 합니다.
- `import "fmt"`: 콘솔 출력에 사용하는 표준 패키지를 불러옵니다.
- `func main()`: Go 프로그램의 시작점입니다. `go run`이나 `go build` 시 이 함수부터 실행됩니다.
- `fmt.Println("Hello, World!")`: 콘솔에 문자열을 출력합니다.

## 실행 방법
1. 터미널에서 실행
```bash
go run hello.go
```
출력
```
Hello, World!
```
> 이 방법은 소스 코드를 직접 실행하며, 실행 파일을 만들지 않습니다.

2. 실행 파일 만들기
```bash
go build hello.go
./hello
```
> 이 명령은 실행 가능한 바이너리를 생성합니다.