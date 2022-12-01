package errors

type MylError interface {
	// Satisfy generic error interface
	error
	Code() string
	Message() string
	OrigErrs() []error
}

func New(code, message string, err error) MylError {
	var errs []error
	if err != nil {
		errs = append(errs, err)
	}
	return newBaseError(
		code, message, errs,
	)
}

type ComandError interface {
	MylError
	Token() string
}

func NewCommandError(err MylError, token string) ComandError {
	return newCommandError(err,  token)
}