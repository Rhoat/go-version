package version

import (
	"fmt"
	"runtime"
)

// Info contains all the versioning information for a package/tool
type Info struct {
	ApplicationName string `json:"application_name"`
	CommitHash      string `json:"commit_hash"`
	BuildDate       string `json:"build_date"`
	BuildOS         string `json:"build_os"`
	BuildARCH       string `json:"build_arch"`
	BuildTarget     string `json:"build_target"`
	BuildTag        string `json:"build_number"`
	BuildVersion    string `json:"build_version"`
}

var (
	Version = Info{
		ApplicationName: ApplicationName, // defined via LDFLags at compile-time
		CommitHash:      CommitHash,      // defined via LDFLags at compile-time
		BuildDate:       BuildDate,       // defined via LDFLags at compile-time
		BuildTag:        BuildTag,        // defined via LDFLags at compile-time
		BuildVersion:    BuildVersion,    // defined via LDFLags at compile-time
		BuildOS:         runtime.GOOS,
		BuildARCH:       runtime.GOARCH,
	}
)

func (vi Info) String() string {
	return fmt.Sprintf("%s v%s %s (%s) %s - BuildDate: %s", vi.ApplicationName, vi.BuildVersion, vi.BuildTag, vi.CommitHash, GetTarget(), vi.BuildDate)
}

// Stringer interface for the version
func String() string {
	return Version.String()
}

func GetTarget() string {
	return fmt.Sprintf("%s/%s", Version.BuildOS, Version.BuildARCH)
}
