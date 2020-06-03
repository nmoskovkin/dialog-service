package main

import (
	"database/sql"
	"dialogService/app/controller"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"strconv"
)

func main() {
	loadEnvFile()
	config := createConfigFromEnvVars()
	dbMaster := establishDbConnection(config.MysqlDSN)
	defer dbMaster.Close()
	migrateDatabase(dbMaster, "app/migrations")
	router := initRouter(dbMaster)
	startWebServer(router, config.Port)
}

func establishDbConnection(mysqlDsn string) *sql.DB {
	db, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

func loadEnvFile() {
	_ = godotenv.Load()
}

func createConfigFromEnvVars() *Config {
	config, err := NewConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}

func migrateDatabase(db *sql.DB, path string) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		"mysql", driver)

	if err != nil {
		panic(err.Error())
	}

	err = m.Up()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		panic(err.Error())
	}
}

func initRouter(dbMaster *sql.DB) *mux.Router {
	handlerFactory := controller.NewHandlerFactory()

	router := mux.NewRouter()
	router.HandleFunc(
		"/messages",
		handlerFactory.CreateHandler(controller.CreateMessagePostHandler(dbMaster)).ServeHTTP,
	).Methods(http.MethodPost)

	return router
}

func startWebServer(router *mux.Router, port uint16) {
	err := http.ListenAndServe(":"+strconv.Itoa(int(port)), router)
	if err != nil {
		panic(err.Error())
	}
}
