package helpers

import "runtime"

// Get the OS currently being used
func GetOS() string {
	os := runtime.GOOS
	return os
}
