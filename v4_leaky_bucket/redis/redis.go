package redis

import (
	"time"
	"github.com/garyburd/redigo/redis"
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

func New(network, address string) (*Storage, error) {
	s := &Storage{
		pool: redis.NewPool(func (redis.Conn, error) {
			return redis.Dial(network, address)
		}, 5)}

	conn := s.pool.Get()
	defer s.pool.Close()
	if _, err := conn.Do("PING"); err != nil {
		return nil, err
	}

	return s, nil
}