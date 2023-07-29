package vars

import (
	"encoding/json"
	"fmt"
)

const (
	Version           = "3.0.0"
	VersionPrerelease = "dev"
)

type VersionInfo struct {
	Version   string `json:"version"`
	CommitId  string `json:"commit_id"`
	BuildTime string `json:"build_time"`
}

func GetVersion() string {
	version := VersionInfo{}
	if VersionPrerelease != "" {
		version.Version = fmt.Sprintf("%s%s-%s", "v", Version, VersionPrerelease)

	} else {
		version.Version = fmt.Sprintf("%s%s", "v", Version)
	}

	version.CommitId = GITCOMMIT
	version.BuildTime = BUILDTIME
	b, err := json.MarshalIndent(version, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	return string(b)
}
