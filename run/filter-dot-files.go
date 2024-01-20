package run

import (
	"fmt"
	"strings"
)

func filterFiles(files []string, path string) []string {
	filtered := []string{}

	fmt.Println("filtering files...")

	for _, file := range files {
		if strings.Contains(file, "/.") ||
			strings.HasPrefix(file, ".") {
			continue
		}
		filtered = append(filtered, file)
	}

	return filtered
}
