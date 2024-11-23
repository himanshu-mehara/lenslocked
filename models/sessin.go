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
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	row := ss.DB.QueryRow(`
	insert into sessions (user_id , token_hash)
	values ($1,$2) on conflict (user_id) do
	update
	set token_hash = $2 
	returning id;'
	, session.UserID, session.TokenHash`)
	err = row.Scan(&session.ID)

	// if err == sql.ErrNoRows {
	// row := ss.DB.QueryRow(`
	// 	INSERT INTO sessions (user_id, token_hash)
	// 	values ($1,$2)
	// 	returning id;
	// 	`, session.UserID, session.TokenHash)
	// 	err = row.Scan(&session.ID)
	// }
	if err != nil {
		return nil, fmt.Errorf("create : %w", err)
	}
	return &session, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)
	var user User
	row := ss.DB.QueryRow(`
		select users.id,
		users.email,
		users.password_hash 
		from sessions 
		join users on users.id = sessions.user_id
		where sessions.token_hash = $1;`, tokenHash)

	err := row.Scan(&user.ID,&user.Email,&user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}
	// var user User
	// row = ss.DB.QueryRow(`
	// select email,password_hash
	// from users where id = $1;
	// `, user.ID)
	// err = row.Scan(&user.Email, &user.PasswordHash)
	// if err != nil {
	// 	return nil, fmt.Errorf("user: %w",err)
	// }
	return &user, nil
}

func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
	delete from sessions
	where token_hash = $1;`,tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w",err)
	}
	return nil
}

func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	return base64.URLEncoding.EncodeToString(tokenHash[:])

}

// type TokenManager struct {}

// func (tm TokenManager) New() (token,tokenHash string, err error) {
// 	return
// }
