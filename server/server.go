package server

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"kino_backend/comments_microservice_client"
	"kino_backend/delivery/chat_support_delivery"
	"kino_backend/delivery/comments_service_delivery"
	"kino_backend/delivery/films_delivery"
	"kino_backend/delivery/sessions_delivery"
	"kino_backend/delivery/sessions_service_delivery"
	"kino_backend/delivery/tickets_delivery"
	"kino_backend/delivery/users_delivery"
	"kino_backend/metrics"
	"kino_backend/repository"
	"kino_backend/session_microservice_client"
	"kino_backend/sessions"
	"kino_backend/useCase"
	"kino_backend/utilits/middleware"
	"log"
	"net/http"
	"os"
)

type Server struct {
	routing *mux.Router
}

func CreateServer(database *sqlx.DB, Sesredis *sessions.SessionManager) (*Server, error) {
	server := new(Server)

	//l := logger.InitLogger()
	//defer l.Sync()

	authConnStr := flag.String("auth_connstr", "localhost:8081", "auth-service connection string")
	commentConnStr := flag.String("comment_connstr", "localhost:8082", "comment-service connection string")
	flag.Parse()

	var err error
	r := mux.NewRouter()
	Access := new(middleware.AccessLogger)
	Access.StdLogger = log.New(os.Stdout, "STD ", log.LUTC|log.Lshortfile)
	chat := useCase.InitChat()
	sesmic := session_microservice_client.ConnectSessionManager(*authConnStr)
	commic := comments_microservice_client.ConnectCommentsManager(*commentConnStr)

	fuc := useCase.NewFilmUseCase(repository.NewFilmRepository(database))
	tuc := useCase.NewTicketUseCase(repository.NewTicketRepository(database))
	uuc := useCase.NewUserUseCase(repository.NewUserRepository(database))
	suc := useCase.NewSessionUseCase(repository.NewSessionsRepository(Sesredis.RedisConn))
	scuc := useCase.NewSupportChatsUseCase(repository.NewSupportChatRepository(database), chat)
	scuc.Run()

	apif := films_delivery.NewMyHandlerFilms(fuc)
	apit := tickets_delivery.NewMyHandlerTicket(tuc)
	apis := sessions_delivery.NewMyHandlerFilms(suc, uuc)
	apiu := users_delivery.NewMyHandlerUser(uuc, suc)
	apisc := chat_support_delivery.NewMyHandlerCS(scuc)
	apism := sessions_service_delivery.NewMyHandlerSessions(sesmic, uuc)
	apicm := comments_service_delivery.NewMyHandlerFilms(commic)
	prometheus.MustRegister(metrics.AccessHits)

	r = r.PathPrefix("/api/").Subrouter()

	r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/profile", apiu.ProfileHandler)

	r.HandleFunc("/profile", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apiu.ProfileHandler, sesmic)))))
	r.HandleFunc("/films", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.ProfileFilmHandler, sesmic)))))
	r.HandleFunc("/film/{id}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.ProfileOneFilm, sesmic)))))
	r.HandleFunc("/allfilms", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.ProfileAllFilms, sesmic)))))
	r.HandleFunc("/allfilms/today", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.FilmsForToday, sesmic)))))
	r.HandleFunc("/allfilms/soon", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.FilmsForSoon, sesmic)))))
	r.HandleFunc("/ticket", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apit.ProfileTicketHandler, sesmic)))))
	r.HandleFunc("/session", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apis.ProfileSessionsHandler, sesmic)))))
	r.HandleFunc("/authorized", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apis.ProfileAuth, sesmic)))))
	r.HandleFunc("/support", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apisc.SupportChat, sesmic)))))
	r.HandleFunc("/sessionservice", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apism.ProfileSessionsMicroserviceHandler, sesmic)))))
	r.HandleFunc("/commentservice", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apicm.CommentsHandler, sesmic)))))
	r.HandleFunc("/commentByID/{id}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apicm.CommentsByIDHandler, sesmic)))))
	r.HandleFunc("/commentByFilm/{film}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apicm.CommentsByFilmHandler, sesmic)))))
	r.HandleFunc("/commentByUsername/{username}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apicm.CommentsByUsernameHandler, sesmic)))))
	r.HandleFunc("/createmoviesession", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.MovieSession, sesmic)))))
	r.HandleFunc("/get_movie_sessions_times_for_today/{movie_id}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.GetTimesMovieSessionsForToday, sesmic)))))
	r.HandleFunc("/get_seats/{ms_id}", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.GetSeatsByMSID, sesmic)))))
	r.HandleFunc("/film_vote", middleware.RecoverMiddleware(metrics.CountHitsMiddleware(middleware.CorsMiddleware(
		middleware.SessionMiddleware(apif.PostVote, sesmic)))))

	err = nil
	server.routing = r

	return server, err
}

func (s *Server) RunServer() {
	http.ListenAndServe(":8080", s.routing)
}
