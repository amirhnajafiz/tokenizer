package stdout

import "fmt"

func (s Stdout) printToConcole() error {
	for _, item := range s.inputs {
		fmt.Println(item)
	}

	return nil
}
