package validator

import (
	fmt "fmt"
	"strings"
)

// MultiError model
type MultiError struct {
	errs map[string]string
}

// NewMultiError constructor
func NewMultiError() *MultiError {
	return &MultiError{errs: make(map[string]string)}
}

// Append error to multierror
func (m *MultiError) Append(key string, err error) {
	if err != nil {
		if val, ok := m.errs[key]; ok {
			m.errs[key] = fmt.Sprintf("%s; %s", val, err.Error())
		} else {
			m.errs[key] = err.Error()
		}
	}
}

// HasError check if err is exist
func (m *MultiError) HasError() bool {
	return len(m.errs) != 0
}

// ToMap return list map of error
func (m *MultiError) ToMap() map[string]string {
	return m.errs
}

// Error implement error from multiError
func (m *MultiError) Error() string {
	var str []string
	for i, s := range m.errs {
		str = append(str, fmt.Sprintf("%s: %s", i, s))
	}
	return strings.Join(str, "\n")
}
