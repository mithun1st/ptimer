package utils

import (
	"fmt"
	"ptimer/pkg"
	"strings"
)

func ProgressBar(
	progressPercent float32,
	progressAnsiColor string,
	diffAnsiColor string,
	fillAnsiColor string,
	textAnsiColor string,
) {

	size, _ := pkg.TerminalSize()
	widthSize := max(int(size.Col)-40, 0)

	var progress int = int(progressPercent * float32(widthSize))

	//* Start
	pkg.Log("\r[")
	pkg.Log(progressAnsiColor, strings.Repeat(" ", progress))

	//* Middle
	if progressPercent != 1 {
		pkg.Log(diffAnsiColor, " ")
	}

	//* End
	pkg.Log(fillAnsiColor, strings.Repeat(" ", widthSize-progress))
	pkg.Log("]")

	percent := int(progressPercent * 100)
	pkg.Log(textAnsiColor, fmt.Sprintf(" %02d%% ", percent))
	pkg.Log("|")
}
