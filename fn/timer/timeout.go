package timer

import "time"

func Deadline(correction float32, threshold time.Duration) func(time.Time, bool) time.Time {
	return func(deadline time.Time, present bool) time.Time {
		if !present {
			return deadline
		}
		now := time.Now()
		timeout := deadline.Sub(now)
		corrected := time.Duration(float32(timeout) * correction)
		if timeout-corrected > threshold {
			return now.Add(timeout - threshold)
		}
		return now.Add(corrected)
	}
}

func Timeout(correction float32, threshold time.Duration) func(time.Duration, bool) time.Duration {
	return func(timeout time.Duration, present bool) time.Duration {
		if !present {
			return timeout
		}
		corrected := time.Duration(float32(timeout) * correction)
		if timeout-corrected > threshold {
			return timeout - threshold
		}
		return corrected
	}
}
