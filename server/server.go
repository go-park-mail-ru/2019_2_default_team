package server

import (
	"github.com/gorilla/mux"
	"kino_backend/db"
	"kino_backend/delivery/films_delivery"
	"kino_backend/delivery/sessions_delivery"
	"kino_backend/delivery/tickets_delivery"
	"kino_backend/delivery/users_delivery"
	fpostgres "kino_backend/repository/films_repository/Postgres"
	"kino_backend/repository/session_repository/Redis"
	tpostgres "kino_backend/repository/tickets_repository/Postgres"
	upostgres "kino_backend/repository/users_repository/Postgres"
	"kino_backend/useCase/films_usecase"
	"kino_backend/useCase/sessions_usecase"
	"kino_backend/useCase/tickets_usecase"
	"kino_backend/useCase/users_usecase"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
	"os"
)

type Server struct{
	routing *mux.Router
}

func CreateServer() (*Server ,error){
	server := new(Server)

	//l := logger.InitLogger()
	//defer l.Sync()

	var err error
	r := mux.NewRouter()
	Access := new(middleware.AccessLogger)
	Access.StdLogger = log.New(os.Stdout, "STD ", log.LUTC|log.Lshortfile)

	fuc := films_usecase.NewFilmUseCase(fpostgres.NewFilmRepository(db.Db))
	tuc := tickets_usecase.NewTicketUseCase(tpostgres.NewTicketRepository(db.Db))
	uuc := users_usecase.NewUserUseCase(upostgres.NewUserRepository(db.Db))
	suc := sessions_usecase.NewSessionUseCase(Redis.NewSessionsRepository(Redis.Rd))

	apif := films_delivery.NewMyHandlerFilms(fuc)
	apiu := users_delivery.NewMyHandlerUser(uuc)
	apit := tickets_delivery.NewMyHandlerTicket(tuc)
	apis := sessions_delivery.NewMyHandlerFilms(suc)



	r.HandleFunc("/profile", apiu.ProfileHandler)


	r.HandleFunc("/profile", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apiu.ProfileHandler))))
	r.HandleFunc("/film", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.ProfileFilmHandler))))
	r.HandleFunc("/ticket", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apit.ProfileTicketHandler))))
	r.HandleFunc("/session", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apis.ProfileSessionsHandler))))
	//r.HandleFunc("/logging", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.SessionHandler))))
	//r.HandleFunc("/profile", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileHandler))))
	//r.HandleFunc("/film", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileFilmHandler))))
	//r.HandleFunc("/ticket", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileTicketHandler))))



	err = nil
	server.routing = r

	return server, err
}

func (s *Server)RunServer() {
	http.ListenAndServe(":8080", s.routing)
}
