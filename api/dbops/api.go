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
	if err == sql.ErrNoRows{
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