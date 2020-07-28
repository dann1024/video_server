package dbops

import (
	"testing"
)

var tempvid string

func clearTables() {
	dbConn.Exec("delete from users")
	dbConn.Exec("delete from  video_info")
	dbConn.Exec("delete from  comments")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("ReGet", testRegetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("abc123", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}

}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("abc123")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}

}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("abc123", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("abc123")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestGetVideoInfo(t *testing.T) {
	clearTables()
	t.Run("PrepPateUser", testAddUser)
	t.Run("AddVideo", testAddNewView)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)

}

func testAddNewView(t *testing.T) {
	vi, err := AddNewView(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideo: %v", err)
	}
}

func testReGetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of GetVideo: %v", err)
	}
}
