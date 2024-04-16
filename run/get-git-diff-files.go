package run

import (
	"eslint-git-diff/structs"
	"fmt"
	"os/exec"
	"strings"
)

// Returns list of git diff files.
// One line, space separated.
func GetGitDiffFiles(flg structs.Flg) string {
	gitArguments := strings.Fields("diff --name-only --diff-filter=d " + flg.Branch)

	out, err := exec.Command("git", gitArguments...).CombinedOutput()

	if err != nil {
		fmt.Printf("GIT DIFF ERROR!\n%s\n", string(out))
		panic(err)
	}

	filesOneByLine := strings.ReplaceAll(string(out), "\r\n", "\n")
	filesFiltered := filterFiles(strings.Split(filesOneByLine, "\n"), flg.Watch)

	filesInOneLine := strings.Join(filesFiltered, " ")

	fmt.Println(filesInOneLine)

	return filesInOneLine
}
