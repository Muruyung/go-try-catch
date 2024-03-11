package gotry

import (
	"fmt"
	"runtime"
	"strings"
)

type stackInteractor struct {
	listInfo []traceInfo
}

func getStackTrace() StackTrace {
	listInfo := make([]traceInfo, 0)

	for i := 4; ; i++ {
		// if i == 5 || i == 6 {
		// 	continue
		// }

		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		listInfo = append(listInfo, traceInfo{
			pc:   pc,
			file: file,
			line: line,
		})
	}

	listInfo = listInfo[:len(listInfo)-2]
	return &stackInteractor{
		listInfo: listInfo,
	}
}

func (trace *stackInteractor) Print() {
	fmt.Println(trace.String())
}

func (trace *stackInteractor) String() string {
	var info string

	for _, val := range trace.listInfo {
		if strings.Contains(val.String(), "github.com/Muruyung/go-try-catch") {
			continue
		}
		info += val.String()
	}

	return info
}
