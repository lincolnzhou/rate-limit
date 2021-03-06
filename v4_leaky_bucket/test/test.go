package test

import (
	"github.com/lincolnzhou/rate-limit/v4_leaky_bucket"
	"math"
	"testing"
	"time"
	"sync"
)

func CreateTest(s v4_leaky_bucket.Storage) func(*testing.T) {
	return func(t *testing.T) {
		now := time.Now()
		bucket, err := s.Create("test-bucket", 100, time.Minute)
		if err != nil {
			t.Fatal(err)
		}
		if capacity := bucket.Capacity(); capacity != 100 {
			t.Fatalf("expected capacity of %d, got %d", 100, capacity)
		}
		e := float64(1 * time.Second)
		if error := float64(bucket.Reset().Sub(now.Add(time.Minute))); math.Abs(error) > e {
			t.Fatalf("expected reset time close to %s, got %s", now.Add(time.Minute), bucket.Reset())
		}
	}
}

func AddTest(s v4_leaky_bucket.Storage) func(*testing.T) {
	return func(t *testing.T) {
		bucket, err := s.Create("test-bucket", 10, time.Minute)
		if err != nil {
			t.Fatal(err)
		}

		addAndTestRemaining := func(add, remaining uint) {
			if state, err := bucket.Add(add); err != nil {
				t.Fatal(err)
			} else if bucket.Remaining() != state.Remaining {
				t.Fatalf("epected bucket and state remaining to match, bucket is %d, state is %d", bucket.Remaining(), state.Remaining)
			} else if bucket.Remaining() != remaining {
				t.Fatalf("expected %d remaining, got %d", remaining, bucket.Remaining())
			}
		}

		addAndTestRemaining(1, 9)
		addAndTestRemaining(3, 6)
		addAndTestRemaining(6, 0)

		if _, err := bucket.Add(1); err == nil {
			t.Fatalf("expected ErrorFull, received no error")
		} else if err != v4_leaky_bucket.ErrorFull {
			t.Fatalf("expected ErrorFull, received %v", err)
		}
	}
}

func AddResetTest(s v4_leaky_bucket.Storage) func(*testing.T) {
	return func(t *testing.T) {
		bucket, err := s.Create("test-bucket", 1, time.Millisecond)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := bucket.Add(1); err != nil {
			t.Fatal(err)
		}
		time.Sleep(2*time.Millisecond)
		if state, err := bucket.Add(1); err != nil {
			t.Fatal(err)
		} else if (state.Remaining != 0) {
			t.Fatalf("expected full bucket, got %d", state.Remaining)
		} else if (state.Reset.Unix() < time.Now().Unix()) {
			t.Fatal("reset time is in the past")
		}
	}
}

func ThreadSafeAddTest(s v4_leaky_bucket.Storage) func(*testing.T) {
	return func(t *testing.T) {
		n := 100
		k := 50

		bucket, err := s.Create("test-bucket", uint(n), time.Minute)
		if err != nil {
			t.Fatal(err)
		}

		var wgErrors sync.WaitGroup
		errors := make(chan error)
		wgErrors.Add(1)
		go func() {
			defer wgErrors.Done()
			count := 0
			for err := range errors {
				count++
				if err != v4_leaky_bucket.ErrorFull {
					t.Fatalf("got an error that is not ErrorFull: %s", err)
				}
			}

			if count != k {
				t.Fatalf("got %d errors, expected %d", count, k)
			}
		}()



		wgErrors.Wait()
		close(errors)
	}
}
