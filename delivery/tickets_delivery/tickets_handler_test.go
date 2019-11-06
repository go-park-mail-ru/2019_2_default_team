package tickets_delivery

//
//func TestPostBuyTicket (t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userCRUD := tmocks.NewMockUseCase(ctrl)
//	var ticketJSON = `{"user_d" : 1, "film_id":1}`
//	u := &models.RegisterTicket{
//		UserID:1,
//		FilmID: 1,
//	}
//	//var ctx context.Context
//	tick := models.RegisterTicket{UserID:1, FilmID: 1}
//	pfout := models.Ticket{
//		TicketID: 1,
//		RegisterTicket: tick,
//
//	}
//
//	l := logger.InitLogger()
//	defer l.Sync()
//
//	rec := httptest.NewRecorder()
//	req := httptest.NewRequest(http.MethodPost, "/ticket", strings.NewReader(ticketJSON))
//	//e := echo.New()
//	//c := e.NewContext(req, rec)
//	c := req.Context()
//
//	userCRUD.EXPECT().PostTicket(c, u).Return(pfout, nil).Times(1)
//
//	handler := &Handler{useCase:userCRUD}
//
//	handler.postBuyTicket(rec, req)
//	assert.Equal(t, http.StatusOK, rec.Code)
//
//}
//
//func TestGetProfileTicket (t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	userCRUD := fmocks.NewMockUseCase(ctrl)
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
