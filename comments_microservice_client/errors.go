package comments_microservice_client

import (
	"errors"
)

var (
	ErrConnRefused     = errors.New("no session in database")
	ErrCommentNotFound = errors.New("comment not found")
)

//Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty" db:"username"`
//FilmTitle            string   `protobuf:"bytes,2,opt,name=FilmTitle,proto3" json:"FilmTitle,omitempty" db:"film_title"`
//Text                 string   `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty" db:"comment"`
//ID                   uint64   `protobuf:"varint,4,opt,name=ID,proto3" json:"ID" db:"id"`
