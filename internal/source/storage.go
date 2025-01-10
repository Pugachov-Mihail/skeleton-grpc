package source

import (
	"database/sql"
	"fmt"
	"log/slog"
	conf "sceleton/internal/config"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath conf.Storage, log *slog.Logger) (*Storage, error) {
	//TODO при использовании поправить лог
	log = log.With("connect db", "sceleton-grpc")
	connDb := fmt.Sprintf("postgres://" + storagePath.UserDb + ":" + storagePath.PassDb + "@" + storagePath.Host +
		":" + storagePath.PortDb + "/" + storagePath.DbName + "?sslmode=disable")

	db, err := sql.Open("postgres", connDb)

	if err = db.Ping(); err != nil {
		log.Error("Ошибка базы ", err)
		return nil, fmt.Errorf("%s: %v", "postgre", err)
	}

	return &Storage{db: db}, nil
}
