package utils

import (
	"fmt"
	"ptimer/internal/config"
	"ptimer/internal/enum"
	"ptimer/pkg"
	"strconv"
	"strings"
)

func ChooseOption[T1 any](title string, items []T1, build func(T1) string, exitTitle string) *T1 {

	// Print title
	pkg.Log(enum.Black.Bg(), title)

	// Print items
	elements := make([]string, len(items))
	for i, model := range items {
		elements[i] = fmt.Sprintf("%d. %s", i+1, build(model))
	}
	pkg.Log("\n", strings.Join(elements, "\n")+"\n")

	// Print 0. quite text
	if exitTitle != "" {
		pkg.Log(enum.Red, "0. ", exitTitle, "\n")
	}

	// Print input title
	str := fmt.Sprintf("👉 Select option (%d-%d):", 1, len(items))
	pkg.Log(enum.Underline, enum.Black, str, enum.Raw, " ")

	// Input UI
	input, _ := config.InputByCli()
	index, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		pkg.LogError((err))
		pkg.Log("\n")
		return ChooseOption(title, items, build, exitTitle)
	}

	// Press 0 to quite
	if exitTitle != "" && index == 0 {
		return nil
	}

	// Range check
	if index-1 < 0 || len(items) < int(index) {
		pkg.LogWarning("Wrong option!")
		pkg.Log("\n")
		return ChooseOption(title, items, build, exitTitle)
	}

	pkg.Log("\n")
	return &items[index-1]
}
