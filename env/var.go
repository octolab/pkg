package env

import "strings"

// Environment variables relate to the built-in runtime package.
// See https://pkg.go.dev/runtime/#hdr-Environment_Variables.
const (
	Go111Module = "GO111MODULE"
	GoArch      = "GOARCH"
	GoDebug     = "GODEBUG"
	GoGC        = "GOGC"
	GoMaxProcs  = "GOMAXPROCS"
	GoOS        = "GOOS"
	GoPath      = "GOPATH"
	GoRace      = "GORACE"
	GoRoot      = "GOROOT"
	GoTraceback = "GOTRACEBACK"
)

// Must does the same as the New but throws panic
// if something went wrong.
func Must(key, value string) Variable {
	// it will be changed in the future
	v, _ := New(key, value)
	return v
}

// New returns a new Variable on top of the passed key/value pair.
// It returns error if something went wrong.
func New(key, value string) (Variable, error) {
	// it will be changed in the future
	return Variable{name: key, value: value}, nil
}

// From converts an environment in the form "key=value" into Environment.
// It skips invalid entries silently.
//
// See the example.
func From(environ []string) Environment {
	env := make(Environment, 0, len(environ))
	for _, kv := range environ {
		pos := strings.Index(kv, sep)
		if pos == -1 {
			continue
		}
		if v, err := New(kv[:pos], kv[pos+1:]); err == nil {
			env = append(env, v)
		}
	}
	return env
}

// Variable represents an environment variable.
type Variable struct {
	name, value string
}

// Name returns an environment variable name.
func (v Variable) Name() string {
	return v.name
}

// String returns a string representation of an
// environment variable in the form "key=value".
func (v Variable) String() string {
	return v.name + sep + v.value
}

// Value returns an environment variable value.
func (v Variable) Value() string {
	return v.value
}

// Environment represents a set of environment variables.
type Environment []Variable

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
func (env Environment) Environ() []string {
	entries := make([]string, 0, len(env))
	for _, v := range env {
		entries = append(entries, v.String())
	}
	return entries
}

// Lookup retrieves the value of the environment variable named
// by the key. If the variable is present in the environment, the
// value (which may be empty) is returned, and the boolean is true.
// Otherwise, the returned value will be empty, and the boolean will
// be false.
func (env Environment) Lookup(key string) (Variable, bool) {
	for _, v := range env {
		if v.name == key {
			return v, true
		}
	}
	return Variable{}, false
}

const sep = "="
