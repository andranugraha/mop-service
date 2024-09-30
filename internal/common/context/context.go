package contextutil

import (
	"context"
	"github.com/empnefsi/mop-service/internal/common/strings"
)

func GetUserID(ctx context.Context) uint64 {
	if ctx == nil {
		return 0
	}
	if v := ctx.Value("user"); v != nil {
		if user, ok := v.(*strings.Claims); ok {
			return user.UserID
		}
	}
	return 0
}

func GetMerchantID(ctx context.Context) uint64 {
	if ctx == nil {
		return 0
	}
	if v := ctx.Value("user"); v != nil {
		if user, ok := v.(*strings.Claims); ok {
			return user.MerchantID
		}
	}
	return 0
}

func GetRole(ctx context.Context) uint32 {
	if ctx == nil {
		return 0
	}
	if v := ctx.Value("user"); v != nil {
		if user, ok := v.(*strings.Claims); ok {
			return user.Role
		}
	}
	return 0
}

func GetEmail(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v := ctx.Value("user"); v != nil {
		if user, ok := v.(*strings.Claims); ok {
			return user.Email
		}
	}
	return ""
}
