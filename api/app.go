package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	v1 "rvgdb-service/api/routes/v1"
	"rvgdb-service/pkg/company"

	_ "github.com/lib/pq"
	_ "rvgdb-service/pkg/config"
)

func main() {
	db, err := initDb()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	cService := company.NewCompanyService(db)

	// Setup routes
	router := gin.Default()
	v1.CompanyRouter(router, cService)

	// Start server
	err = router.Run()
	if err != nil {
		log.Println(err)
	}
}

func initDb() (*sql.DB, error) {
	dbinfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.Get("database.host"),
		viper.GetInt("database.port"),
		viper.Get("database.user"),
		viper.Get("database.password"),
		viper.Get("database.dbname"),
		viper.Get("database.sslmode"),
	)

	return sql.Open("postgres", dbinfo)
}
