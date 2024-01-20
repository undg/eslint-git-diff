package run

import (
	"strings"
)

func filterFiles(files []string, path string) []string {
	filtered := []string{}

	for _, file := range files {
		// @TODO (undg) 2024-01-20: read it from eslintignore
		isIgnored := //
			strings.Contains(file, "pnpm-lock.yaml") ||
				strings.Contains(file, "yarn.lock") ||
				strings.Contains(file, "package.json") ||
				strings.Contains(file, "package-lock.json")

		isNotInWatchPath := path != "." && !strings.HasPrefix(file, path)

		isDotFile := strings.Contains(file, "/.") ||
			strings.HasPrefix(file, ".")

		if isIgnored ||
			isNotInWatchPath ||
			isDotFile {
			continue
		}
		filtered = append(filtered, file)
	}

	return filtered
}
