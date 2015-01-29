package data

import (
	"fmt"
	"strings"

	"upper.io/db"
	"upper.io/db/mongo"
)

var DB db.Database

type DBConf struct {
	Database string   `toml:"database"`
	Hosts    []string `toml:"hosts"`
}

func (cf *DBConf) ConnectionUrl() string {
	return fmt.Sprintf("mongodb://%s/%s", strings.Join(cf.Hosts, ","), cf.Database)
}

func NewDBSession(conf DBConf) (db.Database, error) {
	connUrl, err := mongo.ParseURL(conf.ConnectionUrl())
	if err != nil {
		return nil, err
	}
	d, err := db.Open(mongo.Adapter, connUrl)
	if err != nil {
		return nil, err
	}
	DB = d
	setupCollections()
	return d, nil
}

func setupCollections() {
	UserCollection, _ = DB.Collection("users")
	TodoCollection, _ = DB.Collection("todos")
}
