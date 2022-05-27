package log

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var tmpl *template.Template
var print []printType
var statistics *statType

func init() {
	var err error
	tmpl, err = template.ParseFiles("static/index.html")
	if err != nil {
		panic(err)
	}
	statistics = &statType{
		Start:      time.Now().Format(layout),
		CountError: 0,
		CountLog:   0,
		Fields:     map[string]interface{}{},
	}
}

func addPrint(l *logType) {
	var atr []string
	for _, v := range l.fields {
		atr = append(atr, fmt.Sprintf("%s=%s", v.key, v.val))
	}

	statistics.CountLog++

	if l.types == "ERRO" || l.types == "WARN" {
		statistics.CountError++
	}

	print = append(print, printType{
		sysDate: l.data,
		Date:    l.data.Format(layout),
		Types:   l.types,
		Atr:     atr,
		Message: l.message,
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, print)
}

func HandlerStat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(statistics)
}

func GetStatUrl() string {
	return "/admin/stat"
}

func GetHandlerUrl() string {
	return "/admin/log"
}
