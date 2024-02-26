package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"
)

type Sessions map[string]Session

type Session struct {
	userId     int
	expiration time.Time
}

func (session Session) isExpired() bool {
	return session.expiration.Before(time.Now())
}

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func InitializeSessionManager() Sessions {
	return make(Sessions)
}

func (sessionManager Sessions) Add(userId int) (http.Cookie, error) {
	expiration := time.Now().Add(time.Minute * 30)
	token, err := generateToken()
	if err != nil {
		return http.Cookie{}, err
	}
	session := Session{
		userId:     userId,
		expiration: expiration,
	}
	cookie := http.Cookie{
		Expires: expiration,
		Name:    "session_token",
		Value:   token,
	}
	sessionManager[token] = session
	return cookie, nil
}

func (sessionManager Sessions) Remove(token string) {
	delete(sessionManager, token)
}
