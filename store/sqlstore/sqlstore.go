package sqlstore

import "github.com/adetunjii/netflakes/port"

type Stores struct {
	movie port.MovieStore
}

type SqlStore struct {
	db     port.DB
	logger port.Logger
	stores Stores
}

var _ port.SqlStore = (*SqlStore)(nil)

func New(db port.DB, logger port.Logger) *SqlStore {
	sqlstore := &SqlStore{
		db:     db,
		logger: logger,
	}

	sqlstore.stores.movie = newMovieStore(sqlstore)
	return sqlstore
}

func (s *SqlStore) Movie() port.MovieStore {
	return s.stores.movie
}
