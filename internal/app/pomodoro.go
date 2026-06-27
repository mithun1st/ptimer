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

func _viewLog(numOfSession int, pomodoro model.PomodoroEntity, log map[enum.TimerType]time.Duration) {

	// Title
	pkg.Log(pomodoro.TitleWithSessionName(), "\n")
	pkg.Log(enum.Underline, fmt.Sprintf("Session No:%s %d\n", enum.Raw, numOfSession))

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

	size, _ := pkg.TerminalSize()
	widthSize := max(int(size.Col)-20, 0)

	buildBar := func(t enum.TimerType) {
		percent := int((float32(log[t].Seconds()) / float32(totalSecond)) * float32(widthSize))
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

	var numOfSession int = 0
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
		numOfSession++
		sessions := []enum.TimerType{enum.Work}
		if numOfSession%constant.DefaultSession == 0 {
			sessions = append(sessions, enum.LongBreak)
		} else {
			sessions = append(sessions, enum.ShortBreak)
		}

		for i := 0; i < len(sessions); i++ {
			session := sessions[i]
			var duration time.Duration = session.GetDuration(item)

			// Input
			sessionTitle := fmt.Sprintf(
				"%s%s/%s%s(%s): %s",
				item.Icon, item.Name,
				session.Info().Icon, session.Info().Title, session.GetDuration(item), session.Info().StartTitle,
			)
			confirm := utils.ChooseOption(sessionTitle, arr, buildFnc, "End Focus")
			if confirm == nil {
				pkg.Log("\n")
				_viewLog(numOfSession, item, sessionLog)
				//END Focus
				focusEnable = false
				break
			}

			// Start Timer
			if *confirm == enum.YesSilently {
				timer.TimerStart(duration, session.Info())
			} else if *confirm == enum.YesNotify {
				timer.TimerStart(duration, session.Info())
				err := pkg.NotificationWithSound(session.Info().Title+session.Info().Icon, session.Info().EndTitle)
				if err != nil {
					// pkg.LogError(err)
				}
			} else if *confirm == enum.ViewLog {
				i--
				_viewLog(numOfSession, item, sessionLog)
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
		return e.Title()
	}

	for {
		item := utils.ChooseOption("Which Pomodoro mode would you like to start?", pomodoros, buildFnc, "Exit")
		if item == nil {
			pkg.Log(enum.Cyan, constant.DeveloperInfo)
			break //EXIT
		}
		_pomodoro(*item)
	}
}
