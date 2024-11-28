package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	"webdev/rand"
)

const (
	DefaultResetDuration = 1 * time.Hour
)

type PasswordReset struct {
	ID        int
	UserID    int
	Token     string
	TokenHash string
	ExpiresAt time.Time
}

type PasswordResetService struct {
	DB            *sql.DB
	BytesPerToken int
	Duration      time.Duration
	// Now func() time.Time
}

func (service *PasswordResetService) Create(email string) (*PasswordReset, error) {
	email = strings.ToLower(email)
	var UserID int
	row := service.DB.QueryRow(`
	select id from users where email = $1
	;`, email)
	err := row.Scan(&UserID)
	if err != nil {

		return nil, fmt.Errorf("create; %w", err)
	}
	// return nil, fmt.Errorf("implement password reser service create")
	bytesPerToken := service.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)

	}
	duration := service.Duration
	if duration == 0 {
		duration = DefaultResetDuration
	}
	pwReset := PasswordReset{
		UserID:    UserID,
		Token:     token,
		TokenHash: service.hash(token),
		ExpiresAt: time.Now().Add(duration),
	}
	row = service.DB.QueryRow(`
	insert into password_resets (user_id , token_hash,expires_at)
	values ($1,$2,$3) on conflict (user_id) do
	update
	set token_hash = $2 ,expires_at = $3
	returning id;`,
		pwReset.UserID, pwReset.TokenHash, pwReset.ExpiresAt)
	err = row.Scan(&pwReset.ID)
	if err != nil {
		return nil, fmt.Errorf("create : %w", err)
	}
	return &pwReset, nil
}

func (service *PasswordResetService) Consume(token, newPassword string) (*User, error) {
	return nil, fmt.Errorf("implement password reset service ")
}

func (service *PasswordResetService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])

}
