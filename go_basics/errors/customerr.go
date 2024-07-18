package main

type CustomErr struct {
	msg string
}

func NewCustomErr(msg string) *CustomErr {
	return &CustomErr{msg: msg}
}

func (e CustomErr) Error() string {
	return e.msg
}

var _ error = (*CustomErr)(nil) // позволяет проверить, что структура CustomerErr соответствует интерфейсу error
