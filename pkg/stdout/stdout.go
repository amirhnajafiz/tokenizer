package exporter

// Stdout is a struct that is used for displaying
// the application output in multiple interfaces.
type Stdout struct{}

func (s Stdout) Print(input ...string) error {
	return nil
}
