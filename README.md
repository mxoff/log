#log from OFF

## Install
```sh
go get github.com/mxoff/log
```
---
## Examples
```go
log.Logger().WithField("port", ":8080").Debugln("Start server")

r.HandleFunc(log.GetUrlHandler(), log.GetUrlGeneral)
r.HandleFunc(log.GetUrlLog(), log.HandlerLog)
r.HandleFunc(log.GetUrlStat(), log.HandlerStat)
```
задать новую длину в консоли
```go
log.NewCountLenStrLog(55)
```
---
## View
```sh
/admin/stat
/admin/log
```
log.SetGeneralUrl("/tool/")
log.GetUrlHandler()

---
## Stat
```go
log.Stat().StatFieldChange("Field", 123)
log.Stat().StatFieldChange("New", "123")
log.Stat().StatFieldChange("Field", "123")
```