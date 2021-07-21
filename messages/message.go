package messages

import (
	"fmt"

	"github.com/NooBeeID/go-logging/colors"
)

func SysInfo(msg string, args ...interface{}) {
	stringMsg := "[ INFO ] " + msg + "\n"
	printMsg(string(colors.Cyan), stringMsg, args...)
}

func SysError(msg string, args ...interface{}) {
	stringMsg := "[ ERROR ] " + msg + "\n"
	printMsg(string(colors.Red), stringMsg, args...)
}

func SysWarning(msg string, args ...interface{}) {
	stringMsg := "[ WARNING ] " + msg + "\n"
	printMsg(string(colors.Yellow), stringMsg, args...)
}

func SysMessage(msg string, args ...interface{}) {
	stringMsg := "[ MESSAGE ] " + msg + "\n"
	printMsg(string(colors.White), stringMsg, args...)
}

func SysLog(msg string, args ...interface{}) {
	stringMsg := "[ LOG ] " + msg + "\n"
	printMsg(string(colors.White), stringMsg, args...)
}

func NewLog(logType string, color string, msg string, args ...interface{}) {
	stringMsg := "[ " + logType + " ] " + msg + "\n"
	printMsg(string(color), stringMsg, args...)
}

func Print(msg string, args ...interface{}) {
	printMsg(string(colors.White), msg, args...)
}

func PrintWithColor(color string, msg string, args ...interface{}) {
	if color == "" {
		color = colors.White
	}
	printMsg(string(color), msg, args...)
}

func Println(msg string, args ...interface{}) {
	printMsg(string(colors.White), msg+"\n", args...)
}

func PrintWithColorln(color string, msg string, args ...interface{}) {
	if color == "" {
		color = colors.White
	}
	printMsg(string(color), msg+"\n", args...)
}

func clearColor() {
	fmt.Print(string(colors.Reset))
}

func printMsg(color, msg string, args ...interface{}) {
	fmt.Print(string(color))
	fmt.Printf(msg, args...)
	clearColor()
}
