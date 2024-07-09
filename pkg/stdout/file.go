package stdout

import (
	"fmt"
	"os"
)

func (s Stdout) exportToFile() error {
	file, err := os.Create(s.Path)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}

	for _, item := range s.inputs {
		if _, err := file.Write([]byte(item)); err != nil {
			return fmt.Errorf("failed to write content: %v", err)
		}
	}

	return nil
}
