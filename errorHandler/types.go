package errorHandler

type BadRequestError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type UnAuthorizedError struct {
	Message string
}
type NotFoundError struct {
	Message string
}

type ForbiddenError struct {
	Message string
}

func (err *BadRequestError) Error() string {
	return err.Message
}

func (err *InternalServerError) Error() string {
	return err.Message
}

func (err *NotFoundError) Error() string {
	return err.Message
}

func (err *UnAuthorizedError) Error() string {
	return err.Message
}

func (err *ForbiddenError) Error() string {
	return err.Message
}
