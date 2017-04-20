package v4_leaky_bucket

import (
	"errors"
	"time"
)

var (
	ErrorFull = errors.New("add exceeds free capacity")
)

type Bucket interface {
	Capacity() uint
	Remaining() uint
	Reset() time.Time
	Add(uint) (BucketStat, error)
}

type BucketStat struct {
	Capacity uint
	Remaining uint
	Reset time.Time
}

type Storage interface {
	Create(name string, capacity uint, rate time.Duration) (Bucket, error)
}
