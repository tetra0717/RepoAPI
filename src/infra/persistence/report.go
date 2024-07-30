package persistence

import (
    "database/sql"
    "repo-api/src/domain/model"
    "repo-api/src/domain/repository"
    "log"
    "fmt"
)

func NewReportPersistence() repository.IReportRepository {
    return &reportPersistence{}
}

type reportPersistence struct{}

func (r *reportPersistence) Insert(DB *sql.DB, ID, AuthorID string, Count int, Title, Style, Language string) error {
	var authorExists bool
	authorQuery := "SELECT COUNT(*) > 0 FROM users WHERE id = ?"
	err := DB.QueryRow(authorQuery, AuthorID).Scan(&authorExists)
	if err != nil {
		return fmt.Errorf("failed to check author existence: %w", err)
	}
	if !authorExists {
	    return fmt.Errorf("author does not exist")
	}

	query := "INSERT INTO reports (id, author_id, count, title, style, language) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = DB.Exec(query, ID, AuthorID, Count, Title, Style, Language)
	if err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}

	return nil
}

func (r *reportPersistence) Eject(DB *sql.DB, ID string) error { 
    var existingID string
    checkQuery := "SELECT id FROM reports WHERE id = ?"
    err := DB.QueryRow(checkQuery, ID).Scan(&existingID)
    if err != nil {
        if err == sql.ErrNoRows {
            return fmt.Errorf("report not found")
        }
        return fmt.Errorf("failed to check report existence: %w", err)
    }
    
    query := "DELETE FROM reports WHERE id = ?"
    _, err = DB.Exec(query, ID)
    if err != nil {
        return fmt.Errorf("failed to delete report: %w", err)
    }
    return nil
}

func (r *reportPersistence) GetByID(DB *sql.DB, ID string) (model.Report, error) {
    var report model.Report
    query := "SELECT id, author_id, count, title, style, language FROM reports WHERE id = ?"
    err := DB.QueryRow(query, ID).Scan(&report.ID, &report.AuthorID, &report.Count, &report.Title, &report.Style, &report.Language)
    if err != nil {
        if err == sql.ErrNoRows {
            return report, nil  
        }
        return report, fmt.Errorf("failed to get report by ID: %w", err)
    }
    return report, nil
}

