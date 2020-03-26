package timer

import "time"

// Deadline return a time corrector.
//
// It useful for go.octolab.org/toolkit/protocol/http/header.Deadline
// and go.octolab.org/toolkit/protocol/http/middleware.Deadline.
//
//  mux := http.NewServeMux()
//
//  sla := middleware.Deadline(time.Second, timer.Deadline(0.95, time.Millisecond))
//  handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
//  mux.Handle("/", sla(handler))
//
//  log.Fatal(http.ListenAndServe(":80", mux))
//
// It also useful for the built-in context package.
//
//  func(rw http.ResponseWriter, req *http.Request) {
//  	ctx := req.Context()
//
//  	corrector := timer.Deadline(0.95, time.Millisecond)
//  	ctx, cancel := context.WithDeadline(req.Context(), corrector(ctx.Deadline()))
//  	defer cancel()
//
//  	select {
//  	case <-ctx.Done():
//  		message := http.StatusText(http.StatusRequestTimeout)
//  		http.Error(rw, message, http.StatusRequestTimeout)
//  	case result := <-rpc.Call(ctx):
//  		unsafe.Ignore(json.NewEncoder(rw).Encode(result))
//  	}
//  }
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

// Timeout return a duration corrector.
//
// It useful for go.octolab.org/toolkit/protocol/http/header.Timeout
// and go.octolab.org/toolkit/protocol/http/middleware.Timeout.
//
//  mux := http.NewServeMux()
//
//  sla := middleware.Timeout(time.Second, timer.Timeout(0.95, time.Millisecond))
//  handler := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
//  mux.Handle("/", sla(handler))
//
//  log.Fatal(http.ListenAndServe(":80", mux))
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
