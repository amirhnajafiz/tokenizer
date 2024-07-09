package stdout

// Stdout is a struct that is used for displaying
// the application output in multiple interfaces.
type Stdout struct {
	// case flags
	ClipboardFlag bool
	FileFlag      bool
	// output path
	Path string
	// in module params
	inputs []string
}

// Print method is the main handler of stdout module.
func (s Stdout) Print(input ...string) error {
	s.inputs = input

	// case on existing flags
	switch {
	case s.FileFlag:
		return s.exportToFile()
	case s.ClipboardFlag:
		return s.copyToClipboard()
	default:
		return s.printToConcole()
	}
}
