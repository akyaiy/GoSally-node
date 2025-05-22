package database

import (
	"GoSally/internal/logger"
	"database/sql"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"strings"
)

type SQLiteDriver struct {
	db   *sql.DB
	_src string
}

func ensureDBPath(src string) error {
	dir := filepath.Dir(src[strings.Index(src, "/"):])
	return os.MkdirAll(dir, 0755)
}

func (s *SQLiteDriver) OpenDB(src string) error {
	var err error
	s._src = src
	err = ensureDBPath(s._src)
	if err != nil {
		logger.DatabaseLog.Error("Error creating database directory", "db_source", s._src, "err", err)
	}

	logger.DatabaseLog.Debug("Opening a database", "db_source", s._src)
	s.db, err = sql.Open("sqlite", s._src)
	if err != nil {
		logger.DatabaseLog.Error("Error opening database", "db_source", s._src, "err", err)
		return err
	}
	logger.DatabaseLog.Debug("The database is open", "db_source", s._src)
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS sessions (
        session_id TEXT PRIMARY KEY,
        data BLOB
    );`
	_, err = s.db.Exec(createTableSQL)
	if err != nil {
		logger.DatabaseLog.Error("Error executing command in database", "db_source", s._src, "err", err)
		return err
	}
	return nil
}

func (s *SQLiteDriver) CloseDB() error {
	logger.DatabaseLog.Debug("Closing a database", "db_source", s._src)
	err := s.db.Close()
	if err != nil {
		logger.DatabaseLog.Error("Error closing database", "db_source", s._src, "err", err)
		return err
	}
	logger.DatabaseLog.Debug("The database is closed", "db_source", s._src)
	return nil
}

func (s *SQLiteDriver) InitSession(id string, data []byte) error {
	_, err := s.db.Exec("INSERT INTO sessions(session_id, data) VALUES(?, ?)", id, data)
	if err != nil {
		logger.DatabaseLog.Error("Error executing command in database", "db_source", s._src, "err", err)
		return err
	}
	return nil
}

func (s *SQLiteDriver) QuerySession(id string) (data []byte, err error) {
	row := s.db.QueryRow("SELECT data FROM sessions WHERE session_id=?", id)
	err = row.Scan(&data)
	if err != nil {
		logger.DatabaseLog.Error("Error while querying the database", "db_source", s._src, "err", err)
		return nil, err
	}
	return data, nil
}

func (s *SQLiteDriver) CloseSession(id string) error {
	_, err := s.db.Exec("DELETE FROM sessions WHERE session_id=?", id)
	if err != nil {
		logger.DatabaseLog.Error("Error executing command in database", "db_source", s._src, "err", err)
		return err
	}
	return nil
}
