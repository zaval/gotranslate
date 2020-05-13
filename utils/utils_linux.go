package utils

import (
	"fmt"
	"os/exec"
)

func GetSelectedText() string {
	cmd := exec.Command("xsel")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

func ShowNotification(text string) {
	cmd := exec.Command("notify-send", text)
	cmd.CombinedOutput()
}
