package main

import (
	"flag"
	"fmt"

	run "git-diff-files/run"
	structs "git-diff-files/structs"
)

func main() {
	var files string
	var command []string

	var flg = structs.Flg{}

	{

		flag.BoolVar(&flg.Fix, "fix", false, "add --fix flag to eslint")
		flag.BoolVar(&flg.Fix, "f", false, "add --fix flag to eslint")

		flag.BoolVar(&flg.Eslint, "eslint", true, "run eslint with listed files")
		flag.BoolVar(&flg.Eslint, "e", true, "run eslint with listed files")

		flag.BoolVar(&flg.NoEslint, "no-eslint", false, "dont run eslint, only list files")

		flag.StringVar(&flg.Branch, "branch", "origin/dev", "branch to check files against")
		flag.StringVar(&flg.Branch, "b", "origin/dev", "branch to check files against")

		flag.BoolVar(&flg.Fetch, "fetch", false, "Before run git fetch")

		flag.Parse()
	}

	if flg.Fetch {
		run.GitFetch()
	}

	files = run.GetFiles(flg)

	if flg.NoEslint {
		flg.Eslint = false
	}

	if flg.Eslint {
		run.Eslint(command, files, flg)
	} else {
		fmt.Println("\n" + files)
	}

}
