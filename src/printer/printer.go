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

func Red(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	fmt.Fprint(printWriter, aurora.Red(output))
}

func Blue(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	fmt.Fprint(printWriter, aurora.Blue(output))
}

func Yellow(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	fmt.Fprint(printWriter, aurora.Yellow(output))
}

func Green(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	fmt.Fprint(printWriter, aurora.Green(output))
}

func Fatalln(format string, args ...interface{}) {
	output := fmt.Sprintf(format, args...)
	fmt.Fprint(printWriter, aurora.Red(output))
	panic(output)
}

func Printf(format string, args ...interface{}) {
	fmt.Fprintf(printWriter, format, args...)
}

func Println(format string, args ...interface{}) {
	args = append(args, newLine)
	fmt.Fprintf(printWriter, format, args)
}
