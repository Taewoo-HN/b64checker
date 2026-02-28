package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Base64 디코더 & URL 체크")
	w.Resize(fyne.NewSize(520, 350))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Base64 문자열을 입력하세요...")
	input.SetMinRowsVisible(3)

	result := widget.NewLabel("")
	result.Wrapping = fyne.TextWrapWord

	btn := widget.NewButton("디코딩 & URL 확인", func() {
		text := input.Text
		if text == "" {
			result.SetText("입력값이 없습니다.")
			return
		}

		decoded, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			result.SetText(fmt.Sprintf("⚠️ 디코딩 실패: 올바른 Base64 형식이 아닙니다.\n(%v)", err))
			return
		}

		decodedStr := string(decoded)
		output := fmt.Sprintf("✅ 디코딩 결과:\n%s\n\n", decodedStr)

		u, err := url.ParseRequestURI(decodedStr)
		if err != nil || u.Scheme == "" || u.Host == "" {
			output += "ℹ️ URL 형식이 아니므로 접속 확인을 건너뜁니다."
			result.SetText(output)
			return
		}

		result.SetText(output + "🔍 접속 확인 중...")

		go func() {
			client := http.Client{Timeout: 5 * time.Second}
			resp, err := client.Head(decodedStr)
			if err != nil {
				fyne.Do(func() {
					result.SetText(output + fmt.Sprintf("❌ 접속 오류: %v", err))
				})
				return
			}
			resp.Body.Close()

			if resp.StatusCode < 400 {
				fyne.Do(func() {
					result.SetText(output + fmt.Sprintf("🌐 접속 성공! (상태 코드: %d)", resp.StatusCode))
				})
			} else {
				fyne.Do(func() {
					result.SetText(output + fmt.Sprintf("🚫 접속 실패 (상태 코드: %d)", resp.StatusCode))
				})
			}
		}()
	})

	w.SetContent(container.NewVBox(
		widget.NewLabel("Base64 디코더 & URL 체크 도구"),
		input,
		btn,
		widget.NewSeparator(),
		result,
	))

	w.ShowAndRun()
}
