package main

import (
  "log"
	"github.com/gin-gonic/gin"
	"repo-api/src/infra/persistence"
	"repo-api/src/application"
	"repo-api/src/presentation/rest"
	"repo-api/src/infra"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()


  userPresistence := persistence.NewUserPersistence() 
  userApp := application.NewUserApp(userPresistence)
  userHandler := rest.NewUserHandler(db, userApp)

  reportPresistence := persistence.NewReportPersistence()
  reportApp := application.NewReportApp(reportPresistence)
  reportHandler := rest.NewReportHandler(db, reportApp)
  
  router := gin.Default()
  router.POST("/user", userHandler.HandleRegisterUser)
  router.GET("/user", userHandler.HandleGet)
  router.PUT("/user", userHandler.HandleUpdate)
  router.POST("/report", reportHandler.HandleRegisterReport)
  router.GET("/report", reportHandler.HandleGet)
  router.PUT("/report", reportHandler.HandleUpdate)
  router.DELETE("/report", reportHandler.HandleEject)
  router.Run(":8080")
}
