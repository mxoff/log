package log

import (
	"time"
)

func init() {
	go cleaner()
}

func cleaner() {
	for {
		time.Sleep(1 * time.Hour)
		for i := range print {
			if print[i].sysDate.AddDate(0, 3, 0).Unix() < time.Now().Unix() {
				//fmt.Println("DELL ", i, "   ", print[i].sysDate.Add(time.Second).Unix(), "==", time.Now().Unix(), print[i].Message)
				print = append(print[:i], print[i+1:]...)
				break
			}
		}
	}
}
