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

    var files []string
	var command []string

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

	// git fetch
	{
		if flgFetch {
			out, err := exec.Command("git", "statu").CombinedOutput()
			must(out, err)
			fmt.Printf("output of git fetch:\n%s\n", string(out))
		}
	}

    // eslint
	{
		if flgEslint {
			command = append(command, "eslint_d")
			if flgFix {
				command = append(command, "--fix")
			}
		} else {
			command = append(command, "echo")
		}
	}

    // get files
    {
        gitArguments := strings.Fields("diff --name-only --diff-filter=d " + flgBranch)

        out, err := exec.Command("git", gitArguments...).CombinedOutput()
        must(out, err)
        fmt.Println(string(out))
        files = strings.Fields(string(out))
    }

    // run finall command
    {
        cmd := command[0]
        args := append(command[1:], files...)
        out, err := exec.Command(cmd, args...).CombinedOutput()

        must(out, err)
        
        fmt.Println(string(out))
    }
}
