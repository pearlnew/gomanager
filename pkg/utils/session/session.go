package session

var GlobalSession *Manager

var MaxLifeTime int64 = 15*60*60
var SESSION_USER = "user"

func InitSession(sessionName string) {
	var err error
	GlobalSession, err = NewSessionManager("memory", sessionName, MaxLifeTime)
	if err == nil {
		go GlobalSession.GC()
	}
}

