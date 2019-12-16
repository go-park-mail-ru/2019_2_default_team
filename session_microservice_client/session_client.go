package session_microservice_client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"kino_backend/logger"
	"time"
)

// nolint: golint
type SessionManager struct {
	smc      SessionManagerClient
	grpcConn *grpc.ClientConn
}

func ConnectSessionManager(address string) *SessionManager {
	grpcConn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(30*time.Second), // nolint: megacheck
	)
	if err != nil {
		logger.Panicf("failed to connect to sessionManager at address %v: %v", address, err)
	}

	smc := NewSessionManagerClient(grpcConn)

	logger.Infof("Successfully connected to sessionManager: %v", address)

	return &SessionManager{smc: smc, grpcConn: grpcConn}
}

func (sm *SessionManager) Create(sID string, uID uint) (bool, error) {
	if sm.grpcConn == nil {
		return false, ErrConnRefused
	}

	success, err := sm.smc.Create(
		context.Background(),
		&Session{UID: uint64(uID), SID: sID},
	)

	if err != nil {
		return false, err
	}
	return success.Success, nil
}

func (sm *SessionManager) Get(sID string) (uint, error) {
	if sm.grpcConn == nil {
		return 0, ErrConnRefused
	}
	s, err := sm.smc.Get(
		context.Background(),
		&SessionID{SID: sID},
	)
	if err != nil {
		s, _ := status.FromError(err)
		if s.Message() == ErrKeyNotFound.Error() {
			fmt.Println("error2")
			return 0, ErrKeyNotFound
		}
		return 0, err
	}
	return uint(s.UID), nil
}

func (sm *SessionManager) Delete(sID string) error {
	if sm.grpcConn == nil {
		return ErrConnRefused
	}

	_, err := sm.smc.Delete(
		context.Background(),
		&SessionID{SID: sID},
	)
	return err
}

func (sm *SessionManager) Close() error {
	if sm.grpcConn == nil {
		return ErrConnRefused
	}

	err := sm.grpcConn.Close()
	sm.grpcConn = nil
	return err
}
