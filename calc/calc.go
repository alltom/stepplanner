package calc

import (
	"time"
)

func GetSteps(stepGoal int, now, dayStart, dayEnd time.Time) int {
	progress := now.Sub(dayStart).Seconds() / dayEnd.Sub(dayStart).Seconds()
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	return int(progress * float64(stepGoal))
}
