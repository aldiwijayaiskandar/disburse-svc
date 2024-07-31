package constants

type ErrorCode int

const (
	NoError ErrorCode = iota
	InternalServerError
	InvalidRequest
	NotFound
	Unauthorized
)

var errorMessages = map[ErrorCode]string{
	NoError:             "no error",
	InternalServerError: "internal server error",
	InvalidRequest:      "invalid request",
	NotFound:            "not found",
	Unauthorized:        "unauthorized",
}

func (e ErrorCode) Error() string {
	return errorMessages[e]
}
