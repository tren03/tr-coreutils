package todo

import (
	"database/sql"
	"log/slog"
)

type IRepo interface {
	CreateTodo(*Todo) error
}

type Repo struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewRepo(db *sql.DB, logger *slog.Logger) IRepo {
	return &Repo{
		db:     db,
		logger: logger,
	}
}

func (r *Repo) CreateTodo(todo *Todo) error {
	r.logger.Info("Mocking the create call for now")
	return nil
}
