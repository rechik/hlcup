// +build db_use_cmap

package db

import (
	"github.com/golang/groupcache/singleflight"

	"github.com/ei-grad/hlcup/models"
)

// DB is inmemory database optimized for its task
type DB struct {
	users     *models.UserMap
	locations *models.LocationMap
	visits    *models.VisitMap

	locationMarks *models.LocationMarksMap
	userVisits    *models.UserVisitsMap

	sf singleflight.Group
}

// New creates new DB
func New() *DB {
	return &DB{
		users:         models.NewUserMap(509),
		locations:     models.NewLocationMap(509),
		visits:        models.NewVisitMap(509),
		locationMarks: models.NewLocationMarksMap(509),
		userVisits:    models.NewUserVisitsMap(509),
	}
}