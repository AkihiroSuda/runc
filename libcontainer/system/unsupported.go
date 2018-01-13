// +build !linux

package system

import (
	"os"
)

// RunningInUserNS is a stub for non-Linux systems
// Always returns false
func RunningInUserNS() bool {
	return false
}

// GetParentNSeuid returns the euid within the parent user namespace
// Always returns os.Geteuid on non-linux
func GetParentNSeuid() int {
	return os.Geteuid()
}
