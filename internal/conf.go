package internal

import (
	"fmt"
	"os"
)

const baseFile = "conf.txt"

func getBaseFile() string {
	tmp := os.Getenv("TK_PATH")
	if len(tmp) == 0 {
		return baseFile
	}

	return fmt.Sprintf("%s/%s", tmp, baseFile)
}

// CheckFile returns true if conf file exists
func CheckFile() bool {
	_, err := os.Stat(getBaseFile())
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
	_, err := os.Create(getBaseFile())

	return err
}
