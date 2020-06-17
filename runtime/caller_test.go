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

	type expected struct {
		name                  string
		pkg, receiver, method string
	}

	tests := map[string]struct {
		caller   func() CallerInfo
		expected expected
	}{
		"direct method call": {
			new(structure).callA,
			expected{
				"go.octolab.org/runtime_test.structure.callA",
				"go.octolab.org/runtime_test", "structure", "callA",
			},
		},
		"proxy method call": {
			new(structure).proxyCallA,
			expected{
				"go.octolab.org/runtime_test.structure.callA",
				"go.octolab.org/runtime_test", "structure", "callA",
			},
		},
		"direct method call, pointer": {
			new(structure).callB,
			expected{
				"go.octolab.org/runtime_test.(*structure).callB",
				"go.octolab.org/runtime_test", "*structure", "callB",
			},
		},
		"proxy method call, pointer": {
			new(structure).proxyCallB,
			expected{
				"go.octolab.org/runtime_test.(*structure).callB",
				"go.octolab.org/runtime_test", "*structure", "callB",
			},
		},
		"deep dive call": {
			new(structure).callC,
			expected{
				"go.octolab.org/runtime_test.structure.callC.func2",
				"go.octolab.org/runtime_test", "structure", "callC.func2",
			},
		},
		"call by function type": {
			function(Caller).call,
			func() expected {
				if ahead(t, Version(), go112) {
					// https://golang.org/doc/go1.12#runtime
					return expected{
						"go.octolab.org/runtime_test.function.call-fm",
						"go.octolab.org/runtime_test", "function", "call-fm",
					}
				}
				return expected{
					"go.octolab.org/runtime_test.function.call",
					"go.octolab.org/runtime_test", "function", "call",
				}
			}(),
		},
		"call by primitive type": {
			new(integer).call,
			func() expected {
				if ahead(t, Version(), go112) {
					// https://golang.org/doc/go1.12#runtime
					return expected{
						"go.octolab.org/runtime_test.integer.call-fm",
						"go.octolab.org/runtime_test", "integer", "call-fm",
					}
				}
				return expected{
					"go.octolab.org/runtime_test.integer.call",
					"go.octolab.org/runtime_test", "integer", "call",
				}
			}(),
		},
		"direct function call": {
			callA,
			expected{
				"go.octolab.org/runtime_test.callA",
				"go.octolab.org/runtime_test", "", "callA",
			},
		},
		"direct function call (alt)": {
			altCallA,
			expected{
				"go.octolab.org/runtime_test.altCallA",
				"go.octolab.org/runtime_test", "", "altCallA",
			},
		},
		"proxy function call": {
			proxyCallA,
			expected{
				"go.octolab.org/runtime_test.callA",
				"go.octolab.org/runtime_test", "", "callA",
			},
		},
		"proxy function call (alt)": {
			altProxyCallA,
			expected{
				"go.octolab.org/runtime_test.altCallA",
				"go.octolab.org/runtime_test", "", "altCallA",
			},
		},
		"lambda call": {
			callB,
			func() expected {
				if ahead(t, Version(), go112) {
					// https://golang.org/doc/go1.12#runtime
					return expected{
						"go.octolab.org/runtime_test.callB",
						"go.octolab.org/runtime_test", "", "callB",
					}
				}
				return expected{
					"go.octolab.org/runtime_test.callB.func1",
					"go.octolab.org/runtime_test", "", "callB.func1",
				}
			}(),
		},
		"lambda call (alt)": {
			altCallB,
			expected{
				"go.octolab.org/runtime_test.altCallB.func1",
				"go.octolab.org/runtime_test", "", "altCallB.func1",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			caller := test.caller()
			pkg, receiver, method := caller.Meta()
			assert.Equal(t, test.expected, expected{caller.Name, pkg, receiver, method})
			assert.Equal(t, pkg, caller.PackageName())
			assert.Equal(t, receiver, caller.ReceiverName())
			assert.Equal(t, method, caller.MethodName())
		})
	}
}

// BenchmarkCaller/direct_function_call-4         	 3079887	       376 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/direct_function_call_(alt)-4   	 1883188	       628 ns/op	     216 B/op	       2 allocs/op
// BenchmarkCaller/proxy_function_call-4          	 3146109	       377 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/proxy_function_call_(alt)-4    	 1847580	       642 ns/op	     216 B/op	       2 allocs/op
// BenchmarkCaller/lambda_call-4                  	 2390946	       506 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCaller/lambda_call_(alt)-4            	 1188115	       946 ns/op	     280 B/op	       3 allocs/op
func BenchmarkCaller(b *testing.B) {
	benchmarks := []struct {
		name   string
		caller func() CallerInfo
	}{
		{"direct function call", callA},
		{"direct function call (alt)", altCallA},
		{"proxy function call", proxyCallA},
		{"proxy function call (alt)", altProxyCallA},
		{"lambda call", callB},
		{"lambda call (alt)", altCallB},
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

// helpers

type structure struct{}

//go:noinline prevent to inline in proxyCallA
func (structure) callA() CallerInfo {
	return Caller()
}

func (caller structure) proxyCallA() CallerInfo {
	return caller.callA()
}

//go:noinline prevent to inline in proxyCallB
func (*structure) callB() CallerInfo {
	return Caller()
}

func (caller *structure) proxyCallB() CallerInfo {
	return caller.callB()
}

//go:noinline prevent to inline lambda2 and Caller
func (structure) callC() CallerInfo {
	var lambda1, lambda2 func() CallerInfo
	lambda1 = func() CallerInfo {
		return lambda2()
	}
	lambda2 = func() CallerInfo { //nolint:gocritic
		return Caller()
	}
	return lambda1()
}

type function func() CallerInfo

func (fn function) call() CallerInfo {
	return fn()
}

type integer int

func (integer) call() CallerInfo {
	return Caller()
}

//go:noinline prevent to inline in proxyCallA
func callA() CallerInfo {
	return Caller()
}

func proxyCallA() CallerInfo {
	return callA()
}

func callB() CallerInfo {
	return func() CallerInfo { //nolint:gocritic
		return Caller()
	}()
}

func altCaller() CallerInfo {
	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	return CallerInfo{Name: f.Name(), File: file, Line: line}
}

//go:noinline
func altCallA() CallerInfo {
	return altCaller()
}

func altProxyCallA() CallerInfo {
	return altCallA()
}

func altCallB() CallerInfo {
	return func() CallerInfo { //nolint:gocritic
		return altCaller()
	}()
}
