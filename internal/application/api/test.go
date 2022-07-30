package api

import (
	"time"

	"github.com/google/uuid"
)

type session struct {
	username string
	expires  time.Time
}

type Session interface {
	Set(key, value interface{}) error
	Get(sid string) (session, error)
	Delete(sid string) error
	SessionID() string
}

func (session *session) sessionID() string {
	return uuid.New().String() //Need to sign this with a secret
}

/*

sessionToken := uuid.New().String() //Need to sign this with a secret
	ExpiresAt := time.Now().Add(time.Second * 180)

	//declare the object to be stored in the session
	session := session{
		username: creds.Username,
		expires:  ExpiresAt,
	}
	//Store the session in the redis cache
	err = a.redis.Set(ctx, sessionToken, session, time.Second*180)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("Unable to store session in redis")
		return
	}
	//fmt.Println(sessions)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: ExpiresAt,
	}) */
