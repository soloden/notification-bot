package apperror

import "errors"

type AppError struct {
	Code string `json:"code,omitempty"`
	Err  error  `json:"-"`
}

func (err *AppError) Error() string {
	return err.Err.Error()
}

func NewAppError(code, message string) *AppError {
	return &AppError{
		Code: code,
		Err:  errors.New(message),
	}
}
