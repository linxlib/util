package file

import (
	"os"
	"os/exec"
	"path/filepath"
)

// Exists checks whether given <path> exist.
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}
func Rename(srcFile string, dstFile string) error {
	var cmd *exec.Cmd
	cmd = exec.Command("mv", srcFile, dstFile)
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

func GetCurrentPath() string {
	if ex, err := os.Executable(); err == nil {
		return filepath.Dir(ex)
	}
	return "."
}
