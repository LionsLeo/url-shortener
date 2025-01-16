package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Db struct {
	Redis *redis.Client
}

func InitDb() (*Db, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	db := Db{Redis: rdb}
	var ctx = context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		//! error connecting to redis cache.
		return nil, err
	}
	return &db, nil
}

func (db *Db) Close() {
	db.Redis.Conn().Close()
}
