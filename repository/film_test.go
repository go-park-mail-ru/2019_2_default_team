package repository

import(
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
	"kino_backend/models"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID uint = 1
	userp := models.UserPassword{"emailtest@mail.ru", "pass"}
	user := models.User{UserID:elemID, UserPassword: userp}
	text := ""
	ava := &text

	rows := sqlmock.
		NewRows([]string{"user_id", "email", "password", "nickname", "avatar"})
	expect := []models.Profile{
		{
			user,
			"nicktest",
			ava,
		},
	}
	for _, item := range expect {
		rows = rows.AddRow(item.UserID, item.Email, item.Password,
			item.Nickname, item.Avatar)
	}

	//repo := NewFilmRepository(db)

	mock.
		ExpectQuery("SELECT user_id email password nickname avatar FROM user_profile WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	item, err := GetFilmProfileByID(elemID)

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