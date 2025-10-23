package frames

import (
	"math/rand"
	"time"
)

const titleFrame = `
 _____ _____ ____ _____   _     _     _   __      __
/  ___|  ___/ ___|_   _| | |   (_)   | |  \ \    / /
| |   | |__ \___ \ | |   | |   | |   | |   \ \  / /
| |   |  __| ___) || |   | |   | |   | |    \ \/ /
| |___| |___/____/ | |   | |___| |___| |___  \  /
\_____|_____|_____/|_|   \_____|_____|_____|  \/
`

var titleFrames []string

func generateGlitchedFrames(input string, count int) []string {
	frames := make([]string, count)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	glitchChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+', '-', '='}

	for i := 0; i < count; i++ {
		glitched := []rune(input)
		for j, _ := range glitched {
			if r.Float32() < 0.05 { // 5% chance to glitch
				glitched[j] = glitchChars[r.Intn(len(glitchChars))]
			}
		}
		frames[i] = string(glitched)
	}
	return frames
}

func init() {
	titleFrames = generateGlitchedFrames(titleFrame, 20)
	FrameMap["title"] = DefaultFrameType(titleFrames)
}
