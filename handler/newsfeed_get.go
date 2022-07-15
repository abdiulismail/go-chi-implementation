package handler

import (
	"encoding/json"
	"gochi/newsfeed"
	"net/http"
)

func NewsfeedGet(feed newsfeed.Getter) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		items := feed.GetAll()
		_ = json.NewEncoder(writer).Encode(items)
	}
}