package errs

var (
	ErrWalletDoesNotExist  = NewErrorWithCode(404, "ErrWalletDoesNotExist")
	ErrValidationFailed = NewErrorWithCode(400, "ErrValidationFailed")
	ErrInsufficientBalance = NewErrorWithCode(400, "ErrInsufficientBalance")
	ErrWalletBalanceLimitExceeded = NewErrorWithCode(400, "ErrWalletBalanceLimitExceeded")
)
