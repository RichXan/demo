package xerror

import (
	"fmt"
)

type Error struct {
	code    int
	message string
}

var codes = map[int]string{}

func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = message
	return &Error{code: code, message: message}
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.message
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code:%d, Messsage:%s", e.code, e.message)
}
