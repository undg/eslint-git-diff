package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
)

type Flg struct {
	fix      bool
	eslint   bool
	noEslint bool
	branch   string
	fetch    bool
}

func main() {
	var files string
	var command []string

	var flg = Flg{}

	{

		flag.BoolVar(&flg.fix, "fix", false, "add --fix flag to eslint")
		flag.BoolVar(&flg.fix, "f", false, "add --fix flag to eslint")

		flag.BoolVar(&flg.eslint, "eslint", true, "run eslint with listed files")
		flag.BoolVar(&flg.eslint, "e", true, "run eslint with listed files")

		flag.BoolVar(&flg.noEslint, "no-eslint", false, "dont run eslint, only list files")

		flag.StringVar(&flg.branch, "branch", "origin/dev", "branch to check files against")
		flag.StringVar(&flg.branch, "b", "origin/dev", "branch to check files against")

		flag.BoolVar(&flg.fetch, "fetch", false, "Before run git fetch")

		flag.Parse()
	}

	if flg.fetch {
		runGitFetch()
	}

	files = getFiles(flg)

	if flg.noEslint {
		flg.eslint = false
	}

	if flg.eslint {
		runEslint(command, files, flg)
	} else {
		fmt.Println("\n" + files)
	}

}

func getFiles(flg Flg) string {
	gitArguments := strings.Fields("diff --name-only --diff-filter=d " + flg.branch)

	out, err := exec.Command("git", gitArguments...).CombinedOutput()

	if err != nil {
		fmt.Printf("ERROR!\n%s\n", string(out))
		panic(err)
	}

	return string(out)
}

func runGitFetch() {
	out, err := exec.Command("git", "fetch").CombinedOutput()

	if err != nil {
		fmt.Printf("ERROR!\n%s\n", string(out))
		panic(err)
	}
}

func runEslint(command []string, files string, flg Flg) {
	command = append(command, "eslint_d")
	if flg.fix {
		command = append(command, "--fix")
	}

	// run finall command
	commands := command[0]
	args := append(command[1:], strings.Fields(files)...)
	cmd := exec.Command(commands, args...)

	f, err := pty.Start(cmd)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)
}
