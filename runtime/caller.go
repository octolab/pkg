package runtime

import (
	"path"
	"runtime"
	"strings"
)

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

func (info CallerInfo) Meta() (pkg, receiver, method string) {
	base, raw := path.Split(info.Name)
	parts := strings.Split(raw, ".")
	if len(parts) == 3 {
		return base + parts[0], strings.Trim(parts[1], "()"), parts[2]
	}
	return base + parts[0], "", parts[1]
}
