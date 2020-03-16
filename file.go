package main

import (
	"os"
)

// isExist check if file or directory exists
func isExist(path string) (bool, error) {
	var err error

	if _, err = appFs.Stat(path); !os.IsNotExist(err) {
		return true, err
	}

	return false, err
}
