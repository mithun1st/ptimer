package timer

import (
	"fmt"
	"ptimer/internal/enum"
	"ptimer/internal/utils"
	"ptimer/pkg"
	"time"
)

func _runThread(t time.Duration, style enum.TimerTypeEntity) {

	totalSeconds := int(t.Seconds())

	for seconds := totalSeconds; seconds >= 0; seconds-- {
		minutes := seconds / 60
		secs := seconds % 60

		// Progress bar
		fraction := float32(totalSeconds-seconds) / float32(totalSeconds)

		utils.ProgressBar(
			fraction,
			style.StartColor.Bg(),
			style.MiddleColor.Bg(),
			style.EndColor.Bg(),
			string(style.StartColor),
		)

		// Count down
		str := fmt.Sprintf(" %02d:%02d/%v %s\t", minutes, secs, t, style.Icon)
		pkg.Log(style.EndColor, str)

		// beep sound
		if seconds < 3 {
			pkg.Log(enum.BeepSound)
		}

		// Await
		time.Sleep(time.Second)
	}
	pkg.Log("\n")
}

func TimerStart(duration time.Duration, style enum.TimerTypeEntity) {
	waitTime := time.Now().Add(duration)

	str1 := string(style.Icon) + " " + string(style.Title) + "\t\t" +
		"⏰ " + waitTime.Format(time.TimeOnly)

	pkg.Log(style.StartColor, str1, "\n")
	_runThread(duration, style)

	str2 := "🏁 " + style.EndTitle + "\n"
	pkg.Log(style.EndColor, str2)
}
