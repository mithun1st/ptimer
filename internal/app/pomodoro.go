package app

import (
	"fmt"
	"ptimer/internal/app/timer"
	"ptimer/internal/constant"
	"ptimer/internal/enum"
	"ptimer/internal/model"
	"ptimer/internal/utils"
	"ptimer/pkg"
	"time"
)

func _viewLog(cycle int, pomodoro model.PomodoroEntity, log map[enum.TimerType]time.Duration) {

	// Title
	title2 := fmt.Sprintf(
		"(%s %s, %s %s, %s %s).",
		enum.Work.GetDuration(pomodoro), enum.Work.Info().Title,
		enum.ShortBreak.GetDuration(pomodoro), enum.ShortBreak.Info().Title,
		enum.LongBreak.GetDuration(pomodoro), enum.LongBreak.Info().Title,
	)
	pkg.Log(pomodoro.Name+" Focus ", title2, "\n")
	pkg.Log(enum.Underline, fmt.Sprintf("CYCLE: %d\n", cycle))

	// Duration Info
	buildTitle := func(t enum.TimerType, d time.Duration) {
		color := string(t.Info().EndColor)
		if t == enum.Work {
			color = string(t.Info().StartColor)
		}
		durationStr := fmt.Sprintf("%s", d)
		pkg.Log(color, t.Info().Icon, " ", t.Info().Title, ": ", durationStr, "\n")
	}
	workType := enum.Work
	buildTitle(workType, log[workType])
	shortBreakType := enum.ShortBreak
	buildTitle(shortBreakType, log[shortBreakType])
	longBreakType := enum.LongBreak
	buildTitle(longBreakType, log[longBreakType])

	// Graph
	totalSecond := int((log[workType] + log[shortBreakType] + log[longBreakType]).Seconds())
	limit := constant.ProgressBarSize
	buildBar := func(t enum.TimerType) {
		percent := int((float32(log[t].Seconds()) / float32(totalSecond)) * float32(limit))
		color := t.Info().EndColor.Bg()
		if t == enum.Work {
			color = t.Info().StartColor.Bg()
		}

		for i := range percent {
			if percent/2 == i {
				pkg.Log(enum.Blink, color, t.Info().Icon)
			} else {
				pkg.Log(color, " ")
			}
		}
	}

	pkg.Log("[")
	buildBar(workType)
	buildBar(shortBreakType)
	buildBar(longBreakType)
	pkg.Log("]\n\n")
}

func _pomodoro(item model.PomodoroEntity) {

	var cycle int = 0
	var focusEnable = true
	var sessionLog map[enum.TimerType]time.Duration = map[enum.TimerType]time.Duration{
		enum.Work:       time.Duration(0),
		enum.ShortBreak: time.Duration(0),
		enum.LongBreak:  time.Duration(0),
	}

	// Input
	arr := []enum.ConfirmType{enum.YesSilently, enum.YesNotify, enum.ViewLog}
	buildFnc := func(e enum.ConfirmType) string {
		return string(e)
	}

	for focusEnable {
		cycle++
		sessions := []enum.TimerType{enum.Work}
		if cycle%constant.DefaultSession == 0 {
			sessions = append(sessions, enum.LongBreak)
		} else {
			sessions = append(sessions, enum.ShortBreak)
		}

		for i := 0; i < len(sessions); i++ {
			session := sessions[i]
			var duration time.Duration = session.GetDuration(item)

			// Input
			inputTitle := session.Info().Icon + session.Info().Title + ": " + session.Info().StartTitle
			confirm := utils.ChooseOption(inputTitle, arr, buildFnc, "End Focus")
			if confirm == nil {
				pkg.Log("\n")
				_viewLog(cycle, item, sessionLog)
				//END Focus
				focusEnable = false
				break
			}

			// Start Timer
			if *confirm == enum.YesSilently {
				timer.TimerStart(duration, session.Info())
			} else if *confirm == enum.YesNotify {
				timer.TimerStart(duration, session.Info())
				pkg.NotificationWithSound(session.Info().Title+session.Info().Icon, session.Info().EndTitle)
			} else if *confirm == enum.ViewLog {
				i--
				_viewLog(cycle, item, sessionLog)
				continue
			}
			pkg.Log("\n")

			sessionLog[session] = duration + sessionLog[session]
		}
	}
}

func RunPomodoro() {

	pomodoros := []model.PomodoroEntity{
		model.QuickFocus(),
		model.DeepFocus(),
		model.TurboFocus(),
		// model.Test(),
	}
	buildFnc := func(e model.PomodoroEntity) string {
		return e.Name + " Focus: " + e.DurationTitle()
	}

	for {
		item := utils.ChooseOption("Which Pomodoro mode would you like to start?", pomodoros, buildFnc, "Exit")
		if item == nil {
			break //EXIT
		}
		_pomodoro(*item)
	}
}
