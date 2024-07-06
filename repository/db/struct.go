package db

import (
	"github.com/SchoolAF/exodus/config"
	mongo "github.com/kamagasaki/go-utils/mongo"
	"os"
)

var MongoString = os.Getenv("MONGOSTRCONNECT")

var DBHLCYN = mongo.DBInfo{
	DBString: MongoString,
	DBName:   config.HalcyonMDB,
}
