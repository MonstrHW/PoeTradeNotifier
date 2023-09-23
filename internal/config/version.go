package config

import "strings"

var (
	tag        string
	date       string
	hashCommit string
	note       string
)

func GetAppVersion() string {
	values := []string{
		tag,
		date,
		hashCommit,
		note,
	}

	var version []string
	for _, v := range values {
		if v != "" {
			version = append(version, v)
		}
	}

	if len(version) == 0 {
		return "version not present"
	}

	return strings.Join(version, "\n")
}
