package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"auctionkuy.wildangbudhi.com/domain/v1/auth"
	"github.com/go-redis/redis/v8"
)

type sessionRepository struct {
	redisDB *redis.Client
	ctx     context.Context
}

func NewSessionRepository(redisDB *redis.Client) auth.SessionRepository {
	return &sessionRepository{
		redisDB: redisDB,
		ctx:     context.Background(),
	}
}

func (repo *sessionRepository) IsSessionExist(key string) (bool, error) {

	var err error

	var ctx context.Context
	var ctxCancel context.CancelFunc

	ctx, ctxCancel = context.WithTimeout(repo.ctx, time.Second*10)
	defer ctxCancel()

	var value int64

	value, err = repo.redisDB.Exists(ctx, key).Result()

	if err != nil {
		log.Println(err)
		return false, fmt.Errorf("Service Unavailable")
	}

	if value == 1 {
		return true, nil
	}

	return false, nil

}

func (repo *sessionRepository) SetSession(key string, data *auth.Session, expiration time.Duration) error {

	var err error

	var ctx context.Context
	var ctxCancel context.CancelFunc

	ctx, ctxCancel = context.WithTimeout(repo.ctx, time.Second*10)
	defer ctxCancel()

	var value []byte

	value, err = json.Marshal(data)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable")
	}

	err = repo.redisDB.Set(ctx, key, string(value), expiration).Err()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable")
	}

	return nil

}

func (repo *sessionRepository) GetSession(key string) (*auth.Session, error) {

	var err error

	var ctx context.Context
	var ctxCancel context.CancelFunc

	ctx, ctxCancel = context.WithTimeout(repo.ctx, time.Second*10)
	defer ctxCancel()

	var value string

	value, err = repo.redisDB.Get(ctx, key).Result()

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Service Unavailable")
	}

	var data *auth.Session = new(auth.Session)

	err = json.Unmarshal([]byte(value), data)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Service Unavailable")
	}

	return data, nil
}

func (repo *sessionRepository) RemoveSession(key string) error {

	var err error

	var ctx context.Context
	var ctxCancel context.CancelFunc

	ctx, ctxCancel = context.WithTimeout(repo.ctx, time.Second*10)
	defer ctxCancel()

	err = repo.redisDB.Del(ctx, key).Err()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable")
	}

	return nil

}

func (repo *sessionRepository) ExtendSessionExpiration(key string, expiration time.Duration) error {

	var err error

	var ctx context.Context
	var ctxCancel context.CancelFunc

	ctx, ctxCancel = context.WithTimeout(repo.ctx, time.Second*10)
	defer ctxCancel()

	err = repo.redisDB.Expire(ctx, key, expiration).Err()

	if err != nil {
		log.Println(err)
		return fmt.Errorf("Service Unavailable")
	}

	return nil

}
