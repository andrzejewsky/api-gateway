package configuration

// ArrayFlags type for command line parameters
type ArrayFlags []string

// String required from the interface
func (i *ArrayFlags) String() string {
	return "my string representation"
}

// Set modify an array
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
