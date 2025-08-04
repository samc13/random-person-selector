package logging

import "fmt"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Bold    = "\033[1m"
)

// Println with color
func PrintRed(a ...interface{})    { fmt.Println(Red + fmt.Sprint(a...) + Reset) }
func PrintGreen(a ...interface{})  { fmt.Println(Green + fmt.Sprint(a...) + Reset) }
func PrintYellow(a ...interface{}) { fmt.Println(Yellow + fmt.Sprint(a...) + Reset) }
func PrintBlue(a ...interface{})   { fmt.Println(Blue + fmt.Sprint(a...) + Reset) }
func PrintBold(a ...interface{})   { fmt.Println(Bold + fmt.Sprint(a...) + Reset) }

// Printf with color
func PrintfRed(format string, a ...interface{})    { fmt.Printf(Red+format+Reset+"\n", a...) }
func PrintfGreen(format string, a ...interface{})  { fmt.Printf(Green+format+Reset+"\n", a...) }
func PrintfYellow(format string, a ...interface{}) { fmt.Printf(Yellow+format+Reset+"\n", a...) }
func PrintfBlue(format string, a ...interface{})   { fmt.Printf(Blue+format+Reset+"\n", a...) }
func PrintfBold(format string, a ...interface{})   { fmt.Printf(Bold+format+Reset+"\n", a...) }

func Info(a ...interface{}) {
	PrintGreen("[INFO ] ", fmt.Sprint(a...))
}
func Debug(a ...interface{}) {
	PrintYellow("[DEBUG] ", fmt.Sprint(a...))
}
func Error(a ...interface{}) {
	PrintRed("[ERROR] ", fmt.Sprint(a...))
}

func Infof(format string, a ...interface{}) {
	PrintfGreen("[INFO ] "+format, a...)
}
func Debugf(format string, a ...interface{}) {
	PrintfYellow("[DEBUG] "+format, a...)
}
func Errorf(format string, a ...interface{}) {
	PrintfRed("[ERROR] "+format, a...)
}
