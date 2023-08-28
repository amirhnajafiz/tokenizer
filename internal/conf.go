package internal

import "os"

const baseFile = "conf.txt"

// CheckFile returns true if conf file exists
func CheckFile() bool {
	_, err := os.Stat(baseFile)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

// InitFile creates conf file
func InitFile() error {
	return nil
}
