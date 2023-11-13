package main

import (
	anillkas "anillkas/getdata"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetTodoData(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	idIndex := -1
	for i, segment := range path {
		if segment == "todo" && i+1 < len(path) {
			idIndex = i + 1
			break
		}
	}
	if idIndex == -1 || idIndex >= len(path) {
		http.Error(w, "ID parameter not found", http.StatusBadRequest)
		return
	}

	id := path[idIndex]
	intId, _ := strconv.Atoi(id)
	listeData := anillkas.GetTodo()
	for i := 0; i < len(listeData.TodoData); i++ {
		if listeData.TodoData[i].Id == intId {
			jsonData, _ := json.Marshal(listeData.TodoData[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}
func GetCommentData(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	idIndex := -1
	for i, segment := range path {
		if segment == "comment" && i+1 < len(path) {
			idIndex = i + 1
			break
		}
	}
	if idIndex == -1 || idIndex >= len(path) {
		http.Error(w, "ID parameter not found", http.StatusBadRequest)
		return
	}

	id := path[idIndex]
	intId, _ := strconv.Atoi(id)
	listeData := anillkas.GetComment()
	for i := 0; i < len(listeData.CommentData); i++ {
		if listeData.CommentData[i].Id == intId {
			jsonData, _ := json.Marshal(listeData.CommentData[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}
func GetPostData(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")

	idIndex := -1
	for i, segment := range path {
		if segment == "post" && i+1 < len(path) {
			idIndex = i + 1
			break
		}
	}
	if idIndex == -1 || idIndex >= len(path) {
		http.Error(w, "ID parameter not found", http.StatusBadRequest)
		return
	}

	id := path[idIndex]
	intId, _ := strconv.Atoi(id)
	listeData := anillkas.GetPost()
	for i := 0; i < len(listeData.PostData); i++ {
		if listeData.PostData[i].Id == intId {
			jsonData, _ := json.Marshal(listeData.PostData[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}
	}
}

func AnasayfaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Anasayfa")
}

func main() {
	http.HandleFunc("/", AnasayfaHandler)
	http.HandleFunc("/todo/", GetTodoData)
	http.HandleFunc("/comment/", GetCommentData)
	http.HandleFunc("/post/", GetPostData)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
