package itemcategory

import (
	"github.com/go-redis/redis/v8"
)

type cache struct {
	client *redis.Client
}
