package language

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	go_helpers_error "github.com/adamnasrudin03/go-helpers/error"
	"github.com/adamnasrudin03/go-helpers/net/url"
	"golang.org/x/text/language"
)

var (
	LangID = language.Indonesian.String()
	LangEn = language.English.String()
)

const (
	Auto = "auto"
)

func defaultSourceLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return Auto
	}
	return s
}

func defaultTargetLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return LangID
	}
	return s
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	defer go_helpers_error.PanicRecover("helpers-Translate")

	var (
		translation []interface{}
		text        []string
	)

	encodedSource := url.QueryEscape(source)
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		defaultSourceLang(sourceLang), defaultTargetLang(targetLang), encodedSource)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &translation)
	if err != nil {
		return "", err
	}

	if len(translation) > 0 {
		inner := translation[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		return strings.Join(text, ""), nil
	}

	return "", errors.New("no translated data in response")
}
