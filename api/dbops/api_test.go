package dbops

import (
	"testing"
)

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
