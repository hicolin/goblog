package bootstrap

import (
	"github.com/gorilla/mux"
	"goblog/app/http/models/article"
	"goblog/app/http/models/user"
	"goblog/pkg/config"
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

	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&article.Article{},
	)
}
