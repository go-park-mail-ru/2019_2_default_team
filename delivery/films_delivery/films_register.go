package films_delivery

import (
	"net/http"
	"sync"
	"kino_backend/useCase"
)

type MyHandlerFilms struct{
	mu      *sync.Mutex
	useCase useCase.FilmsUseCase
}

func NewMyHandlerFilms(uc useCase.FilmsUseCase) *MyHandlerFilms {
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

