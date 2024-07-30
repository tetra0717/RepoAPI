package application

import (
    "database/sql"
    "fmt"
    "repo-api/src/domain/repository"
    "repo-api/src/domain/model"
)

type UserApp interface {
    Register(DB *sql.DB, ID, Name string) error
    Get(DB *sql.DB, ID string) (model.User, error)
    Update(DB *sql.DB, ID, Name string) error
}

func NewUserApp(ur repository.IUserRepository) UserApp {
    return &userApp{
        userRepository: ur,
    }
}

type userApp struct {
    userRepository repository.IUserRepository
}

func (u *userApp) Register(DB *sql.DB, ID, Name string) error {
    err := u.userRepository.Insert(DB, ID, Name)
    if err != nil {
        return fmt.Errorf("failed to insert user: %w", err)
    }
    return nil
}

func (u *userApp) Get(DB *sql.DB, ID string) (model.User, error) {
    user, err := u.userRepository.GetByID(DB, ID)
    if err != nil {
        if err == sql.ErrNoRows {
            return model.User{}, fmt.Errorf("user not found: %w", err)
        }
        return model.User{}, fmt.Errorf("failed to get user by ID: %w", err)
    }
    return user, nil
}

func (u *userApp) Update(DB *sql.DB, ID, Name string) error {
    err := u.userRepository.UpdateNameByID(DB, ID, Name)
    if err != nil {
        if err.Error() == "user not found" {
            return err
        }
        return fmt.Errorf("failed to update user name by ID: %w", err)
    }
    return nil
}
