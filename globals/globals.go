package globals

import (
	"cloud-sek/models"
	"database/sql"
)

var (
	Config *models.Config
	DbConn *sql.DB
	Cache  *models.PostCache
)
