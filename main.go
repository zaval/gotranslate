package main

import (
	"flag"
	"fmt"
	"goTranslate/icon"
	"goTranslate/utils"
	"runtime"

	"github.com/getlantern/systray"
)

func onReady() {
	systray.SetTitle("goTranslate")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	mTranslate := systray.AddMenuItem("Translate", "translate the text")
	go func() {
		for {
			<-mTranslate.ClickedCh
			translated, err := utils.Translate(utils.GetSelectedText())
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

	var translateText string
	var selectedText bool
	flag.StringVar(&translateText, "t", "", "translate text")
	flag.BoolVar(&selectedText, "s", false, "translate currently selected text")
	flag.Parse()

	if translateText != "" {
		translated, err := utils.Translate(translateText)
		if err == nil {
			utils.ShowNotification(translated)
		}
		return
	}
	if selectedText {
		translateText = utils.GetSelectedText()
		translated, err := utils.Translate(translateText)
		if err == nil {
			utils.ShowNotification(translated)
		}
		return
	}

	if runtime.GOOS == "windows" {
		go utils.RegisterHotkey()
	}

	systray.Run(onReady, onExit)

}
