package pkg

import (
	"fmt"
	"os/exec"
	"runtime"
)

func NotificationWithSound(title string, message string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "darwin": // macOS
		script := fmt.Sprintf(`display notification "%s" with title "%s" sound name "default"`, message, title)
		cmd = exec.Command("osascript", "-e", script)

	case "linux":
		cmd = exec.Command("notify-send", title, message)

	case "windows":
		script := `& {[System.Windows.Forms.MessageBox]::Show('` + message + `', '` + title + `')}`
		cmd = exec.Command("powershell", "-command", script)

	default:
		return
	}

	cmd.Run()
}
