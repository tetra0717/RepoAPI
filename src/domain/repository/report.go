package repository

import (
    "database/sql"
    "repo-api/src/domain/model"
)

type IReportRepository interface {
    Insert(DB *sql.DB, ID, AuthorID string, Count int, Title, Style, Language string) error
    Eject(DB *sql.DB, ID string) error
    
    GetByID(DB *sql.DB, ID string) (model.Report, error)
    GetByAuthorID(DB *sql.DB, AuthorID string) ([]model.Report, error)
    GetByTitle(DB *sql.DB, AuthorID, Title string) ([]model.Report, error)
    GetByStyle(DB *sql.DB, AuthorID, Style string) ([]model.Report, error)
    GetByLanguage(DB *sql.DB, AuthorID, Language string) ([]model.Report, error)
    
    UpdateCount(DB *sql.DB, ID string, Count int) error
    UpdateTitle(DB *sql.DB, ID, Title string) error
    UpdateStyle(DB *sql.DB, ID, Style string) error
    UpdateLanguage(DB *sql.DB, ID, Language string) error
}
