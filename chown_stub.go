//go:build !linux && !darwin
// +build !linux,!darwin

package lumberjack

import (
	"os"
)

func chown(_ string, _ os.FileInfo) error {
	return nil
}
