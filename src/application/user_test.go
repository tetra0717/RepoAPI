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

func TestUserApp_Get(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserApp := mocks.NewMockUserApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name     string
        id       string
        expected model.User
        err      error
    }{
        {"Success - Valid User", "user1", model.User{ID: "user1", Name: "John Doe"}, nil},
        {"Error - User Not Found", "nonexistent", model.User{}, errors.New("user not found")},
        {"Error - Empty ID", "", model.User{}, errors.New("invalid ID")},
        {"Error - Database Error", "dberror", model.User{}, errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockUserApp.EXPECT().Get(db, tc.id).Return(tc.expected, tc.err)

            user, err := mockUserApp.Get(db, tc.id)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tc.expected, user)
            }
        })
    }
}

func TestUserApp_Register(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserApp := mocks.NewMockUserApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name string
        id   string
        userName string
        err  error
    }{
        {"Success - Valid User", "user1", "John Doe", nil},
        {"Error - Empty ID", "", "John Doe", errors.New("invalid ID")},
        {"Error - Empty Name", "user2", "", errors.New("invalid name")},
        {"Error - Duplicate ID", "duplicate", "Jane Doe", errors.New("user already exists")},
        {"Error - Database Error", "dberror", "DB Error", errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockUserApp.EXPECT().Register(db, tc.id, tc.userName).Return(tc.err)

            err := mockUserApp.Register(db, tc.id, tc.userName)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
            }
        })
    }
}

func TestUserApp_Update(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserApp := mocks.NewMockUserApp(ctrl)
    db := &sql.DB{}

    testCases := []struct {
        name string
        id   string
        newName string
        err  error
    }{
        {"Success - Valid Update", "user1", "Jane Doe", nil},
        {"Error - User Not Found", "nonexistent", "New Name", errors.New("user not found")},
        {"Error - Empty ID", "", "New Name", errors.New("invalid ID")},
        {"Error - Empty Name", "user2", "", errors.New("invalid name")},
        {"Error - Database Error", "dberror", "DB Error", errors.New("database error")},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            mockUserApp.EXPECT().Update(db, tc.id, tc.newName).Return(tc.err)

            err := mockUserApp.Update(db, tc.id, tc.newName)

            if tc.err != nil {
                assert.Error(t, err)
                assert.Equal(t, tc.err.Error(), err.Error())
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
