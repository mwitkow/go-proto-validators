package plugin

import (
	"fmt"
	"strconv"
)

const uuidPattern = "^([a-fA-F0-9]{8}-" +
	"[a-fA-F0-9]{4}-" +
	"[%s]" +
	"[a-fA-F0-9]{3}-" +
	"[8|9|aA|bB][a-fA-F0-9]{3}-" +
	"[a-fA-F0-9]{12})?$"

// getUUIDRegex returns a regex to validate that a string is in UUID
// format. The version parameter specified the UUID version. If version is 0,
// the returned regex is valid for any UUID version
func getUUIDRegex(version int) (string, error) {
	if version < 0 || version > 5 {
		return "", fmt.Errorf("UUID version should be between 0-5, Got %d", version)
	}

	if version == 0 {
		return fmt.Sprintf(uuidPattern, "1-5"), nil
	}

	return fmt.Sprintf(uuidPattern, strconv.Itoa(version)), nil
}
