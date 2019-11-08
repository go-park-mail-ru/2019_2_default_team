package films_delivery

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"kino_backend/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"kino_backend/models"
	"kino_backend/logger"
)

func TestPostSignupProfileFilm (t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userCRUD := repository.NewMockFilmsRepository(ctrl)
	var userJSON = `{"title" : "ToStars", "description":"space", "director" : "someone", "mainactor" : "Pitt", "admin_id" : 1}`
	u := &models.RegisterProfileFilm{
		Title:"ToStars",
		Description:"space",
		Director:"someone",
		MainActor:"Pitt",
		AdminID: 1,
	}
	//var ctx context.Context
	film := models.Film{FilmID:1}
	pfout := models.ProfileFilm{
		Film: film,
		Title:"ToStars",
		Description:"space",
		Director:"someone",
		MainActor:"Pitt",
		AdminID: 1,
	}

	l := logger.InitLogger()
	defer l.Sync()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/film", strings.NewReader(userJSON))
	//e := echo.New()
	//c := e.NewContext(req, rec)
	c := req.Context()

	userCRUD.EXPECT().CreateNewFilm(c, pfout)

	handler := &Handler{useCase: userCRUD}

	handler.postSignupProfileFilm(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestGetProfileFilm (t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userCRUD := fmocks.NewMockUseCase(ctrl)
	var userJSON = `{"reqtitle" : "Joker"}`
	u := &models.RequestProfileFilm{
		Title:"Joker",
	}
	//var ctx context.Context
	film := models.Film{FilmID:1}
	pfout := models.ProfileFilm{
		Film: film,
		Title:"Joker",
		Description:"Absolutely madness",
		Director:"Todd",
		MainActor:"Phoenix",
		AdminID: 1,
	}

	l := logger.InitLogger()
	defer l.Sync()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/film", strings.NewReader(userJSON))

	c := req.Context()

	userCRUD.EXPECT().GetFilm(c, u).Return(pfout, nil).Times(1)

	handler := &Handler{useCase: userCRUD}

	handler.getProfileFilm(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

}


var usersApi UserHandlers