/*
Fancy Errors

This package wraps the errors used in dsc throughout with a stacktrace, but follows the
	standard error package api, with the addition of the Cause method

TODO determine if this should be in a seperate go repo altogether
	(should anyone want it with the coloration)
*/

package fancy_errors

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// TBD want to handle specific sql driver errors?

type StackTraceError struct {
	frame     runtime.Frame
	TimeStamp time.Time
	Err       error
}

func New(errMsg string) error {

	err := errors.New(errMsg)

	return addStackTrace(3, err)
}

func Wrap(err error) error {
	return addStackTrace(3, err)
}

func addStackTrace(offset int, err error) error {

	var frame [3]uintptr

	runtime.Callers(offset, frame[:])
	frames := runtime.CallersFrames(frame[:])
	errorFrame, _ := frames.Next()

	newStackTraceError := StackTraceError{
		frame:     errorFrame,
		TimeStamp: time.Now().UTC(),
		Err:       err,
	}

	return newStackTraceError
}

func (this StackTraceError) Cause() string {

	if this.Err != nil {
		errMsg := fmt.Sprintf("%s\n\t%s\n\t%s:%d", this.Error(), this.frame.Function,
			this.frame.File, this.frame.Line)
		return errMsg
	}

	// nonErrMsg shouldn't really occur if an actual error is generated, but can be used to
	// 		return a stack trace directly
	nonErrMsg := fmt.Sprintf("\n\t%s\n\t%s:%d", this.frame.Function, this.frame.File,
		this.frame.Line)

	return nonErrMsg
}

func (this StackTraceError) Error() string {
	return this.Err.Error()
}
