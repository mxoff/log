package main

import "fmt"

func (l *logType) WithField(key string, value interface{}) *logType {
	l.fields = append(l.fields, fieldType{key: key, val: fmt.Sprint(value)})
	return l
}

func (l *logType) WithError(err error) *logType {
	if err != nil {
		l.fields = append(l.fields, fieldType{
			key: "error",
			val: err.Error(),
		})
	}
	return l
}
