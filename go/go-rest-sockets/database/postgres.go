package database

import (
	"context"
	"database/sql"
	"fmt"
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
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.ID, user.Email, user.Password,)
	return err
}

func (r *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)", post.Id, post.PostContent, post.UserId,)
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

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	// user := &models.User{}
	// err := r.db.QueryRowContext(ctx, "SELECT id, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email)
	// return user, err

	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}



func (r *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	var post = models.Post{}
	
	err := r.db.QueryRowContext(ctx, "SELECT id, created_at, post_content, user_id FROM posts WHERE id = $1", id).Scan(&post.Id, &post.CreatedAt, &post.PostContent, &post.UserId)
	
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return nil, err
	}

	return &post, nil
}

func (r *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "UPDATE posts SET post_content = $1 WHERE id = $2", post.PostContent, post.Id)

	return err
}

func (r *PostgresRepository) DeletePost(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1", id)

	return err
}

func (r *PostgresRepository) ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, created_at, post_content, user_id FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2", 2, page * 2)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var posts []*models.Post
	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.CreatedAt, &post.PostContent, &post.UserId); err == nil {
			posts = append(posts, &post)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
