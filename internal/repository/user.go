package repository

import (
	"context"
	"koda-b6-backend2/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	rows, err := r.db.Query(ctx, "SELECT email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, "SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *UserRepository) Update(ctx context.Context, email string, updatedUser models.User) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET password = $1 WHERE email = $2", updatedUser.Password, email)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, email string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE email = $1", email)
	return err
}
