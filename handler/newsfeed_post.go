package handler

import (
	"encoding/json"
	"gochi/newsfeed"
	"net/http"
)

func NewsfeedPost(feed newsfeed.Adder) http.HandlerFunc{
	return func(writer http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		_ = json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		//_, _ = writer.Write([]byte("good job"))
	}
}