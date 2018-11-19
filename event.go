package main

import "regexp"

const (
	EventCodeFormat = `fc6b([a-z0-9]{4})`
)

var (
	EventCodeRegExp = regexp.MustCompile(EventCodeFormat)
)
