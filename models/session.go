package models

import (
	"sync"
)

type Session struct {
	SessionID string `json:"session_id" example:"ef84d238-47ef-4452-9536-99380db79911"`
}

type Sessions struct {
	sync.Mutex
	Sessions map[string]uint
}

type Authorization struct {
	Authorized bool
}

type Success struct {
	Success bool `json:"Success"`
}
