package algorithms

type AlgorithmResponse struct {
	Allow      bool
	RetryAfter *int64
	Remaining  int
}

type IAlgorithm interface {
	Execute(key string) AlgorithmResponse
}
