package sessions_test

import (
	"bytes"
	"github.com/akyaiy/GoSally-node/internal/database"
	"github.com/akyaiy/GoSally-node/internal/database/sqlite"
	"github.com/akyaiy/GoSally-node/internal/logger"
	_ "modernc.org/sqlite"
	"os"
	"testing"
)

var (
	testDriver database.DBSessions = &sqlite_driver.Driver{}

	sessionId = "id12345"
	reqData   = []byte("hello world")
	ansData   []byte

	err error
)

func TestMain(m *testing.M) {
	logger.InitLog("true")

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = testDriver.OpenDB("file:" + dir + "/database/db.sqlite")
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}

	code := m.Run()

	err = testDriver.CloseDB()
	if err != nil {
		panic("Error closing DB:" + err.Error())
	}

	os.Exit(code)
}

func TestSQLSessionInit(t *testing.T) {
	err = testDriver.InitSession(sessionId, reqData)
	if err != nil {
		t.Fatal("Error initing session:", err)
	}
}

func TestSQLSessionQuery(t *testing.T) {
	ansData, err = testDriver.QuerySession(sessionId)
	if err != nil {
		t.Fatal("Error getting session:", err)
	}

	if !bytes.Equal(reqData, ansData) {
		t.Errorf("data mismath: got %q, want %q", ansData, reqData)
	}
}

func TestSQLSessionClose(t *testing.T) {
	err = testDriver.CloseSession(sessionId)
	if err != nil {
		t.Fatal("Error closing session:", err)
	}
}
