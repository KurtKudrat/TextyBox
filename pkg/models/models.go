package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Texty struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
