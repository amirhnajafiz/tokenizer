package stdout

import (
	"fmt"

	"golang.design/x/clipboard"
)

func (s Stdout) copyToClipboard() error {
	// init returns an error if the package is not ready for use.
	err := clipboard.Init()
	if err != nil {
		return fmt.Errorf("clipboard package is not supported on your system: %v", err)
	}

	for _, item := range s.inputs {
		clipboard.Write(clipboard.FmtText, []byte(item))
	}

	return nil
}
