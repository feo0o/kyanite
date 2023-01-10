package app

var (
	majorVer  = "0"
	minorVer  = "0"
	patchVer  = "0"
	gitCommit string
)

func VersionCore() string {
	return majorVer + "." + minorVer + "." + patchVer
}

func VersionWithBuildIdentifier() string {
	return VersionCore() + "+git." + gitCommit
}

// Version full version info with app name, and git commit digits if exist
func Version() string {
	if gitCommit == "" {
		return Name + " v" + VersionCore()
	}
	return Name + " v" + VersionWithBuildIdentifier()
}
