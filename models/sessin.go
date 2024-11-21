package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"webdev/rand"
)

const (
	MinBytesPerToken = 32
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB            *sql.DB
	BytesPerToken int
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)

	}
	session := Session{
		UserID: userID,
		Token:  token,
		TokenHash: ss.hash(token),
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])

}



// type TokenManager struct {} 

// func (tm TokenManager) New() (token,tokenHash string, err error) {
// 	return 
// }
