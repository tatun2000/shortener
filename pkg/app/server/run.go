package server

import (
	"log"
	"shortener/pkg/app/config"

	"github.com/gorilla/mux"
	bolt "go.etcd.io/bbolt"
)

type Server struct {
	bolt   *bolt.DB
	router *mux.Router
	conf   *config.Config
}

func Run() {
	conf := config.GetConf()
	s := Server{
		bolt:   ConnectToBolt(conf),
		router: mux.NewRouter(),
		conf:   conf,
	}
	defer s.bolt.Close()
	s.CreateBucket(s.conf.BucketName)
	s.StartServer()
}

func ConnectToBolt(conf *config.Config) *bolt.DB {
	db, err := bolt.Open(conf.DbName, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
