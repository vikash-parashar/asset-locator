package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/vikash-parashar/asset-locator/logger"
	"github.com/vikash-parashar/asset-locator/models"
	"github.com/vikash-parashar/asset-locator/utils"
)

func (db *DB) GetUserByEmailID(email string) (*models.User, error) {
	logger.InfoLogger.Println(email)
	query := `
        SELECT id, first_name, last_name, phone, email, password, role
        FROM users
        WHERE email = ?;
    `
	user := &models.User{}
	err := db.QueryRow(query, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		logger.ErrorLogger.Printf("Error fetching user by email: %v", err)
		return nil, err
	}
	logger.InfoLogger.Println("User From DB : ", user)
	return user, nil
}

func (db *DB) RegisterUser(user *models.User) error {
	query := `
        INSERT INTO users (first_name, last_name, phone, email, password, role)
        VALUES (?, ?, ?, ?, ?, ?)
    `
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Role)
	if err != nil {
		logger.ErrorLogger.Printf("Error registering user: %v", err)
		return err
	}
	// user.ID, _ = res.LastInsertId()
	return nil
}

func (db *DB) UpdateUserPassword(userID int, newPassword string) error {
	query := `
        UPDATE users
        SET password = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, newPassword, userID)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating user password: %v", err)
		return err
	}
	return nil
}

// GetAllUsers retrieves all active user records.
func (db *DB) GetAllUsers() ([]*models.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role
        FROM users
    `
	rows, err := db.Query(query)
	if err != nil {
		logger.ErrorLogger.Printf("Error fetching all users: %v", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
		if err != nil {
			logger.ErrorLogger.Printf("Error scanning user rows: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		logger.ErrorLogger.Printf("Error iterating over user rows: %v", err)
		return nil, err
	}
	return users, nil
}

// GetUserByResetToken retrieves a user by their reset token.
func (db *DB) GetUserByResetToken(resetToken string) (*models.User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role
        FROM users
        WHERE reset_token = ?
    `
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found by reset token")
		}
		logger.ErrorLogger.Printf("Error fetching user by reset token: %v", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user's information in the database.
func (db *DB) UpdateUser(user *models.User) error {
	query := `
        UPDATE users
        SET first_name = ?, last_name = ?, email = ?, password = ?, role = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// SetResetToken sets the reset token and reset token expiry for a user in the database.
func (db *DB) SetResetToken(userID int, resetToken string, expiryTime time.Time) error {
	query := `
        UPDATE users
        SET reset_token = ?, reset_token_expiry = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, resetToken, expiryTime, userID)
	if err != nil {
		logger.ErrorLogger.Printf("Error setting reset token: %v", err)
		return err
	}
	return nil
}

// ClearResetToken clears the reset token for a user in the database.
func (db *DB) ClearResetToken(userID int) error {
	query := `
        UPDATE users
        SET reset_token = NULL
        WHERE id = ?
    `
	_, err := db.Exec(query, userID)
	if err != nil {
		logger.ErrorLogger.Printf("Error clearing reset token: %v", err)
		return err
	}
	return nil
}

// VerifyResetToken verifies the reset token for a user.
func (db *DB) VerifyResetToken(resetToken string) (*models.User, error) {
	query := `
        SELECT id, first_name, email, reset_token, reset_token_expiry
        FROM users
        WHERE reset_token = ?
    `
	user := &models.User{}
	err := db.QueryRow(query, resetToken).Scan(&user.ID, &user.FirstName, &user.Email, &user.ResetToken, &user.ResetTokenExpiry)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("reset token not found")
		}
		logger.ErrorLogger.Printf("Error fetching user by reset token: %v", err)
		return nil, err
	}
	logger.InfoLogger.Println("user from db for password reset token : ")
	logger.InfoLogger.Println("getting user from token")
	// logger.InfoLogger.Println(user.ResetToken)
	// logger.InfoLogger.Println(user.FirstName)
	// logger.InfoLogger.Println(user.Email)
	// Check if the reset token has expired (optional)
	if utils.IsTokenExpired(user.ResetTokenExpiry) {
		return nil, errors.New("reset token has expired")
	}

	return user, nil
}

// UpdateUserProfile updates a user's profile information in the database.
func (db *DB) UpdateUserProfile(user *models.User) error {
	query := `
        UPDATE users
        SET first_name = ?, last_name = ?, phone = ?, email = ?, password = ?, role = ?
        WHERE id = ?
    `
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Phone, user.Email, user.Password, user.Role, user.ID)
	if err != nil {
		logger.ErrorLogger.Printf("Error updating user profile: %v", err)
		return err
	}
	return nil
}
