package rest

import (
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    "repo-api/src/application"
    "repo-api/src/domain/model"
    "log"
    "errors"
)

type UserHandler interface {
  HandleRegisterUser(c *gin.Context)
  HandleGet(c *gin.Context)
  HandleUpdate(c *gin.Context)
}

func NewUserHandler(db *sql.DB, au application.UserApp) UserHandler {
  return &userHandler{
    database: db,
    userApp: au,
  }
}

type userHandler struct {
  userApp application.UserApp
  database *sql.DB
}

func (u userHandler) HandleRegisterUser(c *gin.Context) {
  var user model.User
  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
    return
  }

  if user.ID == "" || user.Name == "" {
     c.JSON(http.StatusBadRequest, gin.H{"error": "ID and Name are required"})
     return
   }

  if err := u.userApp.Register(u.database, user.ID, user.Name); err != nil {
    log.Printf("Error registering user: %v", err)
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (u userHandler) HandleGet(c *gin.Context) {
  ID := c.Query("id")
  if ID == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
    return
  }
  user, err := u.userApp.Get(u.database, ID)
  if err != nil {
    log.Printf("Error retrieving user: %v", err)
    if errors.Is(err, sql.ErrNoRows) {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
    }
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u userHandler) HandleUpdate(c *gin.Context) {
  var user model.User
  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
    return
  }
  
  if user.ID == "" || user.Name == "" {
     c.JSON(http.StatusBadRequest, gin.H{"error": "ID and Name are required"})
     return
   }

  err := u.userApp.Update(u.database, user.ID, user.Name); 
  if err != nil {
    log.Printf("Error updating user: %v", err)
    if err.Error() == "user not found" {
      c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
    } else {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
    }
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
