package run

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"

	structs "git-diff-files/structs"
)

type Flg = structs.Flg

func GetFiles(flg Flg) string {
	gitArguments := strings.Fields("diff --name-only --diff-filter=d " + flg.Branch)

	out, err := exec.Command("git", gitArguments...).CombinedOutput()

	if err != nil {
		fmt.Printf("ERROR!\n%s\n", string(out))
		panic(err)
	}

	return string(out)
}

func GitFetch() {
	out, err := exec.Command("git", "fetch").CombinedOutput()

	if err != nil {
		fmt.Printf("ERROR!\n%s\n", string(out))
		panic(err)
	}
}

func Eslint(command []string, files string, flg Flg) {
	command = append(command, "eslint_d")
	if flg.Fix {
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
