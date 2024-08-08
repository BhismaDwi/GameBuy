package middlewares

import "time"

type UserLoginRedis struct {
	UserId    int64
	Username  string
	Role      string
	LoginAt   time.Time
	ExpiredAt time.Time
}

var DummyRedis = make(map[string]UserLoginRedis)
