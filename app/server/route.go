package server

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title1", Text: "Text 1"}}
}
func getPosts(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error mashalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error  the posts array"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusAccepted)
	result, err := json.Marshal(post)
	resp.Write(result)
}
