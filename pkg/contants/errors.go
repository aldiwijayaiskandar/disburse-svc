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

var statusCodes = map[ErrorCode]int{
	NoError:             200,
	InternalServerError: 500,
	InvalidRequest:      400,
	NotFound:            404,
	Unauthorized:        401,
}

func (e ErrorCode) Error() string {
	return errorMessages[e]
}

func (e ErrorCode) StatusCode() int {
	return statusCodes[e]
}
