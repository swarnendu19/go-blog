package connect

import (
	"fmt"

	"github.com/swarnendu19/go-blog/config"
	"github.com/swarnendu19/go-blog/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() *gorm.DB {
	dbConfig := config.GetDBConfig()
	dns := fmt.Sprintf(
		"host= %s user=%s password=%s dbname=%s port=%s sslmode=%s Timezone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Port,
		dbConfig.Sslmode,
		dbConfig.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection with database")
	}

	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Post{}, &models.Comment{})

	return db
}

func CloseDbConncetion(db *gorm.DB) {
	dbSql, err := db.DB()
	if err != nil {
		panic("Failed to close the database")
	}
	dbSql.Close()
}
