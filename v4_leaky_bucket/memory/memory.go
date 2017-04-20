package memory

import (
	"time"
	"sync"
	"github.com/lincolnzhou/rate-limit/v4_leaky_bucket"
)

type bucket struct {
	capacity uint
	remaining uint
	reset time.Time // bucket expire time
	rate time.Duration
	mutex sync.Mutex
}

func (b *bucket) Capacity() uint {
	return b.capacity
}

func (b *bucket) Remaining() uint {
	return b.remaining
}

func (b *bucket) Reset() time.Time {
	return b.reset
}

func (b *bucket) Add(amount uint) (v4_leaky_bucket.BucketStat, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if time.Now().After(b.reset) {
		b.reset = time.Now().Add(b.rate)
		b.remaining = b.capacity
	}

	// over bucket, return err full
	if amount > b.remaining {
		return v4_leaky_bucket.BucketStat{Capacity: b.capacity, Reset: b.reset, Remaining: b.remaining}, v4_leaky_bucket.ErrorFull
	}

	b.remaining = b.remaining - amount
	return v4_leaky_bucket.BucketStat{Capacity: b.capacity, Reset: b.reset, Remaining: b.remaining}, nil
}

type Storage struct {
	buckets map[string]*bucket
}

func (s *Storage) Create(name string, capacity uint, rate time.Duration) (v4_leaky_bucket.Bucket, error) {
	b, ok := s.buckets[name]
	if ok {
		return b, nil
	}

	bucket := &bucket{
		capacity: capacity,
		remaining: capacity,
		reset: time.Now().Add(rate),
		rate: rate,
	}
	s.buckets[name] = bucket

	return bucket, nil
}

func New() *Storage {
	return &Storage{
		buckets: make(map[string]*bucket),
	}
}