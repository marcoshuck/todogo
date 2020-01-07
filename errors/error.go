package errors

// Error is a wrapper that contains a Base error and additional fields
type Error struct {
	Code int64
	Status int64
	Message string
	Base *error
}

func NewError(code int64, status int64, message string, base *error) *Error {
	return &Error{
		Code:    code,
		Status:  status,
		Message: message,
		Base:    base,
	}
}