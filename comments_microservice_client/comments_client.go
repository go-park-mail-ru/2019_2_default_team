package comments_microservice_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"kino_backend/logger"
	"time"
)

// nolint: golint
type CommentsManager struct {
	cmc      CommentsManagerClient
	grpcConn *grpc.ClientConn
}

func ConnectCommentsManager(address string) *CommentsManager {
	grpcConn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(30*time.Second), // nolint: megacheck
	)
	if err != nil {
		logger.Panicf("failed to connect to sessionManager at address %v: %v", address, err)
	}

	cmc := NewCommentsManagerClient(grpcConn)

	logger.Infof("Successfully connected to sessionManager: %v", address)

	return &CommentsManager{cmc: cmc, grpcConn: grpcConn}
}

func (cm *CommentsManager) GetComment(cID CommentID) (Comment, error) {
	if cm.grpcConn == nil {
		return Comment{}, ErrConnRefused
	}

	comment, err := cm.cmc.GetComment(
		context.Background(),
		&CommentID{CID: uint64(cID.CID)},
	)

	if err != nil {
		return Comment{}, err
	}
	return *comment, nil
}

func (cm *CommentsManager) CreateComment(comment Comment) (Nothing, error) {
	if cm.grpcConn == nil {
		return Nothing{}, ErrConnRefused
	}
	fmt.Println("comment is   ", comment)

	_, err := cm.cmc.CreateComment(
		context.Background(),
		&Comment{Username: comment.Username, FilmTitle: comment.FilmTitle, Text: comment.Text},
	)

	fmt.Println(err)

	if err != nil {
		return Nothing{}, err
	}
	return Nothing{}, nil
}

func (cm *CommentsManager) GetCommentsByFilmID(cID CommentID) (CommentsResponse, error) {
	if cm.grpcConn == nil {
		return CommentsResponse{}, ErrConnRefused
	}

	comments, err := cm.cmc.GetCommentsByFilmID(
		context.Background(),
		&CommentID{Film: cID.Film},
	)

	if err != nil {
		return CommentsResponse{}, err
	}
	return *comments, nil
}

func (cm *CommentsManager) GetCommentsByUserID(cID CommentID) (CommentsResponse, error) {
	if cm.grpcConn == nil {
		return CommentsResponse{}, ErrConnRefused
	}

	comments, err := cm.cmc.GetCommentsByUserID(
		context.Background(),
		&CommentID{User: cID.User},
	)

	if err != nil {
		return CommentsResponse{}, err
	}
	return *comments, nil
}

func (cm *CommentsManager) DeleteComment(cID CommentID) error {
	if cm.grpcConn == nil {
		return ErrConnRefused
	}

	_, err := cm.cmc.DeleteComment(
		context.Background(),
		&CommentID{CID: uint64(cID.CID)},
	)

	if err != nil {
		return err
	}
	return nil
}
