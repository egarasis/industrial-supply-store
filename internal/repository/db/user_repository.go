package db

import (
	"context"
	"database/sql"
	"errors"
	"industrial-supply-store/internal/model/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// --- FUNGSI WAJIB SESUAI INTERFACE ---

func (r *UserRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `INSERT INTO users (email, password, role) VALUES (?, ?, ?)`
	result, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	lastID, _ := result.LastInsertId()
	user.ID = int(lastID)
	return user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	query := `SELECT id, email, password, role FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*entity.User, error) {
	query := `SELECT id, email, password, role FROM users WHERE id = ?`
	row := r.db.QueryRowContext(ctx, query, id)

	var u entity.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, password, role FROM users WHERE email = ?`
	row := r.db.QueryRowContext(ctx, query, email)

	var u entity.User
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET email = ?, password = ?, role = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Role, user.ID)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	// Catatan: Pastikan ini sesuai dengan tabel yang ingin kamu hapus
	query := `DELETE FROM user_profiles WHERE user_id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// Ganti bagian ini: func (r *userRepository)
// Menjadi: func (r *UserRepository)

func (r *UserRepository) UpdateProfile(ctx context.Context, profile entity.UserProfile) error {
	// Query diperbarui agar mencakup contact_name, phone, dan address
	query := `
        INSERT INTO user_profiles (user_id, full_name, company_name, contact_name, phone, address)
        VALUES (?, ?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE 
            full_name = ?, 
            company_name = ?,
            contact_name = ?,
            phone = ?,
            address = ?;
    `

	_, err := r.db.ExecContext(
		ctx,
		query,
		profile.UserID,
		profile.FullName,
		profile.CompanyName,
		profile.ContactName, // Parameter insert
		profile.Phone,       // Parameter insert
		profile.Address,     // Parameter insert
		profile.FullName,    // Parameter update (jika data user sudah ada)
		profile.CompanyName, // Parameter update (jika data user sudah ada)
		profile.ContactName, // Parameter update (jika data user sudah ada)
		profile.Phone,       // Parameter update (jika data user sudah ada)
		profile.Address,     // Parameter update (jika data user sudah ada)
	)

	return err
}
