package utils

import (
	"path/filepath"
	"runtime"
)

func GetCurrentDir() string {
	_, callerFilepath, _, _ := runtime.Caller(1)
	current := filepath.Dir(callerFilepath)

	return current
}
