package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string){
	//删除session
	//内存删除
	sessionMap.Delete(sid)
	//数据库删除
	dbops.DeleteSession(sid)
}

//加载session
func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(k, v interface{}) bool {
		//因为取出来的v类型为接口 所以要转换一下
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 //过期时间
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsetSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {

	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct{
			//session已经过期 删除数据
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
