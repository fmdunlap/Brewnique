package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"errors"
	"log"
)

func (p *PostgresProvider) GetUser(id int64) (data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE id = $1", id)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return data.User{}, err
	}

	return u, nil
}

func (p *PostgresProvider) GetUserByEmail(email string) (data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE email = $1", email)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return data.User{}, err
	}
	return u, nil
}

func (p *PostgresProvider) GetUserByUserName(userName string) (data.User, error) {
	result := p.db.QueryRow("SELECT id, email, user_name FROM users WHERE user_name = $1", userName)

	u := data.User{}
	err := result.Scan(&u)
	if err != nil {
		return data.User{}, err
	}
	return u, nil
}

func (p *PostgresProvider) ListUsers() ([]data.User, error) {
	rows, err := p.db.Query("SELECT id, email, user_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []data.User{}
	for rows.Next() {
		user := data.User{}
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (p *PostgresProvider) PutUser(user data.User) (data.User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return data.User{}, err
	}

	log.Printf("Creating user with email %s, user name %s", user.Email, user.UserName)

	err = tx.QueryRow("INSERT INTO users (email, user_name) VALUES ($1, $2) RETURNING id", user.Email, user.UserName).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return data.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return data.User{}, err
	}

	return user, nil
}

func (p *PostgresProvider) UpdateUser(user data.User) (data.User, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return data.User{}, err
	}

	existingUser := data.User{}
	err = tx.QueryRow("SELECT id, email, user_name FROM users WHERE id = $1", user.ID).Scan(&existingUser)
	if err != nil {
		tx.Rollback()
		return data.User{}, err
	}

	if existingUser.ID == 0 {
		return data.User{}, errors.New("user not found")
	}

	log.Printf("Updating user email %v -> %v, username %v -> %v", existingUser.Email, user.Email, existingUser.UserName, user.UserName)

	if user.Email == "" {
		user.Email = existingUser.Email
	}

	if user.UserName == "" {
		user.UserName = existingUser.UserName
	}

	err = tx.QueryRow("UPDATE users SET email = $1, user_name = $2 WHERE id = $3", user.Email, user.UserName, user.ID).Scan(&user.ID)
	if err != nil {
		tx.Rollback()
		return data.User{}, err
	}

	err = tx.Commit()
	if err != nil {
		return data.User{}, err
	}

	return user, nil

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
