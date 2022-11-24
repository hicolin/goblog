package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/app/http/models/article"
	"goblog/app/http/models/user"
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/routes"
	"gorm.io/gorm"
	"time"
)

func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}

func SetupDB() {
	db := model.ConnectDB()
	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
