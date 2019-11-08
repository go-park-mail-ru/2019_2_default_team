package repository

import "context"

type SessionsRepository interface{
	Create(sID string, uID uint) (bool, error)
	Get(ctx context.Context, sID string) (uint, error)
	Delete(ctx context.Context, sID string) error
}


