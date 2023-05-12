package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCnt = errors.New("invalid argument count")
)

func Exit(args ...string) error {

	switch len(args) {
		case 0:
			return os.Exit()
		case 1:
			return os.Exit(args[1])
		default:
			return fmt.Errorf("%w: expected 0 or 1 arguments", ErrInvalidArgCount)
	}
}