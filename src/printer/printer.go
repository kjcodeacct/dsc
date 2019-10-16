package printer

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
	colorable "github.com/mattn/go-colorable"
)

var printWriter io.Writer

const tabChar = `    `
const newLine = `\n`

func init() {
	if runtime.GOOS == "windows" {
		printWriter = colorable.NewColorableStdout()
	} else {
		printWriter = os.Stdout
	}
}

func Red(args ...interface{}) {
	fmt.Fprint(printWriter, aurora.Red(args))
}

func Blue(args ...interface{}) {
	fmt.Fprint(printWriter, aurora.Blue(args))
}

func Yellow(args ...interface{}) {
	fmt.Fprint(printWriter, aurora.Yellow(args))
}

func Green(args ...interface{}) {
	fmt.Fprint(printWriter, aurora.Green(args))
}

func Fatalln(args ...interface{}) {
	args = append(args, newLine)
	fmt.Fprint(printWriter, aurora.Red(args), newLine)
	panic(args)
}

func Printf(format string, args ...interface{}) {
	fmt.Fprintf(printWriter, format, args)
}

func Println(format string, args ...interface{}) {
	args = append(args, newLine)
	fmt.Fprintf(printWriter, format, args)
}

func AddedFiles(files []string) {

}

func RemovedFiles(files []string) {

}
