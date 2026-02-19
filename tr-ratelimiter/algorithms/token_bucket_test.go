package algorithms

import "testing"

func mockTimeFunc(a int) func() int64 {
	return func() int64 {
		return int64(a)
	}
}

func TestTokenBucket(t *testing.T) {
	inMemoryTokenBucketRepo := NewInMemoryTokenBucketRepo()
	tokenBucket := NewTokenBucket(5, 5, inMemoryTokenBucketRepo, mockTimeFunc(0))
	got := tokenBucket.Execute("testuser")
	if got.Allow != true {
		t.Errorf("got Allow %t; want false", got.Allow)
	}
}
