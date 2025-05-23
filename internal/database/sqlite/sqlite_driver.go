package sqlite_driver

import (
	"GoSally/internal/logger"
	"database/sql"
	"errors"
	_ "modernc.org/sqlite"
	"os"
	"path/filepath"
	"strings"
)

type Driver struct {
	db   *sql.DB
	_src string

	/*
	 * Locks the database interface.
	 * _DBFileLock true = file is open
	 * _DBFileLock false = file is closed
	 */
	_DBFileLock bool
}

func ensureDBPath(src string) error {
	dir := filepath.Dir(src[strings.Index(src, "/"):])
	return os.MkdirAll(dir, 0755)
}

func (s *Driver) OpenDB(src string) error {
	if s._DBFileLock {
		return errors.New("DB is already unlocked (opened)")
	}

	s._src = src
	err := ensureDBPath(s._src)
	if err != nil {
		logger.DatabaseLog.Error("Error creating database directory", "db_source", s._src, "err", err)
	}

	logger.DatabaseLog.Debug("Opening a database", "db_source", s._src)
	s.db, err = sql.Open("sqlite", s._src)
	if err != nil {
		logger.DatabaseLog.Error("Error opening database", "db_source", s._src, "err", err)
		return err
	}
	logger.DatabaseLog.Debug("The database is opened", "db_source", s._src)
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
	s._DBFileLock = true
	return nil
}

func (s *Driver) CloseDB() error {
	if !s._DBFileLock {
		return errors.New("DB is already locked (closed)")
	}
	logger.DatabaseLog.Debug("Closing a database", "db_source", s._src)
	err := s.db.Close()
	if err != nil {
		logger.DatabaseLog.Error("Error closing database", "db_source", s._src, "err", err)
		return err
	}
	logger.DatabaseLog.Debug("The database is closed", "db_source", s._src)
	s._DBFileLock = false
	return nil
}

func (s *Driver) InitSession(id string, data []byte) error {
	if !s._DBFileLock {
		return errors.New("DB is locked (closed)")
	}
	_, err := s.db.Exec("INSERT INTO sessions(session_id, data) VALUES(?, ?)", id, data)
	if err != nil {
		logger.DatabaseLog.Error("Error executing command in database", "db_source", s._src, "err", err)
		return err
	}
	return nil
}

func (s *Driver) QuerySession(id string) (data []byte, err error) {
	if !s._DBFileLock {
		return nil, errors.New("DB is locked (closed)")
	}
	row := s.db.QueryRow("SELECT data FROM sessions WHERE session_id=?", id)
	err = row.Scan(&data)
	if err != nil {
		logger.DatabaseLog.Error("Error while querying the database", "db_source", s._src, "err", err)
		return nil, err
	}
	return data, nil
}

func (s *Driver) CloseSession(id string) error {
	if !s._DBFileLock {
		return errors.New("DB is locked (closed)")
	}
	_, err := s.db.Exec("DELETE FROM sessions WHERE session_id=?", id)
	if err != nil {
		logger.DatabaseLog.Error("Error executing command in database", "db_source", s._src, "err", err)
		return err
	}
	return nil
}
