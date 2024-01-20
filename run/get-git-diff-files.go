package run

import (
	"eslint-git-diff/structs"
	"fmt"
	"os/exec"
	"strings"
)

func GetGitDiffFiles(flg structs.Flg) string {
	gitArguments := strings.Fields("diff --name-only --diff-filter=d "+flg.Branch)

	out, err := exec.Command("git", gitArguments...).CombinedOutput()

	if err != nil {
		fmt.Printf("GIT DIFF ERROR!\n%s\n", string(out))
		panic(err)
	}

	files := strings.Join(filterFiles(strings.Split(string(out), "\n"), flg.Watch), " ")

	fmt.Println(files)

	return files
}
