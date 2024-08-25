package constant

import "fmt"

var (
	ErrInvalidIdentifierOrPassword = fmt.Errorf("invalid identifier or password")
	ErrFailedToGetUser             = fmt.Errorf("failed to get user")
	ErrFailedToGetMerchant         = fmt.Errorf("failed to get merchant")
)

const (
	ErrCodeInternalServer = int32(10000)
	ErrCodeInvalidParam   = int32(10001)
)
