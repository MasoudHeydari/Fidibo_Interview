package cache

import (
	"encoding/json"
	"fidibo_interview/contract"
	"fidibo_interview/entity"
	"github.com/go-redis/redis/v7"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) contract.BookCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, post *[]entity.Book) error {
	client := cache.getClient()

	// serialize Post object to JSON
	js, err := json.Marshal(post)
	if err != nil {
		return err
	}

	client.Set(key, js, cache.expires)
	return nil
}

func (cache *redisCache) Get(key string) (*[]entity.Book, error) {
	client := cache.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return nil, nil
	}

	books := make([]entity.Book, 0)
	err = json.Unmarshal([]byte(val), &books)
	if err != nil {
		return nil, err
	}

	return &books, nil
}
