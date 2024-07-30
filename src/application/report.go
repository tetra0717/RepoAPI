package application

import (
	"database/sql"
	"fmt"
	"repo-api/src/domain/model"
	"repo-api/src/domain/repository"
)

type ReportApp interface {
	Register(DB *sql.DB, report model.Report) error
	Eject(DB *sql.DB, ID string) error
	Get(DB *sql.DB, ID, AuthorID, Title, Style, Language string) ([]model.Report, error)
	Update(DB *sql.DB, ID string, Count int, Title, Style, Language string) error
}

func NewReportApp(rr repository.IReportRepository) ReportApp {
	return &reportApp{
		reportRepository: rr,
	}
}

type reportApp struct {
	reportRepository repository.IReportRepository
}

func (r reportApp) Register(DB *sql.DB, report model.Report) error {
	err := r.reportRepository.Insert(DB, report.ID, report.AuthorID, report.Count, report.Title, report.Style, report.Language)
	if err != nil {
	    if err.Error() == "author does not exist" {
	        return fmt.Errorf("author does not exist")
        }
		return fmt.Errorf("failed to insert report with ID %s: %w", report.ID, err)
	}
	return nil
}

func (r reportApp) Eject(DB *sql.DB, ID string) error {
	err := r.reportRepository.Eject(DB, ID)
	if err != nil {
		if err.Error() == "report not found" {
			return fmt.Errorf("report not found")
		}
		return fmt.Errorf("failed to eject report with ID %s: %w", ID, err)
	}
	return nil
}

func (r reportApp) Get(DB *sql.DB, ID, AuthorID, Title, Style, Language string) ([]model.Report, error) {
	var reports []model.Report

	if ID != "" {
		var report model.Report
		report, err := r.reportRepository.GetByID(DB, ID)
		if err != nil {	
		    if err.Error() == "author does not exist" {
	            return nil, fmt.Errorf("author does not exist")
            }
			return nil, fmt.Errorf("failed to get report by ID %s: %w", ID, err)
		}
		reports = append(reports, report)
	} else {
		if AuthorID != "" {
			authorIDReports, err := r.reportRepository.GetByAuthorID(DB, AuthorID)
			if err != nil {	
			    if err.Error() == "author does not exist" {
	                return nil, fmt.Errorf("author does not exist")
                }
				return nil, fmt.Errorf("failed to get reports by AuthorID %s: %w", AuthorID, err)
			}
			reports = append(reports, authorIDReports...)
		}

		if Title != "" {
			var titleReports []model.Report
			titleReports, err := r.reportRepository.GetByTitle(DB, AuthorID, Title)
			if err != nil {
                if err.Error() == "author does not exist" {
	                return nil, fmt.Errorf("author does not exist")
                }
				return nil, fmt.Errorf("failed to get reports by Title %s: %w", Title, err)
			}
			reports = intersectReports(reports, titleReports)
		}

		if Style != "" {
			var styleReports []model.Report
			styleReports, err := r.reportRepository.GetByStyle(DB, AuthorID, Style)
			if err != nil {
			    if err.Error() == "author does not exist" {
	                return nil, fmt.Errorf("author does not exist")
                }
				return nil, fmt.Errorf("failed to get reports by Style %s: %w", Style, err)
			}
			reports = intersectReports(reports, styleReports)
		}

		if Language != "" {
			var languageReports []model.Report
			languageReports, err := r.reportRepository.GetByLanguage(DB, AuthorID, Language)
			if err != nil {
			    if err.Error() == "author does not exist" {
			        return nil, fmt.Errorf("author does not exist")
			    }
				return nil, fmt.Errorf("failed to get reports by Language %s: %w", Language, err)
			}
			reports = intersectReports(reports, languageReports)
		}
	}

	if len(reports) == 0 {
	    return nil, fmt.Errorf("report not found")
  }

	return reports, nil
}

func intersectReports(a, b []model.Report) []model.Report {
    set := make(map[string]model.Report)
    for _, report := range a {
        set[report.ID] = report
    }

    var result []model.Report
    for _, report := range b {
        if _, found := set[report.ID]; found {
            result = append(result, report)
        }
    }
    return result
}

func (r reportApp) Update(DB *sql.DB, ID string, Count int, Title, Style, Language string) error {
	if Count != 0 {
		if err := r.reportRepository.UpdateCount(DB, ID, Count); err != nil {
			return fmt.Errorf("failed to update title for report ID %s: %w", ID, err)
		}
	}
	
	if Title != "" {
		if err := r.reportRepository.UpdateTitle(DB, ID, Title); err != nil {
			return fmt.Errorf("failed to update title for report ID %s: %w", ID, err)
		}
	}

	if Style != "" {
		if err := r.reportRepository.UpdateStyle(DB, ID, Style); err != nil {
			return fmt.Errorf("failed to update style for report ID %s: %w", ID, err)
		}
	}

	if Language != "" {
		if err := r.reportRepository.UpdateLanguage(DB, ID, Language); err != nil {
			return fmt.Errorf("failed to update language for report ID %s: %w", ID, err)
		}
	}

	return nil
}


