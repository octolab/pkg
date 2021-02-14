package env

import "strings"

// Environment variables relate to the built-in runtime package.
// See https://pkg.go.dev/runtime/#hdr-Environment_Variables.
const (
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
	return Variable{key: key, value: value}, nil
}

// From converts ...
func From(environ []string) Vars {
	vars := make(Vars, 0, len(environ))
	for _, env := range environ {
		kv := strings.Split(env, sep)
		if len(kv) != 2 {
			continue
		}
		if v, err := New(kv[0], kv[1]); err == nil {
			vars = append(vars, v)
		}
	}
	return vars
}

// Variable ...
type Variable struct {
	key, value string
}

// Key ...
func (v Variable) Key() string {
	return v.key
}

// String ...
func (v Variable) String() string {
	return v.key + sep + v.value
}

// Value ...
func (v Variable) Value() string {
	return v.value
}

// Vars ...
type Vars []Variable

// Environ ...
func (vv Vars) Environ() []string {
	env := make([]string, 0, len(vv))
	for _, v := range vv {
		env = append(env, v.String())
	}
	return env
}

// Lookup ...
func (vv Vars) Lookup(key string) (Variable, bool) {
	key = strings.ToUpper(key)
	for _, v := range vv {
		if v.key == key {
			return v, true
		}
	}
	return Variable{}, false
}

const sep = "="
