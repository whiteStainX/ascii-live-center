package frames

import (
	"strings"
	"time"
)

var Clock = FrameType{
	GetFrame:  getClockFrame,
	GetLength: getClockLength,
	GetSleep:  DefaultGetSleep(),
}

func getClockFrame(i int, terminalWidth int) string {
	t := time.Now().Format(time.RFC3339)
	padding := (terminalWidth - len(t)) / 2
	if padding < 0 {
		padding = 0
	}
	pad := strings.Repeat(" ", padding)
	return pad + t
}

func getClockLength() int {
	return 0
}
