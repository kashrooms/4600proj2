package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCnt = errors.New("invalid argument count")
	WDir, _          = os.Getwd()
)

func PrintWorkingDirectory(args ...string) error {

	switch len(args) {
		case 0: // change to home directory if available
			fmt.Println("Current working directory: ", WDir)
			return nil
		default:
			return fmt.Errorf("%w: expected zero arguments", ErrInvalidArgCount)
		}
}