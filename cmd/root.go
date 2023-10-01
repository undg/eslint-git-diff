package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/spf13/cobra"

	"eslint-git-diff/run"
	structs "eslint-git-diff/structs"
)

var flg = structs.Flg{}

var rootCmd = &cobra.Command{
	Use:   "eslint-git-diff",
	Short: "Run eslint_d on git diff filles",
	Long: `Wrapper around wrapper for eslint with git integration.
                Add few missing options for eslint. Native to your sytem watch mode and awesome eslint_d.
                Lint only files that are pressent in git diff.`,
	Run: func(cmd *cobra.Command, args []string) {
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
			w := watcher.New()
			w.SetMaxEvents(1)
			w.FilterOps(
				watcher.Write,
				watcher.Move,
				watcher.Rename,
				watcher.Create,
				watcher.Remove,
			)

			go func() {
				for {
					select {
					case event := <-w.Event:
						if event.Op == watcher.Move ||
							event.Op == watcher.Rename ||
							event.Op == watcher.Create ||
							event.Op == watcher.Remove {
							files = run.GetGitDiffFiles(flg)
						}

						if flg.Eslint {
							run.Eslint(command, files, flg)
						} else {
							fmt.Println("\n" + files)
							fmt.Println("\nevent:", event)
						}
					case err := <-w.Error:
						fmt.Println(err)
					case <-w.Closed:
						return
					}
				}
			}()

			if err := w.AddRecursive(flg.Watch); err != nil {
				log.Fatalln(err)
			}

			// Trigger 2 events after watcher started
			go func() {
				w.Wait()
				w.TriggerEvent(watcher.Create, nil)
				w.TriggerEvent(watcher.Remove, nil)
			}()

			// Start the watching process - it'll check for changes every 100 ms
			if err := w.Start(time.Millisecond * 100); err != nil {
				log.Fatalln(err)
			}
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
	rootCmd.PersistentFlags().BoolVarP(&flg.Eslint, "eslint", "e", true, "run eslint with listed files")
	rootCmd.PersistentFlags().BoolVarP(&flg.Fix, "fix", "f", false, "add --fix flag to eslint")
	rootCmd.PersistentFlags().BoolVar(&flg.NoEslint, "no-eslint", false, "don't run eslint, list diff files only")
	rootCmd.PersistentFlags().BoolVar(&flg.Fetch, "fetch", false, "Git fetch, on startup")
	rootCmd.PersistentFlags().StringVarP(&flg.Branch, "branch", "b", "origin/dev", "branch to check files against")
	rootCmd.PersistentFlags().StringVarP(&flg.Watch, "watch", "w", "", "Watch the paths for changes")

}
