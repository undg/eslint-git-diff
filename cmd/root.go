package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"eslint-git-diff/run"
	structs "eslint-git-diff/structs"
)

var flg = structs.Flg{}

var stringVersion string

var rootCmd = &cobra.Command{
	Use:   "eslint-git-diff",
	Short: "Run eslint_d on git diff filles",
	Long: `Wrapper around wrapper for eslint with git integration.
                Add few missing options for eslint. Native to your sytem watch mode and awesome eslint_d.
                Lint only files that are pressent in git diff.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flg.Version {
			fmt.Println(stringVersion)
			return
		}

		var files string
		var command []string

		if flg.Fetch {
			run.GitFetch()
		}

		files = run.GetGitDiffFiles(flg)

		if flg.NoEslint {
			flg.Eslint = false
		}

		if flg.Watch != "" {
			run.Watcher(command, files, flg)
		} else {
			if flg.Eslint {
				run.Eslint(command, files, flg)
			} else {
				fmt.Println("\n" + files)
			}
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize()
	rootCmd.PersistentFlags().BoolVarP(&flg.Eslint, "eslint", "e", true, "run eslint")
	rootCmd.PersistentFlags().BoolVarP(&flg.NoEslint, "no-eslint", "x", false, "don't run eslint, list diff files only")
	rootCmd.PersistentFlags().BoolVarP(&flg.Fix, "fix", "f", false, "add --fix flag to eslint")
	rootCmd.PersistentFlags().BoolVarP(&flg.Fetch, "fetch", "u", false, "Git fetch, on startup")
	rootCmd.PersistentFlags().StringVarP(&flg.Branch, "branch", "b", "origin/dev", "branch to check files against")
	rootCmd.PersistentFlags().StringVarP(&flg.Watch, "watch", "w", "", "watch the path for changes. [eslint-git-diff -w src]")
	rootCmd.PersistentFlags().CountVarP(&flg.Verbose, "verbose", "v", "Be verbose [-v, or -vv]")

	// variable for this flag is crated on compolation time. Check goreleaser or shell build script.
	rootCmd.PersistentFlags().BoolVar(&flg.Version, "version", false, "show app version")
}
