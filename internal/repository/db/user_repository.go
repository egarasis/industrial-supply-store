package db

import (
	"context"
	"database/sql"
	"errors"
	"industrial-supply-store/internal/domain"
	"industrial-supply-store/internal/model/entity"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {

	query := `
		INSERT INTO users (email, password, role)
		VALUES (?, ?, ?)
	`

	result, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {

	var user entity.User

	query := `
	SELECT
		id,
		email,
		password,
		role
	FROM users
	WHERE email= ?
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	return &user, err
}

func (r *userRepository) FindAll(ctx context.Context) ([]entity.User, error) {

	query := `
		SELECT id, email, role
		FROM users
		ORDER BY id
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {

		var user entity.User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Role,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {

	query := `
		SELECT id, name, email
		FROM users
		WHERE id = ?
	`

	var user entity.User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Role,
	)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {

	query := `
		UPDATE users
		SET
			email = ?,
			role = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		user.Email,
		user.Role,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {

	query := `
		DELETE FROM users
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}
