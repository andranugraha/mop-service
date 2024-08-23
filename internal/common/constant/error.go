package constant

import "fmt"

var (
	ErrInvalidIdentifierOrPassword = fmt.Errorf("invalid identifier or password")
	ErrFailedToGetUser             = fmt.Errorf("failed to get user")
)

const (
	ErrCodeInternalServer = int32(10000)
	ErrCodeInvalidParam   = int32(10001)
)
