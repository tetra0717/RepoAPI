package repository

import (
    "database/sql"
    "repo-api/src/domain/model"
)

type IUserRepository interface {
    Insert(DB *sql.DB, ID, Name string) error
    GetByID(DB *sql.DB, ID string) (model.User, error)
    UpdateNameByID(DB *sql.DB, ID, Name string) error
}
