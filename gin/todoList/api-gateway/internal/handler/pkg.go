package handler

import "errors"

// PanicIfUserError 定义错误
func PanicIfUserError(err error)  {
	if err != nil {
		err = errors.New("User Service -- "+err.Error())
		panic(err)
	}
}

// PanicIfTaskError 定义错误
func PanicIfTaskError(err error)  {
	if err != nil {
		err = errors.New("Task Service -- "+err.Error())
		panic(err)
	}
}