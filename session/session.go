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

func (session Session) checkExpiration() bool {
	return session.expiration.Before(time.Now())
}

func generateSessionId() (string, error) {
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

func Generate(userId int, sessionManager Sessions) (http.Cookie, error) {
	expiration := time.Now().Add(time.Minute * 30)
	token, err := generateSessionId()
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

func (sessionManager Sessions) deleteSession(token string) {
	delete(sessionManager, token)
}
