package version

import "fmt"

type VersionError struct {
	message string
	err     error
}

func (e *VersionError) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.err)
}
