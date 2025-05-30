package feed

type Usecase struct {
	feC FeedClient
}

func New(feS FeedClient) *Usecase {
	return &Usecase{feC: feS}
}
