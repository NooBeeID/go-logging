# go-logging
go-logging dapat membantu kamu untuk menampilkan `log` dengan penuh warna. go-logging sangat mirip penggunaannya dengan `fmt.Printf()`, sehingga bisa menggunakan format yang sama.

## How to use
```shell
$ go get github.com/NooBeeID/go-logging
```

## Quick Start
### System Messages
untuk mendapatkan Log yang sudah di format sebelumnya, bisa menggunakan :
```go
messages.SysInfo(msg string, args ...interface{})
messages.SysError(msg string, args ...interface{})
messages.SysWarning(msg string, args ...interface{})
messages.SysMessage(msg string, args ...interface{})
messages.SysLog(msg string, args ...interface{})
```

contoh :
```go
func main () {
    messages.SysInfo("Server running di port %v", ":4000")
	messages.SysError("pq: unexpected , near course_title")
	messages.SysWarning("This features will be deprecated soon !")
	messages.SysMessage("Ini default")
	messages.SysLog("Log sederhana")
}
```
yang mana akan menghasilkan :
```bash
[ INFO ] Server running di port :4000
[ ERROR ] pq: unexpected , near course_title
[ WARNING ] This features will be deprecated soon !
[ MESSAGE ] Ini default
[ LOG ] Log sederhana
```