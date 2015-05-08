package main

import (
	rethink "github.com/dancannon/gorethink"
	"log"
	"os"
	"time"
)

func InitDb() *rethink.Session {
	session, err := rethink.Connect(rethink.ConnectOpts{
		Address:  os.Getenv("RETHINKDB_URL"),
		Database: "test",
		MaxIdle:  10,
		Timeout:  time.Second * 10,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = rethink.DbCreate("test").Exec(session)
	if err != nil {
		log.Println(err)
	}

	_, err = rethink.Db("test").TableCreate("bookmarks").RunWrite(session)
	if err != nil {
		log.Println(err)
	}

	return session
}
