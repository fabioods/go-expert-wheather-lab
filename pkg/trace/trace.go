package trace

import "runtime"

type TraceInfo struct {
	File     string
	Line     int
	Function string
}

func GetTrace() TraceInfo {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return TraceInfo{
		File:     frame.File,
		Line:     frame.Line,
		Function: frame.Function,
	}
}
