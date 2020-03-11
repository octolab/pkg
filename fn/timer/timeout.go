package timer

import "time"

// Deadline return a corrector for go.octolab.org/toolkit/protocol/http/header.Deadline.
// It useful for go.octolab.org/toolkit/protocol/http/middleware.Deadline.
//
//  mux := http.NewServeMux()
//
//  {
//  	handler := http.HandlerFunc(...)
//  	sla := middleware.Deadline(time.Second, timer.Deadline(0.95, time.Millisecond))
//  	mux.Handle("/", sla(handler))
//  }
//
//  log.Fatal(http.ListenAndServe(":80", http.TimeoutHandler(mux, time.Second, "...")))
//
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

// Timeout return a corrector for go.octolab.org/toolkit/protocol/http/header.Timeout.
// It useful for go.octolab.org/toolkit/protocol/http/middleware.Timeout.
//
//  mux := http.NewServeMux()
//
//  {
//  	handler := http.HandlerFunc(...)
//  	sla := middleware.Timeout(time.Second, timer.Timeout(0.95, time.Millisecond))
//  	mux.Handle("/", sla(handler))
//  }
//
//  log.Fatal(http.ListenAndServe(":80", http.TimeoutHandler(mux, time.Second, "...")))
//
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
