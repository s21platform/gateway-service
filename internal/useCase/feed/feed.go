package feed

type Usecase struct {
	feC FeedClient
}

func New(feS FeedClient) *Usecase {
	return &Usecase{feC: feS}
}

// func (u *Usecase) CreateUserPost(r *http.Request) (*feed.CreateUserPostOut, error) {
// 	content := r.URL.Query().Get("content")

// 	resp, err := u.feC.CreateUserPost(r.Context(), content)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create post in usecase: %v", err)
// 	}
// 	return resp, nil
// }
