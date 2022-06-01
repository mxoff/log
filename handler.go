package log

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl *template.Template
var print []printType
var statistics *statErrorType

func init() {
	var err error
	tmpl, err = template.New("index").Parse(tmpl_index)
	if err != nil {
		log.Fatalln(err)
	}
	statistics = &statErrorType{
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

func HandlerLog(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, print)
}

func HandlerCommon(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(statistics)
}

func HandlerStat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stat)
}

var generalUrl = "/tool/"

func GetGeneralUrl() string {
	return generalUrl
}

// SetGeneralUrl string must start and end with "/"
func SetGeneralUrl(u string) string {
	if u != "" && (u[:1] == "/" && u[len(u)-1:] == "/") {
		generalUrl = u
	}
	return generalUrl
}

func GetUrlGeneral() string {
	return generalUrl
}

func GetUrlLog() string {
	return generalUrl + "log"
}

func GetUrlStat() string {
	return generalUrl + "stat"
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI != GetUrlLog() &&
			r.RequestURI != GetUrlGeneral() &&
			r.RequestURI != GetUrlStat() {
			addStat(r.RequestURI)
		}

		next.ServeHTTP(w, r)
	})
}
