package main

import (
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/getlantern/systray"
	"goTranslate/icon"
	"goTranslate/utils"
	"regexp"
)

func translate(inputText string) (string, error) {

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

func onReady() {
	systray.SetTitle("goTranslate")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	mTranslate := systray.AddMenuItem("Translate", "translate the text")
	go func() {
		for {
			<-mTranslate.ClickedCh
			translated, err := translate(utils.GetSelectedText())
			if err == nil {
				utils.ShowNotification(translated)
			}
		}

	}()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the translator")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()

}

func onExit() {
	fmt.Println("EXIT")
}

func main() {

	systray.RunWithAppWindow("Translate", 1024, 768, onReady, onExit)
}
