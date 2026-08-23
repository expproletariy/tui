package core

import (
	"errors"
	"fmt"
)

var (
	ErrTInitError = errors.New("ttab init error")

	ErrTInvalidTerminalDescriptor = errors.New("invlaid termanal descriptor")
	ErrTGetSize                   = errors.New("can not get termanal size")
)

func WrapErrors(err1, err2 error) error {
	return fmt.Errorf("%w:%w", err1, err2)
}
