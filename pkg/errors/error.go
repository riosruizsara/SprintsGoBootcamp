package errors

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) Is(target error) bool {
	_, ok := target.(*ValidationError)
	return ok
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

type DuplicateError struct {
	Message string
}

func (e *DuplicateError) Error() string {
	return e.Message
}

func (e *DuplicateError) Is(target error) bool {
	_, ok := target.(*DuplicateError)
	return ok
}

type ConflictError struct {
	Message string
}

func (e *ConflictError) Error() string {
	return e.Message
}

func (e *ConflictError) Is(target error) bool {
	_, ok := target.(*ConflictError)
	return ok
}

type UnknownError struct {
	Message string
}

func (e *UnknownError) Error() string {
	return e.Message
}

func (e *UnknownError) Is(target error) bool {
	_, ok := target.(*UnknownError)
	return ok
}
