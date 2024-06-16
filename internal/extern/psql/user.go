package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
	"errors"
	"log"
)

func (p *PostgresProvider) GetUser(id int64) (*data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE id = $1", id)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (p *PostgresProvider) GetUserByEmail(email string) (*data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE email = $1", email)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (p *PostgresProvider) GetUserByUsername(userName string) (*data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE user_name = $1", userName)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (p *PostgresProvider) ListUsers() ([]*data.User, error) {
	rows, err := p.db.Query("SELECT id, email, user_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*data.User{}
	for rows.Next() {
		user := data.User{}
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (p *PostgresProvider) PutUser(user *data.User) (*data.User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	log.Printf("Creating user with email %s, user name %s", user.Email, user.Username)

	var insertedUser data.User
	err = tx.QueryRow(`
		INSERT INTO users (email, user_name)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at, email, user_name
	`, user.Email, user.Username).Scan(
		&insertedUser.Id,
		&insertedUser.CreatedAt,
		&insertedUser.UpdatedAt,
		&insertedUser.Email,
		&insertedUser.Username,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &insertedUser, nil
}

func (p *PostgresProvider) UpdateUser(user *data.User) (*data.User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var updatedUser data.User
	err = tx.QueryRow(`
		UPDATE users
		SET email = $1, user_name = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, created_at, updated_at, email, user_name
	`, user.Email, user.Username, user.Id).Scan(
		&updatedUser.Id,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
		&updatedUser.Email,
		&updatedUser.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (p *PostgresProvider) DeleteUser(id int64) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	err = tx.QueryRow("DELETE FROM users WHERE id = $1", id).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
