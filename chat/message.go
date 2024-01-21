package main

import (
	"time"
)

// messageはひとつのメッセージを表す
type message struct {
	Name string
	Message string
	When time.Time
}