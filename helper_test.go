package validator

import (
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestStatusFromError(t *testing.T) {
	err := FieldError("myField", fmt.Errorf("Error %s", "here"))
	if code := status.Code(err); code != codes.InvalidArgument {
		t.Fatalf("Wanted %s got %s", codes.InvalidArgument, code)
	}
}
