package model

import (
	"fmt"
	"time"
)

type PomodoroEntity struct {
	Name         string
	WorkDuration time.Duration
	ShortBreak   time.Duration
	LongBreak    time.Duration
}

func (m PomodoroEntity) DurationTitle() string {
	return fmt.Sprintf("%.0f/%.0f/%.0f min", m.WorkDuration.Minutes(), m.ShortBreak.Minutes(), m.LongBreak.Minutes())
}

func QuickFocus() PomodoroEntity {
	workDuration := time.Minute * 25
	shortBreak := time.Minute * 5
	longBreak := time.Minute * 20
	return PomodoroEntity{
		Name:         "⚡️ Quick",
		WorkDuration: workDuration,
		ShortBreak:   shortBreak,
		LongBreak:    longBreak,
	}
}
func DeepFocus() PomodoroEntity {
	workDuration := time.Minute * 50
	shortBreak := time.Minute * 10
	longBreak := time.Minute * 25
	return PomodoroEntity{
		Name:         "🧠 Deep",
		WorkDuration: workDuration,
		ShortBreak:   shortBreak,
		LongBreak:    longBreak,
	}
}
func TurboFocus() PomodoroEntity {
	workDuration := time.Minute * 90
	shortBreak := time.Minute * 15
	longBreak := time.Minute * 30
	return PomodoroEntity{
		Name:         "🚀 Turbo",
		WorkDuration: workDuration,
		ShortBreak:   shortBreak,
		LongBreak:    longBreak,
	}
}

func Test() PomodoroEntity {
	workDuration := time.Second * 5
	shortBreak := time.Second * 2
	longBreak := time.Second * 3
	return PomodoroEntity{
		Name:         "🧪 Test",
		WorkDuration: workDuration,
		ShortBreak:   shortBreak,
		LongBreak:    longBreak,
	}
}
