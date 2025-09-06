package taskTwentyFive

import "time"

func Sleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	_ = <-timer.C
}
