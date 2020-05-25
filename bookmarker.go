package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
)

// Bookmarker parses json files with bookmarks
type Bookmarker struct {
	Bookmarks []Bookmark
}

func (b Bookmarker) bookmarkPath() string {
	var p string

	conf, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	switch o := runtime.GOOS; o {
	case "darwin":
		p = path.Join(conf, "Google/Chrome/Default/Bookmarks")
	case "linux":
		p = path.Join(conf, "google-chrome/Default/Bookmarks")
	default:
		log.Fatalf("%s is not suppoorted", o)
	}

	return p
}

func (b Bookmarker) interface2json(x interface{}) *simplejson.Json {
	j := simplejson.New()
	j.SetPath([]string(nil), x)
	return j
}

// NewJSON returns a new simplejson.Json object
func (b Bookmarker) NewJSON() *simplejson.Json {
	bytes, err := ioutil.ReadFile(b.bookmarkPath())
	if err != nil {
		log.Fatal(err)
	}

	json, err := simplejson.NewJson(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return json
}

// ParseUnixTime parses time format for Chrome bookmarks
// reference: https://stackoverflow.com/a/57903746/12635122
func (b Bookmarker) ParseUnixTime(ut string) time.Time {
	unixTime, _ := strconv.ParseInt(ut+"0", 10, 64) // 17 digits to 18 digits

	maxd := time.Duration(math.MaxInt64).Truncate(100 * time.Nanosecond)
	maxdUnits := int64(maxd / 100) // number of 100-ns units

	t := time.Date(1601, 1, 1, 0, 0, 0, 0, time.UTC)
	for unixTime > maxdUnits {
		t = t.Add(maxd)
		unixTime -= maxdUnits
	}
	if unixTime != 0 {
		t = t.Add(time.Duration(unixTime * 100))
	}

	return t
}

// Search recursively searches for bookmarks
func (b *Bookmarker) Search(j *simplejson.Json, dirPath string) {
	switch t := j.Get("type").MustString(); t {
	case "folder":
		dirPath += fmt.Sprintf("%s/", j.Get("name").MustString())
		for _, c := range j.Get("children").MustArray() {
			b.Search(b.interface2json(c), dirPath)
		}
	case "url":
		dateAdded := b.ParseUnixTime(j.Get("date_added").MustString())
		name := j.Get("name").MustString()
		path := dirPath + name
		url := j.Get("url").MustString()

		b.Bookmarks = append(b.Bookmarks, Bookmark{DateAdded: dateAdded, Name: name, Path: path, URL: url})
	}
}
