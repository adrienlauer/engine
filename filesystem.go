package engine

import "os"

func ensureWritableDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
