package useCase

import "context"

type SessionsUseCase interface{
	Create(ctx context.Context, sID string, uID uint) (bool, error)
	Get(ctx context.Context, sID string) (uint, error)
	Delete(ctx context.Context, sID string) error
}

