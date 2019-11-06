package server

import (
	"2019_2_default_team/db"
	fmdelivery "2019_2_default_team/films/delivery"
	fpostgres "2019_2_default_team/films/repository/Postgres"
	fusecase "2019_2_default_team/films/usecase"
	"2019_2_default_team/handlers"
	"2019_2_default_team/middleware"
	tdelivery "2019_2_default_team/tickets/delivery"
	tpostgres "2019_2_default_team/tickets/repository/Postgres"
	tusecase "2019_2_default_team/tickets/usecase"
	udelivery "2019_2_default_team/users/delivery"
	upostgres "2019_2_default_team/users/repository/Postgres"
	uusecase "2019_2_default_team/users/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	routing *mux.Router
}

func CreateServer() (*Server, error) {
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

func (s *Server) RunServer() {
	http.ListenAndServe(":8080", s.routing)
}
