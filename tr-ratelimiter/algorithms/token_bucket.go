package algorithms

import "github.com/tren03/tr-coreutils/tr-ratelimiter/utils"

/*
Token bucket would take
- refillRate // per minute
- bucketSize
*/

type TokenBucketState struct {
	timestamp int64
	curTokens int
}

// This will be able to handle in-memory and external deps
type ITokenBucketRepo interface {
	Get(key string) (state *TokenBucketState, ok bool)
	Set(key string, state TokenBucketState) // we could return a err here
}

// --- IN MEMORY REPO --- NOTE : NOT SAFE FOR CONCURRENT USE
type InMemoryTokenBucketRepo struct {
	stateMap map[string]TokenBucketState
}

func NewInMemoryTokenBucketRepo() *InMemoryTokenBucketRepo {
	return &InMemoryTokenBucketRepo{
		stateMap: make(map[string]TokenBucketState),
	}
}

func (i *InMemoryTokenBucketRepo) Get(key string) (*TokenBucketState, bool) {
	if val, ok := i.stateMap[key]; ok {
		return &val, ok
	}
	return nil, false
}

func (i *InMemoryTokenBucketRepo) Set(key string, state TokenBucketState) {
	i.stateMap[key] = state
}

type TokenBucket struct {
	RefillRate int // per minute
	BucketSize int
	Repo       ITokenBucketRepo
	TimeNow    func() int64
}

func NewTokenBucket(RefillRate int, BucketSize int, Repo ITokenBucketRepo, TimeNow func ()int64) *TokenBucket {
	return &TokenBucket{
		RefillRate: RefillRate,
		BucketSize: BucketSize,
		Repo:       Repo,
		TimeNow:    TimeNow,
	}

}

func (t *TokenBucket) Execute(key string) AlgorithmResponse {
	// fetch state of key
	state, ok := t.Repo.Get(key)

	// if key not exist
	if !ok {
		remainingTokens := t.BucketSize - 1
		t.Repo.Set(key, TokenBucketState{timestamp: t.TimeNow(), curTokens: remainingTokens})
		return AlgorithmResponse{
			Allow:     true,
			Remaining: remainingTokens,
			// ex - refill rate = 5/min, bucket size = 5
			// so if first req comes in, then remaining = 5-1 - 4 tokens
		}
	}

	// key exists
	// calculate diff betweeen last time stamp and cur
	curTime := t.TimeNow()
	diff := curTime - state.timestamp
	tokensAdded := (int64(t.RefillRate) * diff) / 60                                // since we calc per min
	curTokenCount := min(int64(t.BucketSize), int64(state.curTokens)+tokensAdded-1) // remove one for current request, and calc tokens (cap to max bucket size)
	if curTokenCount <= 0 {
		// do not allow
		// t.ITokenBucketRepo.Set(key, TokenBucketState{timestamp: t.timeNow(), curTokens: 0}) DO NOT UPDATE TIMESTAMP WHEN FALSE, TO NOT PUNISH USER
		retryAfter := int64(60 / t.RefillRate) // time taken for 1 token to be generated
		return AlgorithmResponse{
			Allow:      false,
			RetryAfter: &retryAfter,
			Remaining:  0,
		}
	} else {
		//allow
		t.Repo.Set(key, TokenBucketState{timestamp: t.TimeNow(), curTokens: int(curTokenCount)})
		return AlgorithmResponse{
			Allow:     true,
			Remaining: int(curTokenCount),
		}
	}

}
