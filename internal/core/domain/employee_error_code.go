package domain

type ErrorCode string

const (
	InternalServerError ErrorCode = "EMP9991"
	BadRequestError     ErrorCode = "EMP9992"
)
