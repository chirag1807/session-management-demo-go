package errorhandling

import "net/http"

type CustomError struct {
	StatusCode int
	Message    string
}

func (c CustomError) Error() string {
	return c.Message
}

var (
	ReadBodyError           = CreateCustomError("Could not Read Request Body, Please Provide Valid Body.", http.StatusBadRequest)
	ReadDataError           = CreateCustomError("Could not Decode the Data, Please Provide Valid Data.", http.StatusBadRequest)
	EmailvalidationError    = CreateCustomError("Email Validation Failed, Please Provide Valid Email.", http.StatusBadRequest)
	DuplicateEmailFound     = CreateCustomError("Duplicate Email Found.", http.StatusConflict)
	RegistrationFailedError = CreateCustomError("User Registration Failed.", http.StatusInternalServerError)
	LoginFailedError        = CreateCustomError("User Login Failed.", http.StatusUnauthorized)
	UnauthorizedError       = CreateCustomError("You are Not Authorized to Perform this Action.", http.StatusUnauthorized)
	NoUserFound             = CreateCustomError("No User Found for This Request.", http.StatusNotFound)
	PasswordNotMatch        = CreateCustomError("Password is Incorrect.", http.StatusUnauthorized)
	SessionExpired          = CreateCustomError("Your Session has Expired, Please Login Again to Continue.", 440)
)

func CreateCustomError(Message string, StatusCode int) error {
	return CustomError{
		StatusCode: StatusCode,
		Message:    Message,
	}
}
