package enum

import (
	"ptimer/internal/model"
	"time"
)

// * Color
type AnsiCodeType string

const (
	// FG
	Black  AnsiCodeType = "\033[90m"
	Red    AnsiCodeType = "\033[91m"
	Green  AnsiCodeType = "\033[92m"
	Yellow AnsiCodeType = "\033[93m"
	Blue   AnsiCodeType = "\033[94m"
	Purple AnsiCodeType = "\033[95m"
	Cyan   AnsiCodeType = "\033[96m"
	White  AnsiCodeType = "\033[97m"

	//
	Raw       AnsiCodeType = "\033[0m"
	Underline AnsiCodeType = "\033[4m"
	Blink     AnsiCodeType = "\033[5m"
	BeepSound AnsiCodeType = "\a"
)

func (t AnsiCodeType) Bg() string {
	switch t {
	case Black:
		return "\033[100m"
	case Red:
		return "\033[101m"
	case Green:
		return "\033[102m"
	case Yellow:
		return "\033[103m"
	case Blue:
		return "\033[104m"
	case Purple:
		return "\033[105m"
	case Cyan:
		return "\033[106m"
	case White:
		return "\033[107m"
	}
	return ""
}

// * Confirm Type
type ConfirmType string

const (
	YesSilently ConfirmType = "🔕 Yes (silently)"
	YesNotify   ConfirmType = "📢 Yes (notify)"
	ViewLog     ConfirmType = "📜 Vew Log"
)

// * TimerType
type TimerType int

const (
	Work TimerType = iota + 1
	ShortBreak
	LongBreak
	Countdown
)

type TimerTypeEntity struct {
	StartColor  AnsiCodeType
	MiddleColor AnsiCodeType
	EndColor    AnsiCodeType

	Icon       string
	Title      string
	StartTitle string
	EndTitle   string
}

func (t TimerType) Info() TimerTypeEntity {
	switch t {
	case Work:
		return TimerTypeEntity{
			StartColor:  Blue,
			MiddleColor: Yellow,
			EndColor:    Cyan,
			Icon:        "🍅",
			Title:       "WORK",
			StartTitle:  "Are you ready to lock in? Let's go!",
			EndTitle:    "Fantastic focus! Break time awaits.",
		}
	case ShortBreak:
		return TimerTypeEntity{
			StartColor:  Cyan,
			MiddleColor: Green,
			EndColor:    Yellow,
			Icon:        "⏳",
			Title:       "SHORT BREAK",
			StartTitle:  "Take a short break. Stretch and relax.",
			EndTitle:    "Break finished. Get back to work",
		}
	case LongBreak:
		return TimerTypeEntity{
			StartColor:  Green,
			MiddleColor: Cyan,
			EndColor:    Red,
			Icon:        "☕️",
			Title:       "LONG BREAK",
			StartTitle:  "Long break unlocked, you've completed sessions.",
			EndTitle:    "Reset complete. Let's start another pomodoro.",
		}
	case Countdown:
		return TimerTypeEntity{
			StartColor:  White,
			MiddleColor: Blink,
			EndColor:    Black,
			Icon:        "⏱️",
			Title:       "COUNTDOWN",
			StartTitle:  "Stopwatch running...",
			EndTitle:    "Times up!",
		}
	}
	return TimerTypeEntity{}
}

func (t TimerType) GetDuration(item model.PomodoroEntity) time.Duration {
	switch t {
	case Work:
		return item.WorkDuration
	case ShortBreak:
		return item.ShortBreak
	case LongBreak:
		return item.LongBreak
	}
	return time.Duration(0)
}

// func (tt TimerType) StartColor() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) MiddleColor() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) EndColor() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) Icon() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) Title() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) StartTitle() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
// func (tt TimerType) EndTitle() string {
// 	switch tt {
// 	case Work:
// 	case ShortBreak:
// 	case LongBreak:
// 	case Countdown:
// 	}
// }
