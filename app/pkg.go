package app

import (
	"fmt"
	"os"
	"path/filepath"
)

// build info
var (
	appName          string
	buildVersionName string
	buildVersionCode int
)

func init() {
	appName = filepath.Base(os.Args[0])
	buildVersionName = "1.0.0"
	buildVersionCode = 0
}

func Name() string {
	return appName
}

func VersionName() string {
	return buildVersionName
}

func VersionCode() int {
	return buildVersionCode
}

func Version() string {
	return VersionName() + fmt.Sprintf("%d", VersionCode())
}
