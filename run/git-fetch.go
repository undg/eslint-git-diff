package run

import (
	"fmt"
	"os/exec"
)

func GitFetch() {
	out, err := exec.Command("git", "fetch").CombinedOutput()

	if err != nil {
		fmt.Printf("GIT FETCH ERROR!\n%s\n", string(out))
		panic(err)
	}
}

