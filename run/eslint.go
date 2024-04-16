package run

import (
	structs "eslint-git-diff/structs"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
)

func Eslint(command []string, files string, flg structs.Flg) {
	command = append(command, "eslint_d")
	if flg.Fix {
		command = append(command, "--fix")
	}

	// run finall command
	commands := command[0]
	args := append(command[1:], strings.Fields(files)...)
	cmd := exec.Command(commands, args...)

	cmdReader, err := pty.Start(cmd)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, cmdReader)
}
