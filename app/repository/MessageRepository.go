package repository

import (
	"database/sql"
	"errors"
)
import "github.com/google/uuid"

type MessageWriteRepository struct {
	db *sql.DB
}

func CreateMessageWriteRepository(db *sql.DB) *MessageWriteRepository {
	return &MessageWriteRepository{db: db}
}

func (repository *MessageWriteRepository) Create(id uuid.UUID, from uuid.UUID, to uuid.UUID, message string) error {
	stmt, err := repository.db.Prepare("INSERT INTO messages (id, `from`, `to`, message) VALUES (?,?,?,?)")
	if err != nil {
		return errors.New("failed to create user, error: " + err.Error())
	}

	_, err = stmt.Exec(id.String(), from, to, message)
	if err != nil {
		return errors.New("failed to create user, error: " + err.Error())
	}

	return nil
}
