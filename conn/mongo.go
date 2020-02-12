package conn

import (
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func init() {
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB_NAME")
	fmt.Println("conn info:", host, dbName)
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("session err:", err)
		os.Exit(2)
	}
	db = session.DB(dbName)
}

// GetMongoDB function to return DB connection
func GetMongoDB() *mgo.Database {
	return db
}
