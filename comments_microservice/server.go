package comments_microservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"kino_backend/comments_microservice_client"
	"kino_backend/db"
	"kino_backend/logger"
)

//rpc CreateComment (Comment) returns (Nothing) {}
//rpc GetComment (CommentID) returns (Comment) {}
//rpc GetCommentsByFilmID (CommentID) returns (CommentsResponse) {}
//rpc GetCommentsByUserID (CommentID) returns (CommentsResponse) {}
//rpc DeleteComment (CommentID) returns (Nothing) {}

var Cm *CommentManager
var Db *sqlx.DB

var (
	ErrInser   = errors.New("error while inserting")
	ErrMarshal = errors.New("error while marshalling")
)

type CommentManager struct {
	PostgresConn sqlx.DB
}

func (cm *CommentManager) Close() {
	cm.PostgresConn.Close()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "postgres"
)

func InitDB(address, database string) *CommentManager {
	var err error

	//dsn := "localhost:5432/some-postgres?"
	//Db, err = sqlx.Open("postgres",
	//	"postgres://"+address+"/"+database+"?sslmode=disable")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Db, err = sqlx.Open("postgres", psqlInfo)

	if err != nil {
		logger.Panic(err)
	}

	if err := Db.Ping(); err != nil {
		logger.Panic(err)
	}

	//Db, err := sqlx.Connect("postgres", "user=foo dbname=some-bar sslmode=disable")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	db.MakeMigrations(Db)

	logger.Infof("Successfully connected to %v, database %v", address, database)
	CM := CommentManager{PostgresConn: *Db}

	return &CM
}

func (cm *CommentManager) CreateComment(ctx context.Context, in *comments_microservice_client.Comment) (*comments_microservice_client.Nothing, error) {
	qres := cm.PostgresConn.QueryRowx(`
		INSERT INTO comments (film_title, username, comment)
		VALUES ($1, $2, $3)`, in.FilmTitle, in.Username, in.Text)
	if err := qres.Err(); err != nil {
		fmt.Println("error in")
		return &comments_microservice_client.Nothing{}, ErrInser
	}
	return &comments_microservice_client.Nothing{}, nil
}

func (cm *CommentManager) GetComment(ctx context.Context, in *comments_microservice_client.CommentID) (*comments_microservice_client.Comment, error) {

	res := comments_microservice_client.Comment{}
	qres := cm.PostgresConn.QueryRowx(`
		SELECT id, film_title, username, comment FROM comments
		WHERE id = $1`,
		in.CID)
	if err := qres.Err(); err != nil {
		return &res, err
	}
	err := qres.StructScan(&res)
	fmt.Println(res)
	if err != nil {
		if err == sql.ErrNoRows {
			return &res, ErrMarshal
		}
		return &res, err
	}

	return &res, nil
}

func (cm *CommentManager) GetCommentsByFilmID(ctx context.Context, in *comments_microservice_client.CommentID) (*comments_microservice_client.CommentsResponse, error) {
	res := comments_microservice_client.CommentsResponse{}
	resComments := make([]comments_microservice_client.Comment, 0)
	resOne := comments_microservice_client.Comment{}
	qres, err := cm.PostgresConn.Queryx(`
		SELECT id, film_title, username, comment FROM comments
		WHERE film_title = $1`,
		in.Film)

	if err != nil {
		return &res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)
		resComments = append(resComments, resOne)
	}

	for i, _ := range resComments {
		res.Comments = append(res.Comments, &resComments[i])
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return &res, err
		}
		return &res, err
	}

	return &res, nil

}

func (cm *CommentManager) GetCommentsByUserID(ctx context.Context, in *comments_microservice_client.CommentID) (*comments_microservice_client.CommentsResponse, error) {
	res := comments_microservice_client.CommentsResponse{}
	resOne := comments_microservice_client.Comment{}
	qres, err := cm.PostgresConn.Queryx(`
		SELECT id, film_title, username, comment FROM comments
		WHERE username = $1`,
		in.User)

	if err != nil {
		return &res, err
	}

	for qres.Next() {
		err = qres.StructScan(&resOne)

	}

	if err != nil {
		if err == sql.ErrNoRows {
			return &res, err
		}
		return &res, err
	}

	return &res, nil
}

func (cm *CommentManager) DeleteComment(ctx context.Context, in *comments_microservice_client.CommentID) (*comments_microservice_client.Nothing, error) {

	return &comments_microservice_client.Nothing{}, nil
}
