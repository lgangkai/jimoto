package util

import (
	"fmt"
	"jimotoapi/conf"
)

func CompleteImageUrl(filename string, config *conf.Config) string {
	if config.ImageServer.Local {
		if filename == "" {
			return ""
		}
		return fmt.Sprintf("%v://%v/%v/%v", config.Server.Scheme, config.Server.Addr, config.ImageServer.LocalPath, filename)
	}
	return ""
}

func CompleteImageUrls(filenames []string, config *conf.Config) []string {
	rst := make([]string, len(filenames))
	for i, filename := range filenames {
		rst[i] = CompleteImageUrl(filename, config)
	}
	return rst
}
