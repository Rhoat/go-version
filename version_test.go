package version_test

import (
	"fmt"
	"testing"

	"github.com/pragmaticengineering/go-version"
)

// Table driven test for Printing Version
func TestPrintVersion(t *testing.T) {
	testCases := []struct {
		TestName        string
		Expected        string
		ApplicationName string
		BuildVersion    string
		BuildTag        string
		CommitHash      string
		BuildOS         string
		BuildARCH       string
		BuildDate       string
	}{
		{
			TestName:        "PrintVersion",
			Expected:        "test v1.0.0 Tag (XYZ) windows/amd64 - BuildDate: Oct-21-2024",
			ApplicationName: "test",
			BuildVersion:    "1.0.0",
			BuildTag:        "Tag",
			CommitHash:      "XYZ",
			BuildOS:         "windows",
			BuildARCH:       "amd64",
			BuildDate:       "Oct-21-2024",
		},
	}

	for _, testCase := range testCases {
		version.Version.ApplicationName = testCase.ApplicationName
		version.Version.BuildVersion = testCase.BuildVersion
		version.Version.BuildTag = testCase.BuildTag
		version.Version.CommitHash = testCase.CommitHash
		version.Version.BuildOS = testCase.BuildOS
		version.Version.BuildARCH = testCase.BuildARCH
		version.Version.BuildDate = testCase.BuildDate

		output := fmt.Sprintf("%s", version.Version)
		if output != testCase.Expected {
			t.Errorf("TestName: %s - Found output is not equal (expected %s but got %s)", testCase.TestName, testCase.Expected, output)
		}

		if testCase.Expected != version.String() {
			t.Errorf("TestName: %s - Found output is not equal (expected %s but got %s)", testCase.TestName, testCase.Expected, version.String())
		}
	}
}

// Version: 0.0.0 Build Date: unknown Commit Hash: unknown Build Tag: dev
// unknown v0.0.0 dev (unknown) windows/amd64 - BuildDate: unknown
