package stdout

// Stdout is a struct that is used for displaying
// the application output in multiple interfaces.
type Stdout struct {
	Clipboard bool
	File      bool
	Path      string
}

func (s Stdout) Print(input ...string) error {
	return nil
}
