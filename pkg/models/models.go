package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type texty struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
