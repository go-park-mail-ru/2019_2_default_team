package users_delivery

//
//func TestPostSignupProfile (t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userCRUD := umocks.NewMockUseCase(ctrl)
//	var userJSON = `{"id":6, "email" : "mye111m@mail.ru", "password":"password", "nickname" : "username11"}`
//	user := models.UserPassword{Email: "mye111m@mail.ru",
//		Password: "password",
//	}
//	u := &models.RegisterProfile{
//		Nickname:"username11",
//		UserPassword:user,
//	}
//
//	outuser := models.User{
//		UserID:       1,
//		UserPassword: user,
//	}
//
//	//var ctx context.Context
//	pfout := models.Profile{
//		User: outuser,
//		Nickname: "username11",
//	}
//
//	l := logger.InitLogger()
//	defer l.Sync()
//
//	rec := httptest.NewRecorder()
//	req := httptest.NewRequest(http.MethodPost, "/profile", strings.NewReader(userJSON))
//	//e := echo.New()
//	//c := e.NewContext(req, rec)
//	c := req.Context()
//
//	userCRUD.EXPECT().PostUser(c, u).Return(pfout, nil).Times(1)
//
//	handler := &Handler{useCase:userCRUD}
//
//	handler.postSignupProfile(rec, req)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//}

//func TestGetProfile (t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userCRUD := umocks.NewMockUseCase(ctrl)
//	var userJSON = `{"reqtitle" : "Joker"}`
//	u := &models.RequestProfileFilm{
//		Title:"Joker",
//	}
//	//var ctx context.Context
//	film := models.Film{FilmID:1}
//	pfout := models.ProfileFilm{
//		Film: film,
//		Title:"Joker",
//		Description:"Absolutely madness",
//		Director:"Todd",
//		MainActor:"Phoenix",
//		AdminID: 1,
//	}
//
//	l := logger.InitLogger()
//	defer l.Sync()
//
//	rec := httptest.NewRecorder()
//	req := httptest.NewRequest(http.MethodGet, "/film", strings.NewReader(userJSON))
//
//	c := req.Context()
//
//	userCRUD.EXPECT().GetFilm(c, u).Return(pfout, nil).Times(1)
//
//	handler := &Handler{useCase:userCRUD}
//
//	handler.getProfileFilm(rec, req)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//}
//
