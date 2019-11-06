package handlers

import "sync"

type UserInput struct {
	Username     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64   `json:"id"`
	Username string    `json:"name"`
	Password string    `json:"-"`
}

type UserEdit struct{
	PrevUsername     string `json:"prevname"`
	PrevPassword string `json:"prevpassword"`
	NewUsername     string `json:"newname"`
	NewPassword string `json:"newpassword"`
}

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type MyHandler struct {
	sessions map[string]uint64
	usersAuth    map[string]*User
	users []User
	mu  *sync.Mutex
}

const ImageUploadPath = "./imagesupload"

var uploadFormTmpl = []byte(`
<html>
	<body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		Image: <input type="file" name="my_file">
		<input type="submit" value="Upload">
	</form>
	</body>
</html>
`)
