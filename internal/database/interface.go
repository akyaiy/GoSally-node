package database

type _dbfile interface {
	OpenDB(source string) error
	CloseDB() error
}

type DBSessions interface {
	_dbfile

	InitSession(id string, data []byte) error
	QuerySession(id string) (data []byte, err error)
	CloseSession(id string) error
}
