package internal

import (
	"os"
	"runtime"
	"strings"
)

func ExtractCodeData() (string, int, string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		file = ""
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	name := ""
	if fn != nil {
		fnName := fn.Name()
		name = fnName[strings.LastIndex(fnName, ".")+1:]
	}

	wd, _ := os.Getwd()
	relPath := strings.TrimPrefix(file, wd)

	return relPath, line, name
}

func ExtractCodeDataRecursive() (string, int, string) {
	relPath := ""
	line := 0
	name := ""

	wd, err := os.Getwd()
	if err != nil {
		return relPath, line, name
	}

	var pcs [10]uintptr
	n := runtime.Callers(3, pcs[:])
	frames := runtime.CallersFrames(pcs[:n])
	for {
		frame, more := frames.Next()
		frFile := frame.File
		frLine := frame.Line
		frName := frame.Function

		if strings.HasPrefix(frFile, wd) {
			relPath = strings.TrimPrefix(frFile, wd)
			line = frLine
			name = frName[strings.LastIndex(frName, ".")+1:]
			break
		}

		if !more {
			break
		}
	}

	return relPath, line, name
}