package pkg

import (
	"fmt"
	"ptimer/internal/enum"
)

func Log(ansiCode ...any) {
	for _, e := range ansiCode {
		fmt.Printf("%s", e)
	}
	fmt.Print(enum.Raw)
}

func LogWarning(value string) {
	Log(enum.Yellow.Bg(), string(enum.Black), "WARNING:")
	str := fmt.Sprintf(" %s\n", value)
	Log(enum.Yellow, str)
	Log(enum.BeepSound)
}
func LogError(value any) {
	Log(enum.Red.Bg(), "ERROR:")
	str := fmt.Sprintf(" %s\n", value)
	Log(enum.Red, str)
	Log(enum.BeepSound)
}
