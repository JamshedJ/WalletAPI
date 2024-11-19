package errs

var (
	ErrWalletDoesNotExist  = NewErrorWithCode(404, "ErrWalletDoesNotExist")
	ErrValidationFailed = NewErrorWithCode(400, "ErrValidationFailed")
	ErrWalletBalanceLimitExceeded = NewErrorWithCode(400, "ErrWalletBalanceLimitExceeded")
)
