package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func must(out []byte, err error) {
        if err != nil {
            fmt.Printf("err: %s\n%s", err, out)
            panic(err)
        }
}

func main() {
	var flgFix bool
	var flgEslint bool
	var flgBranch string
	var flgFetch bool

	shFix := ""
	shEslint := ""
	shBranch := ""

	{
		flag.BoolVar(&flgFix, "fix", false, "add --fix flag to eslint")
		flag.BoolVar(&flgFix, "f", false, "add --fix flag to eslint")
		flag.BoolVar(&flgEslint, "eslint", true, "run eslint with listed files")
		flag.BoolVar(&flgEslint, "e", true, "run eslint with listed files")
		flag.StringVar(&flgBranch, "branch", "origin/dev", "branch to check files against")
		flag.StringVar(&flgBranch, "b", "origin/dev", "branch to check files against")
		flag.BoolVar(&flgFetch, "fetch", false, "Before run git fetch")

		flag.Parse()
	}

	if flgEslint && flgFix {
		shFix = "--fix"
	}

	if flgEslint {
		shEslint = "eslint_d "
	}

	if flgFetch {
        cmd := exec.Command("git", "diff")
        out, err := cmd.CombinedOutput()

        must(out, err)

        fmt.Printf("output of git fetch:\n%s\n", string(out))
	}

	shBranch = flgBranch

	fmt.Println(shEslint, "git diff --name-only --diff-filter=dm", shFix, shBranch)

}
