package app

import (
	"math/rand/v2"
	"ptimer/internal/app/timer"
	"ptimer/internal/enum"
	"ptimer/pkg"
	"time"
)

func _generagteRandomColor() (enum.AnsiCodeType, enum.AnsiCodeType, enum.AnsiCodeType) {
	ansiColors := []enum.AnsiCodeType{
		enum.Black,
		enum.Red,
		enum.Green,
		enum.Yellow,
		enum.Blue,
		enum.Purple,
		enum.Cyan,
		enum.White,
	}

	var startColorIndex int
	var middleColorIndex int
	var endColorIndex int

	startColorIndex = rand.IntN(len(ansiColors))
	middleColorIndex = rand.IntN(len(ansiColors))
	endColorIndex = rand.IntN(len(ansiColors))

	if startColorIndex == middleColorIndex || middleColorIndex == endColorIndex || startColorIndex == endColorIndex {
		return _generagteRandomColor()
	}

	return ansiColors[startColorIndex], ansiColors[middleColorIndex], ansiColors[endColorIndex]
}

func RunCountdown(duration time.Duration, enableNotification bool, enableRandomColor bool) {

	var style enum.TimerTypeEntity

	if enableRandomColor {
		startColor, middleColor, endColor := _generagteRandomColor()

		style = enum.TimerTypeEntity{
			StartColor:  startColor,
			MiddleColor: middleColor,
			EndColor:    endColor,
			Icon:        enum.Countdown.Info().Icon,
			Title:       enum.Countdown.Info().Title,
			StartTitle:  enum.Countdown.Info().StartTitle,
			EndTitle:    enum.Countdown.Info().EndTitle,
		}
	} else {
		style = enum.Countdown.Info()
	}

	timer.TimerStart(duration, style)
	if enableNotification {
		pkg.NotificationWithSound(style.Title+style.Icon, style.EndTitle)
	}
}
