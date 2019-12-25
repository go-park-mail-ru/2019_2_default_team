package films_delivery

import (
	"kino_backend/useCase"
	"net/http"
	"sync"
)

type MyHandlerFilms struct {
	mu      *sync.Mutex
	useCase useCase.FilmsUseCase
}

func NewMyHandlerFilms(uc useCase.FilmsUseCase) *MyHandlerFilms {
	return &MyHandlerFilms{
		mu:      &sync.Mutex{},
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

func (apif *MyHandlerFilms) ProfileOneFilm(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getOneFilm(w, r)
}

func (apif *MyHandlerFilms) ProfileAllFilms(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getAllFilms(w, r)
}

func (apif *MyHandlerFilms) MovieSession(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.postCreateMovieSession(w, r)
}

func (apif *MyHandlerFilms) GetTimesMovieSessionsForToday(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getTimesForToday(w, r)
}

func (apif *MyHandlerFilms) GetSeatsByMSID(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getSeatsByMSID(w, r)
}

func (apif *MyHandlerFilms) PostVote(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.postVote(w, r)
}

func (apif *MyHandlerFilms) FilmsForToday(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getFilmsForToday(w, r)
}

func (apif *MyHandlerFilms) FilmsForSoon(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getFilmsForSoon(w, r)
}

func (apif *MyHandlerFilms) FilmsRecommended(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getRecommendedFilms(w, r)
}

func (apif *MyHandlerFilms) TopFilms(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(apif.useCase)

	h.getTopFilms(w, r)
}
