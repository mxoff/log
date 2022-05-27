package log

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

var countLen = 60

const layout = "02.01.2006 15:04:05"

func (l *logType) print() {
	fields := ""
	for _, v := range l.fields {
		fields += fmt.Sprintf(" %s=%s ", addColor(l.color, v.key), v.val)
	}

	l.data = time.Now()
	str := fmt.Sprintf(" %s [%s] %s    %s",
		addColor(l.color, l.types), l.data.Format(layout), fixLen(l.message), fields)

	addPrint(l)
	fmt.Println(str)
}

func addColor(color, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, colors["reset"])
}

func fixLen(str string) string {
	c := utf8.RuneCountInString(str)
	if c > countLen {
		str = str[:countLen]
	}
	if c < countLen {
		str += strings.Repeat(" ", countLen-c)
	}
	return str
}

// default count = 60
func NewCountLenStrLog(count int) {
	if count > 0 {
		countLen = count
	}
}
