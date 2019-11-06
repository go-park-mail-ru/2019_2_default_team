package Postgres

import(
	"context"
	"github.com/jmoiron/sqlx"
	"kino_backend/models"
	"kino_backend/db"
	"kino_backend/users"
)


type UserRepository struct{
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
	return &UserRepository{
		database:db,
	}
}

func (user UserRepository) GetUser(ctx context.Context, params *models.RequestProfile, auth bool, id uint) (models.Profile, error) {

	if params.ID != 0 {
		profile, err := db.GetUserProfileByID(params.ID)
		if err != nil {

			return models.Profile{}, err
			//switch err.(type) {
			//case db.UserNotFoundError:
			//	w.WriteHeader(http.StatusNotFound)
			//	return
			//default:
			//	log.Println(err)
			//	w.WriteHeader(http.StatusInternalServerError)
			//	return
			//}
		}

		return profile, nil

		//w.Header().Set("Content-Type", "application/json")
		//json, err := json.Marshal(profile)
		//if err != nil {
		//	log.Println(err, "in profileMethod")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//fmt.Fprintln(w, string(json))
	} else if params.Nickname != "" {
		profile, err := db.GetUserProfileByNickname(params.Nickname)
		if err != nil {

			return models.Profile{}, err
			//switch err.(type) {
			//case db.UserNotFoundError:
			//	w.WriteHeader(http.StatusNotFound)
			//	return
			//default:
			//	log.Println(err)
			//	w.WriteHeader(http.StatusInternalServerError)
			//	return
			//}
		}

		return profile, nil
		//
		//w.Header().Set("Content-Type", "application/json")
		//json, err := json.Marshal(profile)
		//if err != nil {
		//	log.Println(err, "in profileMethod")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//fmt.Fprintln(w, string(json))
	} else {
		if !auth {
			return models.Profile{}, users.UserNotAuthError{}
		}

		profile, err := db.GetUserProfileByID(id)
		if err != nil {

			return models.Profile{}, err
			//switch err.(type) {
			//case db.UserNotFoundError:
			//	w.WriteHeader(http.StatusNotFound)
			//	return
			//default:
			//	log.Println(err)
			//	w.WriteHeader(http.StatusInternalServerError)
			//	return
			//}
		}
		return profile, nil
	}
}

//func (user UserRepository) GetUserAuth(id uint) (models.Profile, error) {
//	profile, err := db.GetUserProfileByID(id)
//	if err != nil {
//
//		return models.Profile{}, err
//		//switch err.(type) {
//		//case db.UserNotFoundError:
//		//	w.WriteHeader(http.StatusNotFound)
//		//	return
//		//default:
//		//	log.Println(err)
//		//	w.WriteHeader(http.StatusInternalServerError)
//		//	return
//		//}
//	}
//
//	return profile, nil
//}

func (user UserRepository) PostUser(ctx context.Context, u *models.RegisterProfile) (models.Profile, error){
	newU, err := db.CreateNewUser(u)
	if err != nil {
		return models.Profile{}, err
		//if err == db.ErrUniqueConstraintViolation ||
		//	err == db.ErrNotNullConstraintViolation {
		//	w.WriteHeader(http.StatusUnprocessableEntity)
		//	return
		//}
		//fmt.Println("4")
		//logger.Error(err)
		//w.WriteHeader(http.StatusInternalServerError)
		//return
	}

	return newU, nil

}

func (user UserRepository) PutUser(ctx context.Context, id uint, editUser *models.RegisterProfile) (error){

	err := db.UpdateUserByID(id, editUser)

	if err != nil{
		return err
		//switch err.(type) {
		//case db.UserNotFoundError:
		//	w.WriteHeader(http.StatusNotFound)
		//default:
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//}
		//return
	}

	return nil
}

