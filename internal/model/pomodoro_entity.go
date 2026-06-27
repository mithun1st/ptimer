package model

import (
	"fmt"
	"time"
)

type PomodoroEntity struct {
	Icon         string
	Name         string
	WorkDuration time.Duration
	ShortBreak   time.Duration
	LongBreak    time.Duration
}

func (m PomodoroEntity) Title() string {
	return fmt.Sprint(m.Icon, " ", m.Name, " Focus: ", m.WorkDuration.Minutes(), "/", m.ShortBreak.Minutes(), "/", m.LongBreak.Minutes(), " min")
}

func (m PomodoroEntity) TitleWithSessionName() string {
	return fmt.Sprint(m.Icon, " ", m.Name, " (", m.WorkDuration.Minutes(), "m WORK, ", m.ShortBreak.Minutes(), "m SHORT BREAK, ", m.LongBreak.Minutes(), "m LONG BREAK)")
}

func QuickFocus() PomodoroEntity {
	return PomodoroEntity{
		Icon:         "⚡️",
		Name:         "Quick",
		WorkDuration: time.Minute * 25,
		ShortBreak:   time.Minute * 5,
		LongBreak:    time.Minute * 20,
	}
}
func DeepFocus() PomodoroEntity {
	return PomodoroEntity{
		Icon:         "🧠",
		Name:         "Deep",
		WorkDuration: time.Minute * 50,
		ShortBreak:   time.Minute * 10,
		LongBreak:    time.Minute * 25,
	}
}
func TurboFocus() PomodoroEntity {
	return PomodoroEntity{
		Icon:         "🚀",
		Name:         "Turbo",
		WorkDuration: time.Minute * 90,
		ShortBreak:   time.Minute * 15,
		LongBreak:    time.Minute * 30,
	}
}

func Test() PomodoroEntity {
	return PomodoroEntity{
		Icon:         "🧪",
		Name:         "Test",
		WorkDuration: time.Second * 5,
		ShortBreak:   time.Second * 2,
		LongBreak:    time.Second * 3,
	}
}
