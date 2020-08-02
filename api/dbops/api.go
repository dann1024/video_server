package dbops

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES($1,$2)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)

	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd from users where login_name=$1")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("Delete from users where login_name=$1 and pwd=$2")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewView(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04:05") //M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info(id,author_id,name,display_ctime) VALUES($1,$2,$3,$4)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id,name,display_ctime from video_info where id=$1")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var author_id int
	var name string
	var display_ctime string
	err = stmtOut.QueryRow(vid).Scan(&author_id, &name, &display_ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: author_id, Name: name, DisplayCtime: display_ctime}
	return res, nil

}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("Delete from video_info where id=$1")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	t := time.Now()
	timeLayout := "2006-01-02 15:04:05"
	ctime := t.Format(timeLayout) //M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare("Insert INTO comments(id,video_id,author_id,content,time) Values ($1,$2,$3,$4,$5)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content, ctime)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int64) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`select comments.id,users.login_name,comments.content from comments
                                                              inner join users on comments.author_id=users.id
											where comments.video_id=$1 and comments.time>=$2 and comments.time <=$3`)
	if err != nil {
		return nil, err
	}
	timeLayout := "2006-01-02 15:04:05"
	fromStr := time.Unix(from, 0).Format(timeLayout)
	toStr := time.Unix(to, 0).Format(timeLayout)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, fromStr, toStr)

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
