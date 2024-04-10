package version

import (
	"fmt"

	"github.com/Masterminds/semver"
)

// parseVersion takes in a version string and then parses it to a Version object
func parseVersion(ver string) (Version *semver.Version, err error) {
	v, err := semver.NewVersion(ver)
	if err != nil {
		return nil, &VersionError{
			message: "Parsing Version",
			err:     err,
		}
	}
	return v, nil
}

// UpgradeAvailable checks the version of the binary against a semantic version string
// returns true if binary version is older than the semantic version
func UpgradeAvailable(remoteVersion string) (UpgradeAvailable bool, err error) {
	constraint, err := semver.NewConstraint(fmt.Sprintf("<%s", remoteVersion))
	if err != nil {
		return false, err
	}
	semVerCurrent, err := parseVersion(BuildVersion)
	if err != nil {
		return false, err
	}
	UpgradeAvailable, _ = constraint.Validate(semVerCurrent)
	return UpgradeAvailable, nil
}
