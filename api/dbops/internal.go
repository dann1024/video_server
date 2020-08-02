package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsetSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("insert into sessions (session_id, ttl, login_name) values ($1,$2,$3)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)

	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select ttl,login_name from sessions where session_id=$1 ")
	if err != nil {
		return nil, err
	}
	var ttl, uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ttlint, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = ttlint
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

//说明 返回session集合{session_id:SimpleSession}
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("select id,ttl,login_name from sessions")
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.Query(stmtOut)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, ttl, login_name string
		if er := rows.Scan(&id, &ttl, &login_name); er != nil {
			log.Printf("retrieve session err: %s", err)
		}
		if ttlint, terr := strconv.ParseInt(ttl, 10, 64); terr == nil {
			ss := defs.SimpleSession{TTL: ttlint, Username: login_name}
			m.Store(id, ss)
			log.Printf("session id %s, ttl %d", id, ss.TTL)
		}
	}

	return m, nil
}

//session 删除
func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("delete from sessions where id=$1")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(sid)
	if err != nil {
		return err
	}
	return nil
}
