package method

type Method int64

const (
	GET Method = iota + 1
	POST
	PUT
	PATCH
	DELETE
	HEAD
	OPTIONS
)
