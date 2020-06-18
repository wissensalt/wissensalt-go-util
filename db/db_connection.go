package db

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/wissensalt/wissensalt-go-util/common"
	"github.com/wissensalt/wissensalt-go-util/logging"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type ConnectionProperty struct {
	userName string
	password string
	host string
	port string
	schema string
	dbInstance string
	maxOpenConnection int
	maxIdleConnection int
}

type ConnectionService interface {
	ConnectMySql() (*sql.DB, error)
	ConnectPostgreSql() (*sql.DB, error)
}

func (connectionProperty *ConnectionProperty) Init() {
	err := godotenv.Load()
	if err != nil {
		logging.AppLogger.Fatal("Error Loading .env variables")
	} else {
		maxOpenConnection, _ := strconv.Atoi(os.Getenv("db.max-open-connection"))
		maxIdleConnection, _ := strconv.Atoi(os.Getenv("db.max-idle-connection"))

		connectionProperty.userName = os.Getenv("db.username")
		connectionProperty.password = os.Getenv("db.password")
		connectionProperty.host = os.Getenv("db.host")
		connectionProperty.port = os.Getenv("db.port")
		connectionProperty.schema = os.Getenv("db.schema")
		connectionProperty.dbInstance = os.Getenv("db.instance")
		connectionProperty.maxOpenConnection = maxOpenConnection
		connectionProperty.maxIdleConnection = maxIdleConnection
	}
}


func (connectionProperty ConnectionProperty) ConnectMySql() (*sql.DB, error) {
	var bufferURLDB bytes.Buffer
	bufferURLDB.WriteString(connectionProperty.userName)
	bufferURLDB.WriteString(common.Colon)
	bufferURLDB.WriteString(connectionProperty.password)
	bufferURLDB.WriteString(common.At+"tcp"+common.OpenParenthesis)
	bufferURLDB.WriteString(connectionProperty.host)
	bufferURLDB.WriteString(common.Colon)
	bufferURLDB.WriteString(connectionProperty.port)
	bufferURLDB.WriteString(common.CloseParenthesis+common.Slash)
	bufferURLDB.WriteString(connectionProperty.schema)
	bufferURLDB.WriteString(common.Question+"parseTime"+common.Equals+common.True)

	activeDB, err := sql.Open(connectionProperty.dbInstance, bufferURLDB.String())
	if err != nil {
		logging.AppLogger.Fatal("Error Connecting to DB : ", err)
	} else {
		activeDB.SetMaxOpenConns(connectionProperty.maxOpenConnection)
		activeDB.SetMaxIdleConns(connectionProperty.maxIdleConnection)
		activeDB.SetConnMaxLifetime(time.Hour)
	}

	logging.AppLogger.Println("Successfully connected to db : "+connectionProperty.schema)
	return activeDB, err
}

func (connectionProperty ConnectionProperty) ConnectPostgreSql() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connectionProperty.host,
		connectionProperty.port,
		connectionProperty.userName,
		connectionProperty.password,
		connectionProperty.schema)
	activeDB, err := sql.Open(os.Getenv(connectionProperty.dbInstance), dbInfo)
	if err != nil {
		logging.AppLogger.Fatal("Error connection to DB : ", err)
	} else {
		activeDB.SetMaxOpenConns(connectionProperty.maxOpenConnection)
		activeDB.SetMaxIdleConns(connectionProperty.maxIdleConnection)
		activeDB.SetConnMaxLifetime(time.Hour)
	}

	return activeDB, err
}