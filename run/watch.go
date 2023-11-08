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
		fmt.Print("\033[H\033[2J") // clear screen
		fmt.Printf("[%s] Starting watcher...\n", time.Now().Format("15:00:00"))

		go Eslint(command, files, flg)

		for {
			select {
			case event := <-w.Event:
				run(command, files, flg, event)
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

	// Start the watching process - it'll check for changes every 100 ms
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func run(command []string, files string, flg structs.Flg, event watcher.Event) {
	files = GetGitDiffFiles(flg)

	fmt.Print("\033[H\033[2J") // clear screen

	if flg.Eslint {
		switch flg.Verbose {
		case 0:
			fmt.Printf("\n[%s] %s %s", event.ModTime().Format("15:00:00"), event.Op.String(), event.Name())
		case 1:
			fmt.Printf("\n[%s] %s %s\n%s", event.ModTime(), event.Op.String(), event.Name(), files)
		case 2:
			fmt.Println("\nevent:", event, "\ntime:", event.ModTime(), "\ncommand:", command, "\nflg:", flg, "\nfiles:\n", files)
		default:
			fmt.Println("\nevent:", event, "\ntime:", event.ModTime(), "\ncommand:", command, "\nflg:", flg, "\nfiles:\n", files, "Random error generator: wnMethod: No such interface “org.freedesktop.portal.Inhibit” on object at path /org/freedesktop/portal/desktop", "Cow Say: muuuuuuuuu", "Cat Say: meaw")
		}

		go Eslint(command, files, flg)
	} else {
		fmt.Println("\n" + files)
		fmt.Println("\nevent:", event)
	}

}
