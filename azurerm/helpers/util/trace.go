package util

import (
	"runtime"
)

func GetCallFuncName(depth int) string {
	if depth <= 0 {
		depth = 1
	}

	function, _, _, _ := runtime.Caller(depth)

	return runtime.FuncForPC(function).Name()
}
