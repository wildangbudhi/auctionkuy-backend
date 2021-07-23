package auth

import "time"

type Session struct {
	AccessUUID  string `json:"access"`
	RefreshUUID string `json:"refresh"`
}

type SessionRepository interface {
	IsSessionExist(key string) (bool, error)
	SetSession(key string, data *Session, expiration time.Duration) error
	GetSession(key string) (*Session, error)
	RemoveSession(key string) error
	ExtendSessionExpiration(key string, expiration time.Duration) error
}
