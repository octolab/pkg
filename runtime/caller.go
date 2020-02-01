package runtime

import "runtime"

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
