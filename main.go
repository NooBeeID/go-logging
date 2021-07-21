package main

import (
	"github.com/NooBeeID/go-logging/colors"
	"github.com/NooBeeID/go-logging/messages"
)

func main() {
	messages.SysInfo("Server running di port %v", ":4000")
	messages.SysError("pq: unexpected , near course_title")
	messages.SysWarning("This features will be deprecated soon !")
	messages.SysMessage("Ini default")
	messages.SysLog("Log sederhana")
	messages.NewLog("DANGER", colors.Red, "Gawaaaat !!!")

	messages.Print("Print biasa aja \n")
	messages.PrintWithColor(colors.Blue, "Print dengan biru \n")

	messages.Println("Print dengan new line")
	messages.PrintWithColorln(colors.Green, "Print dengan warna + new line")
}
