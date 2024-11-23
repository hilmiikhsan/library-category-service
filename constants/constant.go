package constants

const (
	SuccessMessage                = "success"
	ErrCategoryAlreadyExist       = "category already exist"
	ErrFailedBadRequest           = "failed to parse request"
	ErrAuthorizationIsEmpty       = "authorization is empty"
	ErrInvalidAuthorizationFormat = "invalid authorization format"
	ErrInvalidAuthorization       = "invalid authorization"
	ErrCategoryNotFound           = "category not found"
	ErrParamIdIsRequired          = "param id is required"
	ErrIdIsNotValidUUID           = "id is not valid uuid"
)

const (
	HeaderAuthorization = "Authorization"
	TokenTypeAccess     = "token"
)
