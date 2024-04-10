package version

import (
	"testing"

	semver "github.com/Masterminds/semver"
)

func TestParseVersionError(t *testing.T) {
	testCases := []struct {
		TestName      string
		VersionNumber string
		Expected      error
	}{
		{
			TestName:      "Invalid Version",
			VersionNumber: "N/A",
			Expected: &VersionError{
				message: "Parsing Version",
				err:     semver.ErrInvalidSemVer,
			},
		}}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := parseVersion(tc.VersionNumber)
			if err.Error() != tc.Expected.Error() {
				t.Errorf("Expected: %s Got: %s", tc.Expected, err)
			}

		})

	}

}

func TestParseVersion(t *testing.T) {
	testCases := []struct {
		TestName      string
		VersionNumber string
		Major         int64
		Minor         int64
		Patch         int64
	}{
		{
			"MajorOnly",
			"1",
			1,
			0,
			0,
		},
		{
			"Major.MinorOnly",
			"1.1",
			1,
			1,
			0,
		},
		{
			"Major.Minor.Patch",
			"1.2.3",
			1,
			2,
			3,
		},
		{
			"Major.Minor.Patch",
			"1.2.0",
			1,
			2,
			0,
		},
	}
	for _, testCase := range testCases {
		version, _ := parseVersion(testCase.VersionNumber)
		major, minor, patch := version.Major(), version.Minor(), version.Patch()
		if major != testCase.Major {
			t.Errorf("TestName: %s - Found Major versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Major, major)
		}
		if minor != testCase.Minor {
			t.Errorf("TestName: %s - Found Minor versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Minor, minor)
		}
		if patch != testCase.Patch {
			t.Errorf("TestName: %s - Found Patch versions are not equal (expected %d but got %d)", testCase.TestName, testCase.Patch, patch)
		}
	}
}

func TestUpgradeAvailable(t *testing.T) {
	testCases := []struct {
		TestName      string
		VersionNumber string
		remoteVersion string
		Expected      bool
	}{
		{
			"Version is the same.",
			"1.0.0",
			"1.0.0",
			false,
		},
		{
			"Version is less than Remote Version.",
			"0.0.0",
			"1.0.0",
			true,
		},
		{
			"Version is greater than Remote Version.",
			"1.0.0",
			"0.0.0",
			false,
		},
	}
	for _, testCase := range testCases {
		BuildVersion = testCase.VersionNumber
		available, _ := UpgradeAvailable(testCase.remoteVersion)
		if available != testCase.Expected {
			t.Errorf("TestName: %s - Expected: %t Got: %t", testCase.TestName, testCase.Expected, available)
		}
	}
}
