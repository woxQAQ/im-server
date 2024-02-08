package utils

import (
	"runtime"
	"strconv"

	"github.com/pkg/errors"
)

func WrapWithCallerInfo(err error) error {
	return errors.Wrap(err, printCallerNameAndLine())
}

func printCallerNameAndLine() string {
	pc, _, line, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name() + "()@" + strconv.Itoa(line) + ": "
}
