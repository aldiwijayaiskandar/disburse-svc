package constants

type StatusCode int

const (
	NoStatus StatusCode = iota
	Success
	Error
)

var statusMessage = map[StatusCode]string{
	Success: "success",
	Error:   "error",
}

func (s StatusCode) Status() string {
	return statusMessage[s]
}
