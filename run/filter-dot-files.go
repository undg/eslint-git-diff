package run

import (
	"fmt"
	"strings"
)

func filterFiles(files []string, path string) []string {
	filtered := []string{}

	fmt.Println("filtering files...")

	for _, file := range files {
		// @TODO (undg) 2024-01-20: read it from eslintignore
		hardcoded := //
			strings.Contains(file, "pnpm-lock.yaml") ||
				strings.Contains(file, "yarn.lock") ||
				strings.Contains(file, "package.json") ||
				strings.Contains(file, "package-lock.json")

		if hardcoded ||
			strings.Contains(file, "/.") ||
			strings.HasPrefix(file, ".") {
			continue
		}
		filtered = append(filtered, file)
	}

	return filtered
}
