# b64checker

Base64 인코딩/디코딩 및 URL 접속 확인 도구 (Go + Fyne GUI)

## 기능

- **Base64 인코딩** — 텍스트를 Base64 형식으로 인코딩
- **Base64 디코딩** — Base64 문자열을 원문으로 디코딩
- **URL 접속 확인** — 디코딩 결과가 URL인 경우 실제 접속 가능 여부를 자동으로 확인

## 실행 방법

### 사전 요구사항

- Go 1.21 이상
- Fyne 의존성 (자동 설치됨)

### 빌드 및 실행

```bash
git clone https://github.com/Taewoo-HN/b64checker.git
cd b64checker
go run main.go
```

또는 바이너리로 빌드:

```bash
go build -o b64checker .
./b64checker
```

## 사용법

1. 입력창에 텍스트 또는 Base64 문자열을 붙여넣기
2. **Base64 인코딩** 버튼 — 입력 텍스트를 Base64로 인코딩
3. **디코딩 & URL 확인** 버튼 — Base64 디코딩 후, 결과가 URL이면 접속 여부까지 확인

## 기술 스택

- [Go](https://golang.org/)
- [Fyne v2](https://fyne.io/) — 크로스플랫폼 GUI 프레임워크

## 라이선스

MIT License — [LICENSE](./LICENSE) 파일 참고
