package dblock_test

import (
	"GoSally/internal/database"
	"GoSally/internal/database/sqlite"
	_ "modernc.org/sqlite"
	"os"
	"testing"
)

var (
	testDriver database.DBSessions = &sqlite_driver.Driver{}
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestAttempt1ToClosedDB(t *testing.T) {
	err := testDriver.CloseDB()
	if !(err != nil && err.Error() == "DB is already locked (closed)") {
		t.Fatalf("Unexpected behavior: expected \"%s\", got \"%s\"", "DB is already locked (closed)", err)
	}
}

func TestAttempt2ToClosedDB(t *testing.T) {
	err := testDriver.InitSession("\000", nil)
	if !(err != nil && err.Error() == "DB is locked (closed)") {
		t.Fatalf("Unexpected behavior: expected \"%s\", got \"%s\"", "DB is locked (closed)", err)
	}
}

func TestAttempt3ToClosedDB(t *testing.T) {
	_, err := testDriver.QuerySession("\000")
	if !(err != nil && err.Error() == "DB is locked (closed)") {
		t.Fatalf("Unexpected behavior: expected \"%s\", got \"%s\"", "DB is locked (closed)", err)
	}
}

func TestAttempt4ToClosedDB(t *testing.T) {
	err := testDriver.CloseSession("\000")
	if !(err != nil && err.Error() == "DB is locked (closed)") {
		t.Fatalf("Unexpected behavior: expected \"%s\", got \"%s\"", "DB is locked (closed)", err)
	}
}

func TestAttempt1ToOpenedDB(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = testDriver.OpenDB("file:" + dir + "/database/db.sqlite")
	if err != nil {
		panic("Error opening DB:" + err.Error())
	}

	err = testDriver.OpenDB("\000")
	if !(err != nil && err.Error() == "DB is already unlocked (opened)") {
		t.Fatalf("Unexpected behavior: expected \"%s\", got \"%s\"", "DB is already unlocked (opened)", err)
	}
	err = testDriver.CloseDB()
	if err != nil {
		panic("Error closing DB:" + err.Error())
	}
}
