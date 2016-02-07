package dbconfig

import (
	"gopkg.in/mgo.v2"
	"time"
)

//const MongoDb details
const (
	hosts      = "ds063439.mongolab.com:63439"
	database   = "tsc"
	username   = "qnotify"
	password   = "quezx123"
	collection = "messages"
)

type DB struct {
	Session *mgo.Session
}

func (db *DB) DoDial() (s *mgo.Session, err error) {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	return mgo.DialWithInfo(info)
}

func (db *DB) Name() string {
	return "tsc"
}



//package dbconfig
//
//import (
//	"gopkg.in/mgo.v2"
//	"os"
//)
//
//type DB struct {
//	Session *mgo.Session
//}
//
//func (db *DB) DoDial() (s *mgo.Session, err error) {
//	return mgo.Dial(DBUrl())
//}
//
//func (db *DB) Name() string {
//	return "my_awesome_app"
//}
//
//func DBUrl() string {
//	dburl := os.Getenv("MONGOHQ_URL")
//
//	if dburl == "" {
//		dburl = "localhost"
//	}
//
//	return dburl
//}
