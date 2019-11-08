package useCase

import (
	"context"
	"kino_backend/repository"
)

type SessionsUseCase interface{
	Create(ctx context.Context, sID string, uID uint) (bool, error)
	Get(ctx context.Context, sID string) (uint, error)
	Delete(ctx context.Context, sID string) error
}

type sessionsUseCase struct{
	sessionRepo repository.SessionRepository
}

func NewSessionUseCase(sessionRepo repository.SessionRepository) *sessionsUseCase {
	return &sessionsUseCase{
		sessionRepo: sessionRepo,
	}
}

func (s sessionsUseCase) Create(ctx context.Context, sID string, uID uint) (bool, error){
	return s.sessionRepo.Create(sID, uID)
}

func (s sessionsUseCase) Get(ctx context.Context, sID string) (uint, error){
	return s.sessionRepo.Get(ctx, sID)
}

func (s sessionsUseCase) Delete(ctx context.Context, sID string) error{
	return s.sessionRepo.Delete(ctx, sID)
}