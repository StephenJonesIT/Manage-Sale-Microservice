package common

type AppError struct{
	StatusCode 	int 	`json:"status_code"`
	RootError 	error 	`json:"-"`
	Message 	string 	`json:"message"`
	Log 		string 	`json:"log"`
	Key 		string 	`json:"key"`
}

func NewAppError(statusCode int, rootError error, message string, log string, key string) *AppError{
	return &AppError{
		StatusCode: statusCode,
		RootError: rootError,
		Message: message,
		Log: log,
		Key: key,
	}
}

func (e *AppError) Error() string{
	return e.RootErr().Error()
}

func (e *AppError) RootErr() error{
	if err, ok := e.RootError.(*AppError); ok{
		return err.RootErr()
	}
	return e.RootError
}