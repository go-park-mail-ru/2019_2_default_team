package repository

import(
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
	"kino_backend/models"
)

func TestGetFilm(t *testing.T) {
	db, mock, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID uint = 1
	film := models.Film{elemID}
	foto := ""
	ava := &foto
	rows := sqlmock.
		NewRows([]string{"film_id", "title", "description", "avatar", "director", "mainactor", "admin_id"})
	expect := []models.ProfileFilm{
		{
			film,
			"Joker",
			"testd",
			ava,
			"testd",
			"actor",
			1,
		},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.FilmID, item.Title, item.Description, item.Avatar, item.Director, item.MainActor, item.AdminID)
	}

	filmRepo := NewFilmRepository(sqlxDB)

	mock.
		ExpectQuery("SELECT film_id, title, description, avatar, director, mainactor, admin_id FROM film_profile WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := filmRepo.GetFilmProfileByIDSQL(elemID)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}
}