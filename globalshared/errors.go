package globalshared

import "github.com/golangid/candi/candihelper"

type ErrorResponse struct {
	Code       int
	Message    string
	MultiError candihelper.MultiError
}

func (e *ErrorResponse) Error() string {
	return e.Message
}
