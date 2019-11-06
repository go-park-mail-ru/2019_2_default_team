package delivery

import (
	"kino_backend/films"
	"net/http"
	"sync"
)

type MyHandlerFilms struct{
	mu      *sync.Mutex
	useCase films.UseCase
}

func NewMyHandlerFilms(uc films.UseCase) *MyHandlerFilms {
	return &MyHandlerFilms{
		mu: &sync.Mutex{},
		useCase: uc,
	}
}

func (apif *MyHandlerFilms) ProfileFilmHandler(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	switch r.Method {
	case http.MethodGet:
		h.getProfileFilm(w, r)
	case http.MethodPost:
		h.postSignupProfileFilm(w, r)
	case http.MethodPut:
		h.putEditFilmProfile(w, r)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

