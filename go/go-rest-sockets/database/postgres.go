package database

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Miguelburitica/go-rest-server/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	response, _ := r.db.ExecContext(ctx, "SELECT id FROM users WHERE email = $1", user.Email,)
	if response != nil {
		return errors.New("user already exists")
	}

	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password,)
	return err
}

func (r *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	// user := &models.User{}
	// err := r.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email)
	// return user, err

	rows, err := r.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
