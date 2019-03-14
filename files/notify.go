package files

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
)

// WatchFile is to watch file modified or not
func WatchFile(fileList []string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
			//some condition should be used to stop by done channel
			//done<-true
		}
	}()

	for _, file := range fileList {
		err = watcher.Add(file)
		if err != nil {
			return err
		}
	}
	<-done

	return nil
}
