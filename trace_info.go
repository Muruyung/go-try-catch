package gotry

import (
	"fmt"
	"runtime"
)

type traceInfo struct {
	pc   uintptr
	file string
	line int
}

// String return string format stacktrace info
func (info *traceInfo) String() string {
	fn := runtime.FuncForPC(info.pc)
	return fmt.Sprintf("%s\n\t%s:%d\n", fn.Name(), info.file, info.line)
}
