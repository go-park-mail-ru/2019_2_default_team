package server

import (
	"github.com/gorilla/mux"
	fmdelivery "kino_backend/films/delivery"
	fpostgres "kino_backend/films/repository/Postgres"
	fusecase "kino_backend/films/usecase"
	tpostgres "kino_backend/tickets/repository/Postgres"
	upostgres "kino_backend/users/repository/Postgres"
	tdelivery "kino_backend/tickets/delivery"
	"kino_backend/handlers"
	"kino_backend/middleware"
	udelivery "kino_backend/users/delivery"
	"log"
	"net/http"
	"os"
	"kino_backend/db"
	tusecase "kino_backend/tickets/usecase"
	uusecase "kino_backend/users/usecase"

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

	fuc := fusecase.NewFilmUseCase(fpostgres.NewFilmRepository(db.Db))
	tuc := tusecase.NewTicketUseCase(tpostgres.NewTicketRepository(db.Db))
	uuc := uusecase.NewUserUseCase(upostgres.NewUserRepository(db.Db))
	apif := fmdelivery.NewMyHandlerFilms(fuc)
	apiu := udelivery.NewMyHandlerUser(uuc)
	apit := tdelivery.NewMyHandlerTicket(tuc)


	api := handlers.NewMyHandler()

	r.HandleFunc("/profile", apiu.ProfileHandler)

	r.HandleFunc("/logging", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(api.SessionHandler))))

	r.HandleFunc("/profile", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apiu.ProfileHandler))))
	r.HandleFunc("/film", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.ProfileFilmHandler))))
	r.HandleFunc("/ticket", middleware.RecoverMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apit.ProfileTicketHandler))))
	//r.HandleFunc("/logging", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.SessionHandler))))
	//r.HandleFunc("/profile", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileHandler))))
	//r.HandleFunc("/film", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileFilmHandler))))
	//r.HandleFunc("/ticket", middleware.RecoverMiddleware(middleware.CorsMiddleware(
	//	middleware.SessionMiddleware(api.ProfileTicketHandler))))




	r.HandleFunc("/photodownload", api.GetPhoto)
	r.HandleFunc("/upload", api.UploadPage)

	err = nil
	server.routing = r

	return server, err
}

func (s *Server)RunServer() {
	http.ListenAndServe(":8080", s.routing)
}
