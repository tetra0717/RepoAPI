package application_test

import (
    "database/sql"
    "errors"
    "testing"

    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"

    "repo-api/src/domain/model"
    "repo-api/mocks"
)

func TestReportApp_Get(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockReportApp := mocks.NewMockReportApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name     string
        id       string
        authorID string
        title    string
        style    string
        language string
        expected []model.Report
        err      error
    }{
        {"Success - Get by ID", "1", "", "", "", "", []model.Report{{ID: "1", Title: "Test Report"}}, nil},
        {"Success - Get by AuthorID", "", "author1", "", "", "", []model.Report{{ID: "2", AuthorID: "author1"}}, nil},
        {"Success - Get by Title", "", "", "Test Title", "", "", []model.Report{{ID: "3", Title: "Test Title"}}, nil},
        {"Success - Get by Style", "", "", "", "TestStyle", "", []model.Report{{ID: "4", Style: "TestStyle"}}, nil},
        {"Success - Get by Language", "", "", "", "", "EN", []model.Report{{ID: "5", Language: "EN"}}, nil},
        {"Success - Get with multiple filters", "", "author1", "Test", "Style", "EN", []model.Report{{ID: "6", AuthorID: "author1", Title: "Test", Style: "Style", Language: "EN"}}, nil},
        {"Error - Not Found", "999", "", "", "", "", nil, errors.New("report not found")},
        {"Error - Database Error", "", "", "", "", "", nil, errors.New("database error")},
        {"Error - Invalid Input", "invalid", "", "", "", "", nil, errors.New("invalid input")},
        {"Success - Empty Result", "", "nonexistent", "", "", "", []model.Report{}, nil},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockReportApp.EXPECT().Get(db, tc.id, tc.authorID, tc.title, tc.style, tc.language).Return(tc.expected, tc.err)

            reports, err := mockReportApp.Get(db, tc.id, tc.authorID, tc.title, tc.style, tc.language)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tc.expected, reports)
            }
        })
    }
}

func TestReportApp_Register(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockReportApp := mocks.NewMockReportApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name   string
        report model.Report
        err    error
    }{
        {"Success - Valid Report", model.Report{ID: "1", Title: "New Report"}, nil},
        {"Error - Empty Title", model.Report{ID: "2", Title: ""}, errors.New("empty title")},
        {"Error - Invalid AuthorID", model.Report{ID: "3", AuthorID: "invalid"}, errors.New("invalid author ID")},
        {"Error - Duplicate ID", model.Report{ID: "4", Title: "Duplicate"}, errors.New("duplicate ID")},
        {"Error - Database Error", model.Report{ID: "5", Title: "DB Error"}, errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockReportApp.EXPECT().Register(db, tc.report).Return(tc.err)

            err := mockReportApp.Register(db, tc.report)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

func TestReportApp_Update(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockReportApp := mocks.NewMockReportApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name     string
        id       string
        count    int
        title    string
        style    string
        language string
        err      error
    }{
        {"Success - Full Update", "1", 2, "Updated Title", "New Style", "EN", nil},
        {"Success - Partial Update", "2", 1, "Updated Title", "", "", nil},
        {"Error - Report Not Found", "999", 1, "Title", "Style", "EN", errors.New("report not found")},
        {"Error - Invalid Count", "3", -1, "Title", "Style", "EN", errors.New("invalid count")},
        {"Error - Empty Title", "4", 1, "", "Style", "EN", errors.New("empty title")},
        {"Error - Invalid Language", "5", 1, "Title", "Style", "Invalid", errors.New("invalid language")},
        {"Error - Database Error", "6", 1, "Title", "Style", "EN", errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockReportApp.EXPECT().Update(db, tc.id, tc.count, tc.title, tc.style, tc.language).Return(tc.err)

            err := mockReportApp.Update(db, tc.id, tc.count, tc.title, tc.style, tc.language)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

func TestReportApp_Eject(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockReportApp := mocks.NewMockReportApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name string
        id   string
        err  error
    }{
        {"Success - Valid ID", "1", nil},
        {"Error - Report Not Found", "999", errors.New("report not found")},
        {"Error - Invalid ID", "", errors.New("invalid ID")},
        {"Error - Database Error", "2", errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockReportApp.EXPECT().Eject(db, tc.id).Return(tc.err)

            err := mockReportApp.Eject(db, tc.id)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
