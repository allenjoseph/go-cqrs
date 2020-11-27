package db

import (
	"context"
	"database/sql"

	"go-cqrs/model"
	"go-cqrs/util"
	_ "github.com/lib/pq" // needed for improve Scan func
)

// PostgresRepository struct
type PostgresRepository struct {
	db *sql.DB
}

// OpenConnection to open db connection
func OpenConnection(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	util.FailOnError(err, "Failed to open db connection")

	return &PostgresRepository{db}, nil
}

// InsertWoof implementation for insert a woof record
func (repository *PostgresRepository) InsertWoof(ctx context.Context, woof model.Woof) error {
	query := "INSERT INTO woofs(id, body, created_at) VALUES($1, $2, $3)"
	_, err := repository.db.Exec(query, woof.ID, woof.Body, woof.CreatedAt)
	return err
}

// ListWoofs implementation for list woof records
func (repository *PostgresRepository) ListWoofs(ctx context.Context, offset uint64, limit uint64) ([]model.Woof, error) {
	query := "SELECT * FROM woofs ORDER BY id DESC OFFSET $1 LIMIT $2"
	rows, err := repository.db.Query(query, offset, limit)
	if err != nil {
		return nil, err
	}

	woofs := []model.Woof{}
	for rows.Next() {
		woof := model.Woof{}
		err = rows.Scan(&woof.ID, &woof.Body, &woof.CreatedAt)
		if err == nil {
			woofs = append(woofs, woof)
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return woofs, err
}

// Close implementation for close the db connection
func (repository *PostgresRepository) Close() {
	repository.db.Close()
}
