package v1

// TypeError is a type of custom error
type TypeError uint16

// The list of error codes
const (
	ErrForbidden    TypeError = iota + 10 // 10, ErrForbidden is used when the user do not have the permission to access the resource
	ErrUnauthorized                       // 11, ErrUnauthorized is used when the user is not logged in
	ErrDatabase                           // 12, ErrDatabase is used when there is an error when interacting with the database
	ErrConflict                           // 13, ErrConflict is used when there is a conflict when saving data
	ErrFromUseCase                        // 14, ErrFromUseCase is used when there is an error from the use case
	ErrValidation                         // 15, ErrValidation is used when there is an error with the validation
	ErrNoFound                            // 16, ErrNoFound is used when the data is not found
	ErrUnknown                            // 17, ErrUnknown is used when the error is unknown
)
