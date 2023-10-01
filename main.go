package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"

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

		flag.StringVar(&flg.Watch, "w", "", "Watch the paths for changes")
		flag.StringVar(&flg.Watch, "watch", "", "Watch the paths for changes")

		flag.Parse()
	}

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

}
