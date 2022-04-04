#log from OFF

## Install
```sh
go get github.com/.../log
```
---
## Examples
```go
log.Logger().WithField("port", ":8080").Debugln("Start server")
r.HandleFunc(log.GetHandlerUrl(), log.Handler)
r.HandleFunc(log.GetStatUrl(), log.HandlerStat)
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

---
## Stat
```go
log.Stat().StatFieldChange("Field", 123)
log.Stat().StatFieldChange("New", "123")
log.Stat().StatFieldChange("Field", "123")
```