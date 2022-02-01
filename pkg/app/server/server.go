package server

import (
	"fmt"
	"log"
	"net/http"

	bolt "go.etcd.io/bbolt"
)

func (s *Server) CreateBucket(name string) {
	if s.bolt != nil {
		s.bolt.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucket([]byte(name))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
			return nil
		})
	}
}

func (s *Server) SaveUrl(backet_name string, key string, value string) {
	if s.bolt != nil {
		s.bolt.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(backet_name))
			err := bucket.Put([]byte(key), []byte(value))
			return err
		})
	}
}

func (s *Server) GetUrl(backet_name string, key string, c chan string) {
	if s.bolt != nil {
		s.bolt.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(backet_name))
			value := bucket.Get([]byte(key))
			c <- string(value)
			return nil
		})
	}
}

func (s *Server) StartServer() {
	LongUrlHandler(s)
	ShortUrlHandler(s)
	srv := &http.Server{
		Handler: s.router,
		Addr:    s.conf.HttpServerHost,
	}
	log.Fatal(srv.ListenAndServe())
}
