package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetSelectedText() string {

	var cmd *exec.Cmd
	if os.Getenv("XDG_SESSION_TYPE") == "wayland" {
		cmd = exec.Command("wl-paste", "-p")
	} else {
		cmd = exec.Command("xsel")
	}

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

func RegisterHotkey() {}
