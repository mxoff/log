package log

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type statPage struct {
	Url   string
	Count map[int]statCount
}
type statCount struct {
	All   int
	Hour  int
	Day   int
	Month int
}

type statType struct {
	Pages map[string]statPage
	statCount
}

var stat statType

func init() {
	stat = statType{
		Pages:     map[string]statPage{},
		statCount: statCount{},
	}
}

func addStat(addr string) {
	if addr == "" {
		return
	}
	addr = buildUrl(addr)

	m := getNowMount()

	// обьявление новой страницы
	if _, ok := stat.Pages[addr]; !ok {
		stat.Pages[addr] = statPage{
			Url:   addr,
			Count: make(map[int]statCount),
		}
	}

	stat.statCount = ucgrade(stat.statCount)
	stat.Pages[addr].Count[m] = ucgrade(stat.Pages[addr].Count[m])
}

func ucgrade(sc statCount) statCount {
	reset()
	sc.Hour++
	sc.Day++
	sc.Month++
	sc.All++
	return sc
}

// TODO debug
func GetStat() {
	m := getNowMount()
	fmt.Print(stat.All, "  ")
	for _, v := range stat.Pages {
		fmt.Print(v.Url, " = ", v.Count[m], "     ")
	}
	fmt.Println()
}

func ShowStatAll() statType {
	return stat
}

func ShowStatPage(addr string) statPage {
	return stat.Pages[addr]
}

var old = time.Now()

func reset() {
	now := time.Now()
	d := [3]bool{}
	if old.Hour() != now.Hour() {
		d[0] = true
	}
	if old.Day() != now.Day() {
		d[1] = true
	}
	if old.Month() != now.Month() {
		d[2] = true
	}
	drop(d)
	old = time.Now()
}

func drop(c [3]bool) {
	m := getNowMount()
	for k, p := range stat.Pages {
		count := p.Count[m]
		switch true {
		case c[0]:
			count.Hour = 0
		case c[1]:
			count.Day = 0
		case c[2]:
			count.Month = 0
		}
		p.Count[m] = count
		stat.Pages[k] = p
	}
}

func getNowMount() int {
	mount, _ := strconv.Atoi(time.Now().Format("01"))
	return mount
}

func buildUrl(str string) string {
	ar := strings.Split(str, "?")
	if len(ar) == 2 {
		str = ar[0]
	}

	if str[len(str)-1:] == "/" {
		str = str[:len(str)-1]
	}
	return str
}
