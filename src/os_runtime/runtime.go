package os_runtime

import "runtime"

const Windows = "windows"

// This is a very simplistic check, but can help verify simple system calls, and path structure
// If a operating system is not unix like, and not windows AND runs in golang please submit an issue
func IsUnixBased() bool {

	if runtime.GOOS == Windows {
		return false
	}

	return true
}
