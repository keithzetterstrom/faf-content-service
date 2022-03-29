package repository

import (
	"github.com/keithzetterstrom/faf-content-service/internal/pkg/data"
	db "github.com/keithzetterstrom/faf-content-service/internal/pkg/repository/postgres"
)

type dataRepository struct {
	dbHandler db.Handler
}

func New(dbHandler db.Handler) data.Repository {
	return &dataRepository{
		dbHandler: dbHandler,
	}
}
