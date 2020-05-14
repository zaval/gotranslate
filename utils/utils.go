package utils

import (
	"regexp"

	"github.com/bregydoc/gtranslate"
)

func Translate(inputText string) (string, error) {

	langFrom := "en"
	langTo := "ru"
	re := regexp.MustCompile(`([а-яА-ЯёЁ]{2,})`)
	res := re.FindString(inputText)
	if res != "" {
		langFrom = "ru"
		langTo = "en"
	}

	translated, err := gtranslate.TranslateWithParams(
		inputText,
		gtranslate.TranslationParams{
			From: langFrom,
			To:   langTo,
		},
	)
	if err != nil {
		return "", err
	}

	return translated, nil
}
