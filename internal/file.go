package internal

import (
	"fmt"
	"os"
)

func exportToFile(path string, data ...string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("%v: %v", ErrFileCreation, err)
	}

	for _, item := range data {
		if _, err := file.Write([]byte(item)); err != nil {
			return fmt.Errorf("failed to write content: %v", err)
		}
	}

	return nil
}
