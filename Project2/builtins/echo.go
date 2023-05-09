package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
	HomeDir, _         = os.UserHomeDir()
)

func PrintEcho(args ...string) error {
	switch len(args){
		case 0:
			return fmt.Errorf("%w: expected one or more arguments (echo)", ErrInvalidArgCount)
		default:
			switch args[0] {
				case "-n": // change to home directory if available
					return fmt.Println(args[1:])
				default:
				return fmt.Print(args[0:])
			}
	}
}
