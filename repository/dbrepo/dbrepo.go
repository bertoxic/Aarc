package dbrepo

import (
	"database/sql"

	"github.com/bertoxic/aarc/config"
	"github.com/bertoxic/aarc/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
    DB *sql.DB 
}


func NewPostgresDBRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {

    return &postgresDBRepo{
        App: a,
        DB: conn,
    }
}