package persistence

import (
    "database/sql"
    "repo-api/src/domain/model"
    "repo-api/src/domain/repository"
    "fmt"
)

func NewUserPersistence() repository.IUserRepository {
    return &userPersistence{}
}

type userPersistence struct{}

func (u *userPersistence) Insert(DB *sql.DB, ID, Name string) error {
    var existingID string
    checkQuery := "SELECT id FROM users WHERE id = ?"
    err := DB.QueryRow(checkQuery, ID).Scan(&existingID)
    if existingID != "" {
        return fmt.Errorf("user with ID %s already exists", ID)
    }
    if err != nil && err != sql.ErrNoRows {
        return fmt.Errorf("failed to check user existence: %w", err)
    }

    query := "INSERT INTO users (id, name) VALUES (?, ?)"
    _, err = DB.Exec(query, ID, Name)
    if err != nil {
        return err
    }
    return nil
}

func (u *userPersistence) GetByID(DB *sql.DB, ID string) (model.User, error) {
    var user model.User
    query := "SELECT id, name FROM users WHERE id = ?"
    err := DB.QueryRow(query, ID).Scan(&user.ID, &user.Name)
    if err != nil {
        return user, err
    }
    return user, nil
}

func (u *userPersistence) UpdateNameByID(DB *sql.DB, ID, Name string) error {
    var existingID string
    checkQuery := "SELECT id FROM users WHERE id = ?"
    err := DB.QueryRow(checkQuery, ID).Scan(&existingID)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("user not found")
        }
        return fmt.Errorf("failed to check user existence: %w", err)
    }

    updateQuery := "UPDATE users SET name = ? WHERE id = ?"
    _, err = DB.Exec(updateQuery, Name, ID)
    if err != nil {
        return fmt.Errorf("failed to update user: %w", err)
    }

    return nil
}
