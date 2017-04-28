package memory

import (
	"testing"
	"github.com/lincolnzhou/rate-limit/v4_leaky_bucket/test"
)

func TestCreate(t *testing.T) {
	test.CreateTest(New())(t)
}

func TestAdd(t *testing.T) {
	test.AddTest(New())(t)
}

func TestAddReset(t *testing.T) {
	test.AddResetTest(New())(t)
}