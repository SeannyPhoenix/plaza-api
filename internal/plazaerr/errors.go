package plazaerr

import "fmt"

// DuplicateAccount ... x
func DuplicateAccount(key string, value string) error {
	switch key {
	case "name":
		return ErrNameNotUnique.Errorf("Account exists with name %s", value)
	case "email":
		return ErrEmailNotUnique.Errorf("Account exists with email %s", value)
	}

	return nil
}

type plazaErrorCode string

func (e *plazaErrorCode) Error() string {
	return string(*e)
}

const (
	// ErrNameNotUnique
	ErrNameNotUnique    plazaErrorCode = "NAME_NOT_UNIQUE"
	ErrEmailNotUnique   plazaErrorCode = "EMAIL_NOT_UNIQUE"
	ErrEmailNotVerified plazaErrorCode = "EMAIL_NOT_VERIFIED"
)

type PlazaError struct {
	message string
	code    plazaErrorCode
	cause   error
}

func (e plazaErrorCode) New(message string) *PlazaError {
	return &PlazaError{message: message, code: e}
}

func (e plazaErrorCode) Wrap(message string, err error) *PlazaError {
	return &PlazaError{code: e, message: message, cause: err}
}

func (e plazaErrorCode) Errorf(format string, a ...interface{}) *PlazaError {
	err := fmt.Errorf(format, a...)
	return e.Wrap(err.Error(), err)
}

func (e *PlazaError) Error() string {
	if e.message != "" {
		return e.message
	}

	return string(e.code)
}

func (e *PlazaError) Unwrap() error {
	return e.cause
}

func (e *PlazaError) Is(target error) bool {
	if other, ok := target.(*plazaErrorCode); ok {
		return *other == e.code
	}

	return false
}

func (e *PlazaError) Extensions() map[string]interface{} {
	return map[string]interface{}{"code": e.code}
}
