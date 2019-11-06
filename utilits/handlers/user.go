package handlers

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type HandlerDB struct {
	DB   *sql.DB
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewMyHandler() *MyHandler {
	return &MyHandler{
		sessions: make(map[string]uint64, 0),
		usersAuth: map[string]*User{
			"testUser": {1, "testuser", "test"},
		},
		users: make([]User, 0),
		mu: &sync.Mutex{},
	}
}



func getPhoto(id int) (os.File, error) {
	fileName := strconv.Itoa(id)
	file, err := os.Open("./imagesupload/" + fileName + ".jpg")
	if err != nil {
		log.Printf("An error occurred: %v", err)
		return *file, err
	}
	return *file, nil
}


func Download(file multipart.File, id string) (returnErr error) {
	defer func() {
		err := file.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	tempFile, err := ioutil.TempFile("imagesupload", "upload-*.jpg")
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	defer func() {
		err := tempFile.Close()

		if err != nil && returnErr == nil {
			log.Printf("error: %v", err)
			returnErr = err
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	err = os.Rename(tempFile.Name(), "imagesupload/"+id+".jpg")

	if err != nil {
		log.Printf("An error occurred: %v", err)
		return err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	return nil
}


func (api *MyHandler) UploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1025)
	file, handler, err := r.FormFile("my_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
	session, err := r.Cookie("session_id")
	id := api.sessions[session.Value]
	strid := strconv.Itoa(int(id))
	error := Download(file, strid)
	if error != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (api *MyHandler) GetPhoto(w http.ResponseWriter, r *http.Request){


	authorized := false
	session, err := r.Cookie("session_id")
	if err == nil && session != nil {
		_, authorized = api.sessions[session.Value]
	}

	if authorized {
		id := api.sessions[session.Value]
		file, err := getPhoto(int(id))
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reader := bufio.NewReader(&file)
		bytes := make([]byte, 10<<20)
		_, err = reader.Read(bytes)

		w.Header().Set("content-type", "multipart/form-data;boundary=1")

		_, err = w.Write(bytes)
		if err != nil {
			log.Printf("An error occurred: %v", err)
			w.WriteHeader(500)
			return
		}

		log.Println("Successfully Uploaded File")

	} else {
		w.Write([]byte("not autrorized"))
	}
}
