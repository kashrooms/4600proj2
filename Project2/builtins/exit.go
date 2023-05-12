package builtins

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	ErrInvalidArgCnt = errors.New("invalid argument count")
)

func Exit(args ...string) error {

	exitCode, _ := strconv.Atoi(args[1])

	switch len(args) {
		case 0:
			os.Exit(0)
		case 1:
			os.Exit(exitCode)
	}
	
	return fmt.Errorf("%w: expected 0 or 1 arguments", ErrInvalidArgCount)
}
