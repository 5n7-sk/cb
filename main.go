package main

import (
	"log"
	"strings"

	"github.com/kyokomi/emoji"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/manifoldco/promptui"
	"github.com/pkg/browser"
)

var bookmarkTypes = []string{"bookmark_bar", "other", "synced"}

func main() {
	b := Bookmarker{}
	json := b.NewJSON()
	roots := json.Get("roots")

	for _, bmType := range bookmarkTypes {
		bookmarks := roots.Get(bmType)
		b.Search(bookmarks, "")
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   emoji.Sprint(":bookmark: {{ .Name | cyan }}"),
		Inactive: "    {{ .Name | cyan }}",
		Selected: emoji.Sprint(":bookmark: {{ .Name | cyan }}"),
		Details: `------------------------------
{{ "\U0001F4C5 added at" }}	{{ .DateAdded }}
{{ "\U0001F4DD name" }}	{{ .Name }}
{{ "\U0001F4C1 path" }}	{{ .Path }}
{{ "\U0001F3E0 url" }}	{{ .URL }}`,
	}

	keys := &promptui.SelectKeys{
		Next:     promptui.Key{Code: promptui.KeyNext, Display: promptui.KeyNextDisplay},
		Prev:     promptui.Key{Code: promptui.KeyPrev, Display: promptui.KeyPrevDisplay},
		PageUp:   promptui.Key{Code: promptui.KeyBackward, Display: promptui.KeyBackwardDisplay},
		PageDown: promptui.Key{Code: promptui.KeyForward, Display: promptui.KeyForwardDisplay},
		Search:   promptui.Key{Code: 63, Display: "?"}, // 63 is rune for "?"
	}

	searcher := func(input string, index int) bool {
		bm := b.Bookmarks[index]
		path := strings.ToLower(bm.Path)
		url := strings.ToLower(bm.URL)
		input = strings.ToLower(input)

		return fuzzy.Match(input, path) || fuzzy.Match(input, url)
	}

	prompt := promptui.Select{
		Label:             "Bookmarks",
		Items:             b.Bookmarks,
		Size:              10,
		HideHelp:          true,
		Templates:         templates,
		Keys:              keys,
		Searcher:          searcher,
		StartInSearchMode: true,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	browser.OpenURL(b.Bookmarks[i].URL)
}
