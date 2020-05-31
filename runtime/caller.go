package runtime

import (
	"path"
	"regexp"
	"runtime"
	"strings"
)

var lambda = regexp.MustCompile(`func\d+$`)

// Caller returns information about the current caller.
//
//  func StoreToDatabase(data Payload) error {
//  	defer stats.NewTiming().Send(runtime.Caller().Name)
//
//  	...
//  }
//
func Caller() CallerInfo {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return CallerInfo{f.Name(), file, line}
}

// CallerInfo holds information about a caller.
type CallerInfo struct {
	Name string
	File string
	Line int
}

// Meta returns package, receiver and method names.
func (info CallerInfo) Meta() (pkg, receiver, method string) {
	base, raw := path.Split(info.Name)
	parts := strings.Split(raw, ".")
	if len(parts) == 4 {
		return base + parts[0], strings.Trim(parts[1], "()"), parts[2] + "." + parts[3]
	}
	if len(parts) == 3 {
		if strings.HasPrefix(parts[2], "func") && lambda.MatchString(parts[2]) {
			return base + parts[0], "", parts[1] + "." + parts[2]
		}
		return base + parts[0], strings.Trim(parts[1], "()"), parts[2]
	}
	return base + parts[0], "", parts[1]
}
