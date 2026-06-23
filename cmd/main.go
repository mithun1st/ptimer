package main

import (
	"ptimer/internal/app"
	"ptimer/internal/config"
)

func main() {

	// CountDown
	timeDuration, enableNotification, enableRandomColor := config.InputCountDownTimeByFlag()
	if timeDuration != nil {
		app.RunCountdown(*timeDuration, enableNotification, enableRandomColor)
		return
	}

	// Pomodoro
	app.RunPomodoro()
}
