package myRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go_web_counter/domain/cnt"
	"strconv"
)

type cntRepo struct {
	RCli *redis.Client
}

func NewDataRepository(rCli *redis.Client) cnt.ICntRepository {
	return &cntRepo{
		RCli: rCli,
	}
}

func (dr *cntRepo) Set(ctx context.Context, key string, val int) error {
	s := strconv.Itoa(val)
	return dr.RCli.Set(ctx, key, s, 0).Err()
}

func (dr *cntRepo) CntUp(ctx context.Context, key string) (int, error) {
	val, err := dr.RCli.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	intVal, err := strconv.Atoi(strconv.FormatInt(val, 10))
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func (dr *cntRepo) CntDown(ctx context.Context, key string) (int, error) {
	val, err := dr.RCli.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	intVal, err := strconv.Atoi(strconv.FormatInt(val, 10))
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func (dr *cntRepo) Get(ctx context.Context, key string) (int, error) {
	val, err := dr.RCli.Get(ctx, key).Result()
	if err == redis.Nil {
		return 0, err
	} else if err != nil {
		return 0, err
	}

	cnt, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}
