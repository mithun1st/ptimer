package config

import (
	"flag"
	"fmt"
	"time"
)

func InputCountDownTimeByFlag() (*time.Duration, bool, bool) {
	var seconds int
	var minutes int
	var hours int
	var enableNotification bool
	var enableRandomColor bool

	flag.IntVar(&seconds, "s", 0, "Input time as seconds.")
	flag.IntVar(&minutes, "m", 0, "Input time as minutes.")
	flag.IntVar(&hours, "h", 0, "Input time as hours.")
	flag.BoolVar(&enableNotification, "notify", true, "Notification toggle by true/false")
	flag.BoolVar(&enableRandomColor, "color", false, "Random color style by true/false")

	flag.Parse()

	totalSeconds := seconds + (minutes * 60) + (hours * 3600)

	if totalSeconds <= 0 {
		return nil, enableNotification, enableRandomColor
	}

	duration := time.Second * time.Duration(totalSeconds)

	return &duration, enableNotification, enableRandomColor

}

func InputByCli() (string, error) {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", err
	}
	return input, nil
}
