package main

import (
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
)

/*
Watch a folder for changes and output evaluated entity

- watch   *.lql files for lql
- watch *.json for data
- watch *.csv for data
  - turn into json at change time

- output results to stdout

*/

func (c *Cli) watch() (interface{}, error) {

	l := newLql()

	done := make(chan bool)
	<-done

	return nil, err
}

type datafile struct {
	name     string
	datatype string
}
type lql struct {
	files map[string]string
	data  map[string]*datafile
}

func newLql() *lql {
	l := &lql{}
	l.start()
	return l
}

func (l *lql) watch() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					fn := strings.ToLower(event.Name)
					log.Println("modified file:", fn)
					switch {
					case strings.HasSuffix(fn, ".lql"):
						// edited lql
					}
				}
			case err := <-watcher.Errors:
				log.Println("watch error:", err)
			}
		}
	}()

	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}
}

func (l *lql) print() {

}
