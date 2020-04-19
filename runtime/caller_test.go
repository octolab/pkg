package runtime_test

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "go.octolab.org/runtime"
)

func TestCaller(t *testing.T) {
	ahead := func(t *testing.T, current GoVersion, target struct {
		version GoVersion
		release string
	}) bool {
		if current.Equal(target.version) {
			return true
		}
		if !unstable(current.Raw) {
			return current.Later(target.version)
		}
		prefix := "devel +61170f85e6 "
		layout := "Mon Jan 2 15:04:05 2006 -0700"
		release, _ := time.Parse(layout, target.release)
		control, _ := time.Parse(layout, current.Raw[len(prefix):])
		t.Log(target.release, "->", release, "<->", control, "<-", current.Raw[len(prefix):])
		return control.After(release)
	}

	tests := map[string]struct {
		caller   func() CallerInfo
		expected string
	}{
		"direct caller": {
			callerA,
			"go.octolab.org/runtime_test.callerA",
		},
		"direct caller (alt)": {
			altCallerA,
			"go.octolab.org/runtime_test.altCallerA",
		},
		"chain caller": {
			callerB,
			"go.octolab.org/runtime_test.callerA",
		},
		"chain caller (alt)": {
			altCallerB,
			"go.octolab.org/runtime_test.altCallerA",
		},
		"lambda caller": {
			callerC,
			func() string {
				if ahead(t, Version(), go112) {
					// https://golang.org/doc/go1.12#runtime
					return "go.octolab.org/runtime_test.callerC"
				}
				return "go.octolab.org/runtime_test.callerC.func1"
			}(),
		},
		"lambda caller (alt)": {
			altCallerC,
			func() string {
				return "go.octolab.org/runtime_test.altCallerC.func1"
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.caller().Name)
		})
	}
}

// BenchmarkCaller/direct_caller-4         	 3024363	       392 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/direct_caller_(alt)-4   	 1881876	       651 ns/op	     216 B/op	       2 allocs/op
// BenchmarkCaller/chain_caller-4          	 3015570	       382 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/chain_caller_(alt)-4    	 1732618	       647 ns/op	     216 B/op	       2 allocs/op
// BenchmarkCaller/lambda_caller-4         	 2273673	       516 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/lambda_caller_(alt)-4   	 1224558	       973 ns/op	     280 B/op	       3 allocs/op
func BenchmarkCaller(b *testing.B) {
	benchmarks := []struct {
		name   string
		caller func() CallerInfo
	}{
		{"direct caller", callerA},
		{"direct caller (alt)", altCallerA},
		{"chain caller", callerB},
		{"chain caller (alt)", altCallerB},
		{"lambda caller", callerC},
		{"lambda caller (alt)", altCallerC},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = bm.caller()
			}
		})
	}
}

func alternateCaller() CallerInfo {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	return CallerInfo{Name: f.Name(), File: file, Line: line}
}

//go:noinline
func callerA() CallerInfo {
	return Caller()
}

func callerB() CallerInfo {
	return callerA()
}

func callerC() CallerInfo {
	return func() CallerInfo {
		return Caller()
	}()
}

//go:noinline
func altCallerA() CallerInfo {
	return alternateCaller()
}

func altCallerB() CallerInfo {
	return altCallerA()
}

func altCallerC() CallerInfo {
	return func() CallerInfo {
		return alternateCaller()
	}()
}
