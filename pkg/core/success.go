package core

type DefaultSuccess interface {
	HttpStatus() int
	Msg() string
}
type Success struct {
	httpStatus int
	msg        string
}

func (o Success) HttpStatus() int { return o.httpStatus }
func (o Success) Msg() string     { return o.msg }

func NewSuccess(status int, msg string) Success {
	return Success{
		httpStatus: status, msg: msg,
	}
}
