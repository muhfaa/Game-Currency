package business

import "errors"

var (
	//ErrNotFound Error when item is not found
	ErrNotFound = errors.New("Data was not found")

	//ErrInvalidSpec Error when data given is not valid on update or insert
	ErrInvalidSpec = errors.New("Given spec is not valid")

	//ErrDuplicate Duplicate ID
	ErrDuplicate = errors.New("Duplicate ID")

	//ErrForbidden forbidden
	ErrForbidden = errors.New("forbidden")
)
