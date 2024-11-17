package search

import "net/http"

type Usecase struct {
	sC SearchClient
}

func New(sC SearchClient) *Usecase {
	return &Usecase{sC: sC}
}

func (u *Usecase) Search(r *http.Request) {

}
