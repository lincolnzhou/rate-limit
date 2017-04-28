package redis

import (
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/lincolnzhou/rate-limit/v4_leaky_bucket"
)

type bucket struct {
	name string
	capacity uint
	remaining uint
	reset time.Time
	rate time.Duration
	pool *redis.Pool
}

type Storage struct {
	pool *redis.Pool
}

func (s *Storage) Create(name string, capacity uint, rate time.Duration) (v4_leaky_bucket.BucketStat, error) {

}

func New(network, address string) (*Storage, error) {
	s := &Storage{
		pool: &redis.Pool{
			MaxIdle: 5,
			Dial: func() (redis.Conn, error) {
				return redis.Dial(network, address)
			},
		}}

	conn := s.pool.Get()
	defer s.pool.Close()
	if _, err := conn.Do("PING"); err != nil {
		return nil, err
	}

	return s, nil
}