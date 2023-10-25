package db

import (
	"database/sql"
	"errors"
	"go-server/models"
	"go-server/utils"
	"log"
)

// GetUserByEmailID retrieves a user by their email address.
func (db *DB) GetUserByEmailID(email string) (*models.User, error) {
	query := `
        SELECT * FROM users
        WHERE email = $1
    `
	user := &models.User{}
	err := db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Error fetching user by email: %v", err)
		return nil, err
	}
	return user, nil
}

// RegisterUser creates a new user record in the database.
func (db *DB) RegisterUser(user *models.User) error {
	query := `
        INSERT INTO users (first_name, last_name, email, password, role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

// UpdateUserPassword updates a user's password by their ID.
func (db *DB) UpdateUserPassword(userID int, newPassword string) error {
	query := `
        UPDATE users
        SET password = $2
        WHERE id = $1
    `
	_, err := db.Exec(query, userID, newPassword)
	if err != nil {
		log.Printf("Error updating user password: %v", err)
		return err
	}
	return nil
}

// GetAllUsers retrieves all active user records.
func (db *DB) GetAllUsers() ([]*models.User, error) {
	query := `
        SELECT * FROM users
    `
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error fetching all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			log.Printf("Error scanning user rows: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over user rows: %v", err)
		return nil, err
	}
	return users, nil
}

// GetUserByResetToken retrieves a user by their reset token.
func (db *DB) GetUserByResetToken(resetToken string) (*models.User, error) {
	query := `
        SELECT * FROM users
        WHERE reset_token = $1
    `
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Printf("Error fetching user by reset token: %v", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user's information in the database.
func (db *DB) UpdateUser(user *models.User) error {
	query := `
        UPDATE users
        SET first_name = $1, last_name = $2, email = $3, password = $4, role = $5
        WHERE id = $6
    `
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// SetResetToken sets the reset token for a user in the database.
func (db *DB) SetResetToken(userID int, resetToken string) error {
	// Assuming you have a "users" table with a "reset_token" column.
	_, err := db.Exec("UPDATE users SET reset_token = ? WHERE id = ?", resetToken, userID)
	if err != nil {
		return err
	}
	return nil
}

// ClearResetToken clears the reset token for a user in the database.
func (db *DB) ClearResetToken(userID int) error {
	// Assuming you have a "users" table with a "reset_token" column.
	_, err := db.Exec("UPDATE users SET reset_token = NULL WHERE id = ?", userID)
	if err != nil {
		return err
	}
	return nil
}

// VerifyResetToken verifies the reset token for a user.
func (db *DB) VerifyResetToken(resetToken string) (*models.User, error) {
	query := `
    SELECT id, email, reset_token
    FROM users
    WHERE reset_token = $1
`
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.Email, &user.ResetToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Reset token not found")
		}
		return nil, err
	}

	// Check if the reset token has expired (optional)
	if utils.IsTokenExpired(user.ResetTokenExpiry) {
		return nil, errors.New("Reset token has expired")
	}

	return user, nil
}
