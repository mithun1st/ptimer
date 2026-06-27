package pkg

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func _mocOs(title string, message string) *exec.Cmd {
	script := fmt.Sprintf(`display notification "%s" with title "%s" sound name "default"`, message, title)
	return exec.Command("osascript", "-e", script)
}

func _linux(title string, message string) *exec.Cmd {

	soundCmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
	soundCmd.Run()

	return exec.Command("notify-send", title, message)
}

func _windows(title string, message string) *exec.Cmd {

	scprit := fmt.Sprintf(
		`[void][System.Reflection.Assembly]::LoadWithPartialName('System.Windows.Forms'); `+
			`$icon = New-Object System.Windows.Forms.NotifyIcon; `+
			`$icon.Icon = [System.Drawing.SystemIcons]::Information; `+
			`$icon.BalloonTipIcon = 'Info'; `+
			`$icon.BalloonTipTitle = '%s'; `+
			`$icon.BalloonTipText = '%s'; `+
			`$icon.Visible = $true; `+
			`$icon.ShowBalloonTip(10000);`,
		title, message,
	)
	return exec.Command("powershell", "-NoProfile", "-WindowStyle", "Hidden", "-Command", scprit)
}

func NotificationWithSound(title string, message string) *error {
	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "darwin": // macOS
		cmd = _mocOs(title, message)

	case "linux":
		cmd = _linux(title, message)

	case "windows":
		cmd = _windows(title, message)

	default:
		err := errors.New("unknown os")
		return &err
	}

	err := cmd.Run()
	if err != nil {
		return &err
	}

	return nil
}
