package main

import (
	"time"
)

// Bookmark represents a bookmark
type Bookmark struct {
	DateAdded time.Time
	Name      string
	Path      string
	URL       string
}
