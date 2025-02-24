package data

import "time"

// !All fields in the Movie struct are exported (start with a capital letter)
// which is necessary for them to be visible to Go's encoding/json package.
// Any fields that aren't exported won't be included when encoding a struct
// to JSON
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}
