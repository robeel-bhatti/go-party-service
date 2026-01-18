package internal

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func CacheMiddleware(ctx context.Context, next http.HandlerFunc, cache *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := cache.Get(ctx, r.URL.Path).Result()

		if errors.Is(err, redis.Nil) {
			next(w, r)
			return
		}

		if err != nil {
			// write 500 response here
			return
		}

		_, err = w.Write([]byte(res))
		if err != nil {
			// write 500 response here or just log the error
			return
		}
	}
}
