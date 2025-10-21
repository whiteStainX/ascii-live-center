package frames

import (
	"strings"
	"time"
)

type FrameType struct {
	GetFrame  func(int, int) string
	GetLength func() int
	GetSleep  func() time.Duration
}

// Create a function that returns the next frame, based on length
func DefaultGetFrame(frames []string) func(int, int) string {
	return func(i int, terminalWidth int) string {
		frame := frames[i%(len(frames))]
		lines := strings.Split(frame, "\n")
		maxWidth := 0
		for _, line := range lines {
			if len(line) > maxWidth {
				maxWidth = len(line)
			}
		}

		padding := (terminalWidth - maxWidth) / 2
		if padding < 0 {
			padding = 0
		}
		pad := strings.Repeat(" ", padding)

		var centeredFrame strings.Builder
		for j, line := range lines {
			centeredFrame.WriteString(pad)
			centeredFrame.WriteString(line)
			if j < len(lines)-1 {
				centeredFrame.WriteString("\n")
			}
		}

		return centeredFrame.String()
	}
}

// Create a function that returns frame length
func DefaultGetLength(frames []string) func() int {
	return func() int {
		return len(frames)
	}
}

// Sleep time between frames
func DefaultGetSleep() func() time.Duration {
	return func() time.Duration {
		return time.Millisecond * 70
	}
}

// Given frames, create a FrameType with those frames
func DefaultFrameType(frames []string) FrameType {
	return FrameType{
		GetFrame:  DefaultGetFrame(frames),
		GetLength: DefaultGetLength(frames),
		GetSleep:  DefaultGetSleep(),
	}
}

var FrameMap = map[string]FrameType{
	"batman":          Batman,
	"batman-running":  BNR,
	"bnr":             BNR,
	"can-you-hear-me": Rick,
	"clock":           Clock,
	"coin":            Coin,
	"donut":           Donut,
	"dvd":             Dvd,
	"forrest":         Forrest,
	"hes":             HES,
	"knot":            TorusKnot,
	"nyan":            Nyan,
	"parrot":          Parrot,
	"rick":            Rick,
	"spidyswing":      Spidy,
	"torus-knot":      TorusKnot,
	"purdue":          Purdue,
	"as":              AStrend,
	"bomb":            Bomb,
	"maxwell":         Maxwell,
	"earth":           Earth,
	"kitty":           Kitty,
	"india":           India,
	"brittany":        Brittany,
}
