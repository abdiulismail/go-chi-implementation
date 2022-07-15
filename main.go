package main

import (
	"github.com/go-chi/chi/v5"
	"gochi/handler"
	"gochi/newsfeed"
	"net/http"
)

func main(){
	port := ":3000"
	r := chi.NewRouter()

	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "world",
	})

	r.Get("/newsfeed",handler.NewsfeedGet(feed))
	r.Post("/newsfeed",handler.NewsfeedPost(feed))

	_ = http.ListenAndServe(port, r)
}