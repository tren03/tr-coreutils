package algorithms

import "testing"

func mockTimeFunc(a int) func() int64 {
	return func() int64 {
		return int64(a)
	}
}

func mockTimeIncrement(a int64) func() int64 {
	return func() int64 {
		a = a + 1
		return a
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

func TestTokenBucketFlow(t *testing.T) {
	inMemoryTokenBucketRepo := NewInMemoryTokenBucketRepo()
	timeFunc := mockTimeIncrement(0)
	tokenBucket := NewTokenBucket(2, 1, inMemoryTokenBucketRepo, timeFunc)
	one := tokenBucket.Execute("testuser")
	two := tokenBucket.Execute("testuser")

	if one.Allow != true {
		t.Errorf("got Allow %v; want false", one)
	}
	if two.Allow != false {
		t.Errorf("got Allow %v; want false", two)
	}

}
