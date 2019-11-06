package sessions_usecase

import (
	"kino_backend/repository"
	"context"
)


type SessionsUseCase struct{
	sessionRepo repository.SessionsRepository
}

func NewSessionUseCase(sessionRepo repository.SessionsRepository) *SessionsUseCase{
	return &SessionsUseCase{
		sessionRepo: sessionRepo,
	}
}

func (s SessionsUseCase) Create(ctx context.Context, sID string, uID uint) (bool, error){
	return s.sessionRepo.Create(ctx, sID, uID)
}

func (s SessionsUseCase) Get(ctx context.Context, sID string) (uint, error){
	return s.sessionRepo.Get(ctx, sID)
}

func (s SessionsUseCase) Delete(ctx context.Context, sID string) error{
	return s.sessionRepo.Delete(ctx, sID)
}