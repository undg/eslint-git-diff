package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

func must(out []byte, err error) {
	if err != nil {
		fmt.Printf("ERROR!\n%s\n", string(out))
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

	// var myArguments []string

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

	// if flgEslint {
	// 	myArguments = append(myArguments, "eslint_d")
	// 	if flgFix {
	// 		myArguments = append(myArguments, "--fix")
	// 	}
	// } else {
	// 	myArguments = append(myArguments, "echo")
	// }

    // myArguments = append(myArguments, "$(git status)")

	if flgFetch {
		out, err := exec.Command("git", "statu").CombinedOutput()
		must(out, err)
		fmt.Printf("output of git fetch:\n%s\n", string(out))
	}

	shBranch = flgBranch

	fmt.Println(shEslint, shFix, "git diff --name-only --diff-filter=dm", shBranch)

    // "git diff --name-only --diff-filter=dm", shFix, shBranch

    gitArguments := strings.Fields("diff --name-only --diff-filter=dm HEAD")
    {
        out, _  := exec.Command("git", gitArguments...).CombinedOutput()
        // must(out, err)
        fmt.Println(string(out))
        fmt.Println(gitArguments)
    }

}
