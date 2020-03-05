package connection

import (
	"errors"
	"fmt"
	"github.com/danny2204/KarcisCoAPI/middleware"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var (
	dbName string
	dbPass string
	dbUser string
	dbType string
	dbHost string
	dbPort string
	sslMode string
	ApiPort string
	connectionString string
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("tol")
	}

	dbPort = os.Getenv("dbPort")
	dbHost = os.Getenv("dbHost")
	dbName = os.Getenv("dbName")
	dbUser = os.Getenv("dbUser")
	dbPass = os.Getenv("dbPass")
	dbType = os.Getenv("dbType")
	sslMode = os.Getenv("sslMode")
	ApiPort = os.Getenv("apiPort")

	connectionString = "host=" + dbHost + " user=" + dbUser + " dbname=" + dbName + " sslmode=" + sslMode + " password=" + dbPass + " port=" + dbPort
	fmt.Print(connectionString)
}

func ConnectDatabase()(*gorm.DB, error) {
	return gorm.Open(dbType, connectionString)
}

func ReturnError(s string) error {
	return errors.New(s)
}

func NewRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.LogMiddleware)

	return r
}
