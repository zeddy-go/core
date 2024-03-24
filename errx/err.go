package errx

import "github.com/zeddy-go/zeddy/slicex"

type InfoKey string

const (
	Msg      InfoKey = "msg"
	Err      InfoKey = "err"
	File     InfoKey = "file"
	Line     InfoKey = "line"
	ErrStack InfoKey = "errStack"
	Code     InfoKey = "code"
	Detail   InfoKey = "detail"
)

var reserveKeys = []InfoKey{
	Msg,
	Err,
	File,
	Line,
	ErrStack,
	Code,
	Detail,
}

func (d InfoKey) IsReserved() bool {
	return slicex.Contains(d, reserveKeys)
}
