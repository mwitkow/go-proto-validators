package plugin

import (
	"fmt"
	"strconv"
)

var uuidPattern = "^[a-fA-F0-9]{8}-" +
	"[a-fA-F0-9]{4}-" +
	"[%s]" +
	"[a-fA-F0-9]{3}-" +
	"[8|9|aA|bB][a-fA-F0-9]{3}-" +
	"[a-fA-F0-9]{12}$"

func getUUIDRegex(version int32) (string, error) {
	if version < 0 || version > 5 {
		return "", fmt.Errorf("UUID version should be between 0-5, Got %d", version)
	}
	switch version {
	case 0:
		return fmt.Sprintf(uuidPattern, "1-5"), nil
	default:
		return fmt.Sprintf(uuidPattern, strconv.FormatInt(int64(version), 10)), nil
	}
}
