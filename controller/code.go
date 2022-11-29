package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeError
)

var Codemap = map[ResCode]string{
	CodeSuccess: "success",
	CodeError:   "error",
}
