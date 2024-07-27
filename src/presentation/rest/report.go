package rest

import (
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    "repo-api/src/application"
    "repo-api/src/domain/model"
    "log"
    "github.com/google/uuid"
)

type ReportHandler interface {
    HandleRegisterReport(c *gin.Context)
    HandleEject(c *gin.Context)
    HandleGet(c *gin.Context)
    HandleUpdate(c *gin.Context)
}

func NewReportHandler(db *sql.DB, ar application.ReportApp) ReportHandler {
    return &reportHandler{
        database: db,
        reportApp: ar,
    }
}

type reportHandler struct {
    reportApp application.ReportApp
    database *sql.DB
}

func (r *reportHandler) HandleRegisterReport(c *gin.Context) {
    var report model.Report
    if err := c.BindJSON(&report); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }

    if report.AuthorID == "" || report.Count == 0 || report.Title == "" || report.Style == "" || report.Language == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "AuthorID, Count, Title, Style, and Language are required"})
        return
    }

    report.ID = uuid.New().String()
    
    if err := r.reportApp.Register(r.database, report); err != nil {
        log.Printf("Error retrieving user: %v", err)
        if err.Error() == "author does not exist" {
            c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register report"})
        }
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "Report registered successfully"})
}

func (r *reportHandler) HandleEject(c *gin.Context) {
    ID := c.Query("id")
    if ID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
        return
    }

    if err := r.reportApp.Eject(r.database, ID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to eject report"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Report ejected successfully"})
}

func (r *reportHandler) HandleGet(c *gin.Context) {
    var reports []model.Report
    var err error

    if ID := c.Query("id"); ID != "" {
        var report []model.Report
        report, err = r.reportApp.Get(r.database, ID, "", "", "", "")
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
            return
        } else if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve report"})
            return
        }
        reports = append(reports, report...)
    } else {
        if AuthorID := c.Query("author_id"); AuthorID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "AuthorID is required"})
            return
        }
        authorID := c.Query("author_id")
        title := c.Query("title")
        style := c.Query("style")
        language := c.Query("language")

        reports, err = r.reportApp.Get(r.database, "", authorID, title, style, language)
        if err != nil {
            if err.Error() == "author does not exist" {
                c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
                return
            } else if err.Error() == "report not found" {
                c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
                return
            }
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reports"})
            return
        }
    }
    
    c.JSON(http.StatusOK, gin.H{"reports": reports})
}

func (r *reportHandler) HandleUpdate(c *gin.Context) {
    var report model.Report
    if err := c.BindJSON(&report); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }

    if err := r.reportApp.Update(r.database, report.ID, report.Count, report.Title, report.Style, report.Language); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update report"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Report updated successfully"})
}
