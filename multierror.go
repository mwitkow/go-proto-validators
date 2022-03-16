package validator

import (
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		m.errs[key] = err.Error()
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

func (m *MultiError) RPCError() error {
	grpcStatus := status.New(codes.InvalidArgument, "BAD_REQ")

	for field, message := range m.errs {
		errMsg := strings.Split(message, " ;; ")

		grpcStatus, _ = grpcStatus.WithDetails(&errdetails.ErrorInfo{
			Reason: errMsg[0],
			Domain: field,
			Metadata: map[string]string{
				"value": errMsg[1],
			},
		})
	}

	return grpcStatus.Err()
}
