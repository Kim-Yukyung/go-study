# 01 - Quote Example

이 예제는 Go에서 외부 패키지를 가져와 사용하는 방법을 보여줍니다.

<br>

## 모듈 초기화
```bash
go mod init quote-example
```
- mod: 모듈과 관련된 작업을 하겠다는 의미
- init: 모듈을 초기화한다는 의미
- quote-example: 이 프로젝트(모듈)의 이름
> 이 명령을 실행하면 go.mod 파일이 생성되며, Go는 이 디렉토리를 하나의 프로젝트(모듈)로 인식하게 됩니다.

예시 (go.mod)
```bash
module quote-example

go 1.24.3
```

<br>

## 의존성 정리
```bash
go mod tidy
```
이 명령은 다음 작업을 수행합니다.

- 코드에서 import된 외부 패키지를 자동으로 다운로드
- go.mod, go.sum 파일에 반영
- 사용하지 않는 패키지 제거

<br>

## 실행 방법
```bash
go run .
```
출력
```bash
Don't communicate by sharing memory, share memory by communicating.
```
