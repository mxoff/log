package main

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestTempMain(t *testing.T) {
	go Exemple()
	go func() {
		time.Sleep(3 * time.Minute)
		//time.Sleep(20 * time.Second)
		os.Exit(1)
	}()

	r := mux.NewRouter()
	r.HandleFunc(GetHandlerUrl(), Handler)
	r.HandleFunc(GetStatUrl(), HandlerStat)

	serverUrl := fmt.Sprintf("%s:%d", "localhost", 8080)
	err := http.ListenAndServe(serverUrl, r)
	if err != nil {
		panic(err)
	}
}

func Exemple() {
	Stat().StatFieldChange("qwe", 456)
	Logger().WithField("port", ":8080").Debugln("Start server")
	Logger().WithField("port", 8080).Debugln("Start server")
	time.Sleep(time.Second)
	Logger().
		WithField("key1", 123).
		WithField("key2", "asdasd ").
		WithField("key3", 123123123).
		WithField("key4", 14564523).
		WithField("key5", 129708903).
		Println("много ключей")
	time.Sleep(time.Second)
	Logger().WithField("key", 123).Println("123456789 123456789 123456789 ")
	time.Sleep(time.Second)
	Logger().WithField("www", 456).Println("123456789 123456789 ")
	time.Sleep(time.Second)
	Logger().WithField("rrr", "qweqwe").Println("123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 ")
	time.Sleep(time.Second)
	Logger().WithField("ppp", "lll").Println("qwertyuiop[]asdfghjkl;zxcvbnm,./qwertyuiop[]asdfghjkl;zxcvbnm,./")
	time.Sleep(time.Second)
	Logger().Println("1111")
	time.Sleep(time.Second)
	Stat().StatFieldChange("qwe", "END")
	Logger().WithField("module", "loginn erroe").Warningln("почти ошибка")

	//	paclog.Logger().Errorln(errors.New("Ошибка чтения данных"))
	//	paclog.Logger().Errorln(nil)
	time.Sleep(time.Second)
	qwe := Logger()
	qwe.WithError(errors.New("error parsing")).Errorln("получение пользователя")
	qwe.WithError(errors.New("error parsing")).
		WithError(errors.New("error parsing 222")).
		Errorln("получение пользователя")
	Logger().WithError(nil).Errorln("получение пользователя 22")
	time.Sleep(time.Second)

	Logger().ErrorNoNil(errors.New("Error"))
	Logger().ErrorNoNil(nil)
	time.Sleep(time.Second)
	Logger().WithField("k1", 111).ErrorNoNil(errors.New("Error key"))
	Logger().WithField("k2", 222).ErrorNoNil(nil)
}
