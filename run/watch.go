package run

import (
	structs "eslint-git-diff/structs"
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)


func Watcher(command []string, files string, flg structs.Flg) {
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
					files = GetGitDiffFiles(flg)
				}

				if flg.Eslint {
					Eslint(command, files, flg)
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
}
