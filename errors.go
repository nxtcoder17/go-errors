package errors

import (
	"fmt"
	"github.com/yext/yerrors"
)

func NewEf(err error, msg string, a ...any) error {
	return yerrors.WrapFrame(yerrors.Errorf("%s when %w", fmt.Sprintf(msg, a...), err), 1)
}

func Newf(msg string, a ...any) error {
	if len(a) > 0 {
		return yerrors.WrapFrame(yerrors.Errorf(msg, a...), 1)
	}
	return yerrors.WrapFrame(yerrors.New(msg), 1)
}

func NewE(err error) error {
	return yerrors.WrapFrame(err, 1)
}

func New(msg string) error {
	return yerrors.WrapFrame(yerrors.New(msg), 1)
}
