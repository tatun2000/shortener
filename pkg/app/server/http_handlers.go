package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func LongUrlHandler(s *Server) {
	s.router.HandleFunc("/a/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		hash := MD5(url)[:8]
		s.SaveUrl(s.conf.BucketName, hash, url)
		fmt.Fprintf(w, "Short URL = %s", hash)
	})
}

func ShortUrlHandler(s *Server) {
	s.router.HandleFunc("/s/{short_url}", func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["short_url"]

		myCh := make(chan string)
		go s.GetUrl(s.conf.BucketName, shortUrl, myCh)
		longUrl := <-myCh
		http.Redirect(w, r, longUrl, s.conf.RedirectNum)
	})
}