func (r *reportPersistence) GetByAuthorID(DB *sql.DB, AuthorID string) ([]model.Report, error) {
    var reports []model.Report
    
	var authorExists bool
	authorQuery := "SELECT COUNT(*) > 0 FROM users WHERE id = ?"
	err := DB.QueryRow(authorQuery, AuthorID).Scan(&authorExists)
	if err != nil {
		return reports, fmt.Errorf("failed to check author existence: %w", err)
	}
	if !authorExists {
	    return reports, fmt.Errorf("author does not exist")
	}

    query := "SELECT id, author_id, count, title, style, language FROM reports WHERE author_id = ?"
    rows, err := DB.Query(query, AuthorID)
    if err != nil {
        return reports, fmt.Errorf("failed to get reports by AuthorID: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var report model.Report
        err := rows.Scan(&report.ID, &report.AuthorID, &report.Count, &report.Title, &report.Style, &report.Language)
        if err != nil {
            log.Println("Error scanning report:", err)
            continue
        }
        reports = append(reports, report)
    }
    return reports, nil
}

func (r *reportPersistence) GetByTitle(DB *sql.DB, AuthorID, Title string) ([]model.Report, error) {
    var reports []model.Report
    
	var authorExists bool
	authorQuery := "SELECT COUNT(*) > 0 FROM users WHERE id = ?"
	err := DB.QueryRow(authorQuery, AuthorID).Scan(&authorExists)
	if err != nil {
		return reports, fmt.Errorf("failed to check author existence: %w", err)
	}
	if !authorExists {
	    return reports, fmt.Errorf("author does not exist")
	}

    query := "SELECT id, author_id, count, title, style, language FROM reports WHERE author_id = ? AND title = ?"
    rows, err := DB.Query(query, AuthorID, Title)
    if err != nil {
        return reports, fmt.Errorf("failed to get reports by Title: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var report model.Report
        err := rows.Scan(&report.ID, &report.AuthorID, &report.Count, &report.Title, &report.Style, &report.Language)
        if err != nil {
            log.Println("Error scanning report:", err)
            continue
        }
        reports = append(reports, report)
    }
    return reports, nil
}

func (r *reportPersistence) GetByStyle(DB *sql.DB, AuthorID, Style string) ([]model.Report, error) {
    var reports []model.Report
    
	var authorExists bool
	authorQuery := "SELECT COUNT(*) > 0 FROM users WHERE id = ?"
	err := DB.QueryRow(authorQuery, AuthorID).Scan(&authorExists)
	if err != nil {
		return reports, fmt.Errorf("failed to check author existence: %w", err)
	}
	if !authorExists {
	    return reports, fmt.Errorf("author does not exist")
	}
    
    query := "SELECT id, author_id, count, title, style, language FROM reports WHERE author_id = ? AND style = ?"
    rows, err := DB.Query(query, AuthorID, Style)
    if err != nil {
        return reports, fmt.Errorf("failed to get reports by Style: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var report model.Report
        err := rows.Scan(&report.ID, &report.AuthorID, &report.Count, &report.Title, &report.Style, &report.Language)
        if err != nil {
            log.Println("Error scanning report:", err)
            continue
        }
        reports = append(reports, report)
    }
    return reports, nil
}

func (r *reportPersistence) GetByLanguage(DB *sql.DB, AuthorID, Language string) ([]model.Report, error) {
    var reports []model.Report
    
	var authorExists bool
	authorQuery := "SELECT COUNT(*) > 0 FROM users WHERE id = ?"
	err := DB.QueryRow(authorQuery, AuthorID).Scan(&authorExists)
	if err != nil {
		return reports, fmt.Errorf("failed to check author existence: %w", err)
	}
	if !authorExists {
	    return reports, fmt.Errorf("author does not exist")
	}

    query := "SELECT id, author_id, count, title, style, language FROM reports WHERE author_id = ? AND language = ?"
    rows, err := DB.Query(query, AuthorID, Language)
    if err != nil {
        return reports, fmt.Errorf("failed to get reports by Language: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var report model.Report
        err := rows.Scan(&report.ID, &report.AuthorID, &report.Count, &report.Title, &report.Style, &report.Language)
        if err != nil {
            log.Println("Error scanning report:", err)
            continue
        }
        reports = append(reports, report)
    }
    return reports, nil
}

func (r *reportPersistence) UpdateCount(DB *sql.DB, ID string, Count int) error {
    query := "UPDATE reports SET count = ? WHERE id = ?"
    _, err := DB.Exec(query, Count, ID)
    if err != nil {
        return fmt.Errorf("failed to update report count: %w", err)
    }
    return nil
}

func (r *reportPersistence) UpdateTitle(DB *sql.DB, ID, Title string) error {
    query := "UPDATE reports SET title = ? WHERE id = ?"
    _, err := DB.Exec(query, Title, ID)
    if err != nil {
        return fmt.Errorf("failed to update report title: %w", err)
    }
    return nil
}

func (r *reportPersistence) UpdateStyle(DB *sql.DB, ID, Style string) error {
    query := "UPDATE reports SET style = ? WHERE id = ?"
    _, err := DB.Exec(query, Style, ID)
    if err != nil {
        return fmt.Errorf("failed to update report style: %w", err)
    }
    return nil
}

func (r *reportPersistence) UpdateLanguage(DB *sql.DB, ID, Language string) error {
    query := "UPDATE reports SET language = ? WHERE id = ?"
    _, err := DB.Exec(query, Language, ID)
    if err != nil {
        return fmt.Errorf("failed to update report language: %w", err)
    }
    return nil
}

