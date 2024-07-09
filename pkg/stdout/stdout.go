package stdout

// Stdout is a struct that is used for displaying
// the application output in multiple interfaces.
type Stdout struct {
	Clipboard bool
	File      bool
	Path      string
	inputs    []string
}

// Print method is the main handler of stdout module.
func (s Stdout) Print(input ...string) error {
	s.inputs = input

	switch {
	case s.File:
		return s.exportToFile()
	default:
		return s.printToConcole()
	}
}
