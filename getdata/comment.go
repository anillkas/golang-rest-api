package anillkas

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Comments struct {
	CommentData []Comment `json:"comments"`
}

type Comment struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func GetComment() Comments {
	dosya, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer dosya.Close()

	arr, err := ioutil.ReadAll(dosya)
	if err != nil {
		log.Fatal(err)
	}

	var comments Comments

	json.Unmarshal(arr, &comments)

	return comments

}
