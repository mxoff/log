package main

import "time"

type printType struct {
	sysDate time.Time
	Date    string
	Types   string
	Atr     []string
	Message string
}

type statType struct {
	Start      string
	CountError int
	CountLog   int
	Fields     map[string]interface{}
}

type fieldType struct {
	key string
	val string
}

type logType struct {
	data    time.Time
	types   string
	color   string
	fields  []fieldType
	message string
}

var colors = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"gray":   "\033[37m",
	"white":  "\033[97m",
}
