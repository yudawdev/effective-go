package error

import (
	"errors"
	"strconv"
	"syscall"
)

// MyError 自定义错误结构
type MyError struct {
	Code int
	Msg  string
	Err  error
}

// Error 实现 error 接口的方法
func (m *MyError) Error() string {
	return "code:" + strconv.Itoa(m.Code) + "|" + "msg:" + m.Msg + "|" + "error:" + m.Err.Error()
}

// 返回 error ，请注意： *MyError 是实现了 Error 方法，即实现 error 接口。但是 MyError 类型则没有实现 Error
func testErr() (int, error) {
	return 0, &MyError{
		Code: 0,
		Msg:  "",
		Err:  nil,
	}
}

// 如何判断错误类型？
func testErrType() int {
	i, err := testErr()
	if err == nil {
		return i
	}

	// 判断错误类型
	var e *MyError
	if errors.As(err, &e) && errors.Is(e.Err, syscall.ENOSPC) {
		return 0
	}

	return 0
}
