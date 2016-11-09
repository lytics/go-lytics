package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"strings"

	"github.com/fsnotify/fsnotify"
	lytics "github.com/lytics/go-lytics"
)

func (c *Cli) watch() (interface{}, error) {

	l := newLql(c)
	l.start()

	return nil, nil
}

type datafile struct {
	name string
	lql  string
	data []url.Values
}

func (d *datafile) loadJson(of string) {
	by, err := ioutil.ReadFile("./" + of)
	exitIfErr(err, fmt.Sprintf("Could not read json file %v", of))
	l := make([]map[string]interface{}, 0)
	err = json.Unmarshal(MakeJsonList(by), &l)
	exitIfErr(err, "Invalid json file")

	qsargs := make([]url.Values, 0, len(l))
	for _, row := range l {
		qs, err := lytics.FlattenJsonMap(row)
		if err == nil {
			qsargs = append(qsargs, qs)
		} else {
			log.Printf("Could not convert row to qs? %v   %v\n", row, err)
		}
	}
	d.data = qsargs
}

type lql struct {
	files map[string]*datafile
	w     *fsnotify.Watcher
	c     *Cli
}

func newLql(c *Cli) *lql {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err.Error())
	}
	return &lql{
		c:     c,
		w:     watcher,
		files: make(map[string]*datafile),
	}
}

func (l *lql) start() {
	defer l.w.Close()
	done := make(chan bool)
	l.loadFiles()
	l.watch()
	<-done
}
func (l *lql) print(d *datafile) {
	if len(d.data) == 0 {
		log.Printf("No data found for %v \n\n", d.name)
		return
	}
	if len(d.lql) == 0 {
		log.Printf("No lql found for %v \n\n", d.name)
		return
	}
	for _, qs := range d.data {
		ent, err := l.c.Client.GetQueryTest(qs, d.lql)
		if err != nil {
			log.Printf("Could not evaluate query/entity: %v \n\n", err)
			return
		}

		// Output the user response
		out, err := json.MarshalIndent(ent, "", "  ")
		if err == nil {
			fmt.Printf("\n%v\n\n", string(out))
		}
	}

}

func (l *lql) handleFile(of string, showOutput bool) {
	if strings.Index(of, ".") < 1 {
		return
	}
	f := strings.ToLower(of)
	name := strings.Split(f, ".")[0]
	df, exists := l.files[name]
	if !exists {
		df = &datafile{name: name}
		l.files[name] = df
	}
	switch {
	case strings.HasSuffix(f, ".lql"):
		log.Println("handle lql file ", f)
		by, err := ioutil.ReadFile("./" + of)
		exitIfErr(err, fmt.Sprintf("Could not read file %v", of))
		df.lql = string(by)
	case strings.HasSuffix(f, ".csv"):
		log.Println("handle csv file ", f)
	case strings.HasSuffix(f, ".json"):
		log.Println("handle json file ", f)
		df.loadJson(of)
	default:
		return
	}
	if showOutput {
		l.print(df)
	}
}

func (l *lql) loadFiles() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		l.handleFile(f.Name(), false)
	}
}

func (l *lql) watch() {

	go func() {
		for {
			select {
			case event := <-l.w.Events:
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create {
					fn := strings.ToLower(event.Name)
					fn = strings.Replace(fn, "./", "", 1)
					//log.Println("modified file:", fn)
					l.handleFile(fn, true)
				}
			case err, ok := <-l.w.Errors:
				if !ok {
					log.Fatal("What, no errors channel")
				} else {
					log.Println("watch error:", err)
				}

			}
		}
	}()

	if err := l.w.Add("./"); err != nil {
		log.Fatal(err)
	}
}

// Convert a slice of bytes into an array by ensuring it is wrapped with []
func MakeJsonList(b []byte) []byte {
	if !bytes.HasPrefix(b, []byte{'['}) {
		b = append([]byte{'['}, b...)
		b = append(b, ']')
	}
	return b
}
