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

// defaultSourceLang returns the default source language if the given source language is empty.
// The default source language is "auto".
func defaultSourceLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return Auto
	}
	return s
}

// defaultTargetLang returns the default target language if the given target language is empty.
// The default target language is the Indonesian language.
func defaultTargetLang(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return LangID
	}
	return s
}

// Translate text from source language to target language
// https://cloud.google.com/translate/automl/docs/reference/rest/v3/projects.locations/translateText
func Translate(source, sourceLang, targetLang string) (string, error) {
	// handle panic
	defer go_helpers_error.PanicRecover("helpers-Translate")

	// prepare variables
	var (
		translation []interface{}
		text        []string
	)

	// encode source text
	encodedSource := url.QueryEscape(source)

	// prepare api url
	url := fmt.Sprintf("https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s",
		defaultSourceLang(sourceLang), defaultTargetLang(targetLang), encodedSource)

	// call api
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// unmarshal response body to struct
	err = json.Unmarshal(body, &translation)
	if err != nil {
		return "", err
	}

	// loop through response body
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
