package main

import (
	"flag"
	"fmt"
	"runtime"

	"gotranslate/icon"
	"gotranslate/utils"

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
	flag.StringVar(&translateText, "t", "", "translate text")
	flag.Parse()

	if translateText != "" {
		tt := utils.GetSelectedText()
		translated, err := utils.Translate(tt)
		fmt.Println(translated)
		if err == nil {
			utils.ShowNotification(translated)
		}
		//fmt.Println(translated)
		return
	}

	if runtime.GOOS == "windows" {
		go utils.RegisterHotkey()
	}

	systray.Run(onReady, onExit)

}
