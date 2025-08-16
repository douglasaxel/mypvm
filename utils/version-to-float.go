package utils

import (
	"strconv"
	"strings"
)

func VersionToFloat(version string) float64 {
	version = strings.Replace(version, ".", "", 1)
	versionFloat, err := strconv.ParseFloat(version, 64)

	if err != nil {
		return 0.0
	}

	return versionFloat
}
